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
	"encoding/json"
)

// Smart art node.
type ISmartArtNode interface {

	// Node list.
	getNodes() []SmartArtNode
	setNodes(newValue []SmartArtNode)

	// Gets or sets the link to shapes.
	getShapes() IResourceUriElement
	setShapes(newValue IResourceUriElement)

	// True for and assistant node.
	getIsAssistant() bool
	setIsAssistant(newValue bool)

	// Node text.
	getText() string
	setText(newValue string)

	// Organization chart layout type associated with current node.
	getOrgChartLayout() string
	setOrgChartLayout(newValue string)
}

type SmartArtNode struct {

	// Node list.
	Nodes []SmartArtNode `json:"Nodes,omitempty"`

	// Gets or sets the link to shapes.
	Shapes IResourceUriElement `json:"Shapes,omitempty"`

	// True for and assistant node.
	IsAssistant bool `json:"IsAssistant"`

	// Node text.
	Text string `json:"Text,omitempty"`

	// Organization chart layout type associated with current node.
	OrgChartLayout string `json:"OrgChartLayout"`
}

func (this SmartArtNode) getNodes() []SmartArtNode {
	return this.Nodes
}

func (this SmartArtNode) setNodes(newValue []SmartArtNode) {
	this.Nodes = newValue
}
func (this SmartArtNode) getShapes() IResourceUriElement {
	return this.Shapes
}

func (this SmartArtNode) setShapes(newValue IResourceUriElement) {
	this.Shapes = newValue
}
func (this SmartArtNode) getIsAssistant() bool {
	return this.IsAssistant
}

func (this SmartArtNode) setIsAssistant(newValue bool) {
	this.IsAssistant = newValue
}
func (this SmartArtNode) getText() string {
	return this.Text
}

func (this SmartArtNode) setText(newValue string) {
	this.Text = newValue
}
func (this SmartArtNode) getOrgChartLayout() string {
	return this.OrgChartLayout
}

func (this SmartArtNode) setOrgChartLayout(newValue string) {
	this.OrgChartLayout = newValue
}

func (this *SmartArtNode) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valNodes, ok := objMap["Nodes"]; ok {
		if valNodes != nil {
			var valueForNodes []SmartArtNode
			err = json.Unmarshal(*valNodes, &valueForNodes)
			if err != nil {
				return err
			}
			this.Nodes = valueForNodes
		}
	}
	
	if valShapes, ok := objMap["Shapes"]; ok {
		if valShapes != nil {
			var valueForShapes ResourceUriElement
			err = json.Unmarshal(*valShapes, &valueForShapes)
			if err != nil {
				return err
			}
			this.Shapes = valueForShapes
		}
	}
	
	if valIsAssistant, ok := objMap["IsAssistant"]; ok {
		if valIsAssistant != nil {
			var valueForIsAssistant bool
			err = json.Unmarshal(*valIsAssistant, &valueForIsAssistant)
			if err != nil {
				return err
			}
			this.IsAssistant = valueForIsAssistant
		}
	}
	
	if valText, ok := objMap["Text"]; ok {
		if valText != nil {
			var valueForText string
			err = json.Unmarshal(*valText, &valueForText)
			if err != nil {
				return err
			}
			this.Text = valueForText
		}
	}
	this.OrgChartLayout = "Initial"
	if valOrgChartLayout, ok := objMap["OrgChartLayout"]; ok {
		if valOrgChartLayout != nil {
			var valueForOrgChartLayout string
			err = json.Unmarshal(*valOrgChartLayout, &valueForOrgChartLayout)
			if err != nil {
				return err
			}
			this.OrgChartLayout = valueForOrgChartLayout
		}
	}

    return nil
}
