![](https://img.shields.io/badge/api-v3.0-lightgrey)  [![GitHub license](https://img.shields.io/github/license/aspose-slides-cloud/aspose-slides-cloud-go)](https://github.com/aspose-slides-cloud/aspose-slides-cloud-go)

# Go REST API to Process Presentation in Cloud
This repository contains Aspose.Slides Cloud SDK for Go source code. This SDK allows you to [process & manipulate PPT, PPTX, ODP, OTP](https://products.aspose.cloud/slides/go) using Aspose.slides Cloud REST APIs in your applications.

You may want to check out Aspose free [Powerpoint to PDF](https://products.aspose.app/slides/conversion), [Powerpoint to Word](https://products.aspose.app/slides/conversion/ppt-to-word), [Powerpoint to JPG](https://products.aspose.app/slides/conversion/ppt-to-jpg), [Powerpoint to PNG](https://products.aspose.app/slides/conversion/ppt-to-png), [PDF to Powerpoint](https://products.aspose.app/slides/import/pdf-to-powerpoint), [JPG to Powerpoint](https://products.aspose.app/slides/import/jpg-to-ppt), and [PNG to Powerpoint](https://products.aspose.app/slides/import/png-to-ppt) converters because they are live implementations of popular conversion processes.

## Presentation Processing Features

- Fetch presentation images in any of the supported file formats.
- Copy the layout side or clone the master slide from the source presentation.
- Process slides shapes, slides notes, placeholders, colors & font theme info.
- Programmatically create a presentation from HTML & export it to various formats.
- Merge multiple presentations or split the single presentation into multiple ones.
- Extract and replace text from a specific slide or an entire presentation.

## Read & Write Presentation Formats

**Microsoft PowerPoint:** PPT, PPTX, PPS, PPSX, PPTM, PPSM, POTX, POTM
**OpenOffice:** ODP, OTP, FODP
**Other**: PDF, PDF/A

## Save Presentation As

**Fixed Layout:** XPS
**Images:** JPEG, PNG, BMP, TIFF, GIF, SVG
**Web:** HTML/HTML5
**Other:** MPEG4, SWF (export whole presentations)

## Enhancements in Version 22.9

* Added new **ReplaceFont** and **ReplaceFontOnline** methods that allow replacing presentation fonts.
* Added new **SetEmbeddedFontFromRequest** and **SetEmbeddedFontFromRequestOnline** that allow to embed fonts without uploading them to the storage.
* Added an optional **fontsFolder** parameter to **SetEmbeddedFont** and **SetEmbeddedFontOnline** methods to enable using custom fonts with those methods.
* Added **FontSubstRules** property to **ExportOptions** class to allow specifying font substitution rules in conversion methods.
* Added an optional **group** parameter to **ImportImagesFromSvg** method to allow importing SVG images as a GroupShape.
* Added new **CreateSmartArtNode** and  **DeleteSmartArtNode** methods to manage individual nodes in SmartArt shapes.
* Added **FillFomat**, **LineFormat**, **EffectFormat** and  **ThreeDFormat** properties to **DataPoint** class that allow formatting individual chart data points.

## Enhancements in Version 22.8

* With new **GetFonts** and **GetFontsOnline** methods you can get a list for fonts used in the presentation.
* Added new **SetEmbeddedFont**, **SetEmbeddedFontOnline**, **DeleteEmbeddedFont** and **DeleteEmbeddedFontOnline** methods to embed/unembed presentation fonts.
* Added new **ImportImagesFromSvg** method to import SVG images as individual geometry shapes.

## Enhancements in Version 22.7

* Added **Html**, **Pdf**, **Xps**, **Pptx**, **Odp**, **Otp**, **Ppt**, **Pps**, **Ppsx**, **Pptm**, **Ppsm**, **Potx**, **Pot**, **Potm**, **Svg**, **Fodp**, **Xaml**, **Html5** to the list of allowed values for **SlideExportFormat** enum. You can now export slide notes to those formats.
* Added **Html5** to the list of allowed values for **SlideExportFormat** enum. You can now export individual slides to HTML5.
* Added **Url** to the list of allowed values for **PresentationToMerge.SourceEnum** enum. You can now merge presentations from external URLs.
* New **DeleteUnusedMasterSlides** and **DeleteUnusedMasterSlidesOnline** methods.
* New **SetChartAxis**, **SetChartLegend**, **SetChartWall** methods allow to moduly chart elements more conveniently.
* **UpdateChartSeriesGroup** method was renamed to **SetChartSeriesGroup**.

## Enhancements in Version 22.6
* Added ned **GetSubshapeParagraphEffective**, **GetParagraphEffective** and **GetSubshapePortionEffective**, **GetPortionEffective** methods to retrieve actual format values for paragraphs and portions, whether they are inherited from parent entities or not.
* Password parameter is now optional for **GetProtectionProperties** method. So, you don't need to specify the password to check whether a presentation has a password.
* Added new **ChartSeriesGroup** class and **SeriesGroups** property to **Chart** class to enable managing chart series groups. Added new **UpdateChartSeriesGroup** method.
* Added **HasRoundedCorners** property to **Chart** class.
* Added **InvertIfNegative** property to **OneValueChartDataPoint** class.
* Changed **FormatScheme** class to return actual format values instead of resource links.

## Enhancements in Version 22.5
* Added **Paragraphs** property to **SmartArtNode** class to enable getting and setting text for SmartArt nodes.
* Added **ImageTransformList** property to **PictureFill** class to enable setting image transform effects. Added **ImageTransformEffect** class and subclasses for different kind of of effects.
* Added **PictureFillFormat** property to **VideoFrame** class to enable setting poster image for video frames.
* Added **SlideIndex** and **ShapeIndex** properties to **ResourceUri** class to simplify retrieveing slide & shape indexes for resources.
* Removed redundant **BoxAndWhiskersSeries**, **WaterfallSeries** and **WaterfallChartDataPoint** classes.

## Enhancements in Version 22.4
* Added **TransitionType** and **SlidesTransitionDuration** properties to **VideoExportOptions** class to enable creation videos with transitions.
* Added **DefaultPortionFormat** property to **Paragraph** class. Added new **PortionFormat** class.
* Added **EmbeddedFileBase64Data** and a number of other properties to **OleObjectFrame** class to enable creation of OLE Object frames.
* Added **AccessPermissions** class to support access permissions for PDF export.
* Added **PictureFillformat** property to **AudioFrame** class.
* Added **RowIndex** and **ColumnIndex** properties to **TableCell** class.
* Moved **Width** and **Height** properties from **ExportOptions** base class to the new **ImageExportOptionsBase class**. This is a superclass for **ImageExportOptions**, **GifExportOptions** and **TiffExportOptions** classes.
* Removed redundant **Shapes** property from **ShapeBase** class. It is only left for **GroupShape** class.

## Enhancements in Version 22.3
* Added **ModernSlideComment** class to support modern comments. Also added **SlideCommentBase** as base class for comments.
* Added optional **shapeIndex** parameter to **CreateComment** and **CreateCommentOnline** methods.
* Added **GetParagraphRectangle**, **GetPortionRectangle** method and new **TextBounds** class to get paragraph or portion bounds.
* Added optional **shapeType** parameter for **GetShapes** method. You can now get list of shapes of a particular type (e.g. charts or tables).
* Added **FontFallbackRules** class and **FontFallbackRules** property to **ExportOptions** class.
* Added **LatinFont**, **EastAsianFont** and **ComplexScriptFont** properties to **Portion** class to enable getting and seting portion font name.
* Added **ChartLinesFormat** class; added **MajorGridLinesFormat** and **MinorGridLinesFormat** properties to Axis class.
* Added **HideLegend** boolean property to **Legend** class.

## Enhancements in Version 22.2
* Added **Mpeg4** to the list of allowed values for **ExportFormat** type. You can now convert presentations to video.
* New **HighlightShapeText** and **HighlightShapeRegex** methods.
* New **DeleteUnusedLayoutSlides** and **DeleteUnusedLayoutSlides** methods.
* New **ZoomFrame** and **ZoomObject** classes with a number of subclasses.
* Added **TextFrameFormat** property to **Shape** class to support WordArt.
* Added **XYSeries** class as a supercalss for **ScatterSeries** and **BubbleSeries** methods.
* Added **None** to the list of allowed values for **TimeUnitType** enum type.
* **Level** property of **Category** class is deprecated and will be removed after v22.4.

## Enhancements in Version 22.11

* You can now specify data sources for chart elements. This is done with new **DataSource** type which is use with propertries **Chart.DataSourceForCategories**, **Series.DataSourceForName**, **OneValueSeries.DataSourceForValues**, **XYSeries.DataSourceForXValues**, **XYSeries.DataSourceForYValues**, **BubbleSeries.DataSourceForBubbleSizeValues**.
* You can now specify formulas for data points using new properties **OneValueChartDataPoint.ValueFormula**, **ScatterChartDataPoint.XValueFormula**, **ScatterChartDataPoint.YValueFormula**, **BubbleChartDataPoint.BubbleSizeFormula**.
* Added boolean **UseFrameSize** and **UseFrameRotation** properties to **SvgExportOptions** class.

## Enhancements in Version 22.10

* Added new **GetSlideShowProperties** and  **SetSlideShowProperties** methods to get & set slideshow properties.
* Methods that work with shapes, paragraphs and portions now have an optional **subShape** parameter which can be used to access subshapes (SmartArt elements, grouped shapes). Separate methods for subshapes (**GetSubshapeParagraphs**, **CreateSubshape**, **AlignSubshapes** etc.) have been removed.
* Added boolean **RepeatUntilEndSlide** and **RepeatUntilNextClick** properties to **Effect** class.

## Enhancements in Version 22.1
* New **MathParagraph** property of **Portion** class allows to get and set math formulas. A set of **MathElement** subclasses allows to specify complex mathematical expressions.
* New **DownloadPortionAsMathMl** and **SavePortionAsMathMl** methods allow to export math formulas to MathML format.
* New **HyperlinkClick** and **HyperlinkMouseOver** methods of **ShapeBase** and **Portion** classes enable creation and manipulating hyperlinks for shapes and portions.
* New methods **GetShapeGeometryPath** and **SetShapeGeometryPath** can be used to get or set geometry paths for shapes. You can use lines, arcs and curves to specify custom shape boundaries.
* New **AlignSubshapes** method enables aligning shapes within a shape group.
* New **PlayAcrossSlides** and **RewindAudio** properties are added to **AudioFrame** class.
* Added **InClickSequence** to the list of allowable values for **AudioPlayModePreset** and **AudioPlayModePreset** enum types.

## Licensing
All Aspose.Slides Cloud SDKs are licensed under MIT License.

## How to use the SDK?
The complete source code is available in this repository folder.

## Prerequisites
To use Aspose Slides Cloud SDK for Go you need to register an account with [Aspose Cloud](https://www.aspose.cloud/) and lookup/create App Key and SID at [Cloud Dashboard](https://dashboard.aspose.cloud/#/apps). There is free quota available. For more details, see [Aspose Cloud Pricing](https://purchase.aspose.cloud/pricing).

### Installation

```sh
slides get github.com/aspose-slides-cloud/aspose-slides-cloud-slides
```

### Sample usage

The example code below converts a PowerPoint document to PDF format using asposeslidescloud library:
```slides
        cfg := asposeslidescloud.NewConfiguration()
	cfg.AppSid = "MyClientId"
	cfg.AppKey = "MyClientSecret"
       	apiClient := asposeslidescloud.NewAPIClient(cfg)
        file, _ := ioutil.ReadFile("MyPresentation.pptx")
	r, _, _ := apiClient.SlidesApi.Convert(file, "pdf")
	fmt.Printf("My PDF was saved to %s", r.Name())
```

## Aspose.Slides Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift|Perl|Go|
|---|---|---|---|---|---|---|--|--|--|
| [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-dotnet) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-java) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-php) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-python) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-ruby)  | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-nodejs) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-android) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-swift)|[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-perl) |[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-slides) |
| [NuGet](https://www.nuget.org/packages/Aspose.slides-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Composer](https://packagist.org/packages/aspose/slides-sdk-php) | [PIP](https://pypi.org/project/asposeslidescloud/) | [GEM](https://rubygems.org/gems/aspose_slides_cloud)  | [NPM](https://www.npmjs.com/package/asposeslidescloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Cocoapods](https://cocoapods.org/pods/AsposeslidesCloud)|[Meta Cpan](https://metacpan.org/release/AsposeSlidesCloud-SlidesApi) | [Go.Dev](https://pkg.slides.dev/github.com/aspose-slides-cloud/aspose-slides-cloud-slides/) |

[Product Page](https://products.aspose.cloud/slides/slides) | [Documentation](https://docs.aspose.cloud/display/slidescloud/Home) | [API Reference](https://apireference.aspose.cloud/slides/) | [Code Samples](https://github.com/aspose-slides-cloud/aspose-slides-cloud-slides) | [Blog](https://blog.aspose.cloud/cateslidesry/slides/) | [Free Support](https://forum.aspose.cloud/c/slides) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
