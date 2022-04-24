function square(x){
  return (x*x);
}

function mymap(arr, f){
  var arrNew = [];
  for(i=0;i<arr.length;i++){
    arrNew = arrNew.concat(f(arr[i]));
  }
  return arrNew;
}

result1 = mymap ([1, 2, 3, 4, 5], square);
result2 = mymap ([1, 4, 9, 16, 25], Math.sqrt);
console.log("result1 :", result1);
console.log("result2 :", result2);