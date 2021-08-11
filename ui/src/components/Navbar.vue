<template>
  <tc-navbar autoBackground="true">
    <tc-navbar-item slot="logo" icon="file" name="File" />
    <tc-navbar-item slot="logo" icon="card" name="Dashboard" />
    <tc-navbar-item slot="logo" icon="tools" name="Campaigns" />
    <tc-navbar-item slot="logo" icon="users-duo" name="Groups" />
    <tc-navbar-item slot="logo" icon="user-card" name="Accounts" />
  </tc-navbar>
</template>
<script>
export default {
  name: "Navbar",
  props: {},
  components: {},
};
</script>

<style>
.tc-navbar-item:hover {
  background-color: #00000017 !important;
}
</style>
