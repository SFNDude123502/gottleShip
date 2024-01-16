<script>
  import {HumanMove, SetShip, GetIBoard, GetUBoard, SummonShips} from '../wailsjs/go/main/Game.js'

  let [iBoard, uBoard, alreadyShot] = [[],[],[]]
  // iBoard is the board with your ships, uBoard is the the board with the shots you have taken
  let pos, shipPos
  let shipsPlaced = 0
  let isVert = false
  let [resultMessage, oppMessage, shipMessage, winMessage]= ["","","", ""]

  GetIBoard().then(res => iBoard = res)
  GetUBoard().then(res => uBoard = res)

  const range = (start, stop) => Array.from({length: 1+stop-start}, (value, index) => start + index);

  const takeShot = () => {
    let msgToMessage = {
      0: "Missed a Shot",
      1: "Hit a Ship",
      2: "Sunk a Ship",
      3: "Won!"
    }
    let posInt = functionOfTruth(pos)
    if (alreadyShot.includes(posInt)){
      resultMessage = "You Can't Shoot The Same Spot Twice"
      oppMessage = "N/A"
      return
    }
    alreadyShot.push(posInt)
    HumanMove(posInt).then(result =>{ 
      console.log(result)
      GetIBoard().then(res => iBoard = res)
      GetUBoard().then(res => uBoard = res)
      if (result[0]==4){
        resultMessage = "Error"
      }else if(result[0]==3){
        updateWinner(1)
      }else{
        resultMessage = "You Have "+ msgToMessage[result[0]]
      }
      if (result[1]==4){
        oppMessage = "Error"
      }else if (result[1]==3){
        updateWinner(2)
      }else{
        oppMessage = "Opponent Has "+ msgToMessage[result[1]]
      }
    })
  }
  const setAShip = async () =>{
    let shipInt = functionOfTruth(shipPos)
    await SetShip(shipInt, 1, shipsPlaced, isVert).then(res => {
      shipMessage = res
      shipsPlaced += (shipMessage=="" ? 1 : 0)
    })
    if (shipMessage !== ""){
      return
    }
    GetIBoard().then(res => iBoard = res)
    GetUBoard().then(res => uBoard = res)
    updateInvis()
  }
  const summonArmada = async () =>{
    await SummonShips(1)
    GetIBoard().then(res => iBoard = res)
    GetUBoard().then(res => uBoard = res)
    shipsPlaced = 5
    updateInvis()
  }
  const addInvis = (id) => document.getElementById(id).classList.add("invis")
  const removeInvis = (id) => document.getElementById(id).classList.remove("invis")
  
  const updateInvis = () =>{
    if (shipsPlaced > 0){
      addInvis("f")
    }
    if (shipsPlaced >= 5){
      addInvis("d")
      addInvis("e")
    }
    if (shipsPlaced >= 5){
      ["a","b","c"].forEach((e)=>{removeInvis(e)})
    }
  }
  const functionOfTruth = (str) => {
    let a;
    let letterToInt={
      "a":0,"b":1,"c":2,
      "d":3,"e":4,"f":5,
      "g":6,"h":7,"i":8,
      "j":9,
    }
    if (str.length > 3){
      return 0
    }
    if (str.length == 3 ){
      if (str[1]!="1" || str[2]!="0"){
        return 0;
      }
      a=10
    }else{
      a=parseInt(str[1])
    }
    return 10*letterToInt[str[0]] + a-1
  }

  const updateWinner = (p) => {
    ["a","b","c"].forEach((e)=>{addInvis(e)})
    if (p === 1){
      winMessage = "Won!"
    }else{
      winMessage = "Lost :("
    }
    removeInvis("winText")
  }
  
</script>

<main>
  <div class="others">
    <div class="result invis" id="a">Your Turn: {resultMessage}</div>
    <div class="result invis" id="b">Opponent's Turn: {oppMessage}</div>
    <div class="input-box invis" id="c">
      <input autocomplete="off" bind:value={pos} class="input" type="text"/>
      <button class="btn" on:click={takeShot}>Take Your Shot</button>
    </div>

    <div class="result" id="d">{shipMessage}</div>
    <div class="input-box" id="e">
      <span>Upper/Leftmost Ship Square:</span>
      <br/>
      <input autocomplete="off" bind:value={shipPos} placeholder="ex. a1" class="input"  type="text"/>
      <br/>
      <span class="result" >Vertical?</span>
      <input class="input" type="checkbox" bind:checked={isVert}/>
      <br/>
      <button class="btn" on:click={setAShip}>Place a Ship</button>
      <button class="btn" id="f" on:click={summonArmada}>Random Ships</button>
    </div>
    
    <div class="result invis" id="winText">You Have {winMessage}</div>
  </div> 
  <div class="tallKey">
    {#each ["a","b","c","d","e","f","g","h","i", "j"] as row}
      <div class="row">{row}</div>
    {/each}

    <div class="row" id="empty"></div>
    <div class="row" id="empty"></div>

    {#each  ["a","b","c","d","e","f","g","h","i", "j"] as row}
      <div class="row">{row}</div>
    {/each}
  </div>
  <div class="board">
    {#each uBoard as value}
      <div class="gridItem v{value}">{value}</div>
    {/each}
  </div>
  <div class="key">
    {#each range(1,10) as col}
      <div class="column">{col}</div>
    {/each}
  </div>
  <div class="board">
    {#each iBoard as value}
      <div class="gridItem v{value}">{value}</div>
    {/each}
  </div>
</main>
