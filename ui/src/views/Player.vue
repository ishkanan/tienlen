<template>
  <div v-if="player" :class="$style.viewport">
    <div :class="$style.controls">
      <button v-if="canStart" @click="doStart">Start game</button>
      <button v-if="canPlay" :disabled="!cardsSelected" @click="doPlay">Play cards</button>
      <button v-if="canPass" class="danger" @click="doPass">Pass turn</button>
      <h2 v-if="autoPassing">Your turn will be automatically passed...</h2>
      <h2 v-if="waiting">Waiting for turn...</h2>
      <h2 v-if="winPlace > 0 && !canStart">All done bucko! Have a break.</h2>
    </div>

    <div :class="$style.hand">
      <div v-if="winPlace > 0" :class="$style.placed">
        <h2 :class="$style.note">{{ ordinalisedWinPlace }}</h2>
      </div>

      <Hand v-else-if="showHand" :cards="hand" @selected="onSelected" />

      <CardView
        v-else
        :class="$style.unfaced"
        :card="unfaced"
        :selectable="false"
        :show-face="false"
      />
    </div>

    <div :class="$style.nameBar">
      <BlockIcon v-if="passed" :class="$style.block" />
      <h3 :class="{ [$style.isTurn]: player.isTurn }">{{ player.name }}</h3>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref, watch } from '@vue/composition-api';
import BlockIcon from '~/components/BlockIcon.vue';
import CardView from '~/components/Card.vue';
import Hand from '~/components/Hand.vue';
import { Card, Suit, EventSeverity } from '~/lib/models';
import { ordinalise, startFlashTitle } from '~/lib/utils';
import { requestStartGame, requestTurnPass, requestTurnPlay } from '~/lib/socket';
import { game } from '~/store/game';

export default defineComponent({
  components: { BlockIcon, CardView, Hand },

  setup() {
    const autoPassed = ref(false);
    const autoPassing = ref(false);
    const selectedRanks = ref<number[]>([]);

    const player = computed(() => game.self);
    const winPlace = computed(() => {
      const placed = game.winPlaces.findIndex((p) => p.position === game.self?.position);
      return placed === -1 ? 0 : placed + 1;
    });
    const ordinalisedWinPlace = computed(() => ordinalise(winPlace.value));
    const waiting = computed(
      () => game.isInProgress && (!player.value?.isTurn ?? false) && winPlace.value === 0,
    );
    const canStart = computed(() => !game.isInProgress && game.opponents.length > 0);
    const canPass = computed(
      () =>
        game.isInProgress &&
        !game.isPaused &&
        !game.firstRound &&
        !game.newRound &&
        (player.value?.isTurn ?? false) &&
        !autoPassing.value &&
        !autoPassed.value,
    );
    const canPlay = computed(
      () =>
        game.isInProgress &&
        !game.isPaused &&
        (player.value?.isTurn ?? false) &&
        !autoPassing.value &&
        !autoPassed.value,
    );
    const cardsSelected = computed(() => selectedRanks.value.length > 0);
    const showHand = computed(() => game.isInProgress && winPlace.value === 0);
    const hand = computed(() => game.selfHand);
    const unfaced = computed(() => {
      return {
        suit: Suit.Spades,
        faceValue: 2,
        globalRank: 1,
        suitRank: 1,
      };
    });
    const paused = computed(() => game.isPaused);
    const passed = computed(() => player.value?.isPassed ?? false);
    const newRound = computed(() => game.newRound);

    const doStart = () => requestStartGame();
    const doPass = () => requestTurnPass();
    const doPlay = () => {
      const cards = selectedRanks.value.reduce<Card[]>((memo, rank) => {
        const card = game.selfHand.find((c) => c.globalRank === rank);
        if (!card) return memo;
        memo.push(card);
        return memo;
      }, []);
      requestTurnPlay({ cards });
    };
    const onSelected = (ranks: number[]) => (selectedRanks.value = ranks);

    watch(
      canPlay,
      (val) => {
        // we auto-skip if we have less cards than other players' last played
        // (since it is impossible to beat with less cards)
        if (
          canPlay.value &&
          game.lastPlayed.length > hand.value.length &&
          (!player.value?.lastPlayed ?? false)
        ) {
          autoPassing.value = true;
          window.setTimeout(() => {
            doPass();
            game.pushEvent({
              severity: EventSeverity.Warning,
              runes: [
                { message: 'You were auto-passed as you have less cards than the current hand.' },
              ],
            });
            autoPassing.value = false;
            autoPassed.value = true;
          }, 2000);
          return;
        }

        const name = player.value?.name || '';
        val && startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★');
      },
      { immediate: true },
    );

    watch(
      newRound,
      (val) => {
        if (val) autoPassed.value = false;
      },
      { immediate: true },
    );
    watch(
      paused,
      (val) => {
        const name = player.value?.name || '';
        !val &&
          player.value?.isTurn &&
          startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★');
      },
      { immediate: true },
    );

    return {
      autoPassing,
      canPass,
      canPlay,
      cardsSelected,
      canStart,
      doPass,
      doPlay,
      doStart,
      hand,
      onSelected,
      ordinalisedWinPlace,
      passed,
      player,
      showHand,
      unfaced,
      waiting,
      winPlace,
    };
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  max-height: 100%;
  margin: 20px 0 0 0;
}

.controls {
  width: 100%;
  color: white;
  height: 45px;
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  justify-content: center;

  & button:nth-child(2) {
    margin-left: 80px;
  }

  & h1 {
    color: white;
  }
}

.hand {
  width: 100%;
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  justify-content: center;

  & .placed {
    width: 100px;
    height: 120px;
    background: url(../assets/images/trophy.png);
    background-size: cover;
    background-repeat: no-repeat;
    margin-top: 40px;
    text-align: center;

    & .note {
      background-color: black;
      border-radius: 5px;
      color: #f2f2f2;
      padding: 0px;
      width: 60%;
      margin: 40px auto auto auto;
    }
  }

  & .unfaced {
    margin-top: 40px;
  }
}

.nameBar {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  color: #f2f2f2;

  & .block {
    width: 24px;
    height: 24px;
    fill: red;
    background-color: black;
    border-radius: 12px;
  }

  & h3 {
    margin: 10px 0 10px 0;
    padding: 2px 6px 2px 6px;
  }

  & .isTurn {
    background-color: #f2f2f2;
    border-radius: 5px;
    color: black;
    padding: 2px 6px 2px 6px;
    border: 3px solid blue;
  }
}
</style>
