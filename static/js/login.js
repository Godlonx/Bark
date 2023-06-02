function openPopup() {
    var popup = document.getElementsByClassName('pop-up')[0];
    popup.classList.add('active');
    var page = document.getElementsByClassName('login')[0];
    page.classList.add('blured');
}