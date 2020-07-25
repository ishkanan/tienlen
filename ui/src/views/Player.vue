<template>
  <div
    v-if="player"
    :class="$style.viewport"
  >
    <div :class="$style.controls">
      <button v-if="canStart" @click="doStart">Start game</button>
      <button v-if="canPlay" @click="doPlay">Play cards</button>
      <button v-if="canPass" class="danger" @click="doPass">Pass turn</button>
      <h2 v-if="autoPassing">Your turn will be automatically passed...</h2>
      <h2 v-if="waiting">Waiting for turn...</h2>
      <h2 v-if="winPlace > 0 && !canStart">All done bucko! Have a break.</h2>
    </div>

    <div :class="$style.hand">
      <div v-if="winPlace > 0" :class="$style.placed">
        <h2 :class="$style.note">{{ ordinalisedWinPlace }}</h2>
      </div>

      <ds-hand
        v-else-if="showHand"
        :cards="hand"
        @selected="onSelected"
      />

      <ds-card
        v-else
        :class="$style.unfaced"
        :card="unfaced"
        :selectable="false"
        :show-face="false"
      />
    </div>

    <div :class="$style.nameBar">
      <ds-block-icon v-if="passed" :class="$style.block"/>
      <h3 :class="{ [$style.isTurn]: player.isTurn }">{{ player.name }}</h3>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import BlockIcon from '~/components/BlockIcon.vue';
import CardView from '~/components/Card.vue';
import Hand from '~/components/Hand.vue';
import { Card, Player, Suit, GameEventKind } from '~/lib/models';
import { ordinalise, startFlashTitle } from '~/lib/utils';
import { requestStartGame, requestTurnPass, requestTurnPlay } from '~/lib/socket';
import { game } from '~/store/game';

interface Data {
  autoPassed: boolean;
  autoPassing: boolean;
  selectedRanks: number[];
}

export default Vue.extend({
  components: {
    'ds-block-icon': BlockIcon,
    'ds-card': CardView,
    'ds-hand': Hand,
  },

  data(): Data {
    return {
      autoPassed: false,
      autoPassing: false,
      selectedRanks: [],
    };
  },

  computed: {
    waiting(): boolean {
      return game.isInProgress &&
        (!this.player?.isTurn ?? false) &&
        this.winPlace === 0;
    },
    winPlace(): number {
      const placed = game.winPlaces.findIndex(p => p.position === game.self?.position);
      return placed === -1 ? 0 : placed + 1;
    },
    ordinalisedWinPlace(): string {
      return ordinalise(this.winPlace);
    },
    canStart(): boolean {
      return !game.isInProgress &&
        game.opponents.length > 0;
    },
    canPass(): boolean {
      return game.isInProgress &&
        !game.isPaused &&
        !game.firstRound &&
        !game.newRound &&
        (this.player?.isTurn ?? false) &&
        !this.autoPassing &&
        !this.autoPassed;
    },
    canPlay(): boolean {
      return game.isInProgress &&
        !game.isPaused &&
        (this.player?.isTurn ?? false) &&
        !this.autoPassing &&
        !this.autoPassed;
    },
    showHand(): boolean {
      return game.isInProgress &&
        this.winPlace === 0;
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
    paused(): boolean {
      return game.isPaused;
    },
    passed(): boolean {
      return this.player?.isPassed ?? false;
    },
    newRound(): boolean {
      return game.newRound;
    },
  },

  watch: {
    canPlay: {
      immediate: true,
      handler(value: boolean) {
        // we auto-skip if we have less cards than other players' last played
        // (since it is impossible to beat with less cards)
        if (this.canPlay && game.lastPlayed.length > this.hand.length && (!this.player?.lastPlayed ?? false)) {
          this.autoPassing = true;
          window.setTimeout(() => {
            this.doPass();
            game.showNotification({
              kind: GameEventKind.Warning,
              message: 'Your turn was auto-passed.',
            });
            this.autoPassing = false;
            this.autoPassed = true;
          }, 2000);
          return;
        }

        const name = this.player?.name || '';
        value && startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★');
      },
    },
    newRound: {
      immediate: true,
      handler(value: boolean) {
        if (value) this.autoPassed = false;
      },
    },
    paused: {
      immediate: true,
      handler(value: boolean) {
        const name = this.player?.name || '';
        !value && this.player?.isTurn && startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★');
      },
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

@media (max-width: 1100px) {
  .viewport {
    max-width: 100%;
    max-height: 0;
    min-height: 100%;
    margin: 0;
  }

  .controls {
    height: 35px;
    margin-top: 0;

    & h1 {
      font-size: 1.5em;
    }
  }

  .hand {
    margin-top: 15px;

    & .placed {
      width: 60px;
      height: 75px;

      & .note {
        margin-top: 25px;
      }
    }

    & .unfaced {
      margin-top: 30px;
    }
  }

  .nameBar {
    & .block {
      width: 20px;
      height: 20px;
    }

    & h3 {
      font-size: 1em;
    }
  }

  .note {
    font-size: 1.1em;
    padding-top: 6px;
  }
}
</style>