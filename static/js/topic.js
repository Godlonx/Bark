function commentRedirection(id) {
    window.location = "http://localhost:8080/comment?id="+id
}

function like(val) {
    console.log("pommmme");
    likeInfo = document.getElementById("isLike")
    if (likeInfo.value == val) {
        likeInfo.value = 0
    } else {
        likeInfo.value = val
    }
    console.log(likeInfo.value, val);
    document.getElementById("form").submit();
}