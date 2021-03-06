package primitives

import (
	"chaoticneutraltech/transition/shaders"
	"chaoticneutraltech/transition/textures"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type rectangle struct {
	vertices      []float32
	indices       []uint32
	vbo           uint32
	vao           uint32
	ebo           uint32
	shaderProgram uint32
	texture       uint32
}

func NewRectangle() *rectangle {
	r := new(rectangle)
	r.vertices = []float32{
		// positions          // colors           // texture coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // bottom left
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // top left
	}
	r.indices = []uint32{
		0, 1, 2,
		2, 3, 0,
	}
	r.vbo = 0
	r.vao = 0
	r.ebo = 0
	r.shaderProgram = 0

	return r
}

func (r *rectangle) GetShaderProgram() uint32 {
	return r.shaderProgram
}

func (r *rectangle) GetVertices() []float32 {
	return r.vertices
}

func (r *rectangle) Draw() {

	gl.BindTexture(gl.TEXTURE_2D, r.texture)

	gl.UseProgram(r.shaderProgram)
	gl.BindVertexArray(r.vao)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
}

func (r *rectangle) Init() {

	vertexShader, err := shaders.Compile("shaders/examples/textureVertex.vs", gl.VERTEX_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	fragmentShader, err := shaders.Compile("shaders/examples/textureFragment.fs", gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	r.shaderProgram, err = shaders.CreateShaderProgram([]uint32{vertexShader, fragmentShader})
	if err != nil {
		log.Fatal(err)
	}

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)
	gl.GenBuffers(1, &r.ebo)

	gl.BindVertexArray(r.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(r.vertices), gl.Ptr(r.vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(r.indices), gl.Ptr(r.indices), gl.STATIC_DRAW)

	//Postion
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	//Color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	//Texture
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))
	gl.EnableVertexAttribArray(2)

	texture, err := textures.NewTexture("textures/container.jpg", textures.Jpg)
	if err != nil {
		log.Fatal(err)
	}

	texture.GetTextureDimensions()

	//fmt.Println(texture.GetTextureBytes())

	r.texture = texture.GenerateGlTexture()

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}
