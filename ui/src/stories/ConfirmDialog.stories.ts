import type { Meta, StoryObj } from '@storybook/vue3'

import ConfirmDialog from '../components/ConfirmDialog.vue'

const meta = {
  title: 'ConfirmDialog',
  component: ConfirmDialog,
  tags: ['autodocs'],
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

export const WithDefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some message for giggles',
  },
};

export const WithNonDefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some message for giggles',
    confirmButtonText: 'Confirm!',
    cancelButtonText: 'Cancel!',
  },
};
