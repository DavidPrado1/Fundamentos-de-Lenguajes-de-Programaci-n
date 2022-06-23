from tabulate import tabulate
import time
import math

def partition(array, low, high):
  pivot = array[high]
  i = low - 1
  for j in range(low, high):
    if array[j] <= pivot:
      i = i + 1
      (array[i], array[j]) = (array[j], array[i])
  (array[i + 1], array[high]) = (array[high], array[i + 1])
  return i + 1


def quick_sort(array, low, high):
  if low < high:
    pi = partition(array, low, high)
    quick_sort(array, low, pi - 1)
    quick_sort(array, pi + 1, high)

ns=[100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000,50000]
prom=[]
dvs=[]
arr2=["test1.txt","test2.txt", "test3.txt", "test4.txt", "test5.txt", "test6.txt", "test7.txt", "test8.txt", "test9.txt", "test10.txt", "test11.txt", "test12.txt", "test13.txt", "test14.txt","test15.txt"]
for i in range(0,15):
    lprom=[]
    for k in range(0,5):
        arr = []
        archivo = open(arr2[i], "r", encoding='utf-8')
        for linea in archivo:
            arr.append(int(linea.strip()))
        archivo.close
        inicio = time.time()
        quick_sort(arr, 0, len(arr) - 1)
        fin = time.time()
        lprom.append(fin-inicio)
    promf = 0
    for j in range(0,5):
        promf += lprom[j]
    promf = (promf/5)
    prom.append(promf)
    devst=0
    for m in range(0,5):
        devst += pow(lprom[m]-promf,2)
    devst=math.sqrt(devst/4)
    dvs.append(devst)
 
print(tabulate({'n': ns, 'Promedio': prom, 'Desviacion Estandar': dvs}, headers="keys", tablefmt='fancy_grid'))