package window

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type window struct {
	widht  int
	height int
	title  string
}

func NewWindow(widht int, height int, title string) *window {
	w := new(window)
	w.widht = widht
	w.height = height
	w.title = title

	return w
}

func (w window) Create() *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfwWindow, err := glfw.CreateWindow(w.widht, w.height, w.title, nil, nil)
	if err != nil {
		panic(err)
	}

	return glfwWindow
}
