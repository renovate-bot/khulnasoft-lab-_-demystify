package runtime

import (
	"github.com/spf13/viper"

	"github.com/khulnasoft-lab/demystify/demystify"
)

type Options struct {
	Ci           bool
	Image        string
	Source       demystify.ImageSource
	IgnoreErrors bool
	ExportFile   string
	CiConfig     *viper.Viper
	BuildArgs    []string
}
