<template>
    <div>
        <b-navbar fixed-top transparent="false" type="is-light">
            <template slot="brand" v-if="logo">
                <b-navbar-item tag="router-link" :to="{ path: '/' }">
                <img src="" alt="FlatTrack logo - topbar">
                </b-navbar-item>
            </template>
            <template slot="start">
                <b-navbar-item href="/">
                    Home
                </b-navbar-item>
                <!--
                <b-navbar-item href="https://docs.flattrack.io">
                    Documentation
                </b-navbar-item>
                <b-navbar-item href="https://blog.flattrack.io">
                    Blog
                </b-navbar-item>
                -->
                <b-navbar-dropdown label="Info">
                    <b-navbar-item href="/#/about">
                        About
                    </b-navbar-item>
                    <b-navbar-item href="/#/contact">
                        Contact
                    </b-navbar-item>
                </b-navbar-dropdown>
            </template>
            <template slot="end" v-if="accountButtons">
                <b-navbar-item tag="div">
                <div class="buttons">
                    <b-button type="is-primary" tag="a" href="https://live.flattrack.io/signup">
                        <strong>Sign up</strong>
                    </b-button>
                    <b-button type="is-light" tag="a" href="https://live.flattrack.io/login">
                        Log in
                    </b-button>
                </div>
                </b-navbar-item>
            </template>
        </b-navbar>
    </div>
</template>

<script>
export default {
  name: 'navbar',
  props: {
    'logo': Boolean,
    'accountButtons': Boolean
  }
}
</script>

<style scoped>

</style>
