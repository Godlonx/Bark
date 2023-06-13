function pulse() {
    const logo =  document.getElementById("bark-logo")
    let color1= "#c56eff00"
    let color2= "#9800fec5"
    const obj =  document.getElementById("remember")
    if (logo.value !== "angry") {
        color1 = "#c56eff00";
        color2 = "#9800fec5";
    }else{
        color1 = "#aa0c1900";
        color2 = "#aa0c19c2";
    }

    if (obj.checked) {
        obj.style.boxShadow = "0 0 0.1vw 1vw "+color1;
        obj.style.transition = 'box-shadow 1s'
        
    }else{
        obj.style.boxShadow = "0 0 0.1vw 0vw "+color2;    
        obj.style.transition = 'box-shadow 0.8s'    
    }
  }


function barkColorLogin(){
    var link = document.querySelector("link[rel~='icon']");
    const logo =  document.getElementById("bark-logo")
    const newDiv = document.createElement("div");
    if (logo.value !== "angry") {
        link.href = '../static/img/angry-face.png';
        document.body.style.setProperty('--main-color', '#AA0C18');
        document.body.style.backgroundImage = "url(../static/img/angry-bg.svg)";
        logo.value = "angry"
        logo.src ="../static/img/angry-face.png"
    }else{
        link.href = '../static/img/logo.png';
        document.body.style.setProperty('--main-color', '#9800fe');
        document.body.style.backgroundImage = "url(../static/img/background.svg)";
        logo.value = "calm"
        logo.src ="../static/img/logo.png"
    }
}

function barkColorRegister(){
    var link = document.querySelector("link[rel~='icon']");
    const logo =  document.getElementById("bark-logo")
    const newDiv = document.createElement("div");
    console.log("test");
    if (logo.value !== "cool") {
        link.href = '../static/img/cool.png';
        document.body.style.setProperty('--main-color', '#F29555');
        document.body.style.backgroundImage = "url(../static/img/cool-bg.svg)";
        logo.value = "cool"
        logo.src ="../static/img/cool.png"
    }else{
        link.href = '../static/img/logo.png';
        document.body.style.setProperty('--main-color', '#9800fe');
        document.body.style.backgroundImage = "url(../static/img/background.svg)";
        logo.value = "calm"
        logo.src ="../static/img/logo.png"
    }
}