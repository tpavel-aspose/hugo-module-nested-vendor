package main

import (
	"vendor_bootstrap"

	"github.com/tpavel-aspose/hugo-module-nested-vendor/assets"
)

func main() {
	print("hugo-modelt nested vendor app\n")
	assets.AssetsLibCall()
	vendor_bootstrap.VendorBostrapLibCall()
}
