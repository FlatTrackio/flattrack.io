<template>
  <div>
    <navbar :logo=logo :accountButtons=accountButtons v-if="navbar"/>
    <headerDisplay/>
    <div class="container">
      <section class="section">
        <nav class="breadcrumb is-medium has-arrow-separator" aria-label="breadcrumbs">
          <ul>
            <li><a href="/"> Home </a></li>
            <li class="is-active"><a href="/privacy-policy"> Privacy Policy </a></li>
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
        <h1 class="title is-3"> What information do we collect? </h1>
        <p class="subtitle is-4"> Email addresses per submission. </p>
        <br>

        <h1 class="title is-3"> Is the collection of data opt-in or opt-out? </h1>
        <p class="subtitle is-4"> Opt-in only. Email addresses are only inputted and not collected without consent. </p>
        <br>

        <h1 class="title is-3"> What will this information be used for? </h1>
        <p class="subtitle is-4"> Sending email notifications once FlatTrack is ready. </p>
        <br>

        <h1 class="title is-3"> Does this site have embedded trackers? </h1>
        <p class="subtitle is-4"> We do not use trackers of any kind. </p>
        <br>

        <h1 class="title is-3"> Do we use 'cookies'? </h1>
        <p class="subtitle is-4"> Cookies are small text files stored in the browser. We do not store any cookies. </p>
        <br>

        <h1 class="title is-3"> Do we share data to third-parties? </h1>
        <p class="subtitle is-4"> No. We do not sell, trade, or transfer any kind of data to any external entities. </p>
        <br>

        <h1 class="title is-3"> How is my data protected? </h1>
        <p class="subtitle is-4"> Any data communicated to or from the site is encrypted using Secure Socket Layer (SSL) technology. <br>
            Access to the site is limited, disallowing access to storage.
        </p>
        <br>

        <h1 class="title is-3"> How do I unsubscribe? </h1>
        <p class="subtitle is-4"> Please do not hesitate to contact us via email, for removing your email address from the notify list. </p>
        <br>

        <h1 class="title is-3"> What is there to do if I believe our Privacy Policy has been breached? </h1>
        <p class="subtitle is-4"> Please do not hesitate to contact us via email, we will assist you. </p>
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
