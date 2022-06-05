#version 410 core

in vec3 position;
in vec3 normal;
in vec3 color;

uniform mat4 viewMtrx;
uniform mat4 projMtrx;

out vec3 newColor;
out vec3 vLighting;

void main()
{
    vec3 ambientLight = vec3(0.3, 0.3, 0.3);
    vec3 directionalLightColor = vec3(1, 1, 1);
    vec3 directionalVector = normalize(vec3(0, 1, -1));

    float directional = max(dot(normal.xyz, directionalVector), 0.0);
    vLighting = ambientLight + (directionalLightColor * directional);

    gl_Position = projMtrx * viewMtrx * vec4(position, 1.0);

    newColor = color;
}