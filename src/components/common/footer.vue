<template>
    <div>
        <footer class="footer">
            <div class="content has-text-centered">
                <p>
                    <strong>FlatTrack</strong> and <strong>FlatTrack.io <a :href="`https://gitlab.com/flattrack/flattrack.io/-/tags/${flattrackioVersion}`">({{ flattrackioVersion }})</a></strong> by Caleb Woodbine.
                    <br> The <a href="https://gitlab.com/flattrack/flattrack.io" target="_blank"> source code </a> is licensed
                    <a href="https://www.gnu.org/licenses/gpl-3.0.en.html" target="_blank"> GPL-3.0 </a>.
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
