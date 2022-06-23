#include<iostream>
#include<stdlib.h>
#include<string.h>
#include <ctime>
#include<fstream>
#include<string>
#include <math.h>
#include <algorithm>
using namespace std;

void swap(int* a, int* b){
    int t = *a;
    *a = *b;
    *b = t;
}

void printArray(int arr[], int n){
    int i;
    cout << "[ ";
    for (i = 0; i < n; i++){
        cout << arr[i] << " ";
    }
    cout <<"]"<< endl;
}
 
int partition (int arr[], int low, int high){
    int pivot = arr[high];
    int i = (low - 1);
 
    for (int j = low; j <= high - 1; j++){
        if (arr[j] < pivot){
            i++;
            swap(&arr[i], &arr[j]);
        }
    }
    swap(&arr[i + 1], &arr[high]);
    return (i + 1);
}
 
void quickSort(int arr[], int low, int high){
    if (low < high){
        int pi = partition(arr, low, high);
        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

int main(){
    string arrN[]={"test1.txt","test2.txt", "test3.txt", "test4.txt","test5.txt", "test6.txt", "test7.txt", "test8.txt", "test9.txt", "test10.txt", "test11.txt", "test12.txt", "test13.txt", "test14.txt","test15.txt"};
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
            //if(m==4){
                //printArray(arr,n);
            //}
            t0 = clock();
            quickSort(arr, 0, n - 1);
            t1 = clock();
            double time = (double(t1-t0)/CLOCKS_PER_SEC);
            ltiempos[m]=time;
            //if(m==4){
                //printArray(arr,n);
            //}
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