document.addEventListener("DOMContentLoaded", function (event) {
    document.ondragstart = function () {
        return false
    }

    document.ondrop = function () {
        return false
    }
    document.body.ondragstart = function () {
        return false
    }

    document.body.ondrop = function () {
        return false
    }
})
