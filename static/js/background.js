var c = document.getElementById('c'),
    ctx = c.getContext('2d'),
    cw = c.width = window.innerWidth,
    ch = c.height = window.innerHeight


let radius = Math.round(cw/2)
let ampl = 1

let point = [[-100,200],[cw/2,ch/2+500]]

let waveSpeed1 = 0
let waveSpeed2 = 0
let modWave1 = 1
let modWave2 = 1

function loop(){
    ctx.clearRect(0, 0, cw, ch);
   
    ctx.beginPath(); // Start a new path

    ctx.fillStyle = 'blue';

    ctx.moveTo(-100, -100);
    //ctx.beginPath();

    ctx.fillStyle = "#9800fe";
    ctx.strokeStyle= "#9800fe";
    ctx.bezierCurveTo(point[0][0]+waveSpeed1, point[0][1]-waveSpeed1/2,point[1][0]-waveSpeed2/2, point[1][1]+waveSpeed2,cw, ch);
    ctx.fill();
    
    ctx.closePath();
    ctx.moveTo(cw, ch);
    
    ctx.bezierCurveTo(point[1][0], point[1][1],point[0][0], point[0][1],0, 0);
    ctx.fill();
    ctx.stroke();
    ctx.closePath();
    waveSpeed1+=modWave1

    if (waveSpeed1==200) {
        modWave1 = -1
    }else if (waveSpeed1==0){
        modWave1 = 1
    }

    waveSpeed2+=modWave2

    if (waveSpeed2==200) {
        modWave2 = -1
    }else if (waveSpeed2==0){
        modWave2 = 1
    }

    

    return window.requestAnimationFrame(loop);
}

loop()