#version 410 core

// Model matrix is identity 
uniform mat4 viewMtrx; 
uniform mat4 projMtrx;
uniform float intense;
uniform float value;

in vec3 position;     // Per-vertex position information we will pass in.
in vec3 color;        // Per-vertex color information we will pass in.
in vec3 normal;       // Per-vertex normal information we will pass in.
 
out vec3 v_Position;       // This will be passed into the fragment shader.
out vec4 v_Color;          // This will be passed into the fragment shader.
out vec3 v_Normal;         // This will be passed into the fragment shader.
out float l_intense;
out float l_value;

void main()
{
    mat4 u_MVPMatrix = projMtrx*viewMtrx;
    mat4 u_MVMatrix = viewMtrx;

    // Transform the vertex into eye space.
    v_Position = vec3(u_MVMatrix * vec4(position, 1.0));
 
    // Pass through the color.
    v_Color = vec4(color, 1.0);
 
    // Transform the normal's orientation into eye space.
    v_Normal = vec3(u_MVMatrix * vec4(normal, 0.0));
    
    l_intense = intense;
    l_value = value;
    
    // gl_Position is a special variable used to store the final position.
    // Multiply the vertex by the matrix to get the final point in normalized screen coordinates.
    gl_Position = u_MVPMatrix * vec4(position, 1.0);
}