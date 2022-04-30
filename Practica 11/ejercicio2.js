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

function PrioridadIterativo(arr){
  var result= [];
  for(i=0;i<arr.length;i++){
    if(arr[i].priority==1){
      result.push(arr[i])
    }
  }
  return result;
}

const PrioridadFuncional = tasks.filter(t => t.priority ==1);

console.log("Version Iterativa")
console.log("Tareas con Prioridad 1",PrioridadIterativo(tasks))
console.log(" ")
console.log("Version Funcional")
console.log("Tareas con Prioridad 1:",PrioridadFuncional)