function isEven(num){
  return (num%2 == 0); 
}

function find(arr, f){
  var arrNew = [];
  for(i=0;i<arr.length;i++){
    if(f(arr[i])==true){
      arrNew = arrNew.concat(arr[i]);
    }
  }
  return arrNew;
}

result = find([1, 3, 5, 4, 2], isEven);
console.log("find result :", result);