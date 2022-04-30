function square(x){
  return (x*x);
}

function double(x){
  return (2*x);
}

function compose(f1, f2){
  var composite = R.compose(f1, f2);
  return composite;
}

var f3 = compose(square, double);
var f4 = compose(double, square);
console.log("f3 (5) :", f3(5));
console.log("f4 (5) :", f4(5));
console.log("f3 (10) :", f3(10));
console.log("f4 (10) :", f4(10));