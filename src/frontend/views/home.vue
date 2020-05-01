<template>
  <div>
    <navbar :logo=logo :accountButtons=accountButtons />
    <section class="hero is-info is-medium is-bold">
      <div class="hero-body">
        <div class="container has-text-centered">
          <div class="section">
            <img src="" alt="FlatTrack logo - center" class="image is-256x256" v-if="logo">
            <h1 class="title is-1">
              FlatTrack
            </h1>
            <h2 class="subtitle is-3">
              Keep track of your flat or community house
            </h2>
          </div>
        </div>
      </div>
    </section>
    <div>
      <section class="hero is-light is-bold">
        <div class="hero-body">
          <div class="container">
            <h1 class="title is-2">
              Features
            </h1>
            <h2 class="subtitle is-4">
              What FlatTrack will do
            </h2>
          </div>
          <features/>
        </div>
      </section>
    </div>
    <div>
      <section class="hero is-dark is-small">
        <div class="hero-body">
          <div class="container">
            <h1 class="title is-1">Herd cats no longer!</h1>
            <p class="subtitle is-4">Hand the reminders over to us.</p>
            <img src="" alt="herding cats" v-if="logo">
          </div>
        </div>
      </section>
    </div>
    <div>
      <section class="hero is-small is-bold">
        <div class="hero-body">
          <div class="container">
            <h1 class="title is-1">The Goal of FlatTrack</h1>
            <br>
            <p class="subtitle is-4">We aim to</p>
            <div class="content">
              <dl class="is-large">
                <li>
                  make it easier to manage common things in a flat
                </li>
                <li>
                  provide features which bring a flat into collaboration
                </li>
                <li>
                  have a stream-lined non-techical Free and Open Source solution to organising a flat
                </li>
              </dl>
            </div>
          </div>
        </div>
      </section>
    </div>
    <div>
      <section class="hero is-primary is-small is-bold">
        <div class="hero-body">
          <div class="container">
            <div class="hero-body">
              <h1 class="title is-1">
                A Free and Open Source Community Project
              </h1>
              <br>
              <p class="subtitle is-4">
                Built for the Community by the Community.
                <br><br>
                FlatTrack's source code will be Freely available, licensed under <a href="https://www.gnu.org/licenses/agpl-3.0.en.html" target="_blank">AGPL-3.0.</a>
                <br>
                Contributors welcome!
              </p>
              <b-button type="is-primary" inverted outlined tag="a" target="_blank" href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">What's FOSS?</b-button>
            </div>
          </div>
        </div>
      </section>
    </div>
    <div>
      <section class="hero is-warning is-small">
        <div class="hero-body">
          <div class="container">
            <div class="hero-body">
              <h1 class="title is-1">
                FlatTrack is under construction
              </h1>
              <p class="subtitle is-4">
                Early equiries will be available soon, we look forward to helping you and your flat.
              </p>
              <br>
              <div class="container">
                <b-progress type="is-primary" :value="30" max="100">30%</b-progress>
              </div>
              <section class="section">
                <interested/>
              </section>
            </div>
          </div>
        </div>
      </section>
    </div>
    <ftfooter/>
  </div>
</template>

<script>
export default {
  name: 'home',
  data () {
    return {
      logo: false,
      accountButtons: false
    }
  },
  components: {
    navbar: () => import('@/frontend/components/navbar'),
    features: () => import('@/frontend/components/features'),
    interested: () => import('@/frontend/components/interested'),
    ftfooter: () => import('@/frontend/components/footer')
  }
}
</script>

<style scoped>
</style>
