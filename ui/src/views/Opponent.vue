<template>
  <div
    v-if="opponent"
    :class="$style.viewport"
  >
    <div :class="$style.nameAndCard">
      <h3 :class="{ [$style.isTurn]: opponent.isTurn }">{{ opponent.name }}</h3>
      <div v-if="!opponent.connected" :class="$style.disconnected"/>
      <ds-card
        v-else
        :card="unfaced"
        :selectable="false"
        :show-face="false"
      >
        <template v-slot:info>
          <h1 v-if="opponent.cardsLeft > 0" :class="$style.cardsLeft">
            x {{ opponent.cardsLeft }}
          </h1>
        </template>
      </ds-card>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import CardView from '~/components/Card.vue';
import { Card, Player, Suit } from '~/lib/models';
import { game } from '~/store/game';

export default Vue.extend({
  components: {
    'ds-card': CardView,
  },

  props: {
    position: {
      type: Number,
      required: true,
    },
  },

  computed: {
    unfaced(): Card {
      return {
        suit: this.opponent?.isTurn ? Suit.Hearts : Suit.Spades,
        faceValue: 3,
        globalRank: 52,
        suitRank: 13,
      };
    },
    opponent(): Player | undefined {
      return game.opponents.find(c => c.position === this.position);
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  max-width: 100%;
  max-height: 100%;
}

.nameAndCard {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  & h3 {
    color: #f2f2f2;
    margin-bottom: 10px;
  }

  & .cardsLeft {
    color: #f2f2f2;
    padding-top: 17px;
    text-align: center;
  }

  & .isTurn {
    background-color: #f2f2f2;
    border-radius: 5px;
    color: black;
    padding: 2px 6px 2px 6px;
    border: 3px solid blue;
  }

  & .disconnected {
    width: 80px;
    height: 90px;
    background: url(../assets/images/disconnected.png);
    background-size: contain;
    background-repeat: no-repeat;
  }
}
</style>