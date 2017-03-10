import numpy 
import sys

#x = numpy.array([100, -102, -100, 12312, 1])

x = numpy.array([float(x) for x in sys.argv[1:]])

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

print (mn1, norm)

res = numpy.divide(crossProd(numpy.negative(x), mn1), norm)

print (res)
print (res[:2])
        
