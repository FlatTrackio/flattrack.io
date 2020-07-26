<template>
<div>
  <h1 class="title is-1">News</h1>
  <div>
    <div class="card pointer-cursor-on-hover" @click="goToPost">
      <div class="card-content">
        <div class="media">
          <div class="media-left">
            <b-icon icon="newspaper" size="is-medium"></b-icon>
          </div>
          <div class="media-content">
            <b-skeleton active v-if="pageLoading"></b-skeleton>
            <p class="title is-3" v-if="!pageLoading">{{ post.title }}</p>
            <b-skeleton active v-if="pageLoading"></b-skeleton>
            <p class="subtitle is-5" v-if="!pageLoading">By {{ post.creator }} on {{ post.pubDate }}</p>
          </div>
        </div>
      </div>
      <div class="content">
        <div class="notification">
          <div class="content">
            <b-skeleton active v-if="pageLoading"></b-skeleton>
            <p class="subtitle is-5" v-if="!pageLoading">
              {{ post.contentSnippet }}
              <a :href="post.link">Read more</a>
            </p>
            <b-field grouped v-if="!pageLoading">
              <b-tag>
                Categories
              </b-tag>
              <p class="control">
                <b-taglist>
                  <b-tag type="is-info" v-for="tag in post.categories" v-bind:key="tag">{{ tag }}</b-tag>
                </b-taglist>
              </p>
            </b-field>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
import common from '@/frontend/common/common'

export default {
  name: 'news',
  data () {
    return {
      pageLoading: true,
      post: {}
    }
  },
  methods: {
    goToPost () {
      if (typeof this.post.link === 'undefined') {
        return
      }
      window.location.href = this.post.link
    }
  },
  async beforeMount () {
    common.GetLatestRSSPost().then(resp => {
      console.log({ resp })
      this.post = resp
      this.pageLoading = false
    })
  }
}
</script>

<style scoped>

</style>
