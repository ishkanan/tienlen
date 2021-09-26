<template>
  <div :class="$style.viewport">
    <button @click="doNameInput">Change name</button>
    <button class="danger" @click="doResetConfirm">Reset game</button>
    <InputDialog
      v-if="showInput"
      title="Enter new name"
      :default="playerName"
      @confirm="handleNameInput"
    />
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
import InputDialog from '~/components/InputDialog.vue';
import { requestChangeName, requestResetGame } from '~/lib/socket';
import { game } from '~/store/game';

export default defineComponent({
  components: { ConfirmDialog, InputDialog },

  setup() {
    const playerName = computed(() => game.self?.name ?? '');

    const showInput = ref(false);
    const doNameInput = () => (showInput.value = true);
    const handleNameInput = (confirm: boolean, name: string) => {
      if (confirm) requestChangeName({ name });
      showInput.value = false;
    };

    const showConfirm = ref(false);
    const doResetConfirm = () => (showConfirm.value = true);
    const handleConfirm = (confirm: boolean) => {
      if (confirm) requestResetGame();
      showConfirm.value = false;
    };

    return {
      playerName,
      showInput,
      doNameInput,
      handleNameInput,
      showConfirm,
      doResetConfirm,
      handleConfirm,
    };
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
