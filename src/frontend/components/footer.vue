<template>
    <div>
        <footer class="footer">
            <div class="content has-text-centered">
                <p>
                    <strong>FlatTrack</strong> and <strong>FlatTrack.io <a :href="`https://gitlab.com/flattrack/flattrack.io/-/tags/${flattrackioVersion}`" target="_blank">({{ flattrackioVersion }})</a></strong> by <a href="https://calebwoodbine.com">Caleb Woodbine</a>.
                    <br>
                    <a href="https://gitlab.com/flattrack/flattrack.io" target="_blank"> Source code </a> - <a href="/privacy-policy">Privacy Policy</a>
                </p>
            </div>
        </footer>
    </div>
</template>

<script>
import apiroot from '@/frontend/requests/public/apiroot'

export default {
  name: 'footer',
  data () {
    return {
      flattrackioVersion: '0.0.0'
    }
  },
  beforeMount () {
    apiroot.GetAPIroot().then(resp => {
      this.flattrackioVersion = resp.data.metadata.version
    })
  }
}
</script>

<style scoped>

</style>
