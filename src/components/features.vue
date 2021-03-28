<template>
    <div>
        <div class="container">
            <section class="section is-paddingless">
              <div v-for="feature in features" v-bind:key="feature">
                <div class="card pointer-cursor-on-hover">
                  <div class="card-content">
                    <div class="media">
                      <div class="media-left">
                        <b-icon :icon="feature.icon" size="is-medium"></b-icon>
                      </div>
                      <div class="media-content">
                        <p class="title is-3" :class="feature.unavailable ? 'has-text-grey' : ''">{{ feature.name }}</p>
                        <p class="subtitle is-5" :class="feature.unavailable ? 'has-text-grey' : ''">{{ feature.shortDescription }}</p>
                      </div>
                    </div>
                  </div>
                  <div class="content">
                    <div class="notification">
                      <div class="content">
                        <p class="subtitle is-5" :class="feature.unavailable ? 'has-text-grey' : ''">
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
          description: 'Add what you need to your list. Template a new list from last week\'s list. Find the total price of the whole list.',
          unavailable: false
        },
        {
          name: 'Tasks',
          icon: 'clipboard-list',
          shortDescription: 'Sort what needs to be done',
          description: 'Add your daily, weekly, and monthly tasks for your flat. Receive notifications for when you need them done. Automatically rotate between your flatmates.',
          unavailable: true
        },
        {
          name: 'Board',
          icon: 'view-dashboard-outline',
          shortDescription: 'Share ideas and updates',
          description: 'Post photos, updates, ideas, links, and much more.',
          unavailable: true
        },
        {
          name: 'Shared Calendar',
          icon: 'calendar',
          shortDescription: 'Sync calendar events',
          description: 'Keep up-to-date with events that are happening with your flatmates. Integrates with your existing calendar.',
          unavailable: true
        },
        {
          name: 'Contacts',
          icon: 'account-group',
          shortDescription: 'Sync your flatmates accounts as contacts',
          description: 'Get up-to-date contact information for your flatmates. Integrates with your existing contacts app.',
          unavailable: true
        },
        {
          name: 'Recipes',
          icon: 'food',
          shortDescription: 'Create and share recipes',
          description: 'Save recipes to share in the flat. Make it easy for everyone to have the meals that they love, made by anyone.',
          unavailable: true
        },
        {
          name: 'High Fives',
          icon: 'hand',
          shortDescription: 'Celebrate your flatmates',
          description: 'Send a positive message to a flatmate to say thanks for something.',
          unavailable: true
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>
