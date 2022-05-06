
from array import array
from math import sin, cos
from typing import Any, Callable

from .common import IdRw

class mtrx2(object):
    __slots__ = ["data"]
    
    def __init__(self) -> None:
        self.data = array("f", [x-x for x in range(4)])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f}
{:.4f} {:.4f}
""".format(self.data[0],  self.data[1],  
           self.data[2],  self.data[3])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def __mul__(self, other):
        rt = mtrx2()

        for i in range(2):
            for j in range(2):
                for k in range(2):
                    rt[IdRw(i, j, 2)] = rt[IdRw(i, j, 2)] + self.data[IdRw(k, j, 2)]*other.data[IdRw(i, k, 2)]

        return rt

    def tobytes(self):
        return self.data.tobytes()

def mtrx2SetIdtt() -> mtrx2:
    rt = mtrx2()

    onDiag: Callable[[int], float] = lambda elem: 1.0 if (elem % 4) == 0 else 0.0
    for i, _ in enumerate(rt.data):
        rt[i] = onDiag(i)

    return rt

def mtrx2SetIdttLoop() -> mtrx2:
    rt = mtrx2()

    for i in range(3):
        for j in range(3): 
            if i == j: 
                rt[IdRw(i, j, 3)] = 1.0
            else:
                rt[IdRw(i, j, 3)] = 0.0

    return rt

def mtrx2Rtn(phi: float) -> mtrx2:
    rt = mtrx2()

    cosphi: float
    sinphi: float
	
    sinphi = sin(phi)
    cosphi = cos(phi)

    rt[0] = cosphi
    rt[1] = -sinphi
    rt[2] = -sinphi
    rt[3] = cosphi

    return rt

def mtrx2Det(m: mtrx2) -> float:
    return m[0]*m[3] - m[1]*m[2]

