import type { Meta, StoryObj } from '@storybook/vue3'

import ConfirmDialog from '../components/ConfirmDialog.vue'

const meta = {
  title: 'ConfirmDialog',
  component: ConfirmDialog,
  argTypes: {
    title: { control: 'string' },
    message: { control: 'string' },
    confirmButtonText: { control: 'string' },
    cancelButtonText: { control: 'string' },
    onConfirm: { action: 'clicked' },
  },
  args: { },
} satisfies Meta<typeof ConfirmDialog>

export default meta
type Story = StoryObj<typeof meta>

export const DefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some message for giggles',
  },
};

export const NonDefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some message for giggles',
    confirmButtonText: 'Confirm!',
    cancelButtonText: 'Cancel!',
  },
};
