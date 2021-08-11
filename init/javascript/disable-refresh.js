document.onkeydown = function (e) {
    // disable R key
    if (e.ctrlKey && e.keyCode == 82) {
        return false
    }
    if ((e.which || e.keyCode) == 116 || (e.which || e.keyCode) == 82) {
        e.preventDefault()
    }
}
