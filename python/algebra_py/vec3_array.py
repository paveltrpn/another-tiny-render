from array import array

from .common import *
from .vec3 import vec3
from .mtrx4 import mtrx4
from .mtrx3 import mtrx3

from sys import float_info

class CVec3Arrayf(object):
    __slots__ = ["data", "capacity", "dataSize"]
    
    def __init__(self, arr: list) -> None:
        """
        :type arr: list of float elemets rate by 3
        :type self.capacity: int - number of individual float coordinates. Mast be in rate of 3.
        :type self.dataSize: int - sizeof(self.data).
        """
        if not (len(arr)%3 == 0):
            print("CVec3Arrayf(): WARNING! Initial array lenght is not in rate of 3!")

        self.data = array("f", [elem for elem in arr])
        self.capacity = len(arr)
        self.dataSize = len(arr)*4 # size of array.array("f", [...]) element is 4 bytes 

    def __setitem__(self, i: int, value: vec3):
        self.data[i+0] = value[0]
        self.data[i+1] = value[1]
        self.data[i+2] = value[2]

    def __getitem__(self, i: int) -> vec3:
        return vec3(self.data[i+0], self.data[i+1], self.data[i+2])

    def tobytes(self):
        return self.data.tobytes()

    def applyVec3(self, v: vec3) -> None:
        for arId in range(0, self.capacity, 3):
            self.data[arId+0] = self.data[arId+0] + v[0]
            self.data[arId+1] = self.data[arId+1] + v[1]
            self.data[arId+2] = self.data[arId+2] + v[2]

    def applyMtrx3Naive(self, A: mtrx3) -> None:
        """Perform multiplication of internal array of vector3 by a matrix3x3.

        :type A: matrix 3 by 3
        """
        for arId in range(0, self.capacity, 3):
            for i in range(3): 
                tmp = 0
                for j in range(3):
                    tmp = tmp + A[IdRw(i, j, 3)]*self.data[arId+j]
                self.data[arId+i] = tmp

    def applyMtrx4Naive(self, A: mtrx4) -> None:
        """Perform multiplication of internal array of vector3 by a matrix4x4.

        :type A: matrix 4 by 4
        """
        vec = array("f", [0.0, 0.0, 0.0, 0.0])
        rt = array("f", [0.0, 0.0, 0.0, 0.0])

        for arId in range(0, self.capacity, 3):
            vec[0] = self.data[arId+0]
            vec[1] = self.data[arId+1]
            vec[2] = self.data[arId+2]
            vec[3] = 1.0

            # TODO - развернуть цикл и избавиться от временных переменных
            for i in range(4): 
                tmp = 0
                for j in range(4):
                    tmp = tmp + A[IdRw(i, j, 4)]*vec[j]
                rt[i] = tmp
        
            self.data[arId+0] = rt[0]
            self.data[arId+1] = rt[1]
            self.data[arId+2] = rt[2]

    
    def applyMtrx3(self, A: mtrx3) -> None:
        for arId in range(0, self.capacity, 3):
            x = self.data[arId+0]
            y = self.data[arId+1]
            z = self.data[arId+2]
            
            self.data[arId+0] = x * A[0] + y * A[3] + z * A[6]
            self.data[arId+1] = x * A[1] + y * A[4] + z * A[7]
            self.data[arId+2] = x * A[2] + y * A[5] + z * A[8]

    def applyMtrx4(self, A: mtrx4) -> None:
        """Perform multiplication of internal array of vector3 by a matrix4x4.

        :type A: matrix 4 by 4
        """
        # vec = array("f", [0.0, 0.0, 0.0, 0.0])
        for arId in range(0, self.capacity, 3):
            x = self.data[arId+0]
            y = self.data[arId+1]
            z = self.data[arId+2]
            
            w = A[3] * x + A[7] * y + A[11] * z + A[15]
            if w < float_info.epsilon:
                w = 1.0

            self.data[arId+0] = (A[0] * x + A[4] * y + A[8] * z + A[12]) / w
            self.data[arId+1] = (A[1] * x + A[5] * y + A[9] * z + A[13]) / w
            self.data[arId+2] = (A[2] * x + A[6] * y + A[10] * z + A[14]) / w

