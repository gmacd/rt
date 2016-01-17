package support

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type GlRenderer struct {
	width, height int

	window  *glfw.Window
	program uint32
	vao     uint32
	texture uint32

	framesToRender chan *Frame
	freeFrames     chan *Frame
}

type Frame struct {
	ShouldStop    bool // 2-way message - checked on both ends
	Pixels        []float32
	Width, Height int
}

func NewGlRenderer(width, height int) *GlRenderer {
	glr := &GlRenderer{
		width, height, nil, 0, 0, 0,
		make(chan *Frame, 1),
		make(chan *Frame, 1)}

	return glr
}

func (glr *GlRenderer) Start() {
	go func() {
		glr.initSystem()
		glr.initScreen()

		// Prepare 2 frames to double buffer
		for i := 0; i < 2; i++ {
			glr.freeFrames <- &Frame{
				glr.window.ShouldClose(),
				make([]float32, 4*glr.width*glr.height),
				glr.width, glr.height}
		}

		// Loop until stop is requested
		for nextFrame := range glr.framesToRender {
			if nextFrame.ShouldStop {
				break
			}

			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			gl.UseProgram(glr.program)
			gl.BindVertexArray(glr.vao)

			updateScreenTexture(glr.texture, nextFrame.Pixels, glr.width, glr.height)
			gl.DrawArrays(gl.TRIANGLES, 0, 6)

			glr.window.SwapBuffers()
			glfw.PollEvents()

			// Consider frame rendered, so push back to be reused
			nextFrame.ShouldStop = glr.window.ShouldClose()
			glr.freeFrames <- nextFrame
		}
	}()
}

func (glr *GlRenderer) NextFrameChan() chan *Frame {
	return glr.freeFrames
}

func (glr *GlRenderer) Render(frame *Frame) {
	glr.framesToRender <- frame
}

func (glr *GlRenderer) Window() *glfw.Window { return glr.window }

func (glr *GlRenderer) initSystem() {
	// TODO Call once goroutine created
	runtime.LockOSThread()

	var err error
	if err = glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glr.window, err = glfw.CreateWindow(glr.width, glr.height, "rt", nil, nil)
	if err != nil {
		panic(err)
	}
	glr.window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func (glr *GlRenderer) initScreen() {
	var err error
	glr.program, err = createScreenShader()
	if err != nil {
		panic(err)
	}

	gl.UseProgram(glr.program)

	projection := mgl32.Ortho2D(-1, 1, 1, -1)
	projectionUniform := gl.GetUniformLocation(glr.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	textureUniform := gl.GetUniformLocation(glr.program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	glr.texture = createScreenTexture(glr.width, glr.height)

	gl.GenVertexArrays(1, &glr.vao)
	gl.BindVertexArray(glr.vao)

	var quadVertices = []float32{
		//  X, Y, Z, U, V
		1.0, -1.0, 0.0, 1.0, 0.0,
		-1.0, -1.0, 0.0, 0.0, 0.0,
		1.0, 1.0, 0.0, 1.0, 1.0,
		-1.0, -1.0, 0.0, 0.0, 0.0,
		-1.0, 1.0, 0.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0, 1.0,
	}

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadVertices)*4, gl.Ptr(quadVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(glr.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(glr.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
}

func createScreenShader() (uint32, error) {
	var vertexShader = `
#version 330

uniform mat4 projection;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * vec4(vert, 1);
}
` + "\x00"

	var fragmentShader = `
#version 330

uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

	return newProgram(vertexShader, fragmentShader)
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	if e := gl.GetError(); e != gl.NO_ERROR {
		panic(fmt.Sprintf("before resizing: %d", e))
	}
	gl.CompileShader(shader)

	if e := gl.GetError(); e != gl.NO_ERROR {
		panic(fmt.Sprintf("before resizing: %d", e))
	}
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func createScreenTexture(width, height int) uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA32F,
		int32(width),
		int32(height),
		0,
		gl.RGBA,
		gl.FLOAT,
		nil)

	return texture
}

func updateScreenTexture(texture uint32, pixels []float32, width, height int) {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexSubImage2D(
		gl.TEXTURE_2D,
		0,
		0,
		0,
		int32(width),
		int32(height),
		gl.RGBA,
		gl.FLOAT,
		gl.Ptr(pixels))
}
