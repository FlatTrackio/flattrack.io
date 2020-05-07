<template>
    <div>
        <div class="container">
            <section class="section">
              <div v-for="feature in features" v-bind:key="feature">
                <div class="card pointer-cursor-on-hover">
                  <div class="card-content">
                    <div class="media">
                      <div class="media-left">
                        <b-icon :icon="feature.icon" size="is-medium"></b-icon>
                      </div>
                      <div class="media-content">
                        <p class="title is-3">{{ feature.name }}</p>
                        <p class="subtitle is-5">{{ feature.shortDescription }}</p>
                      </div>
                    </div>
                  </div>
                  <div class="content">
                    <div class="notification">
                      <div class="content">
                        <p class="subtitle is-5">
                          {{ feature.description }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <br />
              </div>
            </section>
        </div>
    </div>
</template>

<script>
export default {
  name: 'features',
  data () {
    return {
      features: [
        {
          name: 'Shopping List',
          icon: 'format-list-checks',
          shortDescription: 'A list of items for your flat\'s weekly shop',
          description: 'Add what you need to your list. Template a new list from last week\'s list. Find the total price of the whole list.'
        },
        {
          name: 'Tasks',
          icon: 'clipboard-list',
          shortDescription: 'Sort what needs to be done',
          description: 'Add your daily, weekly, and monthly tasks for your flat. Receive notifications for when you need them done. Automatically rotate between your flatmates.'
        },
        {
          name: 'Board',
          icon: 'newspaper',
          shortDescription: 'Share ideas and updates',
          description: 'Post photos, updates, ideas, links, and much more.'
        },
        {
          name: 'Shared Calendar',
          icon: 'calendar',
          shortDescription: 'Sync calendar events',
          description: 'Keep up-to-date with events that are happening with your flatmates. Integrates with your existing calendar.'
        },
        {
          name: 'Contacts',
          icon: 'account-group',
          shortDescription: 'Sync your flatmates accounts as contacts',
          description: 'Get up-to-date contact information for your flatmates. Integrates with your existing contacts app.'
        },
        {
          name: 'Recipes',
          icon: 'food',
          shortDescription: 'Create and share recipes',
          description: 'Save recipes to share in the flat. Make it easy for everyone to have the meals that they love, made by anyone.'
        },
        {
          name: 'High Fives',
          icon: 'hand',
          shortDescription: 'Celebrate your flatmates',
          description: 'Send a positive message to a flatmate to say thanks for something.'
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>
