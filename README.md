# Mantle (WIP)

Mantle is an idiomatic [Go](http://golang.org/) wrapper for OpenGL 3.2 Core Profile on OS X.

|:**Notice**|
|----------|
|This was a summer experiment to learn OpenGL. Please see [Go-GL](https://github.com/go-gl) for all your OpenGL Go needs.|

As I've been learning OpenGL, I've decided that I want this to be a wrapper that deviates enough to be idomatic in Go, and a little nicer to use.

At the same time, I'd like this to be moderately performant, so it may not always be practical to follow "the Go Way" (eg. checking for errors after every call would get expensive).

This is **Work in Progress**.

Alternative, more complete bindings:

* [GoGL](https://github.com/chsc/gogl) generated bindings
* [banthar's gl](https://github.com/banthar/gl) is hand-written

mantle is made available under the [Simplified BSD License](http://opensource.org/licenses/bsd-license.php/).

## Target

As my goal is to learn OpenGL, GLSL and 3D graphics, this package won't concern itself with cross-platform support for the foreseeable future.

* Mac OS X 10.8 (though it should work on 10.7 as well)
* OpenGL 3.2 and GLSL 1.5
* Only supporting the Core Profile
* 64-bit only binaries
* Testing with Intel HD Graphics 4000 or Radeon HD 4870

Thankfully OS X appears to have a [fairly consistient story](https://developer.apple.com/graphicsimaging/opengl/capabilities/) for OpenGL.


## Dependencies

[GLFW](http://www.glfw.org/) is being used to setup a window and context, and for keyboard/mouse/joystick input. This makes use of Jim Teeuwen's [GLFW binding](http://go.pkgdoc.org/github.com/jteeuwen/glfw). The necessary static/dynamic libraries can be installed on OS X with [Homebrew](http://mxcl.github.com/homebrew/).

Based on the [Roadmap](http://wiki.glfw.org/wiki/Roadmap_for_GLFW_3), GLFW seems well-suited for an external Go binding. It doesn't try to do too much, but nicely wraps the Objective-C calls in a C API.

While OpenGL on Windows requires extensions to access a modern version of OpenGL, I don't expect I will need [GLEW](http://glew.sourceforge.net/) anytime soon on OS X (though it is also available via Homebrew).

For testing, I'll be making use of [gocheck](http://labix.org/gocheck) ([docs](http://go.pkgdoc.org/launchpad.net/gocheck)).

## Resources

Books:

* [OpenGL SuperBible](http://www.starstonesoftware.com/OpenGL/), 5th ed. covers OpenGL 3.3 ([GLTools](http://code.google.com/p/oglsuperbible5/))
* [OpenGL Shading Language](http://www.amazon.com/OpenGL-Shading-Language-Edition-ebook/dp/B002HMJYC4/), 3rd ed. covers GLSL 1.4
* [3D Math Primer for Graphics and Game Development](http://www.amazon.com/Graphics-Development-Wordware-Library-ebook/dp/B0026A6CJ0/)
* [Texturing and Modeling](http://www.amazon.com/Texturing-Modeling-Third-Edition-Procedural/dp/1558608486/), 3rd ed. for more shader ideas.

Online:

* [Learning Modern 3D Graphics Programming](http://www.arcsynthesis.org/gltut/)
* [OpenGL 3.3 Reference](http://www.opengl.org/sdk/docs/man3/)
* [OpenGL for Mac OS X](https://developer.apple.com/devcenter/mac/resources/opengl/)
* [cgo](http://golang.org/doc/articles/c_go_cgo.html) documentation for Go bindings

Also of interest:

* [crunch](http://code.google.com/p/crunch/), Advanced DXT texture compression and real-time transcoding library (C++)

