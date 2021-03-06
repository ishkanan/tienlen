<template>
  <div :class="$style.viewport">
    <button @click="onScoresOpen">Scores</button>
    <ds-score-board
      v-if="scoresVisible"
      :scores="scores"
      @close="onScoresClose"
    />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import ScoreBoard, { ScoreLine } from '~/components/ScoreBoard.vue';
import { GameState } from '~/lib/messages';
import { game } from '~/store/game';

interface Data {
  scoresVisible: boolean;
}

export default Vue.extend({
  components: {
    'ds-score-board': ScoreBoard,
  },

  data(): Data {
    return {
      scoresVisible: false,
    };
  },

  computed: {
    deltas(): Record<number, number> {
      if (game.isInProgress) return {};
      return game.winPlaces.reduce<Record<number, number>>((memo, player, i) => {
        memo[player.position] = game.opponents.length - i;
        return memo;
      }, {});
    },
    scores(): ScoreLine[] {
      const selfScore = {
        playerName: game.self?.name || 'YOU',
        score: game.self?.score || 0,
        delta: (game.self && this.deltas[game.self.position]) || 0,
      };
      return game.opponents.map(o => ({
        playerName: o.name,
        score: o.score,
        delta: this.deltas[o.position] || 0,
      })).concat(selfScore);
    },
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
}
</style>