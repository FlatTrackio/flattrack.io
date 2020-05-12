<template>
  <div>
    <navbar :logo=logo :accountButtons=accountButtons v-if="navbar"/>
    <headerDisplay/>
    <div class="container">
      <section class="section">
        <nav class="breadcrumb is-medium has-arrow-separator" aria-label="breadcrumbs">
          <ul>
            <li><router-link to="/"> Home </router-link></li>
            <li class="is-active"><router-link to="/privacy-policy"> Privacy Policy </router-link></li>
          </ul>
        </nav>
        <h1 class="title is-1"> Privacy Policy </h1>
        <p class="subtitle is-3"> How data is collected and handled. </p>
        <br>
        <!--
            Notice (what personal information is being collected on the site)
            Choice (what options the customer has about how/whether personal data is collected and used)
            Access (how a customer can see what data has been collected and change/correct it if necessary)
            Security (state how any data that is collected is stored/protected)
            Redress (what customer can do if privacy policy is not met)
        -->
        <div class="container">
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="database"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    What information do we collect?
                  </p>
                  <p class="title">
                    Email addresses per submission.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="pipe"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    Is the collection of data opt-in or opt-out?
                  </p>
                  <p class="title">
                    Opt-in only. Email addresses are only inputted and not collected without consent.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="account-arrow-right"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    What will this information be used for?
                  </p>
                  <p class="title">
                    Sending email notifications once FlatTrack is ready.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="bug"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    Does this site have embedded trackers?
                  </p>
                  <p class="title">
                    We do not use trackers of any kind.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="cookie"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    Do we use 'cookies'?
                  </p>
                  <p class="title">
                    Cookies are small text files stored in the browser. We do not store any cookies.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="share-off"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    Do we share data to third-parties?
                  </p>
                  <p class="title">
                    No. We do not sell, trade, or transfer any kind of data to any external entities.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="share-off"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    How is my data protected?
                  </p>
                  <p class="title">
                    Any data communicated to or from the site is encrypted using Secure Socket Layer (SSL) technology in transport. <br/>
                    Once transported, data is volume-level encrypted by the cloud-provider <br/>
                    Access to the site is limited, disallowing access to storage.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="email-off"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    How do I unsubscribe?
                  </p>
                  <p class="title">
                    Please do not hesitate to contact us via email, for removing your email address from the notify list.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-content">
              <div class="media">
                <div class="media-left">
                  <b-icon
                    size="is-medium"
                    icon="alert"/>
                </div>
                <div class="media-context">
                  <p class="subtitle is-3">
                    What is there to do if I believe our Privacy Policy has been breached?
                  </p>
                  <p class="title">
                    Please do not hesitate to contact us via email, we will assist you.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
    <ftfooter/>
  </div>
</template>

<script>
import navbar from '@/frontend/components/navbar'
import headerDisplay from '@/frontend/components/header-display'
import ftfooter from '@/frontend/components/footer'

export default {
  name: 'privacy-policy',
  data () {
    return {
      logo: false,
      accountButtons: false,
      navbar: true
    }
  },
  components: {
    headerDisplay,
    navbar,
    ftfooter
  }
}
</script>

<style scoped>

</style>
