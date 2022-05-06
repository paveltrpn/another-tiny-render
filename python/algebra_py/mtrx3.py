
from array import array
from math import sin, cos

from typing import Callable

from .common import IdRw
from .vec3 import vec3

class mtrx3(object):
    __slots__ = ["data"]
    
    def __init__(self) -> None:
        self.data = array("f", [x-x for x in range(9)])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f} {:.4f}
{:.4f} {:.4f} {:.4f}
{:.4f} {:.4f} {:.4f}
""".format(self.data[0],  self.data[1],  self.data[2],  
           self.data[3],  self.data[4],  self.data[5],  
           self.data[6],  self.data[7],  self.data[8])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def __mul__(self, other):
        rt = mtrx3()

        for i in range(3):
            for j in range(3):
                for k in range(3):
                    rt[IdRw(i, j, 3)] = rt[IdRw(i, j, 3)]+ self.data[IdRw(k, j, 3)]*other.data[IdRw(i, k, 3)]

        return rt

    def tobytes(self):
        return self.data.tobytes()

def mtrx3SetIdtt() -> mtrx3:
    rt = mtrx3()

    onDiag: Callable[[int], float] = lambda elem: 1.0 if (elem % 4) == 0 else 0.0
    for i, _ in enumerate(rt.data):
        rt[i] = onDiag(i)

    return rt

def mtrx3SetIdttLoop() -> mtrx3:
    rt = mtrx3()

    for i in range(3):
        for j in range(3): 
            if i == j: 
                rt[IdRw(i, j, 3)] = 1.0
            else:
                rt[IdRw(i, j, 3)] = 0.0

    return rt

def mtrx3Mult(a: mtrx3, b: mtrx3) -> mtrx3:
    rt = mtrx3()

    for i in range(3):
        for j in range(3):
            for k in range(3):
                rt[IdRw(i, j, 3)] = rt[IdRw(i, j, 3)]+ a[IdRw(k, j, 3)]*b[IdRw(i, k, 3)]

    return rt

def mtrx3FromYaw(angl: float) -> mtrx3:
    sa: float
    ca: float
    rt = mtrx3()

    sa = sin(angl)
    ca = cos(angl)

    rt[0] = ca
    rt[1] = -sa
    rt[2] = 0.0
    rt[3] = sa  
    rt[4] = ca
    rt[5] = 0.0
    rt[6] = 0.0
    rt[7] = 0.0 
    rt[8] = 1.0

    return rt

def mtrx3FromPitch(angl: float) -> mtrx3:
    sa: float
    ca: float
    rt = mtrx3()

    sa = sin(angl)
    ca = cos(angl)

    rt[0] = 1.0 
    rt[1] = 0.0 
    rt[2] = 0.0
    rt[3] = 0.0 
    rt[4] = ca   
    rt[5] = -sa
    rt[6] = 0.0 
    rt[7] = sa   
    rt[8] = ca

    return rt

def mtrx3FromRoll(angl: float) -> mtrx3:
    sa: float
    ca: float
    rt = mtrx3()

    sa = sin(angl)
    ca = cos(angl)

    rt[0] = ca  
    rt[1] = 0.0 
    rt[2] = sa
    rt[3] = 0.0 
    rt[4] = 1.0
    rt[5] = 0.0
    rt[6] = -sa
    rt[7] = 0.0
    rt[8] = ca

    return rt

def mtrx3FromEuler(yaw: float, pitch: float, roll: float) -> mtrx3:
    cosy = cos(yaw)
    siny = sin(yaw)
    cosp = cos(pitch)
    sinp = sin(pitch)
    cosr = cos(roll)
    sinr = sin(roll)

    rt = mtrx3()

    rt[0] = cosy*cosr - siny*cosp*sinr
    rt[1] = -cosy*sinr - siny*cosp*cosr
    rt[2] = siny * sinp

    rt[3] = siny*cosr + cosy*cosp*sinr
    rt[4] = -siny*sinr + cosy*cosp*cosr
    rt[5] = -cosy * sinp

    rt[6] = sinp * sinr
    rt[7] = sinp * cosr
    rt[8] = cosp

    return rt

def mtrx3MultVec3(m: mtrx3, v: vec3) -> vec3:
	tmp: float = 0.0
	rt = vec3()

	for i in range(3): 
		tmp = 0
		for j in range(3):
			tmp = tmp + m[IdRw(i, j, 3)]*v[j]
		rt[i] = tmp

	return rt

def mtrx3GetTranspose(m: mtrx3) -> mtrx3:
    tmp: float
        
    rt = m

    for i in range(3):
        for j in range(3):
            tmp = rt[IdRw(i, j, 3)]
            rt[IdRw(i, j, 3)] = rt[IdRw(j, i, 3)]
            rt[IdRw(j, i, 3)] = tmp

    return rt
        