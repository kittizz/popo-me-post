if (document.addEventListener) {
    document.addEventListener("contextmenu", (event) => event.preventDefault())
} else {
    document.attachEvent("oncontextmenu", function () {
        alert("You've tried to open context menu")
        window.event.returnValue = false
    })
}
