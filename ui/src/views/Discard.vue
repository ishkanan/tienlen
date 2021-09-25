<template>
  <div :class="[$style.viewport, { [$style.winner]: previousWinner && inLobby }]">
    <h1 v-if="paused" :class="$style.message">Game is paused!</h1>
    <template v-else-if="inLobby && previousWinner === undefined">
      <h1 v-if="needMorePlayers && !canStart" :class="$style.message">Wait for more players...</h1>
      <h1 v-else-if="needMorePlayers && canStart" :class="$style.message">
        Wait for more players
        <br />
        or start the game...
      </h1>
      <h1 v-else :class="$style.message">Start the game...</h1>
    </template>
    <template v-else>
      <div :class="$style.messageAndLastPlayed">
        <h2 v-if="inLobby && previousWinner" :class="$style.message">
          The game has finished! One more?
        </h2>
        <h2 v-else-if="lastPlayedCards.length > 0 && !!lastPlayed">
          {{ lastPlayed.name }} played:
        </h2>
        <div :class="$style.lastPlayed">
          <CardView
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
import { defineComponent, computed } from '@vue/composition-api';
import { sortBy } from 'lodash-es';
import CardView from '~/components/Card.vue';
import { game } from '~/store/game';

export default defineComponent({
  components: { CardView },

  setup() {
    const canStart = computed(() => inLobby.value && game.opponents.length > 0);
    const inLobby = computed(() => game.isInLobby);
    const lastPlayed = computed(() => {
      if (game.self?.lastPlayed) return game.self;
      return game.opponents.find((o) => o.lastPlayed);
    });
    const lastPlayedCards = computed(() => sortBy(game.lastPlayed || [], (c) => 52 - c.globalRank));
    const needMorePlayers = computed(() => game.opponents.length !== 3);
    const paused = computed(() => game.isPaused);
    const previousWinner = computed(() => {
      if (game.self?.wonLastGame) return game.self;
      return game.opponents.find((o) => o.wonLastGame);
    });

    return {
      canStart,
      inLobby,
      lastPlayed,
      lastPlayedCards,
      needMorePlayers,
      paused,
      previousWinner,
    };
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
</style>
