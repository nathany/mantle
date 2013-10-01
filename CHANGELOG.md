# Changelog

## r2013-06-25

* Add notice that I'm not continuing to work on this project at this time.

## r2012-08-22

* A rendering context to cache state and make unit testing easier.

## r2012-08-21

* Determined the 25 additional OpenGL functions needed to draw a triangle.
* Shaders are compiling.

## r2012-08-20

* Mantle is the name.
* Refactored into core/gl and gfx packages.
* Moved error handling to gfx wrapper via gl.SetAfterHook()

## r2012-08-19

* SetClearColor so now the window can be white instead of black.
* Planning on an idiomatic Go wrapper, rather than just a straight binding.

## r2012-08-18

* glClear buffer as my first binding. Now the window isn't garbled when resized!
* Error handling, by way of glGetError, with some helpers.
* Use runtime package to log the caller when reporting errors.
* Build tags to bypass the error checking code for release builds.

## r2012-08-17

* After some days of reviewing Go and OpenGL, I decided to work on a Go binding for OpenGL 3.2 Core Profile on OS X, whilst also learning 3D Graphics.
* Project setup using GLFW, thanks to [Jim Teeuwen's binding](http://go.pkgdoc.org/github.com/jteeuwen/glfw).
