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
              Collaborate with your flatmates
            </h2>
          </div>
        </div>
      </div>
    </section>
    <div :class="deviceIsMobile ? 'has-text-centered' : ''">
      <section class="hero is-light is-bold">
        <div class="hero-body">
          <div class="container">
            <h1 class="title is-2">
              Features
            </h1>
            <h2 class="subtitle is-4">
              What FlatTrack will do
            </h2>
            <features/>
            <p class="subtitle is-6">
              Note: features are subject to change
            </p>
          </div>
        </div>
      </section>
    </div>
    <div :class="deviceIsMobile ? 'has-text-centered' : ''">
      <section class="hero is-small is-bold">
        <div class="hero-body">
          <div class="container">
            <h1 class="title is-1">The Goals of FlatTrack</h1>
            <div class="content">
              <dl class="is-large">
                <li class="subtitle is-5">
                  to make it easier to manage common things in a flat
                </li>
                <li class="subtitle is-5">
                  to provide features which bring a flat into collaboration
                </li>
                <li class="subtitle is-5">
                  to have a stream-lined non-techical Free and Open Source solution for flat collaboration
                </li>
              </dl>
            </div>
          </div>
        </div>
      </section>
    </div>
    <div :class="deviceIsMobile ? 'has-text-centered' : ''">
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
                FlatTrack's source code will be Freely available, licensed under <span style="text-decoration: underline dotted;"><a href="https://www.gnu.org/licenses/agpl-3.0.en.html" target="_blank" rel="noreferrer">AGPL-3.0.</a></span>
                <br>
                Contributors welcome!
              </p>
              <b-button
                type="is-primary"
                inverted
                outlined
                icon-left="help"
                tag="a"
                target="_blank"
                size="is-medium"
                href="https://simple.wikipedia.org/wiki/Free_and_open-source_software">
                What's FOSS?
              </b-button>
              <br/>
              <br/>
              <br/>
              <br/>
              <h1 class="title is-2">Contribution</h1>
              <p class="subtitle is-4">
                FlatTrack is now a public project. If you wish to contribute, please head over to the project page:
              </p>
              <b-button
                type="is-white"
                icon-left="hammer-wrench"
                tag="a"
                target="_blank"
                size="is-medium"
                href="https://gitlab.com/flattrack/flattrack">
                GitLab project page
              </b-button>
            </div>
          </div>
        </div>
      </section>
    </div>
    <div :class="deviceIsMobile ? 'has-text-centered' : ''">
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
import common from '@/frontend/common/common'

export default {
  name: 'home',
  data () {
    return {
      logo: false,
      accountButtons: false,
      deviceIsMobile: false
    }
  },
  methods: {
    CheckDeviceIsMobile () {
      this.deviceIsMobile = common.DeviceIsMobile()
    }
  },
  components: {
    navbar: () => import('@/frontend/components/navbar'),
    features: () => import('@/frontend/components/features'),
    interested: () => import('@/frontend/components/interested'),
    ftfooter: () => import('@/frontend/components/footer')
  },
  async beforeMount () {
    this.CheckDeviceIsMobile()
    window.addEventListener('resize', this.CheckDeviceIsMobile, true)
  }
}
</script>

<style scoped>
</style>
