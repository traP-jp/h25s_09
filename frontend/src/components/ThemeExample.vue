<template>
  <div :class="$style.example">
    <h2 :class="$style.title">CSS Modules テーマサンプル</h2>

    <div :class="$style.section">
      <h3 :class="$style.sectionTitle">ボタンサンプル</h3>
      <div :class="$style.buttonGroup">
        <button :class="styles.btnPrimary">Primary Button</button>
        <button :class="styles.btnSecondary">Secondary Button</button>
        <button :class="styles.btnOutline">Outline Button</button>
      </div>
    </div>

    <div :class="$style.section">
      <h3 :class="$style.sectionTitle">カードサンプル</h3>
      <div :class="styles.card">
        <p>通常のカード</p>
      </div>
      <div :class="styles.cardElevated">
        <p>エレベーション付きカード</p>
      </div>
    </div>

    <div :class="$style.section">
      <h3 :class="$style.sectionTitle">入力フィールドサンプル</h3>
      <input :class="styles.input" type="text" placeholder="テキストを入力" />
      <textarea :class="styles.input" placeholder="長いテキストを入力" rows="3"></textarea>
    </div>

    <div :class="$style.section">
      <h3 :class="$style.sectionTitle">ユーティリティクラスサンプル</h3>
      <div :class="styles.bgPrimaryLight" style="padding: 1rem; margin: 0.5rem 0">
        <p :class="styles.textPrimary">プライマリテキスト</p>
        <p :class="styles.textSecondary">セカンダリテキスト</p>
        <p :class="styles.textTertiary">ターシャリテキスト</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import styles from '@/assets/styles/components.module.scss'
</script>

<style module lang="scss">
.example {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

.title {
  color: var(--color-text-primary);
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 2rem;
  text-align: center;
}

.section {
  margin-bottom: 3rem;
}

.sectionTitle {
  color: var(--color-text-primary);
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1rem;
  border-bottom: 2px solid var(--color-primary-500);
  padding-bottom: 0.5rem;
}

.buttonGroup {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}
</style>
