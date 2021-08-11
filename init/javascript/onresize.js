if (document.body) {
    document.body.onresize = function () {
        var w = window.outerWidth
        var h = window.outerHeight
        ONxResize(w, h)
    }
}
