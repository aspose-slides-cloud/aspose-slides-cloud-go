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

import "encoding/json"
import "reflect"
import "unicode"

var typeRegistry = make(map[string]reflect.Type)
var derivedTypes = make(map[string]string)
var typeDeterminers = make(map[string]map[string]string)

func init() {
	typeRegistry["ApiInfo"] = reflect.TypeOf(ApiInfo{})
	
	typeDeterminers["ApiInfo"] = make(map[string]string)
	typeRegistry["ArrowHeadProperties"] = reflect.TypeOf(ArrowHeadProperties{})
	
	typeDeterminers["ArrowHeadProperties"] = make(map[string]string)
	typeRegistry["Axes"] = reflect.TypeOf(Axes{})
	
	typeDeterminers["Axes"] = make(map[string]string)
	typeRegistry["Axis"] = reflect.TypeOf(Axis{})
	
	typeDeterminers["Axis"] = make(map[string]string)
	typeRegistry["BlurEffect"] = reflect.TypeOf(BlurEffect{})
	
	typeDeterminers["BlurEffect"] = make(map[string]string)
	typeRegistry["ChartCategory"] = reflect.TypeOf(ChartCategory{})
	
	typeDeterminers["ChartCategory"] = make(map[string]string)
	typeRegistry["ChartTitle"] = reflect.TypeOf(ChartTitle{})
	
	typeDeterminers["ChartTitle"] = make(map[string]string)
	typeRegistry["ChartWall"] = reflect.TypeOf(ChartWall{})
	
	typeDeterminers["ChartWall"] = make(map[string]string)
	typeRegistry["CommonSlideViewProperties"] = reflect.TypeOf(CommonSlideViewProperties{})
	
	typeDeterminers["CommonSlideViewProperties"] = make(map[string]string)
	typeRegistry["CustomDashPattern"] = reflect.TypeOf(CustomDashPattern{})
	
	typeDeterminers["CustomDashPattern"] = make(map[string]string)
	typeRegistry["DiscUsage"] = reflect.TypeOf(DiscUsage{})
	
	typeDeterminers["DiscUsage"] = make(map[string]string)
	typeRegistry["Effect"] = reflect.TypeOf(Effect{})
	
	typeDeterminers["Effect"] = make(map[string]string)
	typeRegistry["EffectFormat"] = reflect.TypeOf(EffectFormat{})
	
	typeDeterminers["EffectFormat"] = make(map[string]string)
	typeRegistry["EntityExists"] = reflect.TypeOf(EntityExists{})
	
	typeDeterminers["EntityExists"] = make(map[string]string)
	typeRegistry["ErrorDetails"] = reflect.TypeOf(ErrorDetails{})
	
	typeDeterminers["ErrorDetails"] = make(map[string]string)
	typeRegistry["ExportOptions"] = reflect.TypeOf(ExportOptions{})
	
	typeDeterminers["ExportOptions"] = make(map[string]string)
	typeRegistry["FileVersions"] = reflect.TypeOf(FileVersions{})
	
	typeDeterminers["FileVersions"] = make(map[string]string)
	typeRegistry["FilesList"] = reflect.TypeOf(FilesList{})
	
	typeDeterminers["FilesList"] = make(map[string]string)
	typeRegistry["FilesUploadResult"] = reflect.TypeOf(FilesUploadResult{})
	
	typeDeterminers["FilesUploadResult"] = make(map[string]string)
	typeRegistry["FillFormat"] = reflect.TypeOf(FillFormat{})
	
	typeDeterminers["FillFormat"] = make(map[string]string)
	typeRegistry["FillOverlayEffect"] = reflect.TypeOf(FillOverlayEffect{})
	
	typeDeterminers["FillOverlayEffect"] = make(map[string]string)
	typeRegistry["FontSet"] = reflect.TypeOf(FontSet{})
	
	typeDeterminers["FontSet"] = make(map[string]string)
	typeRegistry["GlowEffect"] = reflect.TypeOf(GlowEffect{})
	
	typeDeterminers["GlowEffect"] = make(map[string]string)
	typeRegistry["GradientFillStop"] = reflect.TypeOf(GradientFillStop{})
	
	typeDeterminers["GradientFillStop"] = make(map[string]string)
	typeRegistry["IShapeExportOptions"] = reflect.TypeOf(IShapeExportOptions{})
	
	typeDeterminers["IShapeExportOptions"] = make(map[string]string)
	typeRegistry["InnerShadowEffect"] = reflect.TypeOf(InnerShadowEffect{})
	
	typeDeterminers["InnerShadowEffect"] = make(map[string]string)
	typeRegistry["Input"] = reflect.TypeOf(Input{})
	
	typeDeterminers["Input"] = make(map[string]string)
	typeRegistry["InputFile"] = reflect.TypeOf(InputFile{})
	
	typeDeterminers["InputFile"] = make(map[string]string)
	typeRegistry["InteractiveSequence"] = reflect.TypeOf(InteractiveSequence{})
	
	typeDeterminers["InteractiveSequence"] = make(map[string]string)
	typeRegistry["Legend"] = reflect.TypeOf(Legend{})
	
	typeDeterminers["Legend"] = make(map[string]string)
	typeRegistry["LineFormat"] = reflect.TypeOf(LineFormat{})
	
	typeDeterminers["LineFormat"] = make(map[string]string)
	typeRegistry["MergingSource"] = reflect.TypeOf(MergingSource{})
	
	typeDeterminers["MergingSource"] = make(map[string]string)
	typeRegistry["ModelError"] = reflect.TypeOf(ModelError{})
	
	typeDeterminers["ModelError"] = make(map[string]string)
	typeRegistry["NormalViewRestoredProperties"] = reflect.TypeOf(NormalViewRestoredProperties{})
	
	typeDeterminers["NormalViewRestoredProperties"] = make(map[string]string)
	typeRegistry["ObjectExist"] = reflect.TypeOf(ObjectExist{})
	
	typeDeterminers["ObjectExist"] = make(map[string]string)
	typeRegistry["OneValueChartDataPoint"] = reflect.TypeOf(OneValueChartDataPoint{})
	
	typeDeterminers["OneValueChartDataPoint"] = make(map[string]string)
	typeRegistry["OrderedMergeRequest"] = reflect.TypeOf(OrderedMergeRequest{})
	
	typeDeterminers["OrderedMergeRequest"] = make(map[string]string)
	typeRegistry["OuterShadowEffect"] = reflect.TypeOf(OuterShadowEffect{})
	
	typeDeterminers["OuterShadowEffect"] = make(map[string]string)
	typeRegistry["OutputFile"] = reflect.TypeOf(OutputFile{})
	
	typeDeterminers["OutputFile"] = make(map[string]string)
	typeRegistry["Pipeline"] = reflect.TypeOf(Pipeline{})
	
	typeDeterminers["Pipeline"] = make(map[string]string)
	typeRegistry["PlotArea"] = reflect.TypeOf(PlotArea{})
	
	typeDeterminers["PlotArea"] = make(map[string]string)
	typeRegistry["PresentationToMerge"] = reflect.TypeOf(PresentationToMerge{})
	
	typeDeterminers["PresentationToMerge"] = make(map[string]string)
	typeRegistry["PresentationsMergeRequest"] = reflect.TypeOf(PresentationsMergeRequest{})
	
	typeDeterminers["PresentationsMergeRequest"] = make(map[string]string)
	typeRegistry["PresetShadowEffect"] = reflect.TypeOf(PresetShadowEffect{})
	
	typeDeterminers["PresetShadowEffect"] = make(map[string]string)
	typeRegistry["ReflectionEffect"] = reflect.TypeOf(ReflectionEffect{})
	
	typeDeterminers["ReflectionEffect"] = make(map[string]string)
	typeRegistry["ResourceBase"] = reflect.TypeOf(ResourceBase{})
	
	typeDeterminers["ResourceBase"] = make(map[string]string)
	typeRegistry["ResourceUri"] = reflect.TypeOf(ResourceUri{})
	
	typeDeterminers["ResourceUri"] = make(map[string]string)
	typeRegistry["ResourceUriElement"] = reflect.TypeOf(ResourceUriElement{})
	
	typeDeterminers["ResourceUriElement"] = make(map[string]string)
	typeRegistry["ScatterChartDataPoint"] = reflect.TypeOf(ScatterChartDataPoint{})
	
	typeDeterminers["ScatterChartDataPoint"] = make(map[string]string)
	typeRegistry["Series"] = reflect.TypeOf(Series{})
	
	typeDeterminers["Series"] = make(map[string]string)
	typeRegistry["SeriesMarker"] = reflect.TypeOf(SeriesMarker{})
	
	typeDeterminers["SeriesMarker"] = make(map[string]string)
	typeRegistry["ShapeImageExportOptions"] = reflect.TypeOf(ShapeImageExportOptions{})
	
	typeDeterminers["ShapeImageExportOptions"] = make(map[string]string)
	typeRegistry["SlideComment"] = reflect.TypeOf(SlideComment{})
	
	typeDeterminers["SlideComment"] = make(map[string]string)
	typeRegistry["SmartArtNode"] = reflect.TypeOf(SmartArtNode{})
	
	typeDeterminers["SmartArtNode"] = make(map[string]string)
	typeRegistry["SoftEdgeEffect"] = reflect.TypeOf(SoftEdgeEffect{})
	
	typeDeterminers["SoftEdgeEffect"] = make(map[string]string)
	typeRegistry["StorageExist"] = reflect.TypeOf(StorageExist{})
	
	typeDeterminers["StorageExist"] = make(map[string]string)
	typeRegistry["StorageFile"] = reflect.TypeOf(StorageFile{})
	
	typeDeterminers["StorageFile"] = make(map[string]string)
	typeRegistry["TableCell"] = reflect.TypeOf(TableCell{})
	
	typeDeterminers["TableCell"] = make(map[string]string)
	typeRegistry["TableColumn"] = reflect.TypeOf(TableColumn{})
	
	typeDeterminers["TableColumn"] = make(map[string]string)
	typeRegistry["TableRow"] = reflect.TypeOf(TableRow{})
	
	typeDeterminers["TableRow"] = make(map[string]string)
	typeRegistry["Task"] = reflect.TypeOf(Task{})
	
	typeDeterminers["Task"] = make(map[string]string)
	typeRegistry["TextItem"] = reflect.TypeOf(TextItem{})
	
	typeDeterminers["TextItem"] = make(map[string]string)
	typeRegistry["AddLayoutSlide"] = reflect.TypeOf(AddLayoutSlide{})
	derivedTypes["AddLayoutSlide"] = "Task"
	typeDeterminers["AddLayoutSlide"] = make(map[string]string)
	typeDeterminers["AddLayoutSlide"]["Type"] = "AddLayoutSlide"
	typeRegistry["AddMasterSlide"] = reflect.TypeOf(AddMasterSlide{})
	derivedTypes["AddMasterSlide"] = "Task"
	typeDeterminers["AddMasterSlide"] = make(map[string]string)
	typeDeterminers["AddMasterSlide"]["Type"] = "AddMasterSlide"
	typeRegistry["AddShape"] = reflect.TypeOf(AddShape{})
	derivedTypes["AddShape"] = "Task"
	typeDeterminers["AddShape"] = make(map[string]string)
	typeDeterminers["AddShape"]["Type"] = "AddShape"
	typeRegistry["AddSlide"] = reflect.TypeOf(AddSlide{})
	derivedTypes["AddSlide"] = "Task"
	typeDeterminers["AddSlide"] = make(map[string]string)
	typeDeterminers["AddSlide"]["Type"] = "AddSlide"
	typeRegistry["Base64InputFile"] = reflect.TypeOf(Base64InputFile{})
	derivedTypes["Base64InputFile"] = "InputFile"
	typeDeterminers["Base64InputFile"] = make(map[string]string)
	typeDeterminers["Base64InputFile"]["Type"] = "Base64"
	typeRegistry["BubbleChartDataPoint"] = reflect.TypeOf(BubbleChartDataPoint{})
	derivedTypes["BubbleChartDataPoint"] = "ScatterChartDataPoint"
	typeDeterminers["BubbleChartDataPoint"] = make(map[string]string)
	typeRegistry["BubbleSeries"] = reflect.TypeOf(BubbleSeries{})
	derivedTypes["BubbleSeries"] = "Series"
	typeDeterminers["BubbleSeries"] = make(map[string]string)
	typeDeterminers["BubbleSeries"]["DataPointType"] = "Bubble"
	typeRegistry["ColorScheme"] = reflect.TypeOf(ColorScheme{})
	derivedTypes["ColorScheme"] = "ResourceBase"
	typeDeterminers["ColorScheme"] = make(map[string]string)
	typeRegistry["Document"] = reflect.TypeOf(Document{})
	derivedTypes["Document"] = "ResourceBase"
	typeDeterminers["Document"] = make(map[string]string)
	typeRegistry["DocumentProperties"] = reflect.TypeOf(DocumentProperties{})
	derivedTypes["DocumentProperties"] = "ResourceBase"
	typeDeterminers["DocumentProperties"] = make(map[string]string)
	typeRegistry["DocumentProperty"] = reflect.TypeOf(DocumentProperty{})
	derivedTypes["DocumentProperty"] = "ResourceBase"
	typeDeterminers["DocumentProperty"] = make(map[string]string)
	typeRegistry["FileVersion"] = reflect.TypeOf(FileVersion{})
	derivedTypes["FileVersion"] = "StorageFile"
	typeDeterminers["FileVersion"] = make(map[string]string)
	typeRegistry["FontScheme"] = reflect.TypeOf(FontScheme{})
	derivedTypes["FontScheme"] = "ResourceBase"
	typeDeterminers["FontScheme"] = make(map[string]string)
	typeRegistry["FormatScheme"] = reflect.TypeOf(FormatScheme{})
	derivedTypes["FormatScheme"] = "ResourceBase"
	typeDeterminers["FormatScheme"] = make(map[string]string)
	typeRegistry["GradientFill"] = reflect.TypeOf(GradientFill{})
	derivedTypes["GradientFill"] = "FillFormat"
	typeDeterminers["GradientFill"] = make(map[string]string)
	typeDeterminers["GradientFill"]["Type"] = "Gradient"
	typeRegistry["HtmlExportOptions"] = reflect.TypeOf(HtmlExportOptions{})
	derivedTypes["HtmlExportOptions"] = "ExportOptions"
	typeDeterminers["HtmlExportOptions"] = make(map[string]string)
	typeRegistry["Image"] = reflect.TypeOf(Image{})
	derivedTypes["Image"] = "ResourceBase"
	typeDeterminers["Image"] = make(map[string]string)
	typeRegistry["ImageExportOptions"] = reflect.TypeOf(ImageExportOptions{})
	derivedTypes["ImageExportOptions"] = "ExportOptions"
	typeDeterminers["ImageExportOptions"] = make(map[string]string)
	typeRegistry["Images"] = reflect.TypeOf(Images{})
	derivedTypes["Images"] = "ResourceBase"
	typeDeterminers["Images"] = make(map[string]string)
	typeRegistry["LayoutSlide"] = reflect.TypeOf(LayoutSlide{})
	derivedTypes["LayoutSlide"] = "ResourceBase"
	typeDeterminers["LayoutSlide"] = make(map[string]string)
	typeRegistry["LayoutSlides"] = reflect.TypeOf(LayoutSlides{})
	derivedTypes["LayoutSlides"] = "ResourceBase"
	typeDeterminers["LayoutSlides"] = make(map[string]string)
	typeRegistry["MasterSlide"] = reflect.TypeOf(MasterSlide{})
	derivedTypes["MasterSlide"] = "ResourceBase"
	typeDeterminers["MasterSlide"] = make(map[string]string)
	typeRegistry["MasterSlides"] = reflect.TypeOf(MasterSlides{})
	derivedTypes["MasterSlides"] = "ResourceBase"
	typeDeterminers["MasterSlides"] = make(map[string]string)
	typeRegistry["Merge"] = reflect.TypeOf(Merge{})
	derivedTypes["Merge"] = "Task"
	typeDeterminers["Merge"] = make(map[string]string)
	typeDeterminers["Merge"]["Type"] = "Merge"
	typeRegistry["NoFill"] = reflect.TypeOf(NoFill{})
	derivedTypes["NoFill"] = "FillFormat"
	typeDeterminers["NoFill"] = make(map[string]string)
	typeDeterminers["NoFill"]["Type"] = "NoFill"
	typeRegistry["NotesSlide"] = reflect.TypeOf(NotesSlide{})
	derivedTypes["NotesSlide"] = "ResourceBase"
	typeDeterminers["NotesSlide"] = make(map[string]string)
	typeRegistry["OneValueSeries"] = reflect.TypeOf(OneValueSeries{})
	derivedTypes["OneValueSeries"] = "Series"
	typeDeterminers["OneValueSeries"] = make(map[string]string)
	typeDeterminers["OneValueSeries"]["DataPointType"] = "OneValue"
	typeRegistry["Paragraph"] = reflect.TypeOf(Paragraph{})
	derivedTypes["Paragraph"] = "ResourceBase"
	typeDeterminers["Paragraph"] = make(map[string]string)
	typeRegistry["Paragraphs"] = reflect.TypeOf(Paragraphs{})
	derivedTypes["Paragraphs"] = "ResourceBase"
	typeDeterminers["Paragraphs"] = make(map[string]string)
	typeRegistry["PathInputFile"] = reflect.TypeOf(PathInputFile{})
	derivedTypes["PathInputFile"] = "InputFile"
	typeDeterminers["PathInputFile"] = make(map[string]string)
	typeDeterminers["PathInputFile"]["Type"] = "Path"
	typeRegistry["PathOutputFile"] = reflect.TypeOf(PathOutputFile{})
	derivedTypes["PathOutputFile"] = "OutputFile"
	typeDeterminers["PathOutputFile"] = make(map[string]string)
	typeDeterminers["PathOutputFile"]["Type"] = "Path"
	typeRegistry["PatternFill"] = reflect.TypeOf(PatternFill{})
	derivedTypes["PatternFill"] = "FillFormat"
	typeDeterminers["PatternFill"] = make(map[string]string)
	typeDeterminers["PatternFill"]["Type"] = "Pattern"
	typeRegistry["PdfExportOptions"] = reflect.TypeOf(PdfExportOptions{})
	derivedTypes["PdfExportOptions"] = "ExportOptions"
	typeDeterminers["PdfExportOptions"] = make(map[string]string)
	typeRegistry["PictureFill"] = reflect.TypeOf(PictureFill{})
	derivedTypes["PictureFill"] = "FillFormat"
	typeDeterminers["PictureFill"] = make(map[string]string)
	typeDeterminers["PictureFill"]["Type"] = "Picture"
	typeRegistry["Placeholder"] = reflect.TypeOf(Placeholder{})
	derivedTypes["Placeholder"] = "ResourceBase"
	typeDeterminers["Placeholder"] = make(map[string]string)
	typeRegistry["Placeholders"] = reflect.TypeOf(Placeholders{})
	derivedTypes["Placeholders"] = "ResourceBase"
	typeDeterminers["Placeholders"] = make(map[string]string)
	typeRegistry["Portion"] = reflect.TypeOf(Portion{})
	derivedTypes["Portion"] = "ResourceBase"
	typeDeterminers["Portion"] = make(map[string]string)
	typeRegistry["Portions"] = reflect.TypeOf(Portions{})
	derivedTypes["Portions"] = "ResourceBase"
	typeDeterminers["Portions"] = make(map[string]string)
	typeRegistry["PptxExportOptions"] = reflect.TypeOf(PptxExportOptions{})
	derivedTypes["PptxExportOptions"] = "ExportOptions"
	typeDeterminers["PptxExportOptions"] = make(map[string]string)
	typeRegistry["RemoveShape"] = reflect.TypeOf(RemoveShape{})
	derivedTypes["RemoveShape"] = "Task"
	typeDeterminers["RemoveShape"] = make(map[string]string)
	typeDeterminers["RemoveShape"]["Type"] = "RemoveShape"
	typeRegistry["RemoveSlide"] = reflect.TypeOf(RemoveSlide{})
	derivedTypes["RemoveSlide"] = "Task"
	typeDeterminers["RemoveSlide"] = make(map[string]string)
	typeDeterminers["RemoveSlide"]["Type"] = "RemoveSlide"
	typeRegistry["ReorderSlide"] = reflect.TypeOf(ReorderSlide{})
	derivedTypes["ReorderSlide"] = "Task"
	typeDeterminers["ReorderSlide"] = make(map[string]string)
	typeDeterminers["ReorderSlide"]["Type"] = "ReoderSlide"
	typeRegistry["ReplaceText"] = reflect.TypeOf(ReplaceText{})
	derivedTypes["ReplaceText"] = "Task"
	typeDeterminers["ReplaceText"] = make(map[string]string)
	typeDeterminers["ReplaceText"]["Type"] = "ReplaceText"
	typeRegistry["RequestInputFile"] = reflect.TypeOf(RequestInputFile{})
	derivedTypes["RequestInputFile"] = "InputFile"
	typeDeterminers["RequestInputFile"] = make(map[string]string)
	typeDeterminers["RequestInputFile"]["Type"] = "Request"
	typeRegistry["ResetSlide"] = reflect.TypeOf(ResetSlide{})
	derivedTypes["ResetSlide"] = "Task"
	typeDeterminers["ResetSlide"] = make(map[string]string)
	typeDeterminers["ResetSlide"]["Type"] = "ResetSlide"
	typeRegistry["ResponseOutputFile"] = reflect.TypeOf(ResponseOutputFile{})
	derivedTypes["ResponseOutputFile"] = "OutputFile"
	typeDeterminers["ResponseOutputFile"] = make(map[string]string)
	typeDeterminers["ResponseOutputFile"]["Type"] = "Response"
	typeRegistry["Save"] = reflect.TypeOf(Save{})
	derivedTypes["Save"] = "Task"
	typeDeterminers["Save"] = make(map[string]string)
	typeDeterminers["Save"]["Type"] = "Save"
	typeRegistry["SaveShape"] = reflect.TypeOf(SaveShape{})
	derivedTypes["SaveShape"] = "Task"
	typeDeterminers["SaveShape"] = make(map[string]string)
	typeDeterminers["SaveShape"]["Type"] = "SaveShape"
	typeRegistry["SaveSlide"] = reflect.TypeOf(SaveSlide{})
	derivedTypes["SaveSlide"] = "Task"
	typeDeterminers["SaveSlide"] = make(map[string]string)
	typeDeterminers["SaveSlide"]["Type"] = "SaveSlide"
	typeRegistry["ScatterSeries"] = reflect.TypeOf(ScatterSeries{})
	derivedTypes["ScatterSeries"] = "Series"
	typeDeterminers["ScatterSeries"] = make(map[string]string)
	typeDeterminers["ScatterSeries"]["DataPointType"] = "Scatter"
	typeRegistry["ShapeBase"] = reflect.TypeOf(ShapeBase{})
	derivedTypes["ShapeBase"] = "ResourceBase"
	typeDeterminers["ShapeBase"] = make(map[string]string)
	typeRegistry["Shapes"] = reflect.TypeOf(Shapes{})
	derivedTypes["Shapes"] = "ResourceBase"
	typeDeterminers["Shapes"] = make(map[string]string)
	typeRegistry["Slide"] = reflect.TypeOf(Slide{})
	derivedTypes["Slide"] = "ResourceBase"
	typeDeterminers["Slide"] = make(map[string]string)
	typeRegistry["SlideAnimation"] = reflect.TypeOf(SlideAnimation{})
	derivedTypes["SlideAnimation"] = "ResourceBase"
	typeDeterminers["SlideAnimation"] = make(map[string]string)
	typeRegistry["SlideBackground"] = reflect.TypeOf(SlideBackground{})
	derivedTypes["SlideBackground"] = "ResourceBase"
	typeDeterminers["SlideBackground"] = make(map[string]string)
	typeRegistry["SlideComments"] = reflect.TypeOf(SlideComments{})
	derivedTypes["SlideComments"] = "ResourceBase"
	typeDeterminers["SlideComments"] = make(map[string]string)
	typeRegistry["Slides"] = reflect.TypeOf(Slides{})
	derivedTypes["Slides"] = "ResourceBase"
	typeDeterminers["Slides"] = make(map[string]string)
	typeRegistry["SolidFill"] = reflect.TypeOf(SolidFill{})
	derivedTypes["SolidFill"] = "FillFormat"
	typeDeterminers["SolidFill"] = make(map[string]string)
	typeDeterminers["SolidFill"]["Type"] = "Solid"
	typeRegistry["SplitDocumentResult"] = reflect.TypeOf(SplitDocumentResult{})
	derivedTypes["SplitDocumentResult"] = "ResourceBase"
	typeDeterminers["SplitDocumentResult"] = make(map[string]string)
	typeRegistry["SvgExportOptions"] = reflect.TypeOf(SvgExportOptions{})
	derivedTypes["SvgExportOptions"] = "ExportOptions"
	typeDeterminers["SvgExportOptions"] = make(map[string]string)
	typeRegistry["SwfExportOptions"] = reflect.TypeOf(SwfExportOptions{})
	derivedTypes["SwfExportOptions"] = "ExportOptions"
	typeDeterminers["SwfExportOptions"] = make(map[string]string)
	typeRegistry["TextItems"] = reflect.TypeOf(TextItems{})
	derivedTypes["TextItems"] = "ResourceBase"
	typeDeterminers["TextItems"] = make(map[string]string)
	typeRegistry["Theme"] = reflect.TypeOf(Theme{})
	derivedTypes["Theme"] = "ResourceBase"
	typeDeterminers["Theme"] = make(map[string]string)
	typeRegistry["TiffExportOptions"] = reflect.TypeOf(TiffExportOptions{})
	derivedTypes["TiffExportOptions"] = "ExportOptions"
	typeDeterminers["TiffExportOptions"] = make(map[string]string)
	typeRegistry["UpdateBackground"] = reflect.TypeOf(UpdateBackground{})
	derivedTypes["UpdateBackground"] = "Task"
	typeDeterminers["UpdateBackground"] = make(map[string]string)
	typeDeterminers["UpdateBackground"]["Type"] = "UpdateBackground"
	typeRegistry["UpdateShape"] = reflect.TypeOf(UpdateShape{})
	derivedTypes["UpdateShape"] = "Task"
	typeDeterminers["UpdateShape"] = make(map[string]string)
	typeDeterminers["UpdateShape"]["Type"] = "UpdateShape"
	typeRegistry["ViewProperties"] = reflect.TypeOf(ViewProperties{})
	derivedTypes["ViewProperties"] = "ResourceBase"
	typeDeterminers["ViewProperties"] = make(map[string]string)
	typeRegistry["XpsExportOptions"] = reflect.TypeOf(XpsExportOptions{})
	derivedTypes["XpsExportOptions"] = "ExportOptions"
	typeDeterminers["XpsExportOptions"] = make(map[string]string)
	typeRegistry["BoxAndWhiskerSeries"] = reflect.TypeOf(BoxAndWhiskerSeries{})
	derivedTypes["BoxAndWhiskerSeries"] = "OneValueSeries"
	typeDeterminers["BoxAndWhiskerSeries"] = make(map[string]string)
	typeDeterminers["BoxAndWhiskerSeries"]["DataPointType"] = "OneValue"
	typeRegistry["Chart"] = reflect.TypeOf(Chart{})
	derivedTypes["Chart"] = "ShapeBase"
	typeDeterminers["Chart"] = make(map[string]string)
	typeDeterminers["Chart"]["Type"] = "Chart"
	typeDeterminers["Chart"]["ShapeType"] = "Chart"
	typeRegistry["DocumentReplaceResult"] = reflect.TypeOf(DocumentReplaceResult{})
	derivedTypes["DocumentReplaceResult"] = "Document"
	typeDeterminers["DocumentReplaceResult"] = make(map[string]string)
	typeRegistry["GeometryShape"] = reflect.TypeOf(GeometryShape{})
	derivedTypes["GeometryShape"] = "ShapeBase"
	typeDeterminers["GeometryShape"] = make(map[string]string)
	typeRegistry["GraphicalObject"] = reflect.TypeOf(GraphicalObject{})
	derivedTypes["GraphicalObject"] = "ShapeBase"
	typeDeterminers["GraphicalObject"] = make(map[string]string)
	typeDeterminers["GraphicalObject"]["Type"] = "GraphicalObject"
	typeDeterminers["GraphicalObject"]["ShapeType"] = "GraphicalObject"
	typeRegistry["GroupShape"] = reflect.TypeOf(GroupShape{})
	derivedTypes["GroupShape"] = "ShapeBase"
	typeDeterminers["GroupShape"] = make(map[string]string)
	typeDeterminers["GroupShape"]["Type"] = "GroupShape"
	typeDeterminers["GroupShape"]["ShapeType"] = "GroupShape"
	typeRegistry["OleObjectFrame"] = reflect.TypeOf(OleObjectFrame{})
	derivedTypes["OleObjectFrame"] = "ShapeBase"
	typeDeterminers["OleObjectFrame"] = make(map[string]string)
	typeDeterminers["OleObjectFrame"]["Type"] = "OleObjectFrame"
	typeDeterminers["OleObjectFrame"]["ShapeType"] = "OleObjectFrame"
	typeRegistry["SlideReplaceResult"] = reflect.TypeOf(SlideReplaceResult{})
	derivedTypes["SlideReplaceResult"] = "Slide"
	typeDeterminers["SlideReplaceResult"] = make(map[string]string)
	typeRegistry["SmartArt"] = reflect.TypeOf(SmartArt{})
	derivedTypes["SmartArt"] = "ShapeBase"
	typeDeterminers["SmartArt"] = make(map[string]string)
	typeDeterminers["SmartArt"]["Type"] = "SmartArt"
	typeDeterminers["SmartArt"]["ShapeType"] = "Diagram"
	typeRegistry["SmartArtShape"] = reflect.TypeOf(SmartArtShape{})
	derivedTypes["SmartArtShape"] = "ShapeBase"
	typeDeterminers["SmartArtShape"] = make(map[string]string)
	typeDeterminers["SmartArtShape"]["Type"] = "SmartArtShape"
	typeDeterminers["SmartArtShape"]["ShapeType"] = "Custom"
	typeRegistry["Table"] = reflect.TypeOf(Table{})
	derivedTypes["Table"] = "ShapeBase"
	typeDeterminers["Table"] = make(map[string]string)
	typeDeterminers["Table"]["Type"] = "Table"
	typeDeterminers["Table"]["ShapeType"] = "Table"
	typeRegistry["WaterfallSeries"] = reflect.TypeOf(WaterfallSeries{})
	derivedTypes["WaterfallSeries"] = "OneValueSeries"
	typeDeterminers["WaterfallSeries"] = make(map[string]string)
	typeDeterminers["WaterfallSeries"]["DataPointType"] = "OneValue"
	typeRegistry["AudioFrame"] = reflect.TypeOf(AudioFrame{})
	derivedTypes["AudioFrame"] = "GeometryShape"
	typeDeterminers["AudioFrame"] = make(map[string]string)
	typeDeterminers["AudioFrame"]["Type"] = "AudioFrame"
	typeDeterminers["AudioFrame"]["ShapeType"] = "AudioFrame"
	typeRegistry["Connector"] = reflect.TypeOf(Connector{})
	derivedTypes["Connector"] = "GeometryShape"
	typeDeterminers["Connector"] = make(map[string]string)
	typeDeterminers["Connector"]["Type"] = "Connector"
	typeRegistry["PictureFrame"] = reflect.TypeOf(PictureFrame{})
	derivedTypes["PictureFrame"] = "GeometryShape"
	typeDeterminers["PictureFrame"] = make(map[string]string)
	typeDeterminers["PictureFrame"]["Type"] = "PictureFrame"
	typeDeterminers["PictureFrame"]["ShapeType"] = "PictureFrame"
	typeRegistry["Shape"] = reflect.TypeOf(Shape{})
	derivedTypes["Shape"] = "GeometryShape"
	typeDeterminers["Shape"] = make(map[string]string)
	typeDeterminers["Shape"]["Type"] = "Shape"
	typeRegistry["VideoFrame"] = reflect.TypeOf(VideoFrame{})
	derivedTypes["VideoFrame"] = "GeometryShape"
	typeDeterminers["VideoFrame"] = make(map[string]string)
	typeDeterminers["VideoFrame"]["Type"] = "VideoFrame"
	typeDeterminers["VideoFrame"]["ShapeType"] = "VideoFrame"
}

