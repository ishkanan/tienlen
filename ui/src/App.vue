<template>
  <div id="app" :class="$style.app">
    <ds-intro v-if="!connected"/>
    <ds-game-table v-else/>
    <ds-toaster/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import Toaster from '~/components/Toaster.vue';
import GameTable from '~/views/GameTable.vue';
import Intro from '~/views/Intro.vue';
import { game, ConnectionState } from '~/store/game';

export default Vue.extend({
  components: {
    'ds-game-table': GameTable,
    'ds-intro': Intro,
    'ds-toaster': Toaster,
  },

  computed: {
    connected(): boolean {
      return game.connState === ConnectionState.Connected;
    },
  },
});
</script>

<style lang="postcss" module>
.app {
  height: 100vh;
  width: 100vw;
  display: flex;
  margin: auto;
  align-items: center;
  justify-content: center;
}
</style>