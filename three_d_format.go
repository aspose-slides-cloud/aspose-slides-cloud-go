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

// ThreeDFormat
type IThreeDFormat interface {

	// Type of a bottom 3D bevel.             
	GetBevelBottom() IShapeBevel
	SetBevelBottom(newValue IShapeBevel)

	// Type of a top 3D bevel.             
	GetBevelTop() IShapeBevel
	SetBevelTop(newValue IShapeBevel)

	// Camera
	GetCamera() ICamera
	SetCamera(newValue ICamera)

	// Contour color
	GetContourColor() string
	SetContourColor(newValue string)

	// Contour width
	GetContourWidth() float64
	SetContourWidth(newValue float64)

	// Depth
	GetDepth() float64
	SetDepth(newValue float64)

	// Extrusion color
	GetExtrusionColor() string
	SetExtrusionColor(newValue string)

	// Extrusion height
	GetExtrusionHeight() float64
	SetExtrusionHeight(newValue float64)

	// Light rig
	GetLightRig() ILightRig
	SetLightRig(newValue ILightRig)

	// Material
	GetMaterial() string
	SetMaterial(newValue string)
}

type ThreeDFormat struct {

	// Type of a bottom 3D bevel.             
	BevelBottom IShapeBevel `json:"BevelBottom,omitempty"`

	// Type of a top 3D bevel.             
	BevelTop IShapeBevel `json:"BevelTop,omitempty"`

	// Camera
	Camera ICamera `json:"Camera,omitempty"`

	// Contour color
	ContourColor string `json:"ContourColor,omitempty"`

	// Contour width
	ContourWidth float64 `json:"ContourWidth,omitempty"`

	// Depth
	Depth float64 `json:"Depth,omitempty"`

	// Extrusion color
	ExtrusionColor string `json:"ExtrusionColor,omitempty"`

	// Extrusion height
	ExtrusionHeight float64 `json:"ExtrusionHeight,omitempty"`

	// Light rig
	LightRig ILightRig `json:"LightRig,omitempty"`

	// Material
	Material string `json:"Material,omitempty"`
}

func NewThreeDFormat() *ThreeDFormat {
	instance := new(ThreeDFormat)
	instance.Material = ""
	return instance
}

func (this *ThreeDFormat) GetBevelBottom() IShapeBevel {
	return this.BevelBottom
}

func (this *ThreeDFormat) SetBevelBottom(newValue IShapeBevel) {
	this.BevelBottom = newValue
}
func (this *ThreeDFormat) GetBevelTop() IShapeBevel {
	return this.BevelTop
}

func (this *ThreeDFormat) SetBevelTop(newValue IShapeBevel) {
	this.BevelTop = newValue
}
func (this *ThreeDFormat) GetCamera() ICamera {
	return this.Camera
}

func (this *ThreeDFormat) SetCamera(newValue ICamera) {
	this.Camera = newValue
}
func (this *ThreeDFormat) GetContourColor() string {
	return this.ContourColor
}

func (this *ThreeDFormat) SetContourColor(newValue string) {
	this.ContourColor = newValue
}
func (this *ThreeDFormat) GetContourWidth() float64 {
	return this.ContourWidth
}

func (this *ThreeDFormat) SetContourWidth(newValue float64) {
	this.ContourWidth = newValue
}
func (this *ThreeDFormat) GetDepth() float64 {
	return this.Depth
}

func (this *ThreeDFormat) SetDepth(newValue float64) {
	this.Depth = newValue
}
func (this *ThreeDFormat) GetExtrusionColor() string {
	return this.ExtrusionColor
}

func (this *ThreeDFormat) SetExtrusionColor(newValue string) {
	this.ExtrusionColor = newValue
}
func (this *ThreeDFormat) GetExtrusionHeight() float64 {
	return this.ExtrusionHeight
}

func (this *ThreeDFormat) SetExtrusionHeight(newValue float64) {
	this.ExtrusionHeight = newValue
}
func (this *ThreeDFormat) GetLightRig() ILightRig {
	return this.LightRig
}

