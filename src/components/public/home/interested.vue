<template>
    <div>
        <b-button type="is-light" rounded size="is-large" @click="sendInterested" :label="interestedMessage"></b-button>
        <br><br>
        <p>Let us know if you're interested</p>
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
        message: `Please enter your email, so we can notify you when the first version is ready.`,
        inputAttrs: {
          placeholder: 'xxxxxx@xxxxxxx.xxx',
          maxlength: 70,
          type: 'email'
        },
        onConfirm: (value) => {
          axios.post('/api/interested', {form: {email: value}}).then(resp => {
            this.interestedMessage = `${resp.data.counter} people are interested`
            console.log(resp)
            Notification.open({
              duration: 8000,
              message: `Thank you for being interested, I've notified the developers.`,
              position: 'is-bottom-right',
              hasIcon: true
            })
          }).catch(err => {
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
