
from array import array
from .vec2 import vec2

class CVec2Arrayf(object):
    """ Represents of a vector 2 array. Argument cap in constructor means nomber of
    vector 2 elements in array, i.e. capacity of array is cap*3 float elements"""

    __slots__ = ["data"]
    
    def __init__(self, cap: int) -> None:
        self.data = array("f", [x-x for x in range(cap*2)])

    def __setitem__(self, key: int, value: vec2):
        "Must set vector 2 element at i-th position as vec2 class or somhow else."
        pass

    def __getitem__(self, i: int) -> vec2:
        "Must return vector 2 element at i-th position as vec2 class or somhow else."
        return vec2()
