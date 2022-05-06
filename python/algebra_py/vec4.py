
from array import array
from sys import float_info

from .common import *
from math import sqrt

class vec4(object):
    __slots__ = ["data"]
    
    def __init__(self, x: float=0.0, y: float=0.0, z: float=0.0, w: float=0.0) -> None:
        self.data = array("f", [x, y, z, w])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f} {:.4f} {:.4f}
""".format(self.data[0],  self.data[1],  self.data[2],  self.data[3])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def tobytes(self):
        return self.data.tobytes()
        
def vec4Lenght(v: vec4) -> float:
	return sqrt(v[0]*v[0] +
		        v[1]*v[1] +
		        v[2]*v[2] +
		        v[3]*v[3])

def vec4Normalize(v: vec4) -> vec4:
    len = vec4Lenght(v)

    rt = v

    if len >= float_info.epsilon: 
        rt[0] = v[0] / len
        rt[1] = v[1] / len
        rt[2] = v[2] / len
        rt[3] = v[3] / len
	
    return rt

def vec4Scale(v: vec4, scale: float) -> vec4:
    rt = v

    rt[0] *= scale
    rt[1] *= scale
    rt[2] *= scale
    rt[3] *= scale

    return rt

def vec4Invert(v: vec4) -> vec4:
    rt = vec4()

    rt[0] = -v[0]
    rt[1] = -v[1]
    rt[2] = -v[2]
    rt[3] = -v[3]

    return rt

def vec4Dot(a: vec4, b: vec4) -> float:
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]

def Vec4Sum(a: vec4, b: vec4) -> vec4:
    rt = vec4()

    rt[0] = a[0] + b[0]
    rt[1] = a[1] + b[1]
    rt[2] = a[2] + b[2]
    rt[3] = a[3] + b[3]

    return rt

def vec4Sub(a: vec4, b: vec4) -> vec4:
    rt = vec4()

    rt[0] = a[0] - b[0]
    rt[1] = a[1] - b[1]
    rt[2] = a[2] - b[2]
    rt[3] = a[3] - b[3]

    return rt
