<template>
  <tc-navbar autoBackground="true">
    <tc-navbar-item
      slot="logo"
      icon="card"
      name="Dashboard"
      @click="onClick('dashboard')"
    />
    <tc-navbar-item
      slot="logo"
      icon="tools"
      name="Campaigns"
      @click="onClick('campaigns')"
    />
    <tc-navbar-item
      slot="logo"
      icon="users-duo"
      name="Groups"
      @click="onClick('groups')"
    />
    <tc-navbar-item
      slot="logo"
      icon="user-card"
      name="Accounts"
      @click="onClick('accounts')"
    />
    <tc-navbar-item
      slot="actions"
      icon="gears"
      name="Settings"
      @click="onClick('settings')"
    />
  </tc-navbar>
</template>
<script>
export default {
  name: "Navbar",
  props: {},
  components: {},
  methods: {
    onClick(name) {
      if (this.$route.name !== name) this.$router.push({ name: name });
    },
  },
};
</script>

<style lang="scss">
.tc-navbar {
  padding-left: 1vw !important;
  padding-right: 1vw !important;
}
.tc-navbar-item:hover {
  background-color: #00000017 !important;
}
</style>
