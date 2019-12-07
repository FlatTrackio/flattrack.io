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
              <li class="is-active"><a href="/about">Contact</a></li>
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
        <p class="subtitle is-3">
            Hold tight there.
        </p>
        <p class="title">
            We're still getting things ready, we'll update this soon.
        </p>
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
import navbar from '@/components/navbar'
import ftfooter from '@/components/footer'
import headerDisplay from '@/components/header-display'

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
