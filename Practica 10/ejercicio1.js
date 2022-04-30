function square(x){
  return (x*x);
}

function double(x){
  return (2*x);
}

function composedValue(square, double, x){
  var composite = R.compose(square, double);
  var result = composite(x)
  return result;
}

console.log(composedValue(square, double, 3));