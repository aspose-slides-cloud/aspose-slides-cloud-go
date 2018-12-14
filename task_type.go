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
// TaskType : Pipeline task type.
type TaskType string

// List of TaskType TaskType
const (
	TaskType_SAVE TaskType = "Save"
	TaskType_SAVE_SLIDE TaskType = "SaveSlide"
	TaskType_SAVE_SHAPE TaskType = "SaveShape"
	TaskType_ADD_SLIDE TaskType = "AddSlide"
	TaskType_ADD_MASTER_SLIDE TaskType = "AddMasterSlide"
	TaskType_ADD_LAYOUT_SLIDE TaskType = "AddLayoutSlide"
	TaskType_REMOVE_SLIDE TaskType = "RemoveSlide"
	TaskType_REODER_SLIDE TaskType = "ReoderSlide"
	TaskType_MERGE TaskType = "Merge"
	TaskType_UPDATE_BACKGROUND TaskType = "UpdateBackground"
	TaskType_RESET_SLIDE TaskType = "ResetSlide"
	TaskType_ADD_SHAPE TaskType = "AddShape"
	TaskType_REMOVE_SHAPE TaskType = "RemoveShape"
	TaskType_UPDATE_SHAPE TaskType = "UpdateShape"
	TaskType_REPLACE_TEXT TaskType = "ReplaceText"
)