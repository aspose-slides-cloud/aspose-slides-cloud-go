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

## Licensing
All Aspose.Slides Cloud SDKs are licensed under MIT License.

## How to use the SDK?
The complete source code is available in this repository folder.

## Prerequisites
To use Aspose Slides Cloud SDK for Go you need to register an account with [Aspose Cloud](https://www.aspose.cloud/) and lookup/create App Key and SID at [Cloud Dashboard](https://dashboard.aspose.cloud/#/apps). There is free quota available. For more details, see [Aspose Cloud Pricing](https://purchase.aspose.cloud/pricing).

### Installation

```sh
go get github.com/aspose-slides-cloud/aspose-slides-cloud-go
```

### Sample usage

The example code below converts a PowerPoint document to PDF format using asposeslidescloud library:
```go
        cfg := asposeslidescloud.NewConfiguration()
	cfg.AppSid = "MyAppSid"
	cfg.AppKey = "MyAppKey"
       	apiClient := asposeslidescloud.NewAPIClient(cfg)
	var request asposeslidescloud.PostSlidesConvertRequest
	request.format = "pdf"
        file, _ := ioutil.ReadFile("MyPresentation.pptx")
	request.document = file
	r, _, _ := apiClient.SlidesApi.PostSlidesConvert(request)
	fmt.Printf("My PDF was saved to %s", r.Name())
```

You can check more [Examples](Examples) of using the SDK.

## Contact Us

Your feedback is very important to us. Please feel free to contact us using our [Support Forums](https://forum.aspose.cloud/c/slides).

## Resources
 
[Website](https://www.aspose.cloud/)  
[Product Home](https://products.aspose.cloud/slides/family)  
[API Reference](https://apireference.aspose.cloud/slides/)  
[Documentation](https://docs.aspose.cloud/display/slidescloud/Home)  
[Blog](https://blog.aspose.cloud/category/slides/)  
 
## Other languages

We generate our SDKs in different languages so you may check if yours is available in our [list](https://github.com/aspose-slides-cloud).
 
If you don't find your language in the list, feel free to request it from us, or use raw REST API requests as you can find it [here](https://products.aspose.cloud/slides/curl).
