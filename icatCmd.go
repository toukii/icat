package icat

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Command = &cobra.Command{
	Use:   "icat",
	Short: "image cat",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		size := len(args)
		if size > 0 {
			viper.Set("input", args[0])
		} else {
			return
		}
		Excute()
	},
}

func init() {
	Command.PersistentFlags().IntP("height", "H", 0, "image height")
	Command.PersistentFlags().IntP("width", "w", 0, "image width")

	viper.BindPFlag("height", Command.PersistentFlags().Lookup("height"))
	viper.BindPFlag("width", Command.PersistentFlags().Lookup("width"))
}

func Excute() error {
	var img image.Image
	imgFile := viper.GetString("input")
	fd, err := os.Open(imgFile)
	if err != nil {
		return err
	}
	if strings.HasSuffix(imgFile, ".png") {
		img, err = png.Decode(fd)
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(imgFile, ".jpg") || strings.HasSuffix(imgFile, ".jpeg") {
		img, err = jpeg.Decode(fd)
		if err != nil {
			return err
		}
	}

	return ICatRect(img, viper.GetInt("height"), viper.GetInt("width"), os.Stdout)
}
