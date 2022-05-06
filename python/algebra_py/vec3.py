
from array import array
from sys import float_info

from .common import *
from math import sqrt

class vec3(object):
    __slots__ = ["data"]
    
    def __init__(self, x: float=0.0, y: float=0.0, z: float=0.0) -> None:
        self.data = array("f", [x, y, z])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f} {:.4f}
""".format(self.data[0],  self.data[1],  self.data[2])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def tobytes(self):
        return self.data.tobytes()

def vec3Lenght(v: vec3) -> float:
	return sqrt(v[0]*v[0] +
		        v[1]*v[1] +
		        v[2]*v[2])

def vec3Normalize(v: vec3) -> vec3:
    len = vec3Lenght(v)

    rt = vec3()

    if len >= float_info.epsilon: 
        rt[0] = v[0] / len
        rt[1] = v[1] / len
        rt[2] = v[2] / len
	
    return rt

def vec3Scale(v: vec3, scale: float) -> vec3:
    rt = v

    rt[0] *= scale
    rt[1] *= scale
    rt[2] *= scale

    return rt

def vec3Invert(v: vec3) -> vec3:
    rt = vec3()

    rt[0] = -v[0]
    rt[1] = -v[1]
    rt[2] = -v[2]

    return rt

def vec3Dot(a: vec3, b: vec3) -> float:
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]

def vec3Sum(a: vec3, b: vec3) -> vec3:
    rt = vec3()

    rt[0] = a[0] + b[0]
    rt[1] = a[1] + b[1]
    rt[2] = a[2] + b[2]

    return rt

def vec3Sub(a: vec3, b: vec3) -> vec3:
    rt = vec3()

    rt[0] = a[0] - b[0]
    rt[1] = a[1] - b[1]
    rt[2] = a[2] - b[2]

    return rt

def vec3Cross(a: vec3, b: vec3) -> vec3:
    rt = vec3()

    rt[0] = a[1]*b[2] - a[2]*b[1]
    rt[1] = a[2]*b[0] - a[0]*b[2]
    rt[2] = a[0]*b[1] - a[1]*b[0]

    return rt
