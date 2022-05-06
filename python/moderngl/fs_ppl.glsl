#version 410 core

// uniform vec3 u_LightPos;       // The position of the light in eye space.
 
in vec3 v_Position;       // Interpolated position for this fragment.
in vec4 v_Color;          // This is the color from the vertex shader interpolated across the
                          // triangle per fragment.
in vec3 v_Normal;         // Interpolated normal for this fragment.

in float l_intense;
in float l_value;

void main()
{
    vec3 u_LightPos = vec3(0.0, 1.5, 0.0);

    // Will be used for attenuation.
    float distance = length(u_LightPos - v_Position);
 
    // Get a lighting direction vector from the light to the vertex.
    vec3 lightVector = normalize(u_LightPos - v_Position);
 
    // Calculate the dot product of the light vector and vertex normal. If the normal and light vector are
    // pointing in the same direction then it will get max illumination.
    float diffuse = max(dot(v_Normal, lightVector), 0.1);
 
    // Add attenuation.
    diffuse = (diffuse * l_intense) * (1.0 / (1.0 + (l_value * distance * distance)));
 
    // Multiply the color by the diffuse illumination level to get final output color.
    gl_FragColor = v_Color * diffuse;
}