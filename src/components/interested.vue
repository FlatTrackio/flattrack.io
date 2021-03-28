<template>
    <div>
        <headerDisplay/>
        <b-button
          type="is-light"
          size="is-medium"
          @click="PostInterested"
          icon-left="heart">
          I'm interested
        </b-button>
        <br><br>
        <p class="subtitle is-5">Let us know if you're interested and wanna be notified when it's ready!</p>
    </div>
</template>

<script>
import axios from 'axios'
import interested from '@//requests/public/interested'
import { DialogProgrammatic as Dialog, ToastProgrammatic as Toast } from 'buefy'

export default {
  name: 'interested-send',
  methods: {
    PostInterested: () => {
      Dialog.prompt({
        title: 'Ready notification',
        hasIcon: true,
        icon: 'email',
        type: 'is-primary',
        message: 'If you would like to be notified when FlatTrack is generally available, submit below and you will be notified. <br><br> Note: By submitting your email address, you agree to the FlatTrack privacy policy.',
        inputAttrs: {
          placeholder: 'Enter you email address',
          maxlength: 70,
          type: 'email',
          icon: 'email'
        },
        onConfirm: (value) => {
          interested.PostInterested(value).then(resp => {
            Toast.open({
              duration: 8000,
              message: resp.data.metadata.response,
              position: 'is-bottom',
              type: 'is-success',
              size: 'is-medium',
              hasIcon: true
            })
          }).catch(err => {
            Toast.open({
              duration: 8000,
              message: `An error has occurred: ${err.response.metadata.response || err}`,
              position: 'is-bottom',
              type: 'is-danger',
              size: 'is-medium',
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
