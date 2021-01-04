![](https://img.shields.io/badge/api-v3.0-lightgrey)  [![GitHub license](https://img.shields.io/github/license/aspose-slides-cloud/aspose-slides-cloud-go)](https://github.com/aspose-slides-cloud/aspose-slides-cloud-go)

# Aspose.Slides Cloud SDK for Go
This repository contains Aspose.Slides Cloud SDK for Go source code. This SDK allows you to work with Aspose.Slides Cloud REST APIs in your Go applications.

## Key Features
* Conversion between various document-related formats (20+ formats supported), including PDF<->PowerPoint conversion
* Download slides and shapes in various formats, including PDF and SVG
* Merge and split PowerPoint presentations
* Access PowerPoint presentation metadata and statistics
* Find and replace
* Full read & write access to Document Object Model, including slides, shapes, paragraphs, portions and many others
* Support of Aspose.Storage API

### New Features & Recent API Changes

#### 20.12
* New property "HttpRequestTimeout" is added to SDK configuration. It allows to specify a custom timeout (in seconds) for HTTP requests made by the SDK.

#### 20.10
* Removed shapeType property from BaseShape class; replaced geometryShapeType with shapeType property for GeometryShape.
* Removed redundant ResourceUriElement class. All properties of ResourceUriElement type are changed to ResourceUri type.

#### 20.9
* New getSections, putSections, putSection, postSection, postSectionMove, deleteSection, deleteSections methods to add, update and delete presentation sections.
* New postSlidesHeaderFooter, postSlideHeaderFooter, getSlideHeaderFooter, postNotesSlideHeaderFooter, getNotesSlideHeaderFooter methods to manage header/footer settings for slides.
* Replaces cateslidesries for parentCateslidesries & level properties in ChartCateslidesry class to properly support Sunburst & Treeview charts.

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
	var request asposeslidescloud.PostSlidesConvertRequest
	request.Format = "pdf"
        file, _ := ioutil.ReadFile("MyPresentation.pptx")
	request.Document = file
	r, _, _ := apiClient.SlidesApi.PostSlidesConvert(request)
	fmt.Printf("My PDF was saved to %s", r.Name())
```

## Aspose.Slides Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift|Perl|Go|
|---|---|---|---|---|---|---|--|--|--|
| [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-dotnet) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-java) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-php) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-python) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-ruby)  | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-nodejs) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-android) | [GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-swift)|[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-perl) |[GitHub](https://github.com/aspose-slides-cloud/aspose-slides-cloud-slides) |
| [NuGet](https://www.nuget.org/packages/Aspose.slides-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Composer](https://packagist.org/packages/aspose/slides-sdk-php) | [PIP](https://pypi.org/project/asposeslidescloud/) | [GEM](https://rubygems.org/gems/aspose_slides_cloud)  | [NPM](https://www.npmjs.com/package/asposeslidescloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-slides-cloud) | [Cocoapods](https://cocoapods.org/pods/AsposeslidesCloud)|[Meta Cpan](https://metacpan.org/release/AsposeSlidesCloud-SlidesApi) | [Go.Dev](https://pkg.slides.dev/github.com/aspose-slides-cloud/aspose-slides-cloud-slides/) | 

[Product Page](https://products.aspose.cloud/slides/slides) | [Documentation](https://docs.aspose.cloud/display/slidescloud/Home) | [API Reference](https://apireference.aspose.cloud/slides/) | [Code Samples](https://github.com/aspose-slides-cloud/aspose-slides-cloud-slides) | [Blog](https://blog.aspose.cloud/cateslidesry/slides/) | [Free Support](https://forum.aspose.cloud/c/slides) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
