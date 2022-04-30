function isEven(num){
  return (num%2 == 0); 
}
const ex = []

function find_recursive (arr , f){
  var arrTmp = arr;
  if(arrTmp.length == 0){
    return p = [];
  }
  else{
    var arrNew = arrTmp.slice(1);
    var result = [];
    if(f(arrTmp[0])==true){
      result[0] = arrTmp[0];
    }
    return ex.concat(result, find_recursive(arrNew, f));
  }
}

result = find_recursive ([1 , 3, 5, 4, 2] , isEven );
console.log(" resuld find_recursive :", result);