
from array import array

from .common import IdRw
from .vec3 import vec3
from .vec4 import vec4
from .mtrx4 import mtrx4
from .mtrx3 import mtrx3

class CVec4Arrayf(object):
    __slots__ = ["data", "capacity", "dataSize"]
    
    def __init__(self, arr: list) -> None:
        """
        :type arr: list of float elemets rate by 4
        :type self.capacity: float - number of individual float coordinates. Mast be in rate of 4.
        :type self.dataSize: int - sizeof(self.data).
        """
        if not (len(arr)%4 == 0):
            print("CVec4Arrayf(): WARNING! Initial array lenght is not in rate of 4!")

        self.data = array("f", [elem for elem in arr])
        self.capacity = len(arr)
        self.dataSize = len(arr)*4 # size of array.array("f", [...]) element is 4 bytes 

    def __setitem__(self, i: int, value: vec4):
        self.data[i+0] = value[0]
        self.data[i+1] = value[1]
        self.data[i+2] = value[2]
        self.data[i+3] = value[3]

    def __getitem__(self, i: int) -> vec4:
        return vec4(self.data[i+0], self.data[i+1], self.data[i+2], self.data[i+3])

    def applyVec3(self, v: vec3) -> None:
        for arId in range(0, self.capacity, 4):
            self.data[arId+0] = self.data[arId+0] + v[0]
            self.data[arId+1] = self.data[arId+1] + v[1]
            self.data[arId+2] = self.data[arId+2] + v[2]

    def applyVec4(self, v: vec4) -> None:
        for arId in range(0, self.capacity, 4):
            self.data[arId+0] = self.data[arId+0] + v[0]
            self.data[arId+1] = self.data[arId+1] + v[1]
            self.data[arId+2] = self.data[arId+2] + v[2]
            self.data[arId+3] = self.data[arId+3] + v[3]

    def applyMtrx4Naive(self, A: mtrx4) -> None:
        """Perform multiplication of internal array of vector4 by a matrix4x4.

        :type A: matrix 4 by 4
        """
        for arId in range(0, self.capacity, 4):
            for i in range(4): 
                tmp = 0
                for j in range(4):
                    tmp = tmp + A[IdRw(i, j, 4)]*self.data[arId+j]
                self.data[arId+i] = tmp

    def applyMtrx4(self, A: mtrx4) -> None:
        """Perform multiplication of internal array of vector3 by a matrix4x4.

        :type A: matrix 4 by 4
        """
        for arId in range(0, self.capacity, 4):
            x = self.data[arId+0]
            y = self.data[arId+1]
            z = self.data[arId+2]
            w = self.data[arId+3]

            self.data[arId+0] = A[0] * x + A[4] * y + A[8] * z + A[12] * w
            self.data[arId+1] = A[1] * x + A[5] * y + A[9] * z + A[13] * w
            self.data[arId+2] = A[2] * x + A[6] * y + A[10] * z + A[14] * w
            self.data[arId+3] = A[3] * x + A[7] * y + A[11] * z + A[15] * w

    def tobytes(self):
        return self.data.tobytes()
