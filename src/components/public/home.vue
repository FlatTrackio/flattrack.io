<template>
  <div>
    <section class="hero is-info is-large is-link is-large-with-navbar">
      <div class="hero-head">
        <nav class="navbar">
          <div class="container">
            <div class="navbar-brand">
              <a class="navbar-item">
                <img src="" alt="Logo">
              </a>
              <span class="navbar-burger burger" data-target="navbarMenuHeroA">
                <span></span>
                <span></span>
                <span></span>
              </span>
            </div>
            <div id="navbarMenuHeroA" class="navbar-menu">
              <div class="navbar-end">
                <a class="navbar-item is-active" href="/">
                  Home
                </a>
                <a class="navbar-item" href="">
                  Sign up
                </a>
                <a class="navbar-item" href="">
                  Community
                </a>
                <a class="navbar-item" href="">
                  Documentation
                </a>
                <span class="navbar-item">
                  <a class="button is-primary is-inverted" href="https://gitlab.com/flattrack/flattrack">
                    <span class="icon">
                      <i class="fab fa-gitlab"></i>
                    </span>
                    <span>Source</span>
                  </a>
                </span>
              </div>
            </div>
          </div>
        </nav>
      </div>
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">
            FlatTrack
          </h1>
          <h2 class="subtitle">
            Keep track of your flat or community house
          </h2>
        </div>
      </div>
    </section>
    <div class="container has-text-centered">
      <section class="section">
        <h1 class="title">
          What is FlatTrack?
        </h1>
        <p>
          FlatTrack is a solution to keep your flat or community house organised.
        </p>
        <!--
          Display and description of tasks (the most general feature)
        -->
      </section>
    </div>
    <div class="container has-text-centered">
      <section class="section">
        <h1 class="title">
          FlatTrack is under construction.
        </h1>
        <p>
          Early equiries will be available soon, we look forward to helping you out soon.
        </p>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'home',
  data () {
    return {
      message: {},
      errors: []
    }
  }
}
</script>

<style scoped>
div h1 {
  text-align: center;
}
</style>
