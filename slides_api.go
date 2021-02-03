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
func (a *SlidesApiService) CopyFile(request CopyFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.SrcPath) == 0 {
		return nil, reportError("Missing required parameter request.srcPath")
	}
	if len(request.DestPath) == 0 {
		return nil, reportError("Missing required parameter request.destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.SrcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.SrcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.DestStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.VersionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(request.DestPath, ""))
	if localVarTempParam := request.SrcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.DestStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.VersionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) CopyFolder(request CopyFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.SrcPath) == 0 {
		return nil, reportError("Missing required parameter request.srcPath")
	}
	if len(request.DestPath) == 0 {
		return nil, reportError("Missing required parameter request.destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/copy/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.SrcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.SrcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.DestStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(request.DestPath, ""))
	if localVarTempParam := request.SrcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.DestStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) CreateFolder(request CreateFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Path) == 0 {
		return nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteChartCategory(request DeleteChartCategoryRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories/{categoryIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	categoryIndexPathStringValue := fmt.Sprintf("%v", request.CategoryIndex)
	if len(categoryIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"categoryIndex"+"}", categoryIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"categoryIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteChartDataPoint(request DeleteChartDataPointRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints/{pointIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", request.SeriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}
	pointIndexPathStringValue := fmt.Sprintf("%v", request.PointIndex)
	if len(pointIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"pointIndex"+"}", pointIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"pointIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteChartSeries(request DeleteChartSeriesRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", request.SeriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteFile(request DeleteFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Path) == 0 {
		return nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.VersionId, "string", "versionId"); err != nil {
		return nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.VersionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteFolder(request DeleteFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Path) == 0 {
		return nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return nil, err
	}
	if request.Recursive != nil {
		if err := typeCheckParameter(*request.Recursive, "bool", "recursive"); err != nil {
			return nil, err
		}
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if request.Recursive != nil {
		localVarQueryParams.Add("Recursive", parameterToString(request.Recursive, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlide(request DeleteNotesSlideRequest) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlideParagraph(request DeleteNotesSlideParagraphRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlideParagraphs(request DeleteNotesSlideParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Paragraphs, "[]int32", "paragraphs"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlidePortion(request DeleteNotesSlidePortionRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlidePortions(request DeleteNotesSlidePortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Portions, "[]int32", "portions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlideShape(request DeleteNotesSlideShapeRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteNotesSlideShapes(request DeleteNotesSlideShapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Shapes, "[]int32", "shapes"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteParagraph(request DeleteParagraphRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteParagraphs(request DeleteParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Paragraphs, "[]int32", "paragraphs"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeletePortion(request DeletePortionRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeletePortions(request DeletePortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Portions, "[]int32", "portions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSection(request DeleteSectionRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", request.SectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.WithSlides != nil {
		if err := typeCheckParameter(*request.WithSlides, "bool", "withSlides"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if request.WithSlides != nil {
		localVarQueryParams.Add("WithSlides", parameterToString(request.WithSlides, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSections(request DeleteSectionsRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Sections, "[]int32", "sections"); err != nil {
		return successPayload, nil, err
	}
	if request.WithSlides != nil {
		if err := typeCheckParameter(*request.WithSlides, "bool", "withSlides"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Sections; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Sections", parameterToString(localVarTempParam, ""))
	}
	if request.WithSlides != nil {
		localVarQueryParams.Add("WithSlides", parameterToString(request.WithSlides, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimation(request DeleteSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimationEffect(request DeleteSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", request.EffectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequence(request DeleteSlideAnimationInteractiveSequenceRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", request.SequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequenceEffect(request DeleteSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", request.SequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", request.EffectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimationInteractiveSequences(request DeleteSlideAnimationInteractiveSequencesRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideAnimationMainSequence(request DeleteSlideAnimationMainSequenceRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideByIndex(request DeleteSlideByIndexRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideShape(request DeleteSlideShapeRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideShapes(request DeleteSlideShapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Shapes, "[]int32", "shapes"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideSubshape(request DeleteSlideSubshapeRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlideSubshapes(request DeleteSlideSubshapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Shapes, "[]int32", "shapes"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Shapes; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Shapes", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlidesCleanSlidesList(request DeleteSlidesCleanSlidesListRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Slides, "[]int32", "slides"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Slides; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Slides", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlidesDocumentProperties(request DeleteSlidesDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSlidesDocumentProperty(request DeleteSlidesDocumentPropertyRequest) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.PropertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.propertyName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.PropertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.DeleteSlidesDocumentProperty
*/
type DeleteSlidesDocumentPropertyRequest struct {
    Name string
    PropertyName string
    Password string
    Folder string
    Storage string
}

/* SlidesApiService Remove background from a slide.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) DeleteSlidesSlideBackground(request DeleteSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSubshapeParagraph(request DeleteSubshapeParagraphRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSubshapeParagraphs(request DeleteSubshapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Paragraphs, "[]int32", "paragraphs"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Paragraphs; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Paragraphs", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSubshapePortion(request DeleteSubshapePortionRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) DeleteSubshapePortions(request DeleteSubshapePortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Portions, "[]int32", "portions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Portions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Portions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.VersionId, "string", "versionId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.VersionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetDiscUsage(request GetDiscUsageRequest) (IDiscUsage, *http.Response, error) {
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
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetFileVersions(request GetFileVersionsRequest) (IFileVersions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFileVersions
	)

	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/version/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetFilesList(request GetFilesListRequest) (IFilesList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFilesList
	)

	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetLayoutSlide(request GetLayoutSlideRequest) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetLayoutSlidesList(request GetLayoutSlidesListRequest) (ILayoutSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetMasterSlide(request GetMasterSlideRequest) (IMasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetMasterSlidesList(request GetMasterSlidesListRequest) (IMasterSlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlide(request GetNotesSlideRequest) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideExists(request GetNotesSlideExistsRequest) (IEntityExists, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IEntityExists
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/exist"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return NotesSlideHeaderFooter*/
func (a *SlidesApiService) GetNotesSlideHeaderFooter(request GetNotesSlideHeaderFooterRequest) (INotesSlideHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlideHeaderFooter
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
    Storage string
    Folder string
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
func (a *SlidesApiService) GetNotesSlideShape(request GetNotesSlideShapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideShapeParagraph(request GetNotesSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideShapeParagraphs(request GetNotesSlideShapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideShapePortion(request GetNotesSlideShapePortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideShapePortions(request GetNotesSlideShapePortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideShapes(request GetNotesSlideShapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetNotesSlideWithFormat(request GetNotesSlideWithFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetParagraphPortion(request GetParagraphPortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetParagraphPortions(request GetParagraphPortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSections(request GetSectionsRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideAnimation(request GetSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ShapeIndex != nil {
		if err := typeCheckParameter(*request.ShapeIndex, "int32", "shapeIndex"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if request.ShapeIndex != nil {
		localVarQueryParams.Add("ShapeIndex", parameterToString(*request.ShapeIndex, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideHeaderFooter(request GetSlideHeaderFooterRequest) (IHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IHeaderFooter
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideShape(request GetSlideShapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideShapeParagraph(request GetSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideShapeParagraphs(request GetSlideShapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideShapes(request GetSlideShapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideSubshape(request GetSlideSubshapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideSubshapeParagraph(request GetSlideSubshapeParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideSubshapeParagraphs(request GetSlideSubshapeParagraphsRequest) (IParagraphs, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraphs
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlideSubshapes(request GetSlideSubshapesRequest) (IShapes, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapes
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesApiInfo() (IApiInfo, *http.Response, error) {
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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


/* SlidesApiService Read presentation info.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Documentstorage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) GetSlidesDocument(request GetSlidesDocumentRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.GetSlidesDocument
*/
type GetSlidesDocumentRequest struct {
    Name string
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Read presentation document properties.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return DocumentProperties*/
func (a *SlidesApiService) GetSlidesDocumentProperties(request GetSlidesDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesDocumentProperty(request GetSlidesDocumentPropertyRequest) (IDocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperty
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.PropertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.propertyName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.PropertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesImageWithDefaultFormat(request GetSlidesImageWithDefaultFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", request.Index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesImageWithFormat(request GetSlidesImageWithFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images/{index}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	indexPathStringValue := fmt.Sprintf("%v", request.Index)
	if len(indexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", indexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"index"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesImages(request GetSlidesImagesRequest) (IImages, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IImages
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/images"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesPlaceholder(request GetSlidesPlaceholderRequest) (IPlaceholder, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPlaceholder
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders/{placeholderIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	placeholderIndexPathStringValue := fmt.Sprintf("%v", request.PlaceholderIndex)
	if len(placeholderIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"placeholderIndex"+"}", placeholderIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"placeholderIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesPlaceholders(request GetSlidesPlaceholdersRequest) (IPlaceholders, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPlaceholders
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/placeholders"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesPresentationTextItems(request GetSlidesPresentationTextItemsRequest) (ITextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITextItems
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/textItems"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.WithEmpty != nil {
		if err := typeCheckParameter(*request.WithEmpty, "bool", "withEmpty"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if request.WithEmpty != nil {
		localVarQueryParams.Add("WithEmpty", parameterToString(request.WithEmpty, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesProtectionProperties(request GetSlidesProtectionPropertiesRequest) (IProtectionProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IProtectionProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/protectionProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlide(request GetSlidesSlideRequest) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlideBackground(request GetSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlideComments(request GetSlidesSlideCommentsRequest) (ISlideComments, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideComments
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/comments"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlideImages(request GetSlidesSlideImagesRequest) (IImages, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IImages
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/images"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlideProperties(request GetSlidesSlidePropertiesRequest) (ISlideProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "withEmpty" (bool) True to incude empty items.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return TextItems*/
func (a *SlidesApiService) GetSlidesSlideTextItems(request GetSlidesSlideTextItemsRequest) (ITextItems, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITextItems
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/textItems"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.WithEmpty != nil {
		if err := typeCheckParameter(*request.WithEmpty, "bool", "withEmpty"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if request.WithEmpty != nil {
		localVarQueryParams.Add("WithEmpty", parameterToString(request.WithEmpty, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesSlidesList(request GetSlidesSlidesListRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesTheme(request GetSlidesThemeRequest) (ITheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ITheme
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesThemeColorScheme(request GetSlidesThemeColorSchemeRequest) (IColorScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IColorScheme
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/colorScheme"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesThemeFontScheme(request GetSlidesThemeFontSchemeRequest) (IFontScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFontScheme
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/fontScheme"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesThemeFormatScheme(request GetSlidesThemeFormatSchemeRequest) (IFormatScheme, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFormatScheme
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/theme/formatScheme"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSlidesViewProperties(request GetSlidesViewPropertiesRequest) (IViewProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IViewProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/viewProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSubshapeParagraphPortion(request GetSubshapeParagraphPortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) GetSubshapeParagraphPortions(request GetSubshapeParagraphPortionsRequest) (IPortions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortions
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) MoveFile(request MoveFileRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.SrcPath) == 0 {
		return nil, reportError("Missing required parameter request.srcPath")
	}
	if len(request.DestPath) == 0 {
		return nil, reportError("Missing required parameter request.destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.SrcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.SrcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.DestStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.VersionId, "string", "versionId"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(request.DestPath, ""))
	if localVarTempParam := request.SrcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.DestStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.VersionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) MoveFolder(request MoveFolderRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.SrcPath) == 0 {
		return nil, reportError("Missing required parameter request.srcPath")
	}
	if len(request.DestPath) == 0 {
		return nil, reportError("Missing required parameter request.destPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/folder/move/{srcPath}"
	srcPathPathStringValue := fmt.Sprintf("%v", request.SrcPath)
	if len(srcPathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", srcPathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"srcPath"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.SrcStorageName, "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.DestStorageName, "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(request.DestPath, ""))
	if localVarTempParam := request.SrcStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.DestStorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) ObjectExists(request ObjectExistsRequest) (IObjectExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IObjectExist
	)

	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/exist/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.VersionId, "string", "versionId"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.VersionId; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("VersionId", parameterToString(localVarTempParam, ""))
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return Paragraph*/
func (a *SlidesApiService) PostAddNewParagraph(request PostAddNewParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewParagraph
*/
type PostAddNewParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
    Position *int32
}

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
 @return Portion*/
func (a *SlidesApiService) PostAddNewPortion(request PostAddNewPortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewPortion
*/
type PostAddNewPortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
    Position *int32
}

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeBase*/
func (a *SlidesApiService) PostAddNewShape(request PostAddNewShapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ShapeToClone != nil {
		if err := typeCheckParameter(*request.ShapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ShapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*request.ShapeToClone, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewShape
*/
type PostAddNewShapeRequest struct {
    Name string
    SlideIndex int32
    Dto IShapeBase
    Password string
    Folder string
    Storage string
    ShapeToClone *int32
    Position *int32
}

/* SlidesApiService Create new shape (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dto" (ShapeBase) Shape DTO.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeBase*/
func (a *SlidesApiService) PostAddNewSubshape(request PostAddNewSubshapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ShapeToClone != nil {
		if err := typeCheckParameter(*request.ShapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ShapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*request.ShapeToClone, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewSubshape
*/
type PostAddNewSubshapeRequest struct {
    Name string
    SlideIndex int32
    Path string
    Dto IShapeBase
    Password string
    Folder string
    Storage string
    ShapeToClone *int32
    Position *int32
}

/* SlidesApiService Creates new paragraph (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return Paragraph*/
func (a *SlidesApiService) PostAddNewSubshapeParagraph(request PostAddNewSubshapeParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewSubshapeParagraph
*/
type PostAddNewSubshapeParagraphRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
    Position *int32
}

/* SlidesApiService Creates new portion (for smart art and group shapes).
 @param name Document name.
 @param slideIndex Slide index.
 @param path Shape path.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
 @return Portion*/
func (a *SlidesApiService) PostAddNewSubshapePortion(request PostAddNewSubshapePortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostAddNewSubshapePortion
*/
type PostAddNewSubshapePortionRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
    Position *int32
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
func (a *SlidesApiService) PostAddNotesSlide(request PostAddNotesSlideRequest) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostChartCategory(request PostChartCategoryRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Category == nil {
		return successPayload, nil, reportError("Missing required parameter request.category")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Category
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostChartDataPoint(request PostChartDataPointRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.DataPoint == nil {
		return successPayload, nil, reportError("Missing required parameter request.dataPoint")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", request.SeriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.DataPoint
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostChartSeries(request PostChartSeriesRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Series == nil {
		return successPayload, nil, reportError("Missing required parameter request.series")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Series
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostCopyLayoutSlideFromSourcePresentation(request PostCopyLayoutSlideFromSourcePresentationRequest) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.CloneFrom) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.cloneFrom")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.CloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.CloneFromStorage, "string", "cloneFromStorage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("CloneFrom", parameterToString(request.CloneFrom, ""))
	localVarQueryParams.Add("CloneFromPosition", parameterToString(request.CloneFromPosition, ""))
	if localVarTempParam := request.CloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("CloneFromStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.CloneFromPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["CloneFromPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostCopyMasterSlideFromSourcePresentation(request PostCopyMasterSlideFromSourcePresentationRequest) (IMasterSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IMasterSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.CloneFrom) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.cloneFrom")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/masterSlides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.CloneFromPassword, "string", "cloneFromPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.CloneFromStorage, "string", "cloneFromStorage"); err != nil {
		return successPayload, nil, err
	}
	if request.ApplyToAll != nil {
		if err := typeCheckParameter(*request.ApplyToAll, "bool", "applyToAll"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("CloneFrom", parameterToString(request.CloneFrom, ""))
	localVarQueryParams.Add("CloneFromPosition", parameterToString(request.CloneFromPosition, ""))
	if localVarTempParam := request.CloneFromStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("CloneFromStorage", parameterToString(localVarTempParam, ""))
	}
	if request.ApplyToAll != nil {
		localVarQueryParams.Add("ApplyToAll", parameterToString(request.ApplyToAll, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.CloneFromPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["CloneFromPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* SlidesApiService Read notes slide info.
 @param document Document data.
 @param slideIndex Slide index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
 @return NotesSlide*/
func (a *SlidesApiService) PostGetNotesSlide(request PostGetNotesSlideRequest) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(request.Document) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide"
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Document
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostGetNotesSlideExists(request PostGetNotesSlideExistsRequest) (IEntityExists, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IEntityExists
	)

	if len(request.Document) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.document")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide/exist"
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/octet-stream",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Document
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostGetNotesSlideWithFormat(request PostGetNotesSlideWithFormatRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Document) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.document")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/slides/{slideIndex}/notesSlide/{format}"
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Document
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* SlidesApiService Creates new paragraph.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param dto Paragraph DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new paragraph in the list. Default is at the end of the list.
 @return Paragraph*/
func (a *SlidesApiService) PostNotesSlideAddNewParagraph(request PostNotesSlideAddNewParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostNotesSlideAddNewParagraph
*/
type PostNotesSlideAddNewParagraphRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Dto IParagraph
    Password string
    Folder string
    Storage string
    Position *int32
}

/* SlidesApiService Creates new portion.
 @param name Document name.
 @param slideIndex Slide index.
 @param shapeIndex Shape index.
 @param paragraphIndex Paragraph index.
 @param dto Portion DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "position" (int32) Position of the new portion in the list. Default is at the end of the list.
 @return Portion*/
func (a *SlidesApiService) PostNotesSlideAddNewPortion(request PostNotesSlideAddNewPortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostNotesSlideAddNewPortion
*/
type PostNotesSlideAddNewPortionRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    ParagraphIndex int32
    Dto IPortion
    Password string
    Folder string
    Storage string
    Position *int32
}

/* SlidesApiService Create new shape.
 @param name Document name.
 @param slideIndex Slide index.
 @param dto Shape DTO.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "shapeToClone" (int32) Optional index for clone shape instead of adding a new one.
     @param "position" (int32) Position of the new shape in the list. Default is at the end of the list.
 @return ShapeBase*/
func (a *SlidesApiService) PostNotesSlideAddNewShape(request PostNotesSlideAddNewShapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ShapeToClone != nil {
		if err := typeCheckParameter(*request.ShapeToClone, "int32", "shapeToClone"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ShapeToClone != nil {
		localVarQueryParams.Add("ShapeToClone", parameterToString(*request.ShapeToClone, ""))
	}
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostNotesSlideAddNewShape
*/
type PostNotesSlideAddNewShapeRequest struct {
    Name string
    SlideIndex int32
    Dto IShapeBase
    Password string
    Folder string
    Storage string
    ShapeToClone *int32
    Position *int32
}

/* SlidesApiService Render shape to specified picture format.
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
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostNotesSlideShapeSaveAs(request PostNotesSlideShapeSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostNotesSlideShapeSaveAs
*/
type PostNotesSlideShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
    FontsFolder string
}

/* SlidesApiService Merge the presentation with other presentations specified in the request parameter.
 @param name Document name.
 @param request PresentationsMergeRequest with a list of presentations to merge.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostPresentationMerge(request PostPresentationMergeRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Request == nil {
		return successPayload, nil, reportError("Missing required parameter request.request")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostPresentationMerge
*/
type PostPresentationMergeRequest struct {
    Name string
    Request IPresentationsMergeRequest
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PostSection(request PostSectionRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.SectionName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.sectionName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("SectionName", parameterToString(request.SectionName, ""))
	localVarQueryParams.Add("SlideIndex", parameterToString(request.SlideIndex, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSectionMove(request PostSectionMoveRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}/move"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", request.SectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("NewPosition", parameterToString(request.NewPosition, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostShapeSaveAs(request PostShapeSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostShapeSaveAs
*/
type PostShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
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
func (a *SlidesApiService) PostSlideAnimationEffect(request PostSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Effect == nil {
		return successPayload, nil, reportError("Missing required parameter request.effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlideAnimationInteractiveSequence(request PostSlideAnimationInteractiveSequenceRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Sequence == nil {
		return successPayload, nil, reportError("Missing required parameter request.sequence")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Sequence
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlideAnimationInteractiveSequenceEffect(request PostSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Effect == nil {
		return successPayload, nil, reportError("Missing required parameter request.effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", request.SequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "position" (int32) The target position at which to create the slide. Add to the end by default.
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
     @param "layoutAlias" (string) Alias of layout slide for new slide. Alias may be the type of layout, name of layout slide or index
 @return Slides*/
func (a *SlidesApiService) PostSlidesAdd(request PostSlidesAddRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.LayoutAlias, "string", "layoutAlias"); err != nil {
		return successPayload, nil, err
	}

	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.LayoutAlias; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("LayoutAlias", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesAdd
*/
type PostSlidesAddRequest struct {
    Name string
    Position *int32
    Password string
    Folder string
    Storage string
    LayoutAlias string
}

/* SlidesApiService Convert presentation from request content to format specified.
 @param document Document data.
 @param format Export format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesConvert(request PostSlidesConvertRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Document) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.document")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Document
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesConvert
*/
type PostSlidesConvertRequest struct {
    Document []byte
    Format string
    Password string
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
func (a *SlidesApiService) PostSlidesCopy(request PostSlidesCopyRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/copy"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Position != nil {
		if err := typeCheckParameter(*request.Position, "int32", "position"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Source, "string", "source"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.SourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.SourceStorage, "string", "sourceStorage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("SlideToCopy", parameterToString(request.SlideToCopy, ""))
	if request.Position != nil {
		localVarQueryParams.Add("Position", parameterToString(*request.Position, ""))
	}
	if localVarTempParam := request.Source; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Source", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.SourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourceStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.SourcePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["SourcePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocument(request PostSlidesDocumentRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.InputPassword, "string", "inputPassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.InputPassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["InputPassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Data
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesDocument
*/
type PostSlidesDocumentRequest struct {
    Name string
    Data []byte
    InputPassword string
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Create presentation document from html.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "html" (string) HTML data.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocumentFromHtml(request PostSlidesDocumentFromHtmlRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromHtml"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Html, "string", "html"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Html
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesDocumentFromHtml
*/
type PostSlidesDocumentFromHtmlRequest struct {
    Name string
    Html string
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Create presentation document from pdf or append pdf to an existing presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pdf" ([]byte) PDF data.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PostSlidesDocumentFromPdf(request PostSlidesDocumentFromPdfRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromPdf"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/octet-stream",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Pdf
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesDocumentFromPdf
*/
type PostSlidesDocumentFromPdfRequest struct {
    Name string
    Pdf []byte
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PostSlidesDocumentFromSource(request PostSlidesDocumentFromSourceRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromSource"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.SourcePath, "string", "sourcePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.SourcePassword, "string", "sourcePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.SourceStorage, "string", "sourceStorage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.SourcePath; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourcePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.SourceStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SourceStorage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.SourcePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["SourcePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesDocumentFromSource
*/
type PostSlidesDocumentFromSourceRequest struct {
    Name string
    SourcePath string
    SourcePassword string
    SourceStorage string
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PostSlidesDocumentFromTemplate(request PostSlidesDocumentFromTemplateRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.TemplatePath) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.templatePath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromTemplate"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Data, "string", "data"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.TemplatePassword, "string", "templatePassword"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.TemplateStorage, "string", "templateStorage"); err != nil {
		return successPayload, nil, err
	}
	if request.IsImageDataEmbedded != nil {
		if err := typeCheckParameter(*request.IsImageDataEmbedded, "bool", "isImageDataEmbedded"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("TemplatePath", parameterToString(request.TemplatePath, ""))
	if localVarTempParam := request.TemplateStorage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("TemplateStorage", parameterToString(localVarTempParam, ""))
	}
	if request.IsImageDataEmbedded != nil {
		localVarQueryParams.Add("IsImageDataEmbedded", parameterToString(request.IsImageDataEmbedded, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.TemplatePassword; len(localVarTempParam) > 0 {
		localVarHeaderParams["TemplatePassword"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Data
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
    Storage string
    Folder string
}

/* SlidesApiService Performs slides pipeline.
 @param pipeline A Pipeline object.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesPipeline(request PostSlidesPipelineRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if request.Pipeline == nil {
		return successPayload, nil, reportError("Missing required parameter request.pipeline")
	}
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
	localVarPostBody = &request.Pipeline
	localVarFiles = request.Files
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlidesPresentationReplaceText(request PostSlidesPresentationReplaceTextRequest) (IDocumentReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentReplaceResult
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.OldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.oldValue")
	}
	if len(request.NewValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/replaceText"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.IgnoreCase != nil {
		if err := typeCheckParameter(*request.IgnoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OldValue", parameterToString(request.OldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(request.NewValue, ""))
	if request.IgnoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(request.IgnoreCase, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlidesReorder(request PostSlidesReorderRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/move"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("NewPosition", parameterToString(request.NewPosition, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlidesReorderMany(request PostSlidesReorderManyRequest) (ISlides, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlides
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/reorder"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.OldPositions, "[]int32", "oldPositions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.NewPositions, "[]int32", "newPositions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.OldPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("OldPositions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.NewPositions; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("NewPositions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "fontsFolder" (string) Custom fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostSlidesSaveAs(request PostSlidesSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesSaveAs
*/
type PostSlidesSaveAsRequest struct {
    Name string
    Format string
    Options IExportOptions
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PostSlidesSetDocumentProperties(request PostSlidesSetDocumentPropertiesRequest) (IDocumentProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Properties == nil {
		return successPayload, nil, reportError("Missing required parameter request.properties")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Properties
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PostSlidesSlideReplaceText(request PostSlidesSlideReplaceTextRequest) (ISlideReplaceResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideReplaceResult
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.OldValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.oldValue")
	}
	if len(request.NewValue) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.newValue")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/replaceText"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.IgnoreCase != nil {
		if err := typeCheckParameter(*request.IgnoreCase, "bool", "ignoreCase"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("OldValue", parameterToString(request.OldValue, ""))
	localVarQueryParams.Add("NewValue", parameterToString(request.NewValue, ""))
	if request.IgnoreCase != nil {
		localVarQueryParams.Add("IgnoreCase", parameterToString(request.IgnoreCase, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "to" (int32) The last slide number for splitting, if is not specified splitting ends at the last slide of the document.
     @param "from" (int32) The start slide number for splitting, if is not specified splitting starts from the first slide of the presentation.
     @param "destFolder" (string) Folder on storage where images are going to be uploaded. If not specified then images are uploaded to same folder as presentation.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "fontsFolder" (string) Custom fonts folder.
 @return SplitDocumentResult*/
func (a *SlidesApiService) PostSlidesSplit(request PostSlidesSplitRequest) (ISplitDocumentResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISplitDocumentResult
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/split"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Format, "string", "format"); err != nil {
		return successPayload, nil, err
	}
	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.To != nil {
		if err := typeCheckParameter(*request.To, "int32", "to"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.From != nil {
		if err := typeCheckParameter(*request.From, "int32", "from"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.DestFolder, "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Format; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Format", parameterToString(localVarTempParam, ""))
	}
	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if request.To != nil {
		localVarQueryParams.Add("To", parameterToString(*request.To, ""))
	}
	if request.From != nil {
		localVarQueryParams.Add("From", parameterToString(*request.From, ""))
	}
	if localVarTempParam := request.DestFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("DestFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSlidesSplit
*/
type PostSlidesSplitRequest struct {
    Name string
    Options IExportOptions
    Format string
    Width *int32
    Height *int32
    To *int32
    From *int32
    DestFolder string
    Password string
    Storage string
    Folder string
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
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return *os.File*/
func (a *SlidesApiService) PostSubshapeSaveAs(request PostSubshapeSaveAsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload *os.File
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if len(request.Format) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PostSubshapeSaveAs
*/
type PostSubshapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    Path string
    ShapeIndex int32
    Format string
    Options IIShapeExportOptions
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
    FontsFolder string
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
func (a *SlidesApiService) PutChartCategory(request PutChartCategoryRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Category == nil {
		return successPayload, nil, reportError("Missing required parameter request.category")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/categories/{categoryIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	categoryIndexPathStringValue := fmt.Sprintf("%v", request.CategoryIndex)
	if len(categoryIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"categoryIndex"+"}", categoryIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"categoryIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Category
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutChartDataPoint(request PutChartDataPointRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.DataPoint == nil {
		return successPayload, nil, reportError("Missing required parameter request.dataPoint")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}/dataPoints/{pointIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", request.SeriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}
	pointIndexPathStringValue := fmt.Sprintf("%v", request.PointIndex)
	if len(pointIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"pointIndex"+"}", pointIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"pointIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.DataPoint
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutChartSeries(request PutChartSeriesRequest) (IChart, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IChart
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Series == nil {
		return successPayload, nil, reportError("Missing required parameter request.series")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/series/{seriesIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	seriesIndexPathStringValue := fmt.Sprintf("%v", request.SeriesIndex)
	if len(seriesIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"seriesIndex"+"}", seriesIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"seriesIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Series
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* SlidesApiService Update a layoutSlide.
 @param name Document name.
 @param slideIndex Slide index.
 @param slideDto Slide update data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return LayoutSlide*/
func (a *SlidesApiService) PutLayoutSlide(request PutLayoutSlideRequest) (ILayoutSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ILayoutSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.SlideDto == nil {
		return successPayload, nil, reportError("Missing required parameter request.slideDto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/layoutSlides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.SlideDto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return NotesSlideHeaderFooter*/
func (a *SlidesApiService) PutNotesSlideHeaderFooter(request PutNotesSlideHeaderFooterRequest) (INotesSlideHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlideHeaderFooter
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutNotesSlideHeaderFooter
*/
type PutNotesSlideHeaderFooterRequest struct {
    Name string
    SlideIndex int32
    Dto INotesSlideHeaderFooter
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Render shape to specified picture format.
 @param name Presentation name.
 @param slideIndex Slide index.
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
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) PutNotesSlideShapeSaveAs(request PutNotesSlideShapeSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Name) == 0 {
		return nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutNotesSlideShapeSaveAs
*/
type PutNotesSlideShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    OutPath string
    Options IIShapeExportOptions
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
    FontsFolder string
}

/* SlidesApiService Merge the presentation with other presentations or some of their slides specified in the request parameter.
 @param name Document name.
 @param request OrderedMergeRequest with a list of presentations and slide indices to merge.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PutPresentationMerge(request PutPresentationMergeRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Request == nil {
		return successPayload, nil, reportError("Missing required parameter request.request")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/merge"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Request
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutPresentationMerge
*/
type PutPresentationMergeRequest struct {
    Name string
    Request IOrderedMergeRequest
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PutSection(request PutSectionRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.SectionName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.sectionName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections/{sectionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	sectionIndexPathStringValue := fmt.Sprintf("%v", request.SectionIndex)
	if len(sectionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sectionIndex"+"}", sectionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sectionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("SectionName", parameterToString(request.SectionName, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSections(request PutSectionsRequest) (ISections, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISections
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Sections == nil {
		return successPayload, nil, reportError("Missing required parameter request.sections")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/sections"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Sections
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSetParagraphPortionProperties(request PutSetParagraphPortionPropertiesRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSetParagraphProperties(request PutSetParagraphPropertiesRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSetSubshapeParagraphPortionProperties(request PutSetSubshapeParagraphPortionPropertiesRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSetSubshapeParagraphProperties(request PutSetSubshapeParagraphPropertiesRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) PutShapeSaveAs(request PutShapeSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Name) == 0 {
		return nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutShapeSaveAs
*/
type PutShapeSaveAsRequest struct {
    Name string
    SlideIndex int32
    ShapeIndex int32
    Format string
    OutPath string
    Options IIShapeExportOptions
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
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
func (a *SlidesApiService) PutSlideAnimation(request PutSlideAnimationRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Animation == nil {
		return successPayload, nil, reportError("Missing required parameter request.animation")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Animation
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideAnimationEffect(request PutSlideAnimationEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Effect == nil {
		return successPayload, nil, reportError("Missing required parameter request.effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/mainSequence/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", request.EffectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideAnimationInteractiveSequenceEffect(request PutSlideAnimationInteractiveSequenceEffectRequest) (ISlideAnimation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideAnimation
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Effect == nil {
		return successPayload, nil, reportError("Missing required parameter request.effect")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/animation/interactiveSequences/{sequenceIndex}/{effectIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	sequenceIndexPathStringValue := fmt.Sprintf("%v", request.SequenceIndex)
	if len(sequenceIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"sequenceIndex"+"}", sequenceIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"sequenceIndex"+"}", "", -1)
	}
	effectIndexPathStringValue := fmt.Sprintf("%v", request.EffectIndex)
	if len(effectIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"effectIndex"+"}", effectIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"effectIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Effect
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideHeaderFooter(request PutSlideHeaderFooterRequest) (IHeaderFooter, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IHeaderFooter
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideSaveAs(request PutSlideSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Name) == 0 {
		return nil, reportError("Missing required parameter request.name")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideShapeInfo(request PutSlideShapeInfoRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlideSubshapeInfo(request PutSlideSubshapeInfoRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "fontsFolder" (string) Custom fonts folder.
 @return */
func (a *SlidesApiService) PutSlidesConvert(request PutSlidesConvertRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Document) == 0 {
		return nil, reportError("Missing required parameter request.document")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/convert/{format}"
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Document
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesConvert
*/
type PutSlidesConvertRequest struct {
    Document []byte
    Format string
    OutPath string
    Password string
    FontsFolder string
}

/* SlidesApiService Update presentation document from html.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "html" (string) HTML data.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PutSlidesDocumentFromHtml(request PutSlidesDocumentFromHtmlRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/fromHtml"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Html, "string", "html"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = request.Html
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesDocumentFromHtml
*/
type PutSlidesDocumentFromHtmlRequest struct {
    Name string
    Html string
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Set footers for all slides in a presentation.
 @param name Document name.
 @param dto HeaderFooter instance.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
 @return Document*/
func (a *SlidesApiService) PutSlidesHeaderFooter(request PutSlidesHeaderFooterRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/headerFooter"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesHeaderFooter
*/
type PutSlidesHeaderFooterRequest struct {
    Name string
    Dto IHeaderFooter
    Password string
    Storage string
    Folder string
}

/* SlidesApiService Update presentation protection properties.
 @param name Document name.
 @param dto The view properties data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ProtectionProperties*/
func (a *SlidesApiService) PutSlidesProtectionProperties(request PutSlidesProtectionPropertiesRequest) (IProtectionProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IProtectionProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/protectionProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesProtectionProperties
*/
type PutSlidesProtectionPropertiesRequest struct {
    Name string
    Dto IProtectionProperties
    Password string
    Folder string
    Storage string
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
		localVarFiles [][]byte
	)

	if len(request.Name) == 0 {
		return nil, reportError("Missing required parameter request.name")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesSaveAs
*/
type PutSlidesSaveAsRequest struct {
    Name string
    OutPath string
    Format string
    Options IExportOptions
    Password string
    Storage string
    Folder string
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
func (a *SlidesApiService) PutSlidesSetDocumentProperty(request PutSlidesSetDocumentPropertyRequest) (IDocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocumentProperty
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.PropertyName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.propertyName")
	}
	if request.Property == nil {
		return successPayload, nil, reportError("Missing required parameter request.property")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/documentproperties/{propertyName}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	propertyNamePathStringValue := fmt.Sprintf("%v", request.PropertyName)
	if len(propertyNamePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", propertyNamePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"propertyName"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Property
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutSlidesSlide(request PutSlidesSlideRequest) (ISlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.SlideDto == nil {
		return successPayload, nil, reportError("Missing required parameter request.slideDto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.SlideDto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "folder" (string) Document folder.
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
 @return SlideBackground*/
func (a *SlidesApiService) PutSlidesSlideBackground(request PutSlidesSlideBackgroundRequest) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Background == nil {
		return successPayload, nil, reportError("Missing required parameter request.background")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/background"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Background
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesSlideBackground
*/
type PutSlidesSlideBackgroundRequest struct {
    Name string
    SlideIndex int32
    Background ISlideBackground
    Folder string
    Password string
    Storage string
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
func (a *SlidesApiService) PutSlidesSlideBackgroundColor(request PutSlidesSlideBackgroundColorRequest) (ISlideBackground, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideBackground
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if len(request.Color) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.color")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/backgroundColor"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("Color", parameterToString(request.Color, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesSlideBackgroundColor
*/
type PutSlidesSlideBackgroundColorRequest struct {
    Name string
    SlideIndex int32
    Color string
    Folder string
    Password string
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
func (a *SlidesApiService) PutSlidesSlideProperties(request PutSlidesSlidePropertiesRequest) (ISlideProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload ISlideProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesSlideProperties
*/
type PutSlidesSlidePropertiesRequest struct {
    Name string
    Dto ISlideProperties
    Password string
    Folder string
    Storage string
}

/* SlidesApiService Set slide size for a presentation.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "storage" (string) Document storage.
     @param "folder" (string) Document folder.
     @param "width" (int32) Slide width.
     @param "height" (int32) Slide height.
     @param "sizeType" (string) Standard slide size type.
     @param "scaleType" (string) Standard slide scale type.
 @return Document*/
func (a *SlidesApiService) PutSlidesSlideSize(request PutSlidesSlideSizeRequest) (IDocument, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IDocument
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slideSize"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if request.Width != nil {
		if err := typeCheckParameter(*request.Width, "int32", "width"); err != nil {
			return successPayload, nil, err
		}
	}
	if request.Height != nil {
		if err := typeCheckParameter(*request.Height, "int32", "height"); err != nil {
			return successPayload, nil, err
		}
	}
	if err := typeCheckParameter(request.SizeType, "string", "sizeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.ScaleType, "string", "scaleType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if request.Width != nil {
		localVarQueryParams.Add("Width", parameterToString(*request.Width, ""))
	}
	if request.Height != nil {
		localVarQueryParams.Add("Height", parameterToString(*request.Height, ""))
	}
	if localVarTempParam := request.SizeType; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("SizeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.ScaleType; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("ScaleType", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

/* Request for SlidesApiService.PutSlidesSlideSize
*/
type PutSlidesSlideSizeRequest struct {
    Name string
    Password string
    Storage string
    Folder string
    Width *int32
    Height *int32
    SizeType string
    ScaleType string
}

/* SlidesApiService Update presentation document properties.
 @param name Document name.
 @param dto The view properties data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "password" (string) Document password.
     @param "folder" (string) Document folder.
     @param "storage" (string) Document storage.
 @return ViewProperties*/
func (a *SlidesApiService) PutSlidesViewProperties(request PutSlidesViewPropertiesRequest) (IViewProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IViewProperties
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/viewProperties"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
     @param "password" (string) Document password.
     @param "folder" (string) Presentation folder.
     @param "storage" (string) Presentation storage.
     @param "scaleX" (float64) X scale ratio.
     @param "scaleY" (float64) Y scale ratio.
     @param "bounds" (string) Shape thumbnail bounds type.
     @param "fontsFolder" (string) Fonts folder.
 @return */
func (a *SlidesApiService) PutSubshapeSaveAs(request PutSubshapeSaveAsRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	)

	if len(request.Name) == 0 {
		return nil, reportError("Missing required parameter request.name")
	}
	if len(request.Path) == 0 {
		return nil, reportError("Missing required parameter request.path")
	}
	if len(request.Format) == 0 {
		return nil, reportError("Missing required parameter request.format")
	}
	if len(request.OutPath) == 0 {
		return nil, reportError("Missing required parameter request.outPath")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/shapes/{path}/{shapeIndex}/{format}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	formatPathStringValue := fmt.Sprintf("%v", request.Format)
	if len(formatPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", formatPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"format"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return nil, err
	}
	if request.ScaleX != nil {
		if err := typeCheckParameter(*request.ScaleX, "float64", "scaleX"); err != nil {
			return nil, err
		}
	}
	if request.ScaleY != nil {
		if err := typeCheckParameter(*request.ScaleY, "float64", "scaleY"); err != nil {
			return nil, err
		}
	}
	if err := typeCheckParameter(request.Bounds, "string", "bounds"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(request.FontsFolder, "string", "fontsFolder"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("OutPath", parameterToString(request.OutPath, ""))
	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
	}
	if request.ScaleX != nil {
		localVarQueryParams.Add("ScaleX", parameterToString(*request.ScaleX, ""))
	}
	if request.ScaleY != nil {
		localVarQueryParams.Add("ScaleY", parameterToString(*request.ScaleY, ""))
	}
	if localVarTempParam := request.Bounds; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Bounds", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.FontsFolder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("FontsFolder", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Options
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
    Password string
    Folder string
    Storage string
    ScaleX *float64
    ScaleY *float64
    Bounds string
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
func (a *SlidesApiService) PutUpdateNotesSlide(request PutUpdateNotesSlideRequest) (INotesSlide, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload INotesSlide
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutUpdateNotesSlideShape(request PutUpdateNotesSlideShapeRequest) (IShapeBase, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IShapeBase
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutUpdateNotesSlideShapeParagraph(request PutUpdateNotesSlideShapeParagraphRequest) (IParagraph, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IParagraph
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) PutUpdateNotesSlideShapePortion(request PutUpdateNotesSlideShapePortionRequest) (IPortion, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IPortion
	)

	if len(request.Name) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.name")
	}
	if request.Dto == nil {
		return successPayload, nil, reportError("Missing required parameter request.dto")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/{name}/slides/{slideIndex}/notesSlide/shapes/{shapeIndex}/paragraphs/{paragraphIndex}/portions/{portionIndex}"
	namePathStringValue := fmt.Sprintf("%v", request.Name)
	if len(namePathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", namePathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"name"+"}", "", -1)
	}
	slideIndexPathStringValue := fmt.Sprintf("%v", request.SlideIndex)
	if len(slideIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"slideIndex"+"}", slideIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"slideIndex"+"}", "", -1)
	}
	shapeIndexPathStringValue := fmt.Sprintf("%v", request.ShapeIndex)
	if len(shapeIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"shapeIndex"+"}", shapeIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"shapeIndex"+"}", "", -1)
	}
	paragraphIndexPathStringValue := fmt.Sprintf("%v", request.ParagraphIndex)
	if len(paragraphIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"paragraphIndex"+"}", paragraphIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"paragraphIndex"+"}", "", -1)
	}
	portionIndexPathStringValue := fmt.Sprintf("%v", request.PortionIndex)
	if len(portionIndexPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"portionIndex"+"}", portionIndexPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"portionIndex"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.Password, "string", "password"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Folder, "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(request.Storage, "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.Folder; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam := request.Storage; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("Storage", parameterToString(localVarTempParam, ""))
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
	if localVarTempParam := request.Password; len(localVarTempParam) > 0 {
		localVarHeaderParams["Password"] = parameterToString(localVarTempParam, "")
	}
	// body params
	localVarPostBody = &request.Dto
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) StorageExists(request StorageExistsRequest) (IStorageExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IStorageExist
	)

	if len(request.StorageName) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.storageName")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/{storageName}/exist"
	storageNamePathStringValue := fmt.Sprintf("%v", request.StorageName)
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
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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
func (a *SlidesApiService) UploadFile(request UploadFileRequest) (IFilesUploadResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFiles [][]byte
	 	successPayload IFilesUploadResult
	)

	if len(request.Path) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.path")
	}
	if len(request.File) == 0 {
		return successPayload, nil, reportError("Missing required parameter request.file")
	}
	// create path and map variables
	localVarPath := a.client.cfg.GetApiUrl() + "/slides/storage/file/{path}"
	pathPathStringValue := fmt.Sprintf("%v", request.Path)
	if len(pathPathStringValue) > 0 {
		localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", pathPathStringValue, -1)
	} else {
		localVarPath = strings.Replace(localVarPath, "/{"+"path"+"}", "", -1)
	}

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(request.StorageName, "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam := request.StorageName; len(localVarTempParam) > 0 {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
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
	localVarPostBody = &request.File
	localVarHttpResponse, responseBytes, err := a.client.makeRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFiles)
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

