# History

## August 2012

### 20.

* Mantle is the name.

### 19.

* SetClearColor so now the window can be white instead of black.
* Planning on an idiomatic Go wrapper, rather than just a straight binding.

### 18.

* glClear buffer as my first binding. Now the window isn't garbled when resized!
* Error handling, by way of glGetError, with some helpers.
* Use runtime package to log the caller when reporting errors.
* Build tags to bypass the error checking code for release builds.

### 17.

* After some days of reviewing Go and OpenGL, I decided to work on a Go binding for OpenGL 3.2 Core Profile on OS X, whilst also learning 3D Graphics.
* Project setup using GLFW, thanks to [Jim Teeuwen's binding](http://go.pkgdoc.org/github.com/jteeuwen/glfw).
