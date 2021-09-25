<template>
  <div :class="$style.viewport">
    <button v-if="isHost" class="danger" @click="doResetConfirm">Reset game</button>
    <ConfirmDialog
      v-if="showConfirm"
      title="Are you sure?"
      message="This will reset EVERYTHING and remove all disconnected players!"
      @confirm="handleConfirm"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from '@vue/composition-api';
import ConfirmDialog from '~/components/ConfirmDialog.vue';
import { requestResetGame } from '~/lib/socket';
import { game } from '~/store/game';

export default defineComponent({
  components: { ConfirmDialog },

  setup() {
    const showConfirm = ref(false);
    const isHost = computed(() => game.self?.position === 1);

    const doResetConfirm = () => (showConfirm.value = true);
    const handleConfirm = (confirm: boolean) => {
      if (confirm) requestResetGame();
      showConfirm.value = false;
    };

    return { isHost, showConfirm, doResetConfirm, handleConfirm };
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-evenly;
}
</style>
