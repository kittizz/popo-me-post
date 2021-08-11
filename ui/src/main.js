import Vue from "vue"
import App from "./App.vue"

Vue.config.productionTip = false
import "tccomponents_vue/lib/tccomponents_vue.css"
import router from "./router"
import store from "./store"

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app")
