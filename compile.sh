cython --embed $1.py && gcc -Ofast -I /usr/include/python3.6m -o $1 $1.c -lpython3.6m -lpthread -lm -lutil -ldl -flto
