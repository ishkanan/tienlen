<template>
  <div id="app" :class="$style.app">
    <Intro v-if="!connected" />
    <Game v-else />
    <Toaster />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from '@vue/composition-api';
import Game from '~/views/Game.vue';
import Intro from '~/views/Intro.vue';
import Toaster from '~/views/Toaster.vue';
import { game, ConnectionState } from '~/store/game';

export default defineComponent({
  components: { Game, Intro, Toaster },

  setup() {
    const connected = computed(() => game.connState === ConnectionState.Connected);

    return { connected };
  },
});
</script>

<style lang="postcss" module>
.app {
  height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
