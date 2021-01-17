package main

import (
	"chaoticneutraltech/transition/primitives"
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 800
const windowHeight = 600

var ()

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Transition", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, 800, 600)
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure global settings
	//gl.Enable(gl.DEPTH_TEST)
	//gl.DepthFunc(gl.LESS)

	//var triangle = primitives.NewTriangle()
	//triangle.Init()

	var rectangle = primitives.NewRectangle()
	rectangle.Init()

	for !window.ShouldClose() {

		//fmt.Println(triangle.GetShaderProgram())
		//fmt.Println(triangle.GetVertices())

		//Process user input
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		//triangle.Draw()
		rectangle.Draw()

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()

		if err := gl.GetError(); err != 0 {
			log.Fatal(err)
		}
	}
}

func processInput(window *glfw.Window) {
	var mouseCallback glfw.MouseButtonCallback = func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		if action == glfw.Release {
			fmt.Println("THIS A TEST")
			fmt.Println()
		}
	}

	var keyboardCallback glfw.KeyCallback = func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			switch key {
			case glfw.KeySpace:
				fmt.Println("Space was pressed")
			}
		}

	}
	window.SetMouseButtonCallback(mouseCallback)
	window.SetKeyCallback(keyboardCallback)
}

func frameBufferSizeCallback(window *glfw.Window, widht int, height int) {
	gl.Viewport(0, 0, int32(widht), int32(height))
}
