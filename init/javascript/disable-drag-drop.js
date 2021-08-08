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
document.addEventListener(
    "click",
    function (e) {
        if (e.which === 2) {
            e.preventDefault()
        }
    },
    false
)

document.addEventListener(
    "mousedown",
    function (e) {
        if (e.which == 2) {
            e.preventDefault()

            return false
        }
    },
    false
)
