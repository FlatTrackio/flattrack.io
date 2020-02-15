<template>
  <div>
    <navbar :logo=logo :accountButtons=accountButtons v-if="navbar"/>
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
    <div class="has-text-centered">
      <a name="moreinfo"></a>
      <section class="hero is-light is-small is-bold">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2" id="moreinfo">
              Features
            </h1>
            <p class="subtitle is-4">
              What FlatTrack will do
            </p>
          </div>
          <features/>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-dark is-small">
        <div class="section">
          <h1 class="title is-1">Herd cats no longer!</h1>
          <p class="subtitle is-4">Hand the reminders over to us.</p>
          <img src="" alt="herding cats" v-if="logo">
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-small is-bold">
        <div class="section">
          <h1 class="title is-1">The Goal of FlatTrack</h1>
          <br>
          <p class="subtitle is-4">We aim to:</p>
          <div class="content">
            <dl class="is-large">
              <li> make it easier to manage common things in a flat </li>
              <li> leverage technology to avoid human conflict over forgetting tasks </li>
              <li> provide features which bring a flat into collaboration (such as the recipes, and calendar) </li>
              <li> have a stream-lined non-techical Free and Open Source solution to organising a flat </li>
            </dl>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-primary is-small is-bold">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-1">
              Free and Open Source Community Project
            </h1>
            <br>
            <p class="subtitle is-4">
              Built for the Community by the Community.
              <br><br>
              FlatTrack's source code will be Freely available, licensed under <a href="https://www.gnu.org/licenses/gpl-3.0.en.html" target="_blank">GPL-3.0.</a>
              <br>
              Contributors welcome!
            </p>
            <b-button type="is-primary" inverted outlined tag="a" target="_blank" href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">What's FOSS?</b-button>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-warning is-small">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-1">
              FlatTrack is under construction
            </h1>
            <p class="subtitle is-4">
              Early equiries will be available soon, we look forward to helping you and your flat.
            </p>
            <br>
            <div class="container">
              <b-progress type="is-primary" :value="25" max="100">25%</b-progress>
            </div>
            <section class="section">
              <interested/>
            </section>
          </div>
        </div>
      </section>
    </div>
    <ftfooter/>
  </div>
</template>

<script>
import navbar from '@/frontend/components/navbar'
import features from '@/frontend/components/features'
import interested from '@/frontend/components/interested'
import ftfooter from '@/frontend/components/footer'

export default {
  name: 'home',
  data () {
    return {
      navbar: true,
      logo: false,
      accountButtons: false
    }
  },
  components: {
    navbar,
    features,
    interested,
    ftfooter
  },
  created () {
  }
}
</script>

<style scoped>
div h1 {
  text-align: center;
}
</style>
