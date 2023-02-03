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
	"encoding/json"
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v23"
)

/*
   Test for call with valid auth data
*/
func TestGoodAuth(t *testing.T) {
	cfg := slidescloud.NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient := slidescloud.NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestBadAuth(t *testing.T) {
	cfg := slidescloud.NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	cfg.AppSid = "invalid"
	testApiClient := slidescloud.NewAPIClient(cfg)
	_, r, e := testApiClient.SlidesApi.GetApiInfo()
	if e == nil {
		t.Errorf("Must have failed.")
		return
	}
	if r == nil {
		t.Errorf("Null response not expected.")
		return
	}
	statusCode := r.StatusCode
	if statusCode != 401 {
		t.Errorf("Unexpected error code: %v.", statusCode)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestGoodAuthToken(t *testing.T) {
	cfg := slidescloud.NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient := slidescloud.NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.AppSid = "invalid"
	testApiClient = slidescloud.NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for call with valid auth data
*/
func TestBadAuthToken(t *testing.T) {
	cfg := slidescloud.NewConfiguration()
	configFile, err := os.Open("testConfig.json")
	if err == nil {
		json.NewDecoder(configFile).Decode(cfg)
	}
	testApiClient := slidescloud.NewAPIClient(cfg)
	_, _, e := testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	cfg.OAuthToken = "invalid"
	testApiClient = slidescloud.NewAPIClient(cfg)
	_, _, e = testApiClient.SlidesApi.GetApiInfo()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
