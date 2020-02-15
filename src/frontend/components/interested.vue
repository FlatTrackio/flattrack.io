<template>
    <div>
        <headerDisplay/>
        <b-button type="is-light" rounded size="is-medium" @click="sendInterested" :label="interestedMessage"></b-button>
        <br><br>
        <p class="subtitle is-5">Let us know if you're interested</p>
    </div>
</template>

<script>
import axios from 'axios'
import { DialogProgrammatic as Dialog, NotificationProgrammatic as Notification } from 'buefy'

export default {
  name: 'interest-send',
  data () {
    return {
      interestedMessage: 'I\'m interested'
    }
  },
  methods: {
    sendInterested: () => {
      Dialog.prompt({
        message: `Enter your email, so we can notify you when the first version is ready and when we are ready. <br><br> Note: By submitting your email address, you agree to FlatTrack's privacy policy.`,
        inputAttrs: {
          placeholder: 'xxxxxx@xxxxxxx.xxx',
          maxlength: 70,
          type: 'email'
        },
        onConfirm: (value) => {
          axios.post(`/api/interested?email=${value}`, { form: { email: value } }).then(resp => {
            console.log(resp)
            Notification.open({
              duration: 8000,
              message: `Thank you for being interested! ${resp.data.metadata.message}`,
              position: 'is-bottom-right',
              type: 'is-light',
              hasIcon: true
            })
          }).catch(err => {
            console.log(err)
            Notification.open({
              duration: 8000,
              message: `An error has occured: ${err}`,
              position: 'is-bottom-right',
              type: 'is-danger',
              hasIcon: true
            })
          })
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
