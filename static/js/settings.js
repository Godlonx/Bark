function openPopup(value) {
    popup = document.getElementById(value)
    back1 = document.getElementById('back1')
    back2 = document.getElementById('back2')
    back1.classList.add('blured')
    back2.classList.add('blured')
    popup.classList.add('popup-Active')
}

function closePopup(value) {
    popup = document.getElementById(value)
    back1 = document.getElementById('back1')
    back2 = document.getElementById('back2')
    back1.classList.remove('blured')
    back2.classList.remove('blured')
    popup.classList.remove('popup-Active')
}
