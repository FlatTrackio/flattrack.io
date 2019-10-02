<template>
  <div>
    <b-navbar transparent="false" type="is-white">
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img src="" alt="FlatTrack logo - topbar">
        </b-navbar-item>
      </template>
      <template slot="start">
        <b-navbar-item href="#">
          Home
        </b-navbar-item>
        <b-navbar-item href="#">
          Documentation
        </b-navbar-item>
        <b-navbar-dropdown label="Info">
          <b-navbar-item href="#">
            About
          </b-navbar-item>
          <b-navbar-item href="#">
            Contact
          </b-navbar-item>
        </b-navbar-dropdown>
      </template>
      <template slot="end">
        <b-navbar-item tag="div">
          <div class="buttons">
            <a class="button is-primary">
              <strong>Sign up</strong>
            </a>
            <a class="button is-light">
              Log in
            </a>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
    <section class="hero is-info is-fullheight is-bold">
      <div class="hero-body">
        <div class="container has-text-centered">
          <div class="section">
            <img src="" alt="FlatTrack logo - center" class="image is-256x256">
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
      <section class="hero is-small">
        <div class="section">
          <div class="hero-body">
            <a name="moreinfo"></a>
            <h1 class="title is-2" id="moreinfo">
              What is FlatTrack?
            </h1>
            <p class="subtitle is-4">
              A solution to keep many areas of your flat or community house organised
            </p>
          </div>
          <div class="container">
            <section class="section">
              <h1 class="title is-1">Features</h1>
              <div class="tile is-ancestor" v-for="feature in features" v-bind:key="feature">
                <div class="tile is-parent">
                  <article class="tile is-child box">
                    <p class="title">{{ feature.name }}
                    </p>
                    <p class="subtitle">{{ feature.shortDescription }}</p>
                    <div class="notification">
                      <div class="content">
                        <p>
                          {{ feature.description }}
                        </p>
                      </div>
                    </div>
                  </article>
                </div>
              </div>
            </section>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-primary is-small">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2">
              <a href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">Free and Open Source - Community Project</a>
            </h1>
            <p class="subtitle is-4">
              FlatTrack's source code will be Freely available, licensed under GPL-3.0
            </p>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-warning is-small">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2">
              FlatTrack is under construction
            </h1>
            <p class="subtitle is-4">
              Early equiries will be available soon, we look forward to helping you out soon
            </p>
          </div>
        </div>
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
      errors: [],
      features: [
        {
          name: 'Tasks',
          shortDescription: 'Self management of tasks',
          description: 'Tasks are assigned and rotated with a frequency of never, daily, weekly, or monthly. Flatmembers are notified about the task assigned to them throughout the cycle. Once the task has been completed, flatmembers mark is as done.'
        },
        {
          name: 'Shopping List',
          shortDescription: 'A list of items for your flat\'s weekly shop',
          description: 'Items are added to the flat-shared shopping list. Items can be marked as weekly, so they populate every week.'
        },
        {
          name: 'Noticeboard',
          shortDescription: 'Post to your flatmates',
          description: 'Send posts of information out to your flatmates, so they know what\'s going on.'
        },
        {
          name: 'Shared Calendar',
          shortDescription: 'Keep track of events',
          description: 'Keep up-to-date with events that are happing with your flatmates. Integrates with your existing calendar.'
        },
        {
          name: 'Recipes',
          shortDescription: 'Create and share recipes',
          description: 'Save recipes to share in the flat. Make it easy for everyone to have the meals that they love, made by anyone.'
        },
        {
          name: 'High Fives',
          shortDescription: 'Celebrate your flatmates',
          description: 'Send a positive message to a flatmate to say thanks for something.'
        }
      ]
    }
  }
}
</script>

<style scoped>
div h1 {
  text-align: center;
}
</style>
