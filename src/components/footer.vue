<template>
    <div>
        <footer class="footer">
            <div class="content has-text-centered">
                <p>
                    <strong>FlatTrack</strong> and <strong>FlatTrack.io <a :href="`https://gitlab.com/flattrack/flattrack.io/-/tags/${flattrackioVersion}`" target="_blank">({{ flattrackioVersion }})</a></strong> by Caleb Woodbine.
                    <br>
                    <a href="https://gitlab.com/flattrack/flattrack.io" target="_blank"> Source code </a> - <a href="/privacy-policy">Privacy Policy</a>
                </p>
            </div>
        </footer>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'footer',
  data () {
    return {
      flattrackioVersion: ''
    }
  },
  created () {
    axios.get('/api').then(resp => {
      this.flattrackioVersion = resp.data.version
    })
  }
}
</script>

<style scoped>

</style>
