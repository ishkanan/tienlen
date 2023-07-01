<script setup lang="ts">
export interface Props {
  title?: string
  message?: string
  confirmButtonText?: string
  cancelButtonText?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Confirm',
  message: '',
  confirmButtonText: 'Yes',
  cancelButtonText: 'No',
})

const emit = defineEmits<{
  (e: 'confirm', value: boolean): void
}>()

const handleClickButton = (confirm: boolean) => {
  emit('confirm', confirm)
}

const handleClickOverlay = () => {
  emit('confirm', false)
}
</script>

<template>
  <transition name="fade">
    <div class="shade" @click="handleClickOverlay">
      <transition name="zoom">
        <div class="dialog">
          <h4 v-if="title">{{ title }}</h4>
          <p v-if="message" class="message">{{ message }}</p>
          <div class="buttonRow">
            <button class="button" @click.stop="() => handleClickButton(true)">
              {{ confirmButtonText }}
            </button>
            <button class="button" @click.stop="() => handleClickButton(false)">
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
  padding: 54px 58px;
  width: 300px;
  border-radius: 8px;
  top: 50%;
  left: 50%;
  transform: translateX(-50%) translateY(-50%);
  position: fixed;
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
  height: 40px;

  a {
    font-style: normal;
    font-weight: 600;
    font-size: 15px;
    line-height: 26px;
    width: 191px;
  }

  &:last-child {
    margin-right: 0;
  }
}
</style>
