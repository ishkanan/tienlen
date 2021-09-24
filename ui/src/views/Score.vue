<template>
  <div :class="$style.viewport">
    <div v-for="scoreLine in scores" :key="scoreLine.playerName" :class="$style.scoreLine">
      <h2>{{ scoreLine.playerName }}</h2>
      <h2 v-if="scoreLine.delta > 0">{{ scoreLine.score }} ( + {{ scoreLine.delta }} )</h2>
      <h2 v-else>{{ scoreLine.score }}</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from '@vue/composition-api';
import { orderBy } from 'lodash-es';
import { game } from '~/store/game';

interface ScoreLine {
  playerName: string;
  score: number;
  delta: number;
}

export default defineComponent({
  setup() {
    const deltas = computed(() => {
      if (game.isInProgress) return {};
      return game.winPlaces.reduce<Record<number, number>>((memo, player, i) => {
        memo[player.position] = game.opponents.length - i;
        return memo;
      }, {});
    });

    const scores = computed<ScoreLine[]>(() => {
      const selfScore = {
        playerName: game.self?.name || 'YOU',
        score: game.self?.score || 0,
        delta: (game.self && deltas.value[game.self.position]) || 0,
      };
      const allScores = game.opponents
        .map((o) => ({
          playerName: o.name,
          score: o.score,
          delta: deltas.value[o.position] || 0,
        }))
        .concat(selfScore);
      return orderBy(allScores, ['score', 'playerName'], ['desc', 'asc']);
    });

    return { scores };
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-evenly;

  & .scoreLine {
    width: 90%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  & h2 {
    margin: 0;
  }
}
</style>
