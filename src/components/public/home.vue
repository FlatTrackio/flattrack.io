<template>
  <div>
    <navbar :logo=logo :accountButtons=accountButtons v-if="navbar"/>
    <section class="hero is-info is-fullheight is-bold">
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
            <b-button tag="a" href="#moreinfo" size="is-large" type="is-light">
              Tell me more
            </b-button>
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
              What is FlatTrack?
            </h1>
            <p class="subtitle is-4">
              A solution to keep many areas of your flat or community house organised
            </p>
          </div>
          <features/>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-dark is-small is-bold">
        <div class="section">
          <h1 class="title is-2">Herd cats no longer!</h1>
          <p class="subtitle is-4">Let us solve the frustration and help with collaboration</p>
          <img src="" alt="herding cats" v-if="logo">
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-small is-bold">
        <div class="section">
          <h1 class="title is-2">The Goal of FlatTrack</h1>
          <p class="subtitle is-4">What we aim to achieve to</p>
          <br>
          <div class="content">
            <p>
              - Make it easier to manage common things in a flat <br>
              - Avoid human conflict over forgetting tasks <br>
              - Provide features which bring a flat into collaboration (such as the recipes, and calendar) <br>
              - Have a Free and Open Source solution to organising a flat <br>
            </p>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-primary is-small is-bold">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2">
              Free and Open Source - Community Project
            </h1>
            <p class="subtitle is-4">
              FlatTrack's source code will be Freely available, licensed under GPL-3.0. <br><br>
              Contributors welcome!
            </p>
            <b-button type="is-primary" inverted outlined tag="a" target="_blank" href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">Tell me more about this</b-button>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-warning is-small is-bold">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2">
              FlatTrack is under construction
            </h1>
            <p class="subtitle is-4">
              Early equiries will be available soon, we look forward to helping you out soon
            </p>
            <br><br>
            <div class="container">
              <progress class="progress is-primary" value="15" max="100">15%</progress>
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
import navbar from '@/components/common/navbar'
import features from '@/components/public/home/features'
import interested from '@/components/public/home/interested'
import ftfooter from '@/components/common/footer'

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
