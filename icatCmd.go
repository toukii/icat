package icat

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/harrydb/go/img/grayscale"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/toukii/goutils"
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
	Command.PersistentFlags().IntP("top", "t", 0, "image top")
	Command.PersistentFlags().IntP("left", "l", 0, "image left")
	Command.PersistentFlags().StringP("base64", "B", "", "base64")
	Command.PersistentFlags().StringP("ext", "x", "png", "ext:png|jpg/jpeg|gif")
	Command.PersistentFlags().StringP("output", "o", "stdout", "output:stdout|filename")
	Command.PersistentFlags().BoolP("gray", "g", false, "gray image")

	viper.BindPFlag("height", Command.PersistentFlags().Lookup("height"))
	viper.BindPFlag("width", Command.PersistentFlags().Lookup("width"))
	viper.BindPFlag("top", Command.PersistentFlags().Lookup("top"))
	viper.BindPFlag("left", Command.PersistentFlags().Lookup("left"))
	viper.BindPFlag("base64", Command.PersistentFlags().Lookup("base64"))
	viper.BindPFlag("ext", Command.PersistentFlags().Lookup("ext"))
	viper.BindPFlag("gray", Command.PersistentFlags().Lookup("gray"))
	viper.BindPFlag("output", Command.PersistentFlags().Lookup("output"))
}

func Excute() error {
	var w io.Writer
	output := viper.GetString("output")
	if output == "stdout" {
		w = NewEncodeWr(os.Stdout, nil)
		// w = NewEncodeWr(os.Stdout, nil)
		w = NewEncodeStdout()
	} else {
		fd, _ := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)
		defer fd.Close()
		w = fd
	}

	if base64Cnt := viper.GetString("base64"); base64Cnt != "" {
		return ICatBase64(base64Cnt, w)
	}

	var r io.Reader
	var img image.Image
	var err error

	imgFile := viper.GetString("input")
	if strings.HasPrefix(imgFile, "http://") || strings.HasPrefix(imgFile, "https://") {
		resp, err := http.Get(imgFile)
		if err != nil {
			return err
		}
		r = resp.Body
	} else {
		fd, err := os.Open(imgFile)
		if err != nil {
			return err
		}
		r = fd
	}

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
	} else if viper.GetString("ext") == "svg" {
		img, err = DecodeSVG(r)
		if err != nil {
			return err
		}
	} else {
		img, err = png.Decode(r)
		if err != nil {
			return err
		}
	}

	if viper.GetBool("gray") {
		img = grayscale.Convert(img, grayscale.ToGrayLightness)

	}
	if output != "stdout" {
		defer CatRect(img, viper.GetInt("height"), viper.GetInt("width"), viper.GetInt("top"), viper.GetInt("left"), NewEncodeWr(os.Stdout, nil))
	}

	return CatRect(img, viper.GetInt("height"), viper.GetInt("width"), viper.GetInt("top"), viper.GetInt("left"), w)
}
