<template>
  <div
    v-if="opponent"
    :class="$style.viewport"
  >
    <div :class="$style.player">
      <span :class="$style.nameBar">
        <ds-block-icon v-if="passed" :class="$style.block"/>
        <h3 :class="{ [$style.isTurn]: opponent.isTurn }">{{ opponent.name }}</h3>
      </span>
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
import BlockIcon from '~/components/BlockIcon.vue';
import CardView from '~/components/Card.vue';
import { Card, Player, Suit } from '~/lib/models';
import { setTitle } from '~/lib/utils';
import { game } from '~/store/game';

export default Vue.extend({
  components: {
    'ds-block-icon': BlockIcon,
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
    paused(): boolean {
      return game.isPaused;
    },
    passed(): boolean {
      return !!this.opponent && this.opponent.isPassed;
    },
  },

  watch: {
    opponent: {
      immediate: true,
      handler(value: Player | undefined) {
        const name = value?.name || '';
        value?.isTurn && setTitle(`Tiến lên || ${name}`);
      },
    },
    paused: {
      immediate: true,
      handler(value: boolean) {
        const name = this.opponent?.name || '';
        !value && this.opponent?.isTurn && setTitle(`Tiến lên || ${name}`);
      },
    },
  },
});
</script>

<style lang="postcss" module>
.viewport {
  max-width: 100%;
  max-height: 100%;
}

.player {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  & .nameBar {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;

    & .block {
      width: 24px;
      height: 24px;
      fill: red;
      background-color: black;
      border-radius: 12px;
    }

    & .isTurn {
      background-color: #f2f2f2;
      border-radius: 5px;
      color: black;
      border: 3px solid blue;
    }
  }

  & h3 {
    color: #f2f2f2;
    margin: 10px 0 10px 0;
    padding: 2px 6px 2px 6px;
  }

  & .cardsLeft {
    color: #f2f2f2;
    padding-top: 17px;
    text-align: center;
  }

  & .disconnected {
    width: 80px;
    height: 80px;
    background: url(../assets/images/disconnected.png);
    background-size: cover;
    background-repeat: no-repeat;
  }
}

@media (max-width: 1100px) {
  .viewport {
    width: 100%;
    height: 100%;
  }

  .player {
    flex-direction: row;
    justify-content: space-between;

    & .nameBar {
      & .block {
        width: 20px;
        height: 20px;
      }
    }

    & h3 {
      margin: 0;
      font-size: 1em;
    }

    & .cardsLeft {
      font-size: 1.1em;
      padding-top: 6px;
    }

    & .disconnected {
      width: 50px;
      height: 50px;
    }
  }
}
</style>