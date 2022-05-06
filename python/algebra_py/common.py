
# multidimensional array mapping, array[i][j]
# row-wise (C, C++):
# (0	1)
# (2	3)

# column-wise (Fortran, Matlab):
# (0	2)
# (1	3)

def IdRw(i: int, j: int, n: int) -> int:
    return (i*n + j)
    