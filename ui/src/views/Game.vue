<template>
  <InGameLayout>
    <template #gameTable>
      <OnePlayerLayout v-if="player && opponents.length === 0">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
      </OnePlayerLayout>
      <TwoPlayerLayout v-if="player && opponents.length === 1">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <OpponentView :position="opponents[0].position" />
        </template>
      </TwoPlayerLayout>
      <ThreePlayerLayout v-if="player && opponents.length === 2">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <OpponentView :position="opponents[0].position" />
        </template>
        <template #opponent2>
          <OpponentView :position="opponents[1].position" />
        </template>
      </ThreePlayerLayout>
      <FourPlayerLayout v-if="player && opponents.length === 3">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <OpponentView :position="opponents[0].position" />
        </template>
        <template #opponent2>
          <OpponentView :position="opponents[1].position" />
        </template>
        <template #opponent3>
          <OpponentView :position="opponents[2].position" />
        </template>
      </FourPlayerLayout>
    </template>
    <template #scoreTable>
      <ScoreView />
    </template>
    <template #activityTable>
      <ActivityView />
    </template>
  </InGameLayout>
</template>

<script lang="ts">
import { defineComponent, computed, watch } from '@vue/composition-api';
import InGameLayout from '~/layouts/InGame.vue';
import OnePlayerLayout from '~/layouts/OnePlayerTable.vue';
import TwoPlayerLayout from '~/layouts/TwoPlayerTable.vue';
import ThreePlayerLayout from '~/layouts/ThreePlayerTable.vue';
import FourPlayerLayout from '~/layouts/FourPlayerTable.vue';
import { Player } from '~/lib/models';
import { setTitle } from '~/lib/utils';
import { game } from '~/store/game';
import ActivityView from '~/views/Activity.vue';
import DiscardView from '~/views/Discard.vue';
import OpponentView from '~/views/Opponent.vue';
import PlayerView from '~/views/Player.vue';
import ScoreView from '~/views/Score.vue';

export default defineComponent({
  components: {
    InGameLayout,
    OnePlayerLayout,
    TwoPlayerLayout,
    ThreePlayerLayout,
    FourPlayerLayout,
    ActivityView,
    DiscardView,
    OpponentView,
    PlayerView,
    ScoreView,
  },

  setup() {
    const positionMap: Record<number, number[]> = {
      1: [2, 3, 4],
      2: [3, 4, 1],
      3: [4, 1, 2],
      4: [1, 2, 3],
    };
    const player = computed(() => game.self);
    const opponents = computed<Player[]>(() => {
      if (!player.value) return [];
      const positions = positionMap[player.value.position];
      if (!positions) return [];
      return positions.reduce<Player[]>((memo, pos) => {
        const player = game.opponents.find((p) => p.position === pos);
        if (!player) return memo;
        memo.push(player);
        return memo;
      }, []);
    });

    const canStart = computed(() => !game.isInProgress && game.opponents.length > 0);
    watch(
      canStart,
      (val) => {
        val && setTitle('Tiến lên || in lobby ...');
      },
      { immediate: true },
    );

    const paused = computed(() => game.isPaused);
    watch(
      paused,
      (val) => {
        val && setTitle('Tiến lên || game paused ...');
      },
      { immediate: true },
    );

    return { player, opponents };
  },
});
</script>

<style lang="postcss" module></style>
