<script setup lang="ts">
// ナビゲーションアイテムのデータは、スクリプト内に配列として持ちます。
const navItems = [
  { to: '/', iconName: 'icon-home', label: 'ホーム' },
  { to: '/achievement', iconName: 'icon-achievement', label: 'アチーブメント' },
  { to: '/profile', iconName: 'icon-user', label: 'プロフィール' },
]
</script>

<template>
  <footer :class="$style.bottomBar">
    <router-link
      v-for="item in navItems"
      :key="item.to"
      :to="item.to"
      :class="$style.tabItem"
      active-class="is-active"
    >
      <i :class="[$style.icon, item.iconName]"></i>
      <span :class="$style.label">{{ item.label }}</span>
    </router-link>
  </footer>
</template>

<style lang="scss" module>
/*
 * 全てのスタイルをこのコンポーネント内にまとめて記述します。
 */

/* 親コンテナのスタイル */
.bottomBar {
  /* --- 位置とサイズ --- */
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  box-sizing: border-box;

  /* --- 見た目と中身のレイアウト --- */
  display: flex;
  justify-content: space-around;
  align-items: center;
  background-color: #048004;
  border-top: 1px solid #095709;
  z-index: 1000;

  /* --- 高さ と セーフエリア対応 --- */
  height: 60px;
  padding-bottom: env(safe-area-inset-bottom);
}

.tabItem {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex-grow: 1;
  height: 100%;
  color: #f6f6f6;
  text-decoration: none;
  transition: color 0.2s;

  &.is-active {
    color: #0d39d9; /* アクティブ時の色 */
    font-weight: 600;
  }
}
.icon {
  font-size: 24px;
}

.label {
  font-size: 20px;
  margin-top: 2px;
}
/*
 * 
active-classで指定したクラス名はグローバルなクラスとして扱われるため、
 * CSSモジュール内からスタイルを当てるには :global() で囲みます。
 */
</style>
