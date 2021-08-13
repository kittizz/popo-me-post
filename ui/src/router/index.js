import Vue from "vue"
import VueRouter from "vue-router"
import Dashboard from "../views/Dashboard.vue"
import Campaigns from "../views/Campaigns.vue"
import Groups from "../views/Groups.vue"
import Accounts from "../views/Accounts.vue"
import Settings from "../views/Settings.vue"

Vue.use(VueRouter)

const routes = [
    {
        path: "/",
        redirect: "/dashboard",
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: Dashboard,
    },
    {
        path: "/campaigns",
        name: "campaigns",
        component: Campaigns,
    },
    {
        path: "/groups",
        name: "groups",
        component: Groups,
    },
    {
        path: "/accounts",
        name: "accounts",
        component: Accounts,
    },
    {
        path: "/settings",
        name: "settings",
        component: Settings,
    },
]

const router = new VueRouter({
    routes,
})

export default router
