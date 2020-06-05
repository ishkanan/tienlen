<template>
  <div :class="[$style.viewport, { [$style.winner]: previousWinner && inLobby }]">
    <h1 v-if="paused" :class="$style.message">
      Game is paused!
    </h1>
    <template v-else-if="inLobby && previousWinner === undefined">
      <h1 v-if="needMorePlayers && !canStart" :class="$style.message">Wait for more players...</h1>
      <h1 v-else-if="needMorePlayers && canStart" :class="$style.message">Wait for more players<br>or start the game...</h1>
      <h1 v-else :class="$style.message">Start the game...</h1>
    </template>
    <template v-else>
      <h1 v-if="inLobby && previousWinner" :class="$style.message">
        {{ previousWinner.name }} won the game!
      </h1>
      <ds-card
        v-for="card in lastPlayed"
        :key="card.globalRank"
        :card="card"
        :selectable="false"
        :show-face="true"
      />
    </template>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { sortBy } from 'lodash-es';
import CardView from '~/components/Card.vue';
import { Card, Player } from '~/lib/models';
import { game } from '~/store/game';

export default Vue.extend({
  components: {
    'ds-card': CardView,
  },

  computed: {
    previousWinner(): Player | undefined {
      return game.opponents.find(o => o.wonLastGame) ??
        (game.self?.wonLastGame ? game.self : undefined);
    },
    canStart(): boolean {
      return this.inLobby && game.opponents.length > 0;
    },
    needMorePlayers(): boolean {
      return game.opponents.length !== 3;
    },
    inLobby(): boolean {
      return game.isInLobby;
    },
    paused(): boolean {
      return game.isPaused;
    },
    lastPlayed(): Card[] {
      return sortBy(game.lastPlayed || [], c => 52 - c.globalRank);
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: hsla(130, 75%, 25%, .80);
  border: black dashed 1px;
  border-radius: 10px;
}

.winner {
  flex-direction: column;
}

.message {
  color: white;
  font-size: 40px;
  text-align: center;
  text-transform: uppercase;
}
</style>