<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
@import "~bulma/sass/utilities/_all";

@import '~@mdi/font/css/materialdesignicons.min.css';

$primary: #00a7d6;
$primary-invert: findColorInvert($primary);

$colors: (
    "white": ($white, $black),
    "black": ($black, $white),
    "light": ($light, $light-invert),
    "dark": ($dark, $dark-invert),
    "primary": ($primary, $primary-invert),
    "info": ($info, $info-invert),
    "success": ($success, $success-invert),
    "warning": ($warning, $warning-invert),
    "danger": ($danger, $danger-invert)
);

$link: #2c77c9;
$link-invert: $black;
$link-focus-border: $primary;

$footer-padding: 1rem 1.5rem 1rem;

@import "~bulma";
@import "~buefy/src/scss/buefy";

</style>

<style>
html {
  scroll-behavior: smooth
}
</style>
