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


function toSetting(){
    window.location = "http://localhost:8080/settings";
}

function barkColor(){
    var link = document.querySelector("link[rel~='icon']");
    const logo =  document.getElementById("bark-logo")
    if (logo.value !== "angry") {
        link.href = '../static/img/angry-face.png';
        document.body.style.setProperty('--main-color', '#AA0C18');
        if (document.title == "Home / Bark") {
            document.body.style.backgroundImage = "url(../static/img/home-bg-angry.svg)";
        }else{
            document.body.style.backgroundImage = "url(../static/img/angry-bg.svg)";
        }
        
        logo.value = "angry"
        logo.src ="../static/img/angry-face.png"
    }else{
        link.href = '../static/img/logo.png';
        document.body.style.setProperty('--main-color', '#9800fe');
        if (document.title == "Home / Bark") {
            document.body.style.backgroundImage = "url(../static/img/home-bg.svg)";
        }else{
            document.body.style.backgroundImage = "url(../static/img/background.svg)";
        }
        logo.value = "calm"
        logo.src ="../static/img/logo.png"
    }
}

window.onload = function registerError(){
    if (document.title == "Register / Bark") {
    errDiv = document.getElementById("err")
    errText = errDiv.innerHTML
    console.log(errDiv.innerHTML);
    if (errDiv.innerHTML != "none") {
        
        switch (errText) {
            case "bad password":
                errDiv.innerHTML = "Your password isn't valid"
                break;
            case "bad email":
                errDiv.innerHTML = "Your email isn't valid"
                break;
            case "bad username":
                errDiv.innerHTML = "Your username isn't valid"
                break;
            case "unequal password":
                errDiv.innerHTML = "Your passwords doesn't match"
                break;
            case "name already used":
                errDiv.innerHTML = "Username already used"
                break;
            case "unequal password":
                errDiv.innerHTML = "Email already used"
                break; 
            default:
                break;
        }
        errDiv.style.animation=" pop-up 3s cubic-bezier(0,.93,0,1) forwards";
        const vanish = new Promise((resolve, reject) => {
            setTimeout(() => {
              resolve('toto');
            }, 5000);
          });
        vanish.then(e =>{ errDiv.style.animation="pop-out 3s cubic-bezier(0,.93,0,1) forwards";})
    }
}
}
