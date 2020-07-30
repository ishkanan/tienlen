<template>
  <div :class="$style.viewport">
    <button @click="onScoresOpen">Scoreboard</button>
    <button class="danger" @click="onLeaveGame">Leave game</button>
    <ds-scores
      v-if="scoresVisible"
      @close="onScoresClose"
    />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { GameState } from '~/lib/messages';
import { requestLeaveGame } from '~/lib/socket';
import { game } from '~/store/game';
import Scores from '~/views/Scores.vue';

interface Data {
  scoresVisible: boolean;
}

export default Vue.extend({
  components: {
    'ds-scores': Scores,
  },

  data(): Data {
    return {
      scoresVisible: false,
    };
  },

  computed: {
    gameState(): GameState {
      return game.gameState;
    },
  },

  watch: {
    gameState: {
      immediate: true,
      handler(value: GameState | undefined, old: GameState | undefined) {
        this.scoresVisible = value === GameState.InLobby && old !== undefined
          ? true
          : this.scoresVisible;
      },
    },
  },

  methods: {
    onLeaveGame() {
      if (window.confirm('Are you sure you want to leave the game?')) {
        requestLeaveGame();
      }
    },
    onScoresOpen() {
      this.scoresVisible = true;
    },
    onScoresClose() {
      this.scoresVisible = false;
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  height: 100%;

  & button:not(:first-child) {
    margin-top: 20px;
  }
}
</style>