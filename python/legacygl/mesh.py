
import glm as glm

#  По три нормали на вершину!!!
boxTris = glm.array(glm.vec3(1.0, 1.0, 1.0),  glm.vec3(-1.0, 1.0, 1.0),  glm.vec3(-1.0, -1.0, 1.0),
					  glm.vec3( 1.0,  1.0, 1.0), glm.vec3(-1.0, -1.0, 1.0), glm.vec3( 1.0, -1.0, 1.0),
					  glm.vec3( 1.0, -1.0, -1.0), glm.vec3( 1.0, -1.0,  1.0), glm.vec3(-1.0, -1.0,  1.0),
					  glm.vec3( 1.0, -1.0, -1.0), glm.vec3(-1.0, -1.0,  1.0), glm.vec3(-1.0, -1.0, -1.0),
					  glm.vec3(-1.0, -1.0,-1.0), glm.vec3(-1.0, -1.0, 1.0), glm.vec3(-1.0,  1.0, 1.0),
					  glm.vec3(-1.0, -1.0, -1.0), glm.vec3(-1.0,  1.0,  1.0), glm.vec3(-1.0,  1.0, -1.0),
					  glm.vec3(-1.0, 1.0, -1.0), glm.vec3( 1.0, 1.0, -1.0), glm.vec3( 1.0,-1.0, -1.0),
					  glm.vec3(-1.0,  1.0, -1.0), glm.vec3( 1.0, -1.0, -1.0), glm.vec3(-1.0, -1.0, -1.0),
					  glm.vec3(1.0, 1.0,-1.0), glm.vec3(1.0, 1.0, 1.0), glm.vec3(1.0,-1.0, 1.0),
					  glm.vec3(1.0,  1.0, -1.0), glm.vec3(1.0, -1.0,  1.0), glm.vec3(1.0, -1.0, -1.0),
					  glm.vec3(-1.0, 1.0,-1.0), glm.vec3(-1.0, 1.0, 1.0), glm.vec3( 1.0, 1.0, 1.0),
					  glm.vec3(-1.0, 1.0, -1.0), glm.vec3( 1.0, 1.0,  1.0), glm.vec3( 1.0, 1.0, -1.0))

boxNrmls = 	glm.array(glm.vec3(0.0, 0.0, 1.0), glm.vec3(0.0, 0.0, 1.0),glm.vec3(0.0, 0.0, 1.0),
						  glm.vec3(0.0, 0.0, 1.0),glm.vec3(0.0, 0.0, 1.0),glm.vec3(0.0, 0.0, 1.0),
						  glm.vec3(0.0,-1.0, 0.0),glm.vec3(0.0,-1.0, 0.0),glm.vec3(0.0,-1.0, 0.0),
						  glm.vec3(0.0,-1.0, 0.0),glm.vec3(0.0,-1.0, 0.0),glm.vec3(0.0,-1.0, 0.0),
						  glm.vec3(-1.0,0.0, 0.0),glm.vec3(-1.0,0.0, 0.0),glm.vec3(-1.0,0.0, 0.0),
						  glm.vec3(-1.0,0.0, 0.0),glm.vec3(-1.0,0.0, 0.0),glm.vec3(-1.0,0.0, 0.0),
						  glm.vec3(0.0, 0.0,-1.0),glm.vec3(0.0, 0.0,-1.0),glm.vec3(0.0, 0.0,-1.0),
						  glm.vec3(0.0, 0.0,-1.0),glm.vec3(0.0, 0.0,-1.0),glm.vec3(0.0, 0.0,-1.0),
						  glm.vec3(1.0, 0.0, 0.0),glm.vec3(1.0, 0.0, 0.0),glm.vec3(1.0, 0.0, 0.0),
						  glm.vec3(1.0, 0.0, 0.0),glm.vec3(1.0, 0.0, 0.0),glm.vec3(1.0, 0.0, 0.0),
						  glm.vec3(0.0, 1.0, 0.0),glm.vec3(0.0, 1.0, 0.0),glm.vec3(0.0, 1.0, 0.0),
						  glm.vec3(0.0, 1.0, 0.0),glm.vec3(0.0, 1.0, 0.0),glm.vec3(0.0, 1.0, 0.0))

boxTexCoords = glm.array(glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0),
						   glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0),
						   glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0),
						   glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0),
						   glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0),
						   glm.vec2(0.0, 0.0), glm.vec2(0.0, 1.0), glm.vec2(1.0, 1.0),
						   glm.vec2(0.0, 0.0), glm.vec2(1.0, 1.0), glm.vec2(1.0, 0.0))

class CMeshList:
    def __init__(self):
        self.meshList = glm.vec3()