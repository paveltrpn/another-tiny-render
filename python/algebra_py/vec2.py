

from array import array

from .common import *
from math import sqrt

class vec2(object):
    __slots__ = ["data"]

    def __init__(self, x: float=0.0, y: float=0.0) -> None:
        self.data = array("f", [x, y])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f}
""".format(self.data[0],  self.data[1])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def tobytes(self):
        return self.data.tobytes()
        
def vec2Lenght(v: vec2) -> float:
	return sqrt(v[0]*v[0] +
		        v[1]*v[1])

def vec2Normalize(v: vec2) -> vec2:
    len = vec2Lenght(v)

    rt = vec2()

    if len != 0.0: 
        rt[0] = v[0] / len
        rt[1] = v[1] / len
	
    return rt

def vec2Scale(v: vec2, scale: float) -> vec2:
    rt = v

    rt[0] *= scale
    rt[1] *= scale

    return rt

def vec2Invert(v: vec2) -> vec2:
    rt = vec2()

    rt[0] = -v[0]
    rt[1] = -v[1]

    return rt

def vec2Dot(a: vec2, b: vec2) -> float:
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]

def vec2Sum(a: vec2, b: vec2) -> vec2:
    rt = vec2()

    rt[0] = a[0] + b[0]
    rt[1] = a[1] + b[1]

    return rt

def vec2Sub(a: vec2, b: vec2) -> vec2:
    rt = vec2()

    rt[0] = a[0] - b[0]
    rt[1] = a[1] - b[1]

    return rt
