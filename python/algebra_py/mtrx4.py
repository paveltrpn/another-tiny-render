
from array import array
from math import sin, cos, tanh, sqrt, fabs, hypot
from sys import float_info

from typing import Callable

from .common import IdRw
from .vec3 import vec3
from .vec4 import vec4

class mtrx4(object):
    __slots__ = ["data"]
    
    def __init__(self) -> None:
        self.data = array("f", [x-x for x in range(16)])

    def __repr__(self) -> str:
        return """
{:.4f} {:.4f} {:.4f} {:.4f}
{:.4f} {:.4f} {:.4f} {:.4f}
{:.4f} {:.4f} {:.4f} {:.4f}
{:.4f} {:.4f} {:.4f} {:.4f}
""".format(self.data[0],  self.data[1],  self.data[2],  self.data[3],
           self.data[4],  self.data[5],  self.data[6],  self.data[7],
           self.data[8],  self.data[9],  self.data[10], self.data[11],
           self.data[12], self.data[13], self.data[14], self.data[15])

    def __setitem__(self, key: int, value: float):
        self.data[key] = value

    def __getitem__(self, key: int) -> float:
        return self.data[key]

    def __mul__(self, other):
        rt = mtrx4()

        for i in range(4):
            for j in range(4):
                for k in range(4):
                    rt[IdRw(i, j, 4)] = rt[IdRw(i, j, 4)]+ self.data[IdRw(k, j, 4)]*other.data[IdRw(i, k, 4)]

        return rt

    def tobytes(self):
        return self.data.tobytes()

def mtrx4SetIdtt() -> mtrx4:
    rt = mtrx4()

    onDiag: Callable[[int], float] = lambda elem: 1.0 if (elem % 5) == 0 else 0.0
    for i, _ in enumerate(rt.data):
        rt[i] = onDiag(i)

    return rt

def mtrx4SetIdttLoop() -> mtrx4:
    rt = mtrx4()

    for i in range(4):
        for j in range(4): 
            if i == j: 
                rt[IdRw(i, j, 4)] = 1.0
            else:
                rt[IdRw(i, j, 4)] = 0.0

    return rt

def mtrx4Mult(a: mtrx4, b: mtrx4) -> mtrx4:
    rt = mtrx4()

    for i in range(4):
        for j in range(4):
            for k in range(4):
                rt[IdRw(i, j, 4)] = rt[IdRw(i, j, 4)]+ a[IdRw(k, j, 4)]*b[IdRw(i, k, 4)]

    return rt

def mtrx4FromEuler(yaw: float, pitch: float, roll: float) -> mtrx4:
    cosy = cos(yaw)
    siny = sin(yaw)
    cosp = cos(pitch)
    sinp = sin(pitch)
    cosr = cos(roll)
    sinr = sin(roll)

    rt = mtrx4()

    rt[0] = cosy*cosr - siny*cosp*sinr
    rt[1] = -cosy*sinr - siny*cosp*cosr
    rt[2] = siny * sinp
    rt[3] = 0.0

    rt[4] = siny*cosr + cosy*cosp*sinr
    rt[5] = -siny*sinr + cosy*cosp*cosr
    rt[6] = -cosy * sinp
    rt[7] = 0.0

    rt[8] = sinp * sinr
    rt[9] = sinp * cosr
    rt[10] = cosp
    rt[11] = 0.0

    rt[12] = 0.0
    rt[13] = 0.0
    rt[14] = 0.0
    rt[15] = 1.0

    return rt

def mtrx4FromAxisAngl(ax: vec3, phi: float) -> mtrx4:
    rt = mtrx4()

    cosphi = cos(phi)
    sinphi = sin(phi)
    vxvy = ax[0] * ax[1]
    vxvz = ax[0] * ax[2]
    vyvz = ax[1] * ax[2]
    vx = ax[0]
    vy = ax[1]
    vz = ax[2]

    rt[0]  = cosphi + (1.0-cosphi)*vx*vx
    rt[1]  = (1.0-cosphi)*vxvy - sinphi*vz
    rt[2]  = (1.0-cosphi)*vxvz + sinphi*vy
    rt[3]  = 0.0

    rt[4]  = (1.0-cosphi)*vxvy + sinphi*vz
    rt[5]  = cosphi + (1.0-cosphi)*vy*vy
    rt[6]  = (1.0-cosphi)*vyvz - sinphi*vx
    rt[7]  = 0.0

    rt[8]  = (1.0-cosphi)*vxvz - sinphi*vy
    rt[9]  = (1.0-cosphi)*vyvz + sinphi*vx
    rt[10] = cosphi + (1.0-cosphi)*vz*vz
    rt[11] = 0.0

    rt[12] = 0.0
    rt[13] = 0.0
    rt[14] = 0.0
    rt[15] = 1.0

    return rt

