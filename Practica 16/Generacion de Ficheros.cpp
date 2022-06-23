#include<iostream>
#include<cstdlib>
#include<ctime>
#include <fstream>
#include <string>
using namespace std;

bool checkrep(int n, int num[],int tam){
  for(int i=0; i<tam; i++)
    if(n == num[i]){
      return true;
    }
  return false;
}

int main(){
  string arr[]={"test1.txt","test2.txt", "test3.txt", "test4.txt", "test5.txt", "test6.txt", "test7.txt", "test8.txt", "test9.txt", "test10.txt", "test11.txt", "test12.txt", "test13.txt", "test14.txt","test15.txt"};
	int arr2[]={100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000,50000};
	for(int j = 0;j<15;j++){
    string nombreArchivo = arr[j];
    ofstream archivo;
	  archivo.open(nombreArchivo.c_str(), fstream::out);
	  srand(time(NULL));
    int tam = arr2[j];
    int n, num[arr2[j]];
    string print;
    string t;
    for(int i=0; i<tam; i++){
      do
        n = 1 + rand() % tam;
      while(checkrep(n, num,tam));
      num[i] = n;
      if(i==tam-1){
          archivo<<n;
      }
      else{
        archivo<<n<<endl;
      }
    }
    archivo.close();
  }
}
