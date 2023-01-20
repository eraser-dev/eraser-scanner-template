package main

import (
	"context"

	"github.com/Azure/eraser/api/unversioned"
	template "github.com/Azure/eraser/pkg/scanners/template"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	log := logf.Log.WithName("scanner").WithValues("provider", "customScanner")

	// create image provider with custom values
	imageProvider := template.NewImageProvider(
		template.WithContext(context.Background()),
		template.WithMetrics(true),
		template.WithDeleteScanFailedImages(true),
		template.WithLogger(log),
	)

	// retrieve list of all non-running, non-excluded images from collector container
	allImages, err := imageProvider.ReceiveImages()
	if err != nil {
		log.Error(err, "unable to retrieve list of images from collector container")
		return
	}

	// scan images with custom scanner
	nonCompliant, failedImages := scan(allImages)

	// send images to eraser container
	if err := imageProvider.SendImages(nonCompliant, failedImages); err != nil {
		log.Error(err, "unable to send non-compliant images to eraser container")
		return
	}

	// complete scan
	imageProvider.Finish()
}

// TODO: implement customized scanner
func scan(allImages []unversioned.Image) ([]unversioned.Image, []unversioned.Image) {
	// scan images and partition into non-compliant and failedImages
	var nonCompliant, failedImages []unversioned.Image

	return nonCompliant, failedImages
}
