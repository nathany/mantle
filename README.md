# GL3: OpenGL Core Profile binding (WIP)

The plan is to learn OpenGL, GLSL and 3D graphics, while also writing an OpenGL 3.2 Core Profile wrapper/binding in [Go](http://golang.org/).

I intend to tackle OpenGL functions as needed, so that I actually understand what they do, and where they are useful. By [documenting my code](http://golang.org/doc/articles/godoc_documenting_go_code.html), this will hopefully prove helpful for others and my future self.

Alternatives: [GoGL](https://github.com/chsc/gogl) generates bindings, [banthar's gl](https://github.com/banthar/gl) is hand-written.

## Target

As my goal is learning 3D Graphics, I'm not going to worry about cross-platform support for the foreseeable future.

* Mac OS X 10.8 (though it should work on 10.7 as well)
* OpenGL 3.2 and GLSL 1.5
* Only supporting the Core Profile (gl3.h).
* 64-bit only binaries
* Testing with Intel HD Graphics 4000 or Radeon HD 4870

Thankfully OS X appears to have a [fairly consistient story](https://developer.apple.com/graphicsimaging/opengl/capabilities/) for OpenGL.

## Dependencies

[GLFW](http://www.glfw.org/) is being used to setup a window and context, and for keyboard/mouse/joystick input. This makes use of Jim Teeuwen's [GLFW binding](http://go.pkgdoc.org/github.com/jteeuwen/glfw). The necessary static/dynamic libraries can be installed on OS X with [Homebrew](http://mxcl.github.com/homebrew/).

Based on the [Roadmap](http://wiki.glfw.org/wiki/Roadmap_for_GLFW_3), GLFW seems well-suited for an external Go binding. It doesn't try to do too much, but nicely wraps the Objective-C calls in a C API.

While OpenGL on Windows requires extensions to access a modern version of OpenGL, I don't expect I will need [GLEW](http://glew.sourceforge.net/) anytime soon on OS X (though it is also available via Homebrew).

## Resources

Books:

* [OpenGL SuperBible](http://www.starstonesoftware.com/OpenGL/), 5th ed. covers OpenGL 3.3 ([GLTools](http://code.google.com/p/oglsuperbible5/))
* [OpenGL Shading Language](http://www.amazon.com/OpenGL-Shading-Language-Edition-ebook/dp/B002HMJYC4/), 3rd ed. covers GLSL 1.4
* [3D Math Primer for Graphics and Game Development](http://www.amazon.com/Graphics-Development-Wordware-Library-ebook/dp/B0026A6CJ0/)
* [Texturing and Modeling](http://www.amazon.com/Texturing-Modeling-Third-Edition-Procedural/dp/1558608486/), 3rd ed. for more shader ideas.

Online:

* [OpenGL 3.3 Reference](http://www.opengl.org/sdk/docs/man3/)
* [cgo](http://golang.org/doc/articles/c_go_cgo.html) documentation for Go bindings

Also of interest:

* [crunch](http://code.google.com/p/crunch/), Advanced DXT texture compression and real-time transcoding library (C++)
