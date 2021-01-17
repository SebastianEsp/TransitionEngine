package textures

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Format int

const (
	Jpeg = iota
	Jpg
	Png
)

type texture struct {
	source string
	bytes  []byte
	format Format
	widht  int
	height int
}

func NewTexture(source string, format Format) (*texture, error) {

	texture := new(texture)
	texture.source = source
	texture.format = format

	f, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var img image.Image

	switch format {
	case Jpeg, Jpg:
		img, err = jpeg.Decode(f)
		if err != nil {
			return nil, err
		}
	case Png:
		img, err = png.Decode(f)
		if err != nil {
			return nil, err
		}
	}

	rgba := image.NewNRGBA(img.Bounds())
	draw.Draw(rgba, img.Bounds(), img, image.ZP, draw.Src)

	texture.bytes = rgba.Pix
	texture.widht = rgba.Bounds().Dx()
	texture.height = rgba.Bounds().Dy()

	return texture, nil
}

func (t texture) GetTextureDimensions() (int, int) {
	return t.widht, t.height
}

func (t texture) GetTextureBytes() []byte {
	return t.bytes
}

func (t texture) GenerateGlTexture() uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	//Set texture wrapping/filtering options
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	//Generate texture
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(t.widht), int32(t.height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(t.bytes))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return tex
}

func (f Format) String() string {
	return [...]string{"jpeg", "jpg", "png"}[f]
}
