import Vue from "vue"
import App from "./App.vue"

Vue.config.productionTip = false
import router from "./router"
import store from "./store"

import * as TCComponents from "tccomponents_vue"
import "tccomponents_vue/lib/tccomponents_vue.css"

for (const component in TCComponents) {
    Vue.component(component, TCComponents[component])
}
new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app")
