<template>
  <div :class="$style.viewport">
    <div :class="$style.board">
      <div
        v-for="scoreLine in sortedScores"
        :key="scoreLine.playerName"
        :class="$style.scoreLine"
      >
        <h2>{{ scoreLine.playerName }}</h2>
        <h2 v-if="scoreLine.delta > 0">{{ scoreLine.score }} (+{{ scoreLine.delta }})</h2>
        <h2 v-else>{{ scoreLine.score }}</h2>
      </div>
      <button @click="onClose">Close</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';
import { sortBy } from 'lodash-es';

export interface ScoreLine {
  playerName: string;
  score: number;
  delta: number;
}

export default Vue.extend({
  props: {
    scores: {
      type: Array as PropType<ScoreLine[]>,
      required: true,
    },
  },

  computed: {
    sortedScores(): ScoreLine[] {
      return sortBy(this.scores, s => s.score * -1);
    },
  },

  methods: {
    onClose() {
      this.$emit('close');
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  background-color: rgba(0, 0, 0, .8);
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 5000;
}

.board {
  width: 400px;
  padding: 30px;
  margin: 100px auto;
  background-color: rgba(255, 255, 255, 0.85);
  border: 4px blue double;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;

  & .scoreLine {
    width: 90%;
    margin-bottom: 30px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  & h2 {
    margin: 0;
  }
}

@media (max-width: 1100px) {
  .board {
    width: 90%;
    max-width: 90%;
    padding: 30px 10px 30px 10px;
    margin: 80px auto;
    border-radius: 0;
  }
}
</style>