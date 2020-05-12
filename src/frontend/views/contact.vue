<template>
  <div>
    <navbar/>
    <headerDisplay/>
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <nav class="breadcrumb is-medium has-arrow-separator" aria-label="breadcrumbs">
            <ul>
              <li><router-link to="/">Home</router-link></li>
              <li class="is-active"><router-link to="/about">Contact</router-link></li>
            </ul>
          </nav>
          <h1 class="title is-1">
            Contact Us
          </h1>
          <h2 class="subtitle is-4">
            You may wish to enquire about something, here's how to contact us.
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
                icon="chat-sleep"/>
            </div>
            <div class="media-context">
              <p class="subtitle is-3">
                Hold tight there.
              </p>
              <p class="title">
                We're still getting things ready, we'll update this soon.
              </p>
            </div>
          </div>
        </div>
      </div>
        <!--
        <div class="card">
            <div class="card-content">
            <p class="subtitle is-3">
                Email
            </p>
            <p class="title">
                Send us an email at:
                <a href="mailto:info@flattrack.io">info@flattrack.io</a>
            </p>
            </div>
        </div>
        -->
    </div>
    <br><br>
    <ftfooter/>
  </div>
</template>

<script>
import navbar from '@/frontend/components/navbar'
import ftfooter from '@/frontend/components/footer'
import headerDisplay from '@/frontend/components/header-display'

export default {
  name: 'contact',
  components: {
    navbar,
    ftfooter,
    headerDisplay
  }
}
</script>

<style scope>

</style>
