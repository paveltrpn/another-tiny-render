
import algebra_py as alg

class CCamera(object):
    __projMtrx: alg.mtrx4
    __viewMtrx: alg.mtrx4

    __target:  alg.vec3
    __up:      alg.vec3
    __eye:     alg.vec3

    side:       alg.vec3
    forward:    alg.vec3

    fov:     float
    aspect:  float
    near:    float
    far:     float

    def __init__(self):
        self.__projMtrx = alg.mtrx4()
        self.__viewMtrx = alg.mtrx4()

        self.__eye = alg.vec3(0.0, 0.0, -1.0)
        self.__target = alg.vec3(0.0, 0.0, 0.0)
        self.__up = alg.vec3(0.0, 1.0, 0.0)

        self.fov = 45.0
        self.aspect = 16.0 / 9.0
        self.near = 0.1
        self.far = 100.0


    @property
    def eye(self) -> alg.vec3:
        return self.__eye


    @eye.setter
    def eye(self, new: alg.vec3) -> None:
        self.__eye = new


    @property
    def target(self) -> alg.vec3:
        return self.__target


    @target.setter
    def target(self, new: alg.vec3) -> None:
        self.__target = new


    @property
    def projMtrx(self) -> alg.mtrx4:
        self.__projMtrx = alg.mtrx4FromPerspective(self.fov, self.aspect, self.near, self.far)
        return self.__projMtrx


    @property
    def viewMtrx(self) -> alg.mtrx4:
        self.__viewMtrx = alg.mtrx4FromLookAt(self.__eye, self.__target, self.__up)
        return self.__viewMtrx


    def calcMoveVectors(self) -> None:
        eyeDir = alg.vec3Sub(self.eye, self.target)

        self.side = alg.vec3Cross(self.__up, eyeDir)
        self.side = alg.vec3Normalize(self.side)

        self.forward = alg.vec3Cross(self.__up, self.side)
        self.forward = alg.vec3Normalize(self.forward)


    def moveViewPointsSideway(self, spd: float) -> None:
        self.calcMoveVectors()
        self.target = alg.vec3Sum(self.target, alg.vec3(self.side[0]*spd, self.side[1]*spd, self.side[2]*spd))
        self.eye    = alg.vec3Sum(self.eye,    alg.vec3(self.side[0]*spd, self.side[1]*spd, self.side[2]*spd))


    def moveViewPointsForward(self, spd: float) -> None:
        self.calcMoveVectors()
        self.target = alg.vec3Sum(self.target, alg.vec3(self.forward[0]*spd, self.forward[1]*spd, self.forward[2]*spd))
        self.eye    = alg.vec3Sum(self.eye,    alg.vec3(self.forward[0]*spd, self.forward[1]*spd, self.forward[2]*spd))
