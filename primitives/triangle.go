package primitives

import (
	"chaoticneutraltech/transition/shaders"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"maze.io/x/math32"
)

type triangle struct {
	vertices      []float32
	vbo           uint32
	vao           uint32
	shaderProgram uint32
}

func NewTriangle() *triangle {
	t := new(triangle)
	t.vertices = []float32{
		// positions         // colors
		0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom left
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // top
	}
	t.vbo = 0
	t.vao = 0
	t.shaderProgram = 0

	return t
}

func (t *triangle) GetShaderProgram() uint32 {
	return t.shaderProgram
}

func (t *triangle) GetVertices() []float32 {
	return t.vertices
}

func (t *triangle) Draw() {
	gl.UseProgram(t.shaderProgram)

	var timeValue float32 = float32(glfw.GetTime())
	var greenValue float32 = math32.Sin(timeValue)/2.0 + 0.5
	vertexColorLocation := gl.GetUniformLocation(t.shaderProgram, gl.Str("ourColor\x00"))
	gl.Uniform4f(vertexColorLocation, 0.0, greenValue, 0.0, 1.0)

	gl.BindVertexArray(t.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (t *triangle) Init() {

	vertexShader, err := shaders.Compile("shaders/examples/rainbowVertex.vs", gl.VERTEX_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	fragmentShader, err := shaders.Compile("shaders/examples/rainbowFragment.fs", gl.FRAGMENT_SHADER)
	if err != nil {
		log.Fatal(err)
	}

	t.shaderProgram, err = shaders.CreateShaderProgram([]uint32{vertexShader, fragmentShader})
	if err != nil {
		log.Fatal(err)
	}

	gl.GenVertexArrays(1, &t.vao)
	gl.GenBuffers(1, &t.vbo)

	gl.BindVertexArray(t.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, t.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(t.vertices)*4, gl.Ptr(t.vertices), gl.STATIC_DRAW)

	//Position attributes
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, nil)
	gl.EnableVertexAttribArray(0)

	//Color attributes
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	//gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE);
}
