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

.pointer-cursor-on-hover {
    cursor: pointer;
}

.card {
    user-select: none;
    border: 1px dashed rgba(0, 0, 0, .0);
    transition: box-shadow 0.4s, border 0.2s;
}

.card:hover {
    background-color: #fbfbfb;
    transition: box-shadow 0.4s, background-color 0.3s, border 0.2s;
    box-shadow: black 0 0px 45px -30px;
    border: 1px solid darkgray;
}

.card:active {
    background-color: #f1f0f0;
    transition: height 0.1s, width 0.1s, margin 0.1s, box-shadow 0.4s;
    height: calc(100% - 8px);
    border: 1px dashed #000000;
    box-shadow: black 0 0px 66px -27px;
    top: 1px;
}
</style>
