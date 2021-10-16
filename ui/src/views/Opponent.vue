<template>
  <div v-if="opponent" :class="$style.viewport">
    <div :class="$style.player">
      <span :class="$style.nameBar">
        <h3 :class="{ [$style.isTurn]: opponent.isTurn }">{{ opponent.name }}</h3>
      </span>

      <div v-if="!opponent.connected" :class="$style.disconnected" />

      <div v-else-if="winPlace > 0" :class="$style.placed">
        <h2 :class="$style.note">{{ ordinalisedWinPlace }}</h2>
      </div>

      <template v-else>
        <CardView :class="$style.card" :card="unfaced" :selectable="false" :show-face="false" />
        <h1 v-if="opponent.cardsLeft > 0" :class="$style.cardsLeft">x {{ opponent.cardsLeft }}</h1>
        <div v-if="passed" :class="$style.passed" />
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, watch } from '@vue/composition-api';
import CardView from '~/components/Card.vue';
import { Suit } from '~/lib/models';
import { ordinalise, setTitle } from '~/lib/utils';
import { game } from '~/store/game';

export default defineComponent({
  components: { CardView },

  props: {
    position: {
      type: Number,
      required: true,
    },
  },

  setup(props) {
    const opponent = computed(() => game.opponents.find((c) => c.position === props.position));
    const unfaced = computed(() => {
      return {
        suit: opponent.value?.isTurn ? Suit.Hearts : Suit.Spades,
        faceValue: 3,
        globalRank: 52,
        suitRank: 13,
      };
    });
    const winPlace = computed(() => {
      const placed = game.winPlaces.findIndex((p) => p.position === props.position);
      return placed === -1 ? 0 : placed + 1;
    });
    const ordinalisedWinPlace = computed(() => ordinalise(winPlace.value));
    const paused = computed(() => game.isPaused);
    const passed = computed(() => opponent.value?.isPassed ?? false);

    watch(
      opponent,
      (val) => {
        const name = val?.name || '';
        val?.isTurn && setTitle(`Tiến lên || ${name}`);
      },
      { immediate: true },
    );

    watch(
      paused,
      (val) => {
        const name = opponent.value?.name || '';
        !val && opponent.value?.isTurn && setTitle(`Tiến lên || ${name}`);
      },
      { immediate: true },
    );

    return {
      opponent,
      ordinalisedWinPlace,
      passed,
      unfaced,
      winPlace,
    };
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

  & .card {
    display: block;
    z-index: 0;
  }

  & .cardsLeft {
    top: -103px;
    width: 80px;
    color: #f2f2f2;
    text-align: center;
    position: relative;
    z-index: 2;
  }

  & .passed {
    top: -201px;
    height: 120px;
    width: 120px;
    background: url(../assets/images/passed.png);
    opacity: 0.7;
    position: relative;
    z-index: 1;
  }

  & .nameBar {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;

    & .isTurn {
      background-color: #f2f2f2;
      border-radius: 5px;
      color: black;
      border: 1px solid black;
      padding: 1px 5px 1px 5px;
    }
  }

  & h3 {
    color: #f2f2f2;
    margin: 10px 0 10px 0;
    padding: 2px 6px 2px 6px;
  }

  & .disconnected {
    width: 80px;
    height: 80px;
    background: url(../assets/images/disconnected.gif);
    background-size: cover;
    background-repeat: no-repeat;
  }

  & .placed {
    width: 100px;
    height: 120px;
    background: url(../assets/images/trophy.png);
    background-size: cover;
    background-repeat: no-repeat;

    & .note {
      background-color: black;
      border-radius: 5px;
      color: #f2f2f2;
      padding: 0px;
      width: 60%;
      margin: 40px auto auto auto;
      text-align: center;
    }
  }
}
</style>
