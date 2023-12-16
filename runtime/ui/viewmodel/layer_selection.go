package viewmodel

import (
	"github.com/khulnasoft-lab/demystify/demystify/image"
)

type LayerSelection struct {
	Layer                                                      *image.Layer
	BottomTreeStart, BottomTreeStop, TopTreeStart, TopTreeStop int
}
