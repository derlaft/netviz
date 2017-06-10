import numpy 
import sys
import itertools
import math

output = 5

# Input format:
# streamID X1 date X2 X3
def input():
    for s in sys.stdin:
        line = s.strip().split()

        if len(line) < 3:
            continue
        yield (line[0], line[1], numpy.array([float(x) for x in line[2:]]))

# http://stackoverflow.com/a/9489115
def crossProd(a,b):
  assert(len(a) == len(b))
  dimension = len(a)
  c = numpy.array([0.0] * dimension)
  for i in range(dimension):
    for j in range(dimension):
      if j != i:
        for k in range(dimension):
          if k != i:
            if k > j:
              c[i] += a[j]*b[k]
            elif k < j:
              c[i] -= a[j]*b[k]
  return c

def gen():
    a = 1
    while True:
        yield a
        yield -a-1
        a += 2

def bforn(n):
    x = 0.0
    while x < 2 * math.pi:
        yield x
        x += 2 * math.pi / n

def do(dot):

    name = dot[0]
    x = dot[2]

    basis = range(len(x))
    #basis = list(itertools.islice(gen(), len(x)))
    #basis = list(itertools.islice(bforn(len(x)), len(x)))

    a = numpy.array(basis)
    b = numpy.array(list(reversed(basis)))

    norm = numpy.linalg.det(numpy.array([[numpy.dot(a, a), numpy.dot(a, b)], [numpy.dot(b, a), numpy.dot(b, b)]]))

    mn1 = numpy.add(
            numpy.add(
                crossProd(a, numpy.dot(a, numpy.dot(b, b))), # a^2 * bb
                numpy.multiply(2, crossProd(a, numpy.dot(b, numpy.dot(a, b))))
            ),  crossProd(b, numpy.dot(b, numpy.dot(a, a)))
        )

    res = numpy.divide(crossProd(numpy.negative(x), mn1), norm)

    return (name, dot[1], *res[:output])
        
if __name__ == "__main__":
    a = []
    for dot in input():
        a.append(do(dot))
        if len(a) >= 10:
            for i in a:
                print(*i)
            a = []
    for i in a:
        print(*i)

