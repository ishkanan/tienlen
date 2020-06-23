<template>
  <div :class="$style.table">
    <ds-one-player v-if="player && opponents.length === 0">
      <template v-slot:player>
        <ds-player/>
      </template>
      <template v-slot:discard>
        <ds-discard/>
      </template>
      <template v-slot:controls>
        <ds-controls/>
      </template>
    </ds-one-player>
    <ds-two-players v-if="player && opponents.length === 1">
      <template v-slot:player>
        <ds-player/>
      </template>
      <template v-slot:opponent1>
        <ds-opponent :position="opponents[0].position"/>
      </template>
      <template v-slot:discard>
        <ds-discard/>
      </template>
      <template v-slot:controls>
        <ds-controls/>
      </template>
    </ds-two-players>
    <ds-three-players v-if="player && opponents.length === 2">
      <template v-slot:player>
        <ds-player/>
      </template>
      <template v-slot:opponent1>
        <ds-opponent :position="opponents[0].position"/>
      </template>
      <template v-slot:opponent2>
        <ds-opponent :position="opponents[1].position"/>
      </template>
      <template v-slot:discard>
        <ds-discard/>
      </template>
      <template v-slot:controls>
        <ds-controls/>
      </template>
    </ds-three-players>
    <ds-four-players v-if="player && opponents.length === 3">
      <template v-slot:player>
        <ds-player/>
      </template>
      <template v-slot:opponent1>
        <ds-opponent :position="opponents[0].position"/>
      </template>
      <template v-slot:opponent2>
        <ds-opponent :position="opponents[1].position"/>
      </template>
      <template v-slot:opponent3>
        <ds-opponent :position="opponents[2].position"/>
      </template>
      <template v-slot:discard>
        <ds-discard/>
      </template>
      <template v-slot:controls>
        <ds-controls/>
      </template>
    </ds-four-players>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import OnePlayer from '~/layouts/OnePlayer.vue';
import TwoPlayers from '~/layouts/TwoPlayers.vue';
import ThreePlayers from '~/layouts/ThreePlayers.vue';
import FourPlayers from '~/layouts/FourPlayers.vue';
import { Player } from '~/lib/models';
import { setTitle } from '~/lib/utils';
import { game } from '~/store/game';
import ControlsView from '~/views/Controls.vue';
import DiscardView from '~/views/Discard.vue';
import OpponentView from '~/views/Opponent.vue';
import PlayerView from '~/views/Player.vue';

const positionMap: Record<number, number[]> = {
  1: [2,3,4],
  2: [3,4,1],
  3: [4,1,2],
  4: [1,2,3],
};

export default Vue.extend({
  components: {
    'ds-one-player': OnePlayer,
    'ds-two-players': TwoPlayers,
    'ds-three-players': ThreePlayers,
    'ds-four-players': FourPlayers,
    'ds-controls': ControlsView,
    'ds-discard': DiscardView,
    'ds-opponent': OpponentView,
    'ds-player': PlayerView,
  },

  computed: {
    opponents(): Player[] {
      if (!this.player) return [];
      const positions = positionMap[this.player.position];
      if (!positions) return [];
      return positions.reduce<Player[]>((memo, pos) => {
        const player = game.opponents.find(p => p.position === pos);
        if (!player) return memo;
        memo.push(player);
        return memo;
      }, []);
    },
    player(): Player | undefined {
      return game.self;
    },
    canStart(): boolean {
      return !game.isInProgress &&
        game.opponents.length > 0;
    },
    paused(): boolean {
      return game.isPaused;
    },
  },

  watch: {
    canStart: {
      immediate: true,
      handler(value: boolean) {
        value && setTitle('Tiến lên || waiting in lobby ...');
      },
    },
    paused: {
      immediate: true,
      handler(value: boolean) {
        value && setTitle('Tiến lên || game paused ...');
      },
    },
  },
});
</script>

<style lang="postcss" module>
.table {
  height: 860px;
  width: 1600px;
  background-color: rgba(48, 112, 16, 0.7);
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  border: 4mm ridge rgba(170, 50, 50, .6);
  border-radius: 5%;
  padding: 20px;
  margin: auto;
}

@media (max-width: 1100px) {
  .table {
    width: 100%;
    height: 100%;
    max-width: 100%;
    max-height: 100%;
    box-shadow: none;
    border: 0;
    border-radius: 0;
    padding: 0;
  }
}
</style>