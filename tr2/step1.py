import numpy 
import sys

#x = numpy.array([100, -102, -100, 12312, 1])

# Input format:
# streamID X1 X2 X3
def input():
    for s in sys.stdin:
        line = s.strip().split()

        if len(line) < 3:
            continue
        yield (line[0], numpy.array([float(x) for x in line[1:]]))

def crossProd(a,b):
  dimension = len(a)
  c = []
  for i in range(dimension):
    c.append(0)
    for j in range(dimension):
      if j != i:
        for k in range(dimension):
          if k != i:
            if k > j:
              c[i] += a[j]*b[k]
            elif k < j:
              c[i] -= a[j]*b[k]
  return c

def do(dot):

    name = dot[0]
    x = dot[1]

    basis = range(len(x))

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

    print (name, res[0], res[1])
        
if __name__ == "__main__":
    for dot in input():
        do(dot)
