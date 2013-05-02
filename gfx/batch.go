package gfx

import (
    "github.com/nathany/mantle/core/gl"
)

// Start populating the array
// void Begin(GLenum primitive, GLuint nVerts, GLuint nTextureUnits = 0);
// Begin(primitiveType=GL_TRIANGLES, nNumVerts=3)
// glGenVertexArrays(1, &vertexArrayObject);
//  glBindVertexArray(vertexArrayObject);

// CopyVertexData3f(M3DVector3f *vVerts)
// glGenBuffers(1, &uiVertexArray);
// glBindBuffer(GL_ARRAY_BUFFER, uiVertexArray);
// glBufferData(GL_ARRAY_BUFFER, sizeof(GLfloat) * 3 * nNumVerts, vVerts, GL_DYNAMIC_DRAW);
func CopyVertexData3f(rc *gl.Context, vertices []float32) {
    // begin
    vertexArrayObject := rc.GenVertexArrays(1)
    vertexArrayObject[0].Bind()

    // copy
    uiVertexArray := rc.GenBuffers(1)
    uiVertexArray[0].Bind(gl.ArrayBuffer)
    gl.ArrayBuffer.SetBufferData(vertices, gl.DynamicDraw)

    // end
    vertexArrayObject[0].Bind()
    rc.EnableVertexAttribArray(0) // vVertex
    uiVertexArray[0].Bind(gl.ArrayBuffer)
    rc.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, 0)
}

// void GLBatch::Draw(void)
// glBindVertexArray(vertexArrayObject);
// glDrawArrays(primitiveType, 0, nNumVerts);
// glBindVertexArray(0);

// Tell the batch you are done
// void End(void);
// glEnableVertexAttribArray(GLT_ATTRIBUTE_VERTEX);
// glBindBuffer(GL_ARRAY_BUFFER, uiVertexArray);
// glVertexAttribPointer(GLT_ATTRIBUTE_VERTEX, 3, GL_FLOAT, GL_FALSE, 0, 0);

// glBindBuffer(GL_ARRAY_BUFFER, uiVertexArray);
// glUnmapBuffer(GL_ARRAY_BUFFER);
