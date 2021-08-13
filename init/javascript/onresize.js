if (document.body) {
    document.body.onresize = function () {
        var w = window.outerWidth
        var h = window.outerHeight
        window.ONxResize(w, h)
    }
}
