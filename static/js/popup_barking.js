let openPopupButton = document.getElementById("open-popup-button")
let popupContainer = document.getElementById("popup-screen")
let cancelButton = document.getElementById("cancel-button")
let barkingButton = document.getElementById("barking-button")
let titleInput = document.getElementById("title-topic")
let textarea = document.getElementById("textarea")

openPopupButton.addEventListener("click", function() {
    popupContainer.style.display = "block"
})

cancelButton.addEventListener("click", function() {
    popupContainer.style.display = "none"
    titleInput.value = ""
    textarea.value = ""
})

barkingButton.addEventListener("click", function() {
    popupContainer.style.display = "none"
    titleInput.value = ""
    textarea.value = ""
})
