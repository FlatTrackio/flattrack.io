<template>
    <div>
      <navbar/>
      <headerDisplay/>
      <section class="hero">
        <div class="hero-body">
          <div class="container">
            <nav class="breadcrumb is-medium has-arrow-separator" aria-label="breadcrumbs">
              <ul>
                <li><a href="/">Home</a></li>
                <li class="is-active"><a href="/about">About Us</a></li>
              </ul>
            </nav>
            <h1 class="title is-1">
              About Us
            </h1>
            <h2 class="subtitle is-4">
              Get to know us a bit more
            </h2>
          </div>
        </div>
      </section>
      <div class="container">
        <div class="card">
          <div class="card-content">
            <div class="media">
              <div class="media-left">
                <b-icon
                  size="is-medium"
                  icon="account-group"/>
              </div>
              <div class="media-context">
                <p class="subtitle is-5">
                  Who are we?
                </p>
                <p class="title is-4">
                  A Free and Open Source development company, started in 2019.
                </p>
              </div>
            </div>
          </div>
        </div>
        <div class="card">
          <div class="card-content">
            <div class="media">
              <div class="media-left">
                <b-icon
                  size="is-medium"
                  icon="map-marker"/>
              </div>
              <div class="media-context">
                <p class="subtitle is-5">
                  Located in
                </p>
                <p class="title is-4">
                  Tauranga, New Zealand.
                </p>
              </div>
            </div>
          </div>
        </div>
        <div class="card">
          <div class="card-content">
            <div class="media">
              <div class="media-left">
                <b-icon
                  size="is-medium"
                  icon="code-braces"/>
              </div>
              <div class="media-context">
                <p class="subtitle is-5">
                  Developed on
                </p>
                <p class="title is-4">
                  <a target="_blank" href="https://gitlab.com/flattrack" rel="noreferrer">GitLab</a>.
                </p>
              </div>
            </div>
          </div>
        </div>
        <div class="card">
          <div class="card-content">
            <div class="media">
              <div class="media-left">
                <b-icon
                  size="is-medium"
                  icon="email"/>
              </div>
              <div class="media-context">
                <p class="subtitle is-5">
                  Contact found
                </p>
                <p class="title is-4">
                  <router-link to="/contact">Here</router-link>.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <br><br>
      <ftfooter/>
    </div>
</template>

<script>
import navbar from '@//components/navbar'
import ftfooter from '@//components/footer'
import headerDisplay from '@//components/header-display'

export default {
  name: 'about',
  components: {
    navbar,
    ftfooter,
    headerDisplay
  }
}
</script>

<style scoped>

</style>
