<template>
  <tc-navbar autoBackground="true">
    <tc-navbar-item slot="logo" icon="file" name="File" />
    <tc-navbar-item slot="logo" icon="card" name="Dashboard" />
    <tc-navbar-item slot="logo" icon="tools" name="Campaigns" />
    <tc-navbar-item slot="logo" icon="users-duo" name="Groups" />
    <tc-navbar-item slot="logo" icon="user-card" name="Accounts" />
    <tc-navbar-item icon="gears" name="Settings" />
  </tc-navbar>
</template>
<script>
import { tcNavbar, tcNavbarItem } from "tccomponents_vue/";
export default {
  name: "Navbar",
  props: {},
  components: {
    tcNavbar,
    tcNavbarItem,
  },
};
</script>

<style>
.tc-navbar-item:hover {
  background-color: #00000017 !important;
}
</style>
