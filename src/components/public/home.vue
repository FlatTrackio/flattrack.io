<template>
  <div>
    <b-navbar transparent="false" type="is-white">
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img src="" alt="FlatTrack logo">
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
    <section class="hero is-info is-medium is-bold">
      <div class="hero-body">
        <div class="container has-text-centered">
          <div class="section">
            <h1 class="title is-1">
              FlatTrack
            </h1>
            <h2 class="subtitle is-3">
              Keep track of your flat or community house
            </h2>
          </div>
        </div>
        <br><br><br><br><br><br><br><br>
      </div>
      <div class="hero-foot">
        <nav class="tabs is-boxed is-fullwidth">
          <div class="container">
            <ul>
              <li class="is-active"><a href="#moreinfo">Tell me more</a></li>
            </ul>
          </div>
        </nav>
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
            <p>
              FlatTrack is a solution to keep your flat or community house organised.
            </p>
            <br>
            <h2 class="title is-3">
              Featuring:
            </h2>
            <ul>
              <div v-for="feature in features" v-bind:key="feature">
                <b-tooltip :label="feature.description" position="is-top">
                  <li>- {{ feature.name }}</li>
                </b-tooltip>
              </div>
            </ul>
          </div>
        </div>
      </section>
    </div>
    <div class="has-text-centered">
      <section class="hero is-primary is-small">
        <div class="section">
          <div class="hero-body">
            <h1 class="title is-2">
              <a href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">Free and Open Source</a>
            </h1>
            <p>
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
              FlatTrack is under construction.
            </h1>
            <p>
              Early equiries will be available soon, we look forward to helping you out soon.
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
          description: 'Self management of tasks'
        },
        {
          name: 'Shopping List',
          description: 'A list of items for your flat\'s weekly shop'
        },
        {
          name: 'Noticeboard',
          description: 'Post to your flatmates'
        },
        {
          name: 'Shared Calendar',
          description: 'Keep track of events'
        },
        {
          name: 'Recipes',
          description: 'Create and share recipes'
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
