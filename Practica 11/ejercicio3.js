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

function DuracionTotalIterativa(arr){
  var result = 0;
  for(i=0;i<arr.length;i++){
    result = result + arr[i].duration;
  }
  return result;
}

const DuracionTotalFuncional = tasks.map(t =>
t.duration).reduce((sum,duration, i, array) => sum + duration, 0);

console.log("Version Iterativa")
console.log("La duracion total de las tareas es de: ",DuracionTotalIterativa(tasks),
" minutos.");
console.log(" ")
console.log("Version Funcional")
console.log("La duracion total de las tareas es de: ",DuracionTotalFuncional,
" minutos.");