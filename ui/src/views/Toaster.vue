<template>
  <div slot-scope="{}"/>
</template>

<script lang="ts">
import Vue from 'vue';
import { ToastObject, ToastPosition } from 'vue-toasted';
import { GameEvent, GameEventKind } from '~/lib/models';
import { game } from '~/store/game';

export default Vue.extend({
  data() {
    return {
      windowWidth: window.innerWidth,
      toastOptions: {
        position: 'bottom-right' as ToastPosition,
        duration: window.innerWidth <= 1100 ? 3000 : 5000,
        keepOnHover: true,
        action: {
          text: 'Close',
          onClick: (_: Event, toastObject: ToastObject) => {
            toastObject.goAway(0);
          },
        },
      },
    };
  },

  computed: {
    events(): GameEvent[] {
      return game.events;
    },
  },

  watch: {
    events: {
      immediate: true,
      handler(value: GameEvent[]) {
        if (value.length === 0) return;
        value.forEach(event => {
          Vue.toasted.show(
            event.message,
            {
              ...this.toastOptions,
              type: event.kind === GameEventKind.Error ? 'error' : 'info',
            },
          );
        });
        game.events = [];
      },
    },
  },

  mounted() {
    window.addEventListener('resize', this.handleWindowResize);
  },

  beforeDestroy() {
    window.removeEventListener('resize', this.handleWindowResize);
  },

  methods: {
    handleWindowResize() {
      this.windowWidth = window.innerWidth;
    },
  },
});
</script>

<style lang="postcss" module>
:global(.toasted) {
  font-size: 16px !important;
}

@media (max-width: 1100px) {
  :global(.toasted) {
    font-size: 13px !important;
    padding: 2px 10px;
    margin: 1px !important;
    max-height: 30px;
  }
}
</style>