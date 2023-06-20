var sidenav = document.getElementById("side-left");
var openBtn = document.getElementById("side-open");

function openNav() {
    if (sidenav.classList[1]=="active") {
        sidenav.classList.remove("active");
        openBtn.classList.remove("active");
    }else{
        sidenav.classList.add("active");
        openBtn.classList.add("active");
    }
    
}