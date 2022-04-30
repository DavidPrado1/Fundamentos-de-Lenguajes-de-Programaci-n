var tasks = [
{'name' : 'Buy milk from the shop',
'duration' : 20 ,
'priority' : 1},
{'name' : 'Clean the house',
'duration' : 120 ,
'priority' : 3},
{'name' : 'Study JS functions',
'duration' : 180 ,
'priority' : 1}
];

function OnlyNameIterativo(arr){
  var result= [];
  for(i=0;i<arr.length;i++){
    result.push(arr[i].name)
  }
  return result;
}

const OnlyNameFuncional = tasks.map(t => t.name);

console.log("Version Iterativa")
console.log("Nombres de las Tareas: ", OnlyNameIterativo(tasks))
console.log(" ")
console.log("Version Funcional")
console.log("Nombres de las Tareas: ", OnlyNameFuncional)