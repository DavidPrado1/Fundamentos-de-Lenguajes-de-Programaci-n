#include<iostream>
#include<stdlib.h>
#include<string.h>
#include <ctime>
#include<fstream>
#include<string>
#include <math.h>
using namespace std;

void insertionSort(int arr[], int n){
    int i, key, j;
    for (i = 1; i < n; i++){
        key = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > key){
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

void printArray(int arr[], int n){
    int i;
    cout << "[ ";
    for (i = 0; i < n; i++){
        cout << arr[i] << " ";
    }
    cout <<"]"<< endl;
}

int main(){
    string arrN[]={"test1.txt","test2.txt", "test3.txt", "test4.txt", "test5.txt", "test6.txt", "test7.txt", "test8.txt", "test9.txt", "test10.txt", "test11.txt", "test12.txt", "test13.txt", "test14.txt","test15.txt"};
	int arr2[]={100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000,50000};
    for(int j=0;j<15;j++){
        double ltiempos[5];
        for(int m=0;m<5;m++){
            unsigned t0, t1;
            int n=arr2[j];
            int arr[n];
            
            ifstream archivo(arrN[j]);
            string texto;

            int i=0;
            while(getline(archivo,texto)){
                arr[i]=stoi(texto);
                i=i+1;
            }
            
            archivo.close();
            //printArray(arr,n);
            t0 = clock();
            insertionSort(arr, n);
            t1 = clock();
            double time = (double(t1-t0)/CLOCKS_PER_SEC);
            ltiempos[m]=time;
            //printArray(arr,n);
        }
        double promf = 0;
        for(int k=0;k<5;k++){
            promf += ltiempos[k];
        }
        promf = promf / 5;
        double dvs = 0;
        for(int p=0;p<5;p++){
            dvs += pow((ltiempos[p]-promf),2);
        }
        dvs=sqrt(dvs/4);
        cout<<"[Array Tamaño:"<<arr2[j]<<"][Promedio de Tiempo de ejecucion:"<<promf<<"][Desviacion Estandar:"<<dvs<<"]"<<endl;
    }	
	
	return 0;
}