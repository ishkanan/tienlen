<script setup lang="ts">
import { ref, computed } from 'vue'

export interface Props {
  title?: string
  message?: string
  default?: string
  confirmButtonText?: string
  cancelButtonText?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Input',
  message: '',
  default: '',
  confirmButtonText: 'Confirm',
  cancelButtonText: 'Cancel',
})

const emit = defineEmits<{
  (e: 'confirm', confirmed: boolean, data: string): void
}>()

const value = ref(props.default)

const validInput = computed(() => value.value !== '')

const handleClickButton = (confirm: boolean) => {
  emit('confirm', confirm, value.value)
}

const handleClickOverlay = () => {
  emit('confirm', false, value.value)
}
</script>

<template>
  <transition name="fade">
    <div class="shade" @click="handleClickOverlay">
      <transition name="zoom">
        <div class="dialog" @click.stop>
          <h4 v-if="title" class="title is-4">{{ title }}</h4>
          <p v-if="message" class="message">{{ message }}</p>
          <input
            v-model="value"
            class="input"
            type="text"
            maxlength="35"
            placeholder="Enter something..."
          />
          <div class="buttonRow">
            <button
              class="button is-info"
              :disabled="!validInput"
              @click.stop="() => handleClickButton(true)"
            >
              {{ confirmButtonText }}
            </button>
            <button class="button is-danger" @click.stop="() => handleClickButton(false)">
              {{ cancelButtonText }}
            </button>
          </div>
        </div>
      </transition>
    </div>
  </transition>
</template>

<style scoped>
.shade {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.4);
}

.dialog {
  background-color: #fff;
  padding: 20px;
  width: 300px;
  border-radius: 8px;
  top: 50%;
  left: 50%;
  transform: translateX(-50%) translateY(-50%);
  position: fixed;

  & h4 {
    margin-top: 0;
  }

  & p {
    background-color: inherit;
  }
}

.message {
  margin-top: 10px;
}

.buttonRow {
  display: flex;
  flex-flow: row nowrap;
  justify-content: right;
  margin-top: 40px;
}

.button {
  margin-right: 12px;

  &:last-child {
    margin-right: 0;
  }
}

.input {
  width: 100%;
  border-radius: 5px;
  padding: 10px;
  font-size: var(--font-size-s);
  transition: all ease-in-out 0.2s;
}
</style>
