const vehicles = [
	  { make: 'Honda', model: 'CR-V', type: 'suv', price: 24045 },
	  { make: 'Honda', model: 'Accord', type: 'sedan', price: 22455 },
	  { make: 'Mazda', model: 'Mazda 6', type: 'sedan', price: 24195 },
	  { make: 'Mazda', model: 'CX-9', type: 'suv', price: 31520 },
	  { make: 'Toyota', model: '4Runner', type: 'suv', price: 34210 },
	  { make: 'Toyota', model: 'Sequoia', type: 'suv', price: 45560 },
	  { make: 'Toyota', model: 'Tacoma', type: 'truck', price: 24320 },
	  { make: 'Ford', model: 'F-150', type: 'truck', price: 27110 },
	  { make: 'Ford', model: 'Fusion', type: 'sedan', price: 22120 },
	  { make: 'Ford', model: 'Explorer', type: 'suv', price: 31660 }
	];

function PreciosAutosSuvIterativo(arr){
  var result= [];
  for(i=0;i<arr.length;i++){
    if(arr[i].type=='suv'){
      result.push(arr[i].price)
    }
  }
  return result;
}

const PreciosAutosSuvFuncional = vehicles.filter(v => v.type =='suv').map(v => v.price);

console.log("Version Iterativa")
console.log("Precios Autos Suv:",PreciosAutosSuvIterativo(vehicles))
console.log(" ")
console.log("Version Funcional")
console.log("Precios Autos Suv:",PreciosAutosSuvFuncional)