def mtrx4FromOffset(src: vec3) -> mtrx4:
    rt = mtrx4SetIdtt()

    rt[3]  = src[0]
    rt[7]  = src[1]
    rt[11] = src[2]

    return rt

def mtrx4FromScale(src: vec3) -> mtrx4:
    rt = mtrx4SetIdtt()

    rt[0]  = src[0]
    rt[5]  = src[1]
    rt[10] = src[2]

    return rt

def mtrx4FromPerspective(fovy: float, aspect: float, near: float, far: float) -> mtrx4:
    rt = mtrx4SetIdtt()
    f = 1.0 / tanh(fovy / sqrt(2))
    nf: float
	
    rt[0] = f / aspect
    rt[1] = 0.0
    rt[2] = 0.0
    rt[3] = 0.0
    rt[4] = 0.0
    rt[5] = f
    rt[6] = 0.0
    rt[7] = 0.0
    rt[8] = 0.0
    rt[9] = 0.0
    rt[11] = -1.0
    rt[12] =  0.0
    rt[13] =  0.0
    rt[15] =  0.0
	
    if (far >= float_info.epsilon):
        nf = 1.0 / (near - far)
        rt[10] = (far + near) * nf
        rt[14] = 2.0 * far * near * nf
    else:
        rt[10] = -1.0
        rt[14] = -2.0 * near

    return rt

def mtrx4FromLookAt(eye: vec3, center: vec3, up: vec3) -> mtrx4:
    out: mtrx4 = mtrx4SetIdtt()

    if (fabs(eye[0]-center[0]) < float_info.epsilon and 
        fabs(eye[1]-center[1]) < float_info.epsilon and 
        fabs(eye[2]-center[2]) < float_info.epsilon):
        return out

    z0 = eye[0] - center[0]
    z1 = eye[1] - center[1]
    z2 = eye[2] - center[2]

    len = 1.0 / hypot(z0, z1, z2); #??? было просто hypot
    z0 *= len
    z1 *= len
    z2 *= len

    x0 = up[1]*z2 - up[2]*z1
    x1 = up[2]*z0 - up[0]*z2
    x2 = up[0]*z1 - up[1]*z0
    len = hypot(x0, x1, x2)
    if (len == 0.0):
        x0 = 0
        x1 = 0
        x2 = 0
    else:
        len = 1.0 / len
        x0 *= len
        x1 *= len
        x2 *= len

    y0 = z1*x2 - z2*x1
    y1 = z2*x0 - z0*x2
    y2 = z0*x1 - z1*x0

    len = hypot(y0, y1, y2)
    if (len == 0.0):
        y0 = 0
        y1 = 0
        y2 = 0
    else:
        len = 1.0 / len
        y0 *= len
        y1 *= len
        y2 *= len

    out[0] = x0
    out[1] = y0
    out[2] = z0
    out[3] = 0.0
    out[4] = x1
    out[5] = y1
    out[6] = z1
    out[7] = 0.0
    out[8] = x2
    out[9] = y2
    out[10] = z2
    out[11] = 0.0
    out[12] = -(x0*eye[0] + x1*eye[1] + x2*eye[2])
    out[13] = -(y0*eye[0] + y1*eye[1] + y2*eye[2])
    out[14] = -(z0*eye[0] + z1*eye[1] + z2*eye[2])
    out[15] = 1.0

    return out

def mtrx4MultVec4(m: mtrx4, v: vec4) -> vec4:
	tmp: float = 0.0
	rt = vec4()

	for i in range(4): 
		tmp = 0
		for j in range(4):
			tmp = tmp + m[IdRw(i, j, 4)]*v[j]
		rt[i] = tmp

	return rt

def mtrx4GetTranspose(m: mtrx4) -> mtrx4:
    tmp: float
        
    rt = m

    for i in range(4):
        for j in range(4):
            tmp = rt[IdRw(i, j, 4)]
            rt[IdRw(i, j, 4)] = rt[IdRw(j, i, 4)]
            rt[IdRw(j, i, 4)] = tmp

    return rt
