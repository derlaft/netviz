#!/usr/bin/python3

from PIL import Image, ImageDraw
import sys, math

h = 640
w = 480

# Input format:
# streamID X1 X2 X3
def input():
    for s in sys.stdin:
        line = s.strip().split()

        if len(line) != 3:
            continue

        yield (line[0], [float(x) for x in line[1:]])

# init vars
points = {}
total = 0

mmax_x = 0
mmax_y = 0
mmin_x = 0
mmin_y = 0


def draw():
    img = Image.new('RGB', (640, 480))
    drw = ImageDraw.Draw(img, 'RGBA')
    print("HERE", points)
    # 2) draw everyting
    for ptx in points.keys():

        arr = [( ( x*w/(mmax_x+mmin_x)), (y*h/(mmax_y+mmin_y))-mmin_y ) for point in points[ptx] for x, y in point]
        print(arr)
        # skip reduced
        if len(arr) < 3:
            continue


        # draw
        drw.polygon(sorted(arr), (255, 255, 255, int(255 * len(arr)/total)))

    del drw

    img.show()
    img.save('out.png', 'PNG')

# What it does:
for new in input():
    print("ADDIN")

    old = points.get(new[0], [])
    # 1) Add new stream point
    points[new[0]] = old + [new[1:]]
    total += 1

    print("eh", new)
    mmax_x = max(mmax_x, new[1][0])
    mmin_x = min(mmin_x, new[1][0])

    mmax_y = max(mmax_y, new[1][1])
    mmin_y = min(mmin_y, new[1][1])

draw()

