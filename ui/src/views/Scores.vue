<template>
  <ds-score-board
    :scores="scores"
    @close="onScoresClose"
  />
</template>

<script lang="ts">
import Vue from 'vue';
import ScoreBoard, { ScoreLine } from '~/components/ScoreBoard.vue';
import { game } from '~/store/game';

export default Vue.extend({
  components: {
    'ds-score-board': ScoreBoard,
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
  },

  methods: {
    onScoresClose() {
      this.$emit('close');
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