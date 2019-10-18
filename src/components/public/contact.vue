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
                <li class="is-active"><a href="/#/about">Contact</a></li>
              </ul>
            </nav>
            <h1 class="title is-1">
              Contact Us
            </h1>
            <h2 class="subtitle is-3">
              You may wish to enquire about something, here's how to contact us
            </h2>
            <section class="section">
                <div class="container">
                    <div class="card">
                        <div class="card-content">
                        <p class="subtitle">
                            Hold tight there.
                        </p>
                        <p class="title">
                            No contact available yet. <br>
                            We're still getting things ready, so we'll update this soon.
                        </p>
                        </div>
                    </div>
                    <!--
                    <div class="card">
                        <div class="card-content">
                        <p class="subtitle">
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
            </section>
          </div>
        </div>
      </section>
      <ftfooter/>
    </div>
</template>

<script>
import navbar from '@/components/common/navbar'
import ftfooter from '@/components/common/footer'
import headerDisplay from '../common/header-display'

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
