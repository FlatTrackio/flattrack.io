<template>
    <div>
        <div class="container">
            <section class="section">
                <h1 class="title is-2">Features</h1>
                <p class="subtitle -is4">What FlatTrack will provide</p>
                <div class="tile is-ancestor" v-for="feature in features" v-bind:key="feature">
                <div class="tile is-parent">
                    <article class="tile is-child box">
                    <p class="title">{{ feature.name }}
                    </p>
                    <p class="subtitle">{{ feature.shortDescription }}</p>
                    <div class="notification">
                        <div class="content">
                            <p>
                                {{ feature.description }}
                            </p>
                        </div>
                    </div>
                    </article>
                </div>
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
          name: 'Tasks',
          shortDescription: 'Self management of tasks',
          description: 'Tasks are assigned and rotated with a frequency of never, daily, weekly, or monthly. Flatmembers are notified about the task assigned to them throughout the cycle. Once the task has been completed, flatmembers mark it as done.'
        },
        {
          name: 'Shopping List',
          shortDescription: 'A list of items for your flat\'s weekly shop',
          description: 'Items are added to the flat-shared shopping list. Items can be marked as weekly, so they populate every week.'
        },
        {
          name: 'Noticeboard',
          shortDescription: 'Post to your flatmates',
          description: 'Post messages onto the flat noticeboard, so your flatmates know what\'s going on.'
        },
        {
          name: 'Shared Calendar',
          shortDescription: 'Keep track of events',
          description: 'Keep up-to-date with events that are happening with your flatmates. Integrate with your existing calendar.'
        },
        {
          name: 'Recipes',
          shortDescription: 'Create and share recipes',
          description: 'Save recipes to share in the flat. Make it easy for everyone to have the meals that they love, made by anyone.'
        },
        {
          name: 'High Fives',
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