func createObjectForType(typeName string, data []byte) (interface{}, error) {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return nil, err
	}
	return createObjectForTypeMap(typeName, typeName, objmap), nil
}

func createObjectForTypeMap(typeName string, baseTypeName string, objMap map[string]*json.RawMessage) interface{} {
	for k, v := range derivedTypes {
		if v == typeName {
			subObject := createObjectForTypeMap(k, typeName, objMap)
			if subObject != nil {
				return subObject
			}
		}
	}
	matches := 0
	if detMap, ok := typeDeterminers[typeName]; ok {
		matches = 0
		for k, v := range detMap {
			objValue, objOk := objMap[k]
			if !objOk {
				objValue, objOk = objMap[lcFirst(k)]
			}
			if !objOk {
				return nil
			}
			if objOk {
				var objString string
			 	err := json.Unmarshal(*objValue, &objString)
				if err == nil {
					if objString != v {
						return nil
					}
				}
			}
			matches++
		}
	}
	if typeName == baseTypeName || matches > 0 {
		if reflectType, ok := typeRegistry[typeName]; ok {
			return reflect.New(reflectType).Interface()
		}
	}
	return nil
}
 
func lcFirst(str string) string {
    for i, v := range str {                                                                                                                                           
        return string(unicode.ToLower(v)) + str[i+1:]
    }  
    return ""
}