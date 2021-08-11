history.pushState(null, null, document.URL)
window.addEventListener("popstate", function () {
    history.pushState(null, null, document.URL)
})
