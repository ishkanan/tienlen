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
      <div :class="$style.messageAndLastPlayed">
        <h2 v-if="inLobby && previousWinner" :class="$style.message">
          {{ previousWinner.name }} won the game!
        </h2>
        <h2 v-else-if="lastPlayedCards.length > 0 && !!lastPlayed">
          {{ lastPlayed.name }} played:
        </h2>
        <div :class="$style.lastPlayed">
          <ds-card
            v-for="card in lastPlayedCards"
            :key="card.globalRank"
            :card="card"
            :selectable="false"
            :show-face="true"
          />
        </div>
      </div>
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
      if (game.self?.wonLastGame) return game.self;
      return game.opponents.find(o => o.wonLastGame);
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
    lastPlayedCards(): Card[] {
      return sortBy(game.lastPlayed || [], c => 52 - c.globalRank);
    },
    lastPlayed(): Player | undefined {
      if (game.self?.lastPlayed) return game.self;
      return game.opponents.find(o => o.lastPlayed);
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
  background-color: rgba(200, 200, 200, 0.4);
  border: black dashed 1px;
  border-radius: 10px;
}

.messageAndLastPlayed {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}

.lastPlayed {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.winner {
  flex-direction: column;
}

.message {
  color: black;
  font-size: 40px;
  text-align: center;
  text-transform: uppercase;
}

@media (max-width: 1100px) {
  .viewport {
    border-radius: 0;
    flex-wrap: wrap;
    margin: 20px 0 0 0;

    & h1 {
      font-size: 1.6em;
    }

    & h2 {
      font-size: 1.3em;
    }
  }
}
</style>