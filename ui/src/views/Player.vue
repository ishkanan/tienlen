<template>
  <div
    v-if="player"
    :class="$style.viewport"
  >
    <div :class="$style.controls">
      <button v-if="canStart" @click="doStart">Start game</button>
      <button v-if="canPlay" @click="doPlay">Play cards!</button>
      <button v-if="canPass" @click="doPass">Pass turn!</button>
    </div>

    <div :class="$style.hand">
      <ds-hand
        v-if="showHand"
        :cards="hand"
        @selected="onSelected"
      />
      <ds-card
        v-else
        :card="unfaced"
        :selectable="false"
        :show-face="false"
      />
    </div>

    <div :class="$style.playerName">
      <h3 :class="{ [$style.isTurn]: player.isTurn }">{{ player.name }}</h3>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import CardView from '~/components/Card.vue';
import Hand from '~/components/Hand.vue';
import { Card, Player, Suit } from '~/lib/models';
import { requestStartGame, requestTurnPass, requestTurnPlay } from '~/lib/socket';
import { game } from '~/store/game';

interface Data {
  selectedRanks: number[];
}

export default Vue.extend({
  components: {
    'ds-card': CardView,
    'ds-hand': Hand,
  },

  data(): Data {
    return {
      selectedRanks: [],
    };
  },

  computed: {
    canStart(): boolean {
      return !game.isInProgress;
    },
    canPass(): boolean {
      return game.isInProgress &&
        !game.firstRound &&
        !game.newRound &&
        !!this.player &&
        this.player.isTurn;
    },
    canPlay(): boolean {
      return game.isInProgress &&
        !!this.player &&
        this.player.isTurn;
    },
    showHand(): boolean {
      return game.isInProgress;
    },
    player(): Player | undefined {
      return game.self;
    },
    hand(): Card[] {
      return game.selfHand;
    },
    unfaced(): Card {
      return {
        suit: Suit.Spades,
        faceValue: 2,
        globalRank: 1,
        suitRank: 1,
      };
    },
  },

  methods: {
    doStart() {
      requestStartGame();
    },
    doPass() {
      requestTurnPass();
    },
    doPlay() {
      const cards = this.selectedRanks.reduce<Card[]>((memo, rank) => {
        const card = game.selfHand.find(c => c.globalRank === rank);
        if (!card) return memo;
        memo.push(card);
        return memo;
      }, []);
      requestTurnPlay({ cards });
    },
    onSelected(ranks: number[]) {
      this.selectedRanks = ranks;
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  max-height: 100%;
  margin: 20px 0 20 0;
}

.controls {
  width: 100%;
  height: 45px;
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  justify-content: center;

  & button:nth-child(2) {
    margin-left: 40px;
  }
}

.hand {
  width: 100%;
  margin-top: 30px;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.playerName {
  display: flex;
  flex-direction: row;
  justify-content: center;
  color: #f2f2f2;

  & h3 {
    margin-top: 10px;
  }
}

.isTurn {
  background-color: #f2f2f2;
  border-radius: 5px;
  color: black;
  padding: 2px 6px 2px 6px;
  border: 3px solid blue;
}
</style>