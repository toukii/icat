package icat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/harrydb/go/img/grayscale"
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
			sp := strings.Split(args[0], ".")
			size := len(sp)
			if size <= 0 {
				return
			}
			viper.Set("ext", sp[size-1])
		} else {
			// return
		}
		if err := Excute(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	Command.PersistentFlags().IntP("height", "H", 0, "image height")
	Command.PersistentFlags().IntP("width", "w", 0, "image width")
	Command.PersistentFlags().StringP("base64", "B", "", "base64")
	Command.PersistentFlags().StringP("ext", "x", "png", "ext:png|jpg/jpeg|gif")
	Command.PersistentFlags().BoolP("gray", "g", false, "gray image")

	viper.BindPFlag("height", Command.PersistentFlags().Lookup("height"))
	viper.BindPFlag("width", Command.PersistentFlags().Lookup("width"))
	viper.BindPFlag("base64", Command.PersistentFlags().Lookup("base64"))
	viper.BindPFlag("ext", Command.PersistentFlags().Lookup("ext"))
	viper.BindPFlag("gray", Command.PersistentFlags().Lookup("gray"))
}

func Excute() error {
	var r io.Reader
	var img image.Image

	if base64Cnt := viper.GetString("base64"); base64Cnt != "" {
		bs, err := base64.StdEncoding.DecodeString(base64Cnt)
		if err != nil {
			return err
		}
		r = bytes.NewReader(bs)
	} else {
		imgFile := viper.GetString("input")
		fd, err := os.Open(imgFile)
		if err != nil {
			return err
		}
		r = fd
	}

	var err error
	if viper.GetString("ext") == "png" {
		img, err = png.Decode(r)
		if err != nil {
			return err
		}
	} else if viper.GetString("ext") == "jpg" || viper.GetString("ext") == "jpeg" {
		img, err = jpeg.Decode(r)
		if err != nil {
			return err
		}
	} else if viper.GetString("ext") == "gif" {
		img, err = gif.Decode(r)
		if err != nil {
			return err
		}
	}

	if viper.GetBool("gray") {
		img = grayscale.Convert(img, grayscale.ToGrayLuminance)
	}

	return ICatRect(img, viper.GetInt("height"), viper.GetInt("width"), os.Stdout)
}
