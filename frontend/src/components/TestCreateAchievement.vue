<template>
  <div class="test-create-achievement">
    <h3>実績作成テスト</h3>
    <div>
      <input
        v-model="achievementName"
        type="text"
        placeholder="実績名を入力"
        @keyup.enter="handleCreate"
      />
      <button @click="handleCreate" :disabled="createMutation.isPending.value">
        {{ createMutation.isPending.value ? '作成中...' : '実績作成' }}
      </button>
    </div>

    <div v-if="createMutation.isError.value" class="error">
      エラー: {{ createMutation.error.value?.message }}
    </div>

    <div v-if="createMutation.isSuccess.value" class="success">
      実績「{{ createMutation.data.value?.name }}」が作成されました！
    </div>
  </div>
</template>

<script setup lang="ts">
import { useCreateAchievement } from '@/lib/composables'
import { ref } from 'vue'

const achievementName = ref('')
const createMutation = useCreateAchievement()

const handleCreate = () => {
  if (achievementName.value.trim()) {
    createMutation.mutate(achievementName.value.trim())
    achievementName.value = ''
  }
}
</script>

<style scoped>
.test-create-achievement {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin: 1rem 0;
}

.test-create-achievement input {
  padding: 0.5rem;
  margin-right: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.test-create-achievement button {
  padding: 0.5rem 1rem;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.test-create-achievement button:disabled {
  background: #6c757d;
  cursor: not-allowed;
}

.error {
  color: #dc3545;
  margin-top: 0.5rem;
}

.success {
  color: #28a745;
  margin-top: 0.5rem;
}
</style>
