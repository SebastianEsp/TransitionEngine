#version 460 core
layout(location = 0) out vec4 FragColor;

layout(location = 1) uniform vec4 ourColor;

void main(){
    FragColor = ourColor;
};
