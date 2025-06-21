<script lang="ts" setup>
import MessageList from '@/components/MessageList.vue'
</script>

<template>
  <div :class="$style.timeline">
    <div :class="$style.container">
      <MessageList :include-replies="false" :use-infinite-scroll="true" />
    </div>
  </div>
</template>

<style lang="scss" module>
.timeline {
  max-width: 800px;
  margin: 0 auto;
  width: 100%;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 0; /* MessageFormがmargin-bottomを持っているため */
}

@media (max-width: 640px) {
  .timeline {
    max-width: 100%;
  }
}
</style>
