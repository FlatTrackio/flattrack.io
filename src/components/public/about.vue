<template>
    <div>
      <navbar/>
      <section class="hero">
        <div class="hero-body">
          <div class="container">
            <nav class="breadcrumb is-medium has-arrow-separator" aria-label="breadcrumbs">
              <ul>
                <li><a href="/">Home</a></li>
                <li class="is-active"><a href="/#/about">About Us</a></li>
              </ul>
            </nav>
            <h1 class="title is-1">
              About Us
            </h1>
            <h2 class="subtitle is-3">
              Get to know us a bit more
            </h2>
          </div>
        </div>
      </section>
      <section class="section">
        <div class="container">
          <div class="card">
            <div class="card-content">
              <p class="subtitle">
                Who are we?
              </p>
              <p class="title">
                A Free and Open Source development company started in 2019.
              </p>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <p class="subtitle">
                Located in
              </p>
              <p class="title">
                Tauranga, New Zealand
              </p>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <p class="subtitle">
                Developed on
              </p>
              <p class="title">
                <a href="https://gitlab.com/flattrack">GitLab</a>
              </p>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <p class="subtitle">
                Contact found
              </p>
              <p class="title">
                <a href="/#/contact">Here</a>
              </p>
            </div>
          </div>
        </div>
      </section>
      <ftfooter/>
    </div>
</template>

<script>
import navbar from '@/components/common/navbar'
import ftfooter from '@/components/common/footer'

export default {
  name: 'about',
  components: {
    navbar,
    ftfooter
  }
}
</script>

<style scoped>

</style>
