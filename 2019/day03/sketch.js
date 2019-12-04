let input;
let wire1 =[];
let wire2 =[];
let minX, maxX,minY,maxY;
let frameCounter = 0;
let fps = 3;
let totalFrames = 0;
let frames1;
let frames2;
let framesP;
let gif;
const FPS = 30, FORMAT = 'gif', WORKERSFOLDER='./', VERBOSE = false, DISPLAY = true,
       FRAMERATE = FPS, FRAMELIMIT = 302
let ccc = new CCapture({ format:'jpg',workersPath:'./', framerate:FPS})

function parseWire(ws){
  let wire= [];
  currPoint = createVector(0,0);
  wire.push(currPoint)
  for (let i= 0; i<ws.length;i++) {
    let steps = int(ws[i].slice(1))
    let nextPoint;
    switch (ws[i][0]) {
      case 'D':
        nextPoint = createVector(currPoint.x,currPoint.y-steps)
        break
      case 'U':
        nextPoint = createVector(currPoint.x, currPoint.y+steps)
        break
      case 'L':
        nextPoint = createVector(currPoint.x-steps, currPoint.y)
        break
      case 'R':
        nextPoint = createVector(currPoint.x+steps, currPoint.y)
        break
    }
    wire.push(nextPoint)
    currPoint = nextPoint
  }
  return wire
}
function preload() {
  input = loadStrings('input03_2019.txt');
}
function setup() {
  frameRate(FPS);
  // put setup code here
  // console.log(input[0])
  let w1s = split(input[0],',')
  let w2s = split(input[1],',')
  frames1 = w1s.length
  frames2 = w2s.length
  totalFrames = Math.max(frames1,frames2)
  wire1 =  parseWire(w1s)
  wire2 = parseWire(w2s)
  minX = 0
  maxX = 0
  minY = 0
  maxY = 0
  for (let i=0;i<wire1.length;i++){
    minX = Math.min(minX,wire1[i].x)
    maxX = Math.max(maxX,wire1[i].x)
    minY = Math.min(minY,wire1[i].y)
    maxY = Math.max(maxY,wire1[i].y)
  }
  for (let i=0;i<wire2.length;i++){
    minX = Math.min(minX,wire2[i].x)
    maxX = Math.max(maxX,wire2[i].x)
    minY = Math.min(minY,wire2[i].y)
    maxY = Math.max(maxY,wire2[i].y)
  }
  console.log(minX,maxX,minY,maxY)
  // console.log(wire1)
  
  ctx = createCanvas(1000,1000)
  background(20);
  ccc.start()
  
}

function draw() {
  
  // put drawing code here
  strokeWeight(3)
  
  if (frameCounter < frames1-1){
    stroke('red');
    //for (let i=dt;i<dt+dpf;i++){
      startX = map(wire1[frameCounter].x,minX,maxX,0,1000)
      startY = map(wire1[frameCounter].y,minY,maxY,1000,0)
      endX = map(wire1[frameCounter+1].x,minX,maxX,0,1000)
      endY = map(wire1[frameCounter+1].y,minY,maxY,1000,0)
      line(startX, startY, endX, endY)
    //}
  }
  if (frameCounter < frames2-1){
    stroke('green')
    //for (let i=dt -wire1.length+2;i<dt-wire1.length+2+dpf;i++){
      //console.log(i)
      startX = map(wire2[frameCounter].x,minX,maxX,0,1000)
      startY = map(wire2[frameCounter].y,minY,maxY,1000,0)
      endX = map(wire2[frameCounter+1].x,minX,maxX,0,1000)
      endY = map(wire2[frameCounter+1].y,minY,maxY,1000,0)
      line(startX, startY, endX, endY)
    }
  //}
  // console.log("HHHH",wire2.length)
  frameCounter++;
  ccc.capture(document.getElementById('defaultCanvas0'))
  let fn = "day03_"+frameCounter+".png";
  console.log(fn)
  if (frameCounter > FRAMELIMIT){
    ccc.stop()
    ccc.save()
    console.log("STOP")
    noLoop()

  }
}