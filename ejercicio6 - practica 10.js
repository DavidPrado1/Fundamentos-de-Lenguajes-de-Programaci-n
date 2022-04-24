function square(x){
  return (x*x);
}

const ex = [];

function mymap_recursive(arr, f){
  var arrTmp = arr;
  if(arrTmp.length == 0){
    return p = [];
  }
  else{
    var arrNew = arrTmp.slice(1);
    var result = [];
    result[0] = f(arrTmp[0]);
    return ex.concat(result, mymap_recursive(arrNew, f));
  }
}

result1 = mymap_recursive ([1, 2, 3, 4, 5], square);
result2 = mymap_recursive ([1, 4, 9, 16, 25], Math.sqrt);
console.log("mymap_recursive result1 :", result1);
console.log("mymap_recursive result2 :", result2);