var personnel = [
{
id: 5 ,
name : " Luke Skywalker ",
pilotingScore : 98 ,
shootingScore : 56 ,
isForceUser : true ,
},
{
id: 82 ,
name : " Sabine Wren ",
pilotingScore : 73 ,
shootingScore : 99 ,
isForceUser : false ,
},
{
id: 22 ,
name : "Zeb Orellios ",
pilotingScore : 20 ,
shootingScore : 59 ,
isForceUser : false ,
},
{
id: 15 ,
name : " Ezra Bridger ",
pilotingScore : 43 ,
shootingScore : 67 ,
isForceUser : true ,
},
{
id: 11 ,
name : " Caleb Dume ",
pilotingScore : 71 ,
shootingScore : 85 ,
isForceUser : true ,
}
];

function crearObj(d, dat1, dat2){
  var obj;
  obj = {
    Nombre : d.name,
    ScoreTotal : dat1+dat2
  }
  return obj;
}

function ScoreTotalIterativo(arr){
  var result= [];
  for(i=0;i<arr.length;i++){
    if(arr[i].isForceUser){
      result.push(crearObj(arr[i],arr[i].pilotingScore, arr[i].shootingScore))
    }
  }
  return result;
}

const ScoreTotalFuncional = personnel.filter(p => p.isForceUser).map(s
=> crearObj(s, s.pilotingScore, s.shootingScore));

console.log("Version Iterativa")
console.log("Score Total de Usuarios de la Fuerza:",ScoreTotalIterativo(personnel))
console.log(" ")
console.log("Version Funcional")
console.log("Score Total de Usuarios de la Fuerza:",ScoreTotalFuncional)