func (this *ThreeDFormat) SetLightRig(newValue ILightRig) {
	this.LightRig = newValue
}
func (this *ThreeDFormat) GetMaterial() string {
	return this.Material
}

func (this *ThreeDFormat) SetMaterial(newValue string) {
	this.Material = newValue
}

func (this *ThreeDFormat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valBevelBottom, ok := objMap["bevelBottom"]; ok {
		if valBevelBottom != nil {
			var valueForBevelBottom ShapeBevel
			err = json.Unmarshal(*valBevelBottom, &valueForBevelBottom)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ShapeBevel", *valBevelBottom)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBevelBottom, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IShapeBevel)
			if ok {
				this.BevelBottom = vInterfaceObject
			}
		}
	}
	if valBevelBottomCap, ok := objMap["BevelBottom"]; ok {
		if valBevelBottomCap != nil {
			var valueForBevelBottom ShapeBevel
			err = json.Unmarshal(*valBevelBottomCap, &valueForBevelBottom)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ShapeBevel", *valBevelBottomCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBevelBottomCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IShapeBevel)
			if ok {
				this.BevelBottom = vInterfaceObject
			}
		}
	}
	
	if valBevelTop, ok := objMap["bevelTop"]; ok {
		if valBevelTop != nil {
			var valueForBevelTop ShapeBevel
			err = json.Unmarshal(*valBevelTop, &valueForBevelTop)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ShapeBevel", *valBevelTop)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBevelTop, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IShapeBevel)
			if ok {
				this.BevelTop = vInterfaceObject
			}
		}
	}
	if valBevelTopCap, ok := objMap["BevelTop"]; ok {
		if valBevelTopCap != nil {
			var valueForBevelTop ShapeBevel
			err = json.Unmarshal(*valBevelTopCap, &valueForBevelTop)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("ShapeBevel", *valBevelTopCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valBevelTopCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(IShapeBevel)
			if ok {
				this.BevelTop = vInterfaceObject
			}
		}
	}
	
	if valCamera, ok := objMap["camera"]; ok {
		if valCamera != nil {
			var valueForCamera Camera
			err = json.Unmarshal(*valCamera, &valueForCamera)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Camera", *valCamera)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCamera, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ICamera)
			if ok {
				this.Camera = vInterfaceObject
			}
		}
	}
	if valCameraCap, ok := objMap["Camera"]; ok {
		if valCameraCap != nil {
			var valueForCamera Camera
			err = json.Unmarshal(*valCameraCap, &valueForCamera)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("Camera", *valCameraCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valCameraCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ICamera)
			if ok {
				this.Camera = vInterfaceObject
			}
		}
	}
	
	if valContourColor, ok := objMap["contourColor"]; ok {
		if valContourColor != nil {
			var valueForContourColor string
			err = json.Unmarshal(*valContourColor, &valueForContourColor)
			if err != nil {
				return err
			}
			this.ContourColor = valueForContourColor
		}
	}
	if valContourColorCap, ok := objMap["ContourColor"]; ok {
		if valContourColorCap != nil {
			var valueForContourColor string
			err = json.Unmarshal(*valContourColorCap, &valueForContourColor)
			if err != nil {
				return err
			}
			this.ContourColor = valueForContourColor
		}
	}
	
	if valContourWidth, ok := objMap["contourWidth"]; ok {
		if valContourWidth != nil {
			var valueForContourWidth float64
			err = json.Unmarshal(*valContourWidth, &valueForContourWidth)
			if err != nil {
				return err
			}
			this.ContourWidth = valueForContourWidth
		}
	}
	if valContourWidthCap, ok := objMap["ContourWidth"]; ok {
		if valContourWidthCap != nil {
			var valueForContourWidth float64
			err = json.Unmarshal(*valContourWidthCap, &valueForContourWidth)
			if err != nil {
				return err
			}
			this.ContourWidth = valueForContourWidth
		}
	}
	
	if valDepth, ok := objMap["depth"]; ok {
		if valDepth != nil {
			var valueForDepth float64
			err = json.Unmarshal(*valDepth, &valueForDepth)
			if err != nil {
				return err
			}
			this.Depth = valueForDepth
		}
	}
	if valDepthCap, ok := objMap["Depth"]; ok {
		if valDepthCap != nil {
			var valueForDepth float64
			err = json.Unmarshal(*valDepthCap, &valueForDepth)
			if err != nil {
				return err
			}
			this.Depth = valueForDepth
		}
	}
	
	if valExtrusionColor, ok := objMap["extrusionColor"]; ok {
		if valExtrusionColor != nil {
			var valueForExtrusionColor string
			err = json.Unmarshal(*valExtrusionColor, &valueForExtrusionColor)
			if err != nil {
				return err
			}
			this.ExtrusionColor = valueForExtrusionColor
		}
	}
	if valExtrusionColorCap, ok := objMap["ExtrusionColor"]; ok {
		if valExtrusionColorCap != nil {
			var valueForExtrusionColor string
			err = json.Unmarshal(*valExtrusionColorCap, &valueForExtrusionColor)
			if err != nil {
				return err
			}
			this.ExtrusionColor = valueForExtrusionColor
		}
	}
	
	if valExtrusionHeight, ok := objMap["extrusionHeight"]; ok {
		if valExtrusionHeight != nil {
			var valueForExtrusionHeight float64
			err = json.Unmarshal(*valExtrusionHeight, &valueForExtrusionHeight)
			if err != nil {
				return err
			}
			this.ExtrusionHeight = valueForExtrusionHeight
		}
	}
	if valExtrusionHeightCap, ok := objMap["ExtrusionHeight"]; ok {
		if valExtrusionHeightCap != nil {
			var valueForExtrusionHeight float64
			err = json.Unmarshal(*valExtrusionHeightCap, &valueForExtrusionHeight)
			if err != nil {
				return err
			}
			this.ExtrusionHeight = valueForExtrusionHeight
		}
	}
	
	if valLightRig, ok := objMap["lightRig"]; ok {
		if valLightRig != nil {
			var valueForLightRig LightRig
			err = json.Unmarshal(*valLightRig, &valueForLightRig)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LightRig", *valLightRig)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLightRig, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILightRig)
			if ok {
				this.LightRig = vInterfaceObject
			}
		}
	}
	if valLightRigCap, ok := objMap["LightRig"]; ok {
		if valLightRigCap != nil {
			var valueForLightRig LightRig
			err = json.Unmarshal(*valLightRigCap, &valueForLightRig)
			if err != nil {
				return err
			}
			vObject, err := createObjectForType("LightRig", *valLightRigCap)
			if err != nil {
				return err
			}
			err = json.Unmarshal(*valLightRigCap, &vObject)
			if err != nil {
				return err
			}
			vInterfaceObject, ok := vObject.(ILightRig)
			if ok {
				this.LightRig = vInterfaceObject
			}
		}
	}
	this.Material = ""
	if valMaterial, ok := objMap["material"]; ok {
		if valMaterial != nil {
			var valueForMaterial string
			err = json.Unmarshal(*valMaterial, &valueForMaterial)
			if err != nil {
				var valueForMaterialInt int32
				err = json.Unmarshal(*valMaterial, &valueForMaterialInt)
				if err != nil {
					return err
				}
				this.Material = string(valueForMaterialInt)
			} else {
				this.Material = valueForMaterial
			}
		}
	}
	if valMaterialCap, ok := objMap["Material"]; ok {
		if valMaterialCap != nil {
			var valueForMaterial string
			err = json.Unmarshal(*valMaterialCap, &valueForMaterial)
			if err != nil {
				var valueForMaterialInt int32
				err = json.Unmarshal(*valMaterialCap, &valueForMaterialInt)
				if err != nil {
					return err
				}
				this.Material = string(valueForMaterialInt)
			} else {
				this.Material = valueForMaterial
			}
		}
	}

	return nil
}
