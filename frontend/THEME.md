# カラーテーマドキュメント

このプロジェクトではCSS Modulesベースのカスタムカラーテーマシステムを実装しています。

## 機能

- ライト/ダークモードの切り替え
- CSS変数によるテーマカスタマイズ
- CSS Modulesによるスコープ化されたスタイル
- Vue 3 Composition APIベースのテーマ管理
- 自動システム設定検出
- ローカルストレージによる設定保存

## 使用方法

### 1. テーマ切り替えボタンの追加

```vue
<template>
  <ThemeToggle />
</template>

<script setup lang="ts">
import ThemeToggle from '@/components/ThemeToggle.vue'
</script>
```

### 2. Composition APIでのテーマ管理

```vue
<script setup lang="ts">
import { useTheme } from '@/composables'

const { isDark, currentTheme, setTheme, toggle } = useTheme()

// テーマ切り替え
const handleToggle = () => {
  toggle()
}

// 特定のテーマに設定
const setDarkMode = () => {
  setTheme('dark')
}
</script>
```

### 3. CSS変数の使用（グローバルスタイル）

```scss
.my-component {
  background-color: var(--color-surface);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border-light);
}
```

### 4. CSS Modulesコンポーネントクラスの使用

```vue
<template>
  <div :class="$style.container">
    <button :class="styles.btnPrimary">Primary Button</button>
    <div :class="styles.card">Card Content</div>
    <input :class="styles.input" type="text" placeholder="Enter text" />
  </div>
</template>

<script setup lang="ts">
import styles from '@/assets/styles/components.module.scss'
</script>

<style module lang="scss">
.container {
  padding: 1rem;
  background-color: var(--color-background);
}
</style>
```

### 5. CSS Modulesでの独自スタイル

```vue
<template>
  <div :class="$style.myComponent">Custom Component</div>
</template>

<style module lang="scss">
.myComponent {
  background-color: var(--color-surface);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border-light);
}
</style>
```

## カラーパレット

### Primary Colors (Blue)

- `--color-primary-50` から `--color-primary-900`
- メインのアクセントカラー

### Secondary Colors (Indigo)

- `--color-secondary-50` から `--color-secondary-900`
- セカンダリアクセントカラー

### Accent Colors (Green)

- `--color-accent-50` から `--color-accent-900`
- 成功状態やポジティブなアクション

### Warning Colors (Orange)

- `--color-warning-50` から `--color-warning-900`
- 警告状態

### Error Colors (Red)

- `--color-error-50` から `--color-error-900`
- エラー状態

### Gray Colors

- `--color-gray-50` から `--color-gray-900`
- テキストやボーダーなど

## セマンティックカラー

- `--color-background`: メインの背景色
- `--color-surface`: カードやモーダルの背景色
- `--color-surface-variant`: バリエーション背景色
- `--color-text-primary`: メインテキスト色
- `--color-text-secondary`: セカンダリテキスト色
- `--color-text-tertiary`: 補助テキスト色
- `--color-border-light/medium/strong`: ボーダー色

## コンポーネントクラス

### ボタン

```html
<button class="btn btn-primary">Primary</button>
<button class="btn btn-secondary">Secondary</button>
<button class="btn btn-outline">Outline</button>
```

### カード

```html
<div class="card">Basic Card</div>
<div class="card card-elevated">Elevated Card</div>
```

### 入力フィールド

```html
<input class="input" type="text" placeholder="Enter text" />
<textarea class="input" placeholder="Enter text"></textarea>
```

## ダークモード対応

ダークモードは `data-theme="dark"` 属性によって制御されます。

```scss
// ライトモード専用スタイル
.my-component {
  background-color: white;
}

// ダークモード専用スタイル
[data-theme='dark'] .my-component {
  background-color: var(--color-gray-800);
}
```

## カスタマイゼーション

新しいカラーやコンポーネントスタイルは `src/assets/styles/theme.scss` で定義できます。

### 新しいカラーの追加

```scss
:root {
  --color-custom-50: #f0f9ff;
  --color-custom-500: #0284c7;
  --color-custom-900: #0c4a6e;
}
```

### 新しいコンポーネントスタイルの追加

```scss
.badge {
  display: inline-flex;
  align-items: center;
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  font-weight: 500;
  background-color: var(--color-primary-100);
  color: var(--color-primary-800);
}
```

## TypeScript サポート

テーマ関連の型定義は `src/lib/theme.ts` で提供されています。

```typescript
import { colors, type ThemeMode, type Theme } from '@/lib/theme'

const primaryColor = colors.primary[500] // '#3b82f6'
```

## ベストプラクティス

1. **CSS変数を使用**: ハードコーディングされた色の代わりに CSS変数を使用
2. **セマンティックな命名**: 色の用途に基づいた変数名を使用
3. **ダークモード対応**: 新しいコンポーネントは両方のテーマで確認
4. **アクセシビリティ**: 十分なコントラスト比を維持
5. **一貫性**: 既存のパターンとスタイルガイドに従う

## ファイル構成

```
src/
├── assets/
│   └── styles/
│       ├── variables.scss           # CSS変数定義（グローバル）
│       └── components.module.scss   # CSS Modulesコンポーネント
├── lib/
│   └── theme.ts                     # テーマ関連の型定義とユーティリティ
├── composables/
│   └── theme.ts                     # Vue 3 Composition API
└── components/
    ├── ThemeToggle.vue              # テーマ切り替えボタン
    └── ThemeExample.vue             # 使用例デモ
```

## CSS Modulesのメリット

1. **スコープ化**: クラス名の衝突を防ぐ
2. **型安全**: TypeScriptでのクラス名補完
3. **保守性**: コンポーネントとスタイルの結合度が高い
4. **パフォーマンス**: 使用されていないスタイルの削除

## 利用可能なCSS Modulesクラス

### ユーティリティクラス

- `styles.textPrimary` - プライマリテキスト色
- `styles.textSecondary` - セカンダリテキスト色
- `styles.textTertiary` - ターシャリテキスト色
- `styles.textInverse` - 逆色テキスト

### 背景色クラス

- `styles.bgPrimary` - プライマリ背景色
- `styles.bgSecondary` - セカンダリ背景色
- `styles.bgAccent` - アクセント背景色
- `styles.bgSurface` - サーフェス背景色
- `styles.bgSurfaceVariant` - サーフェスバリアント背景色

### ボーダークラス

- `styles.borderLight` - 薄いボーダー
- `styles.borderMedium` - 中間ボーダー
- `styles.borderStrong` - 濃いボーダー

### コンポーネントクラス

- `styles.btn` - 基本ボタン
- `styles.btnPrimary` - プライマリボタン
- `styles.btnSecondary` - セカンダリボタン
- `styles.btnOutline` - アウトラインボタン
- `styles.card` - 基本カード
- `styles.cardElevated` - エレベーション付きカード
- `styles.input` - 入力フィールド

### アニメーションクラス

- `styles.fadeIn` - フェードインアニメーション
- `styles.slideInUp` - 下から上にスライドイン

## 使用例

### 基本的な使用方法

```vue
<template>
  <div :class="$style.container">
    <h1 :class="$style.title">My Component</h1>
    <button :class="styles.btnPrimary">Click me</button>
    <div :class="styles.card">
      <p :class="styles.textSecondary">Card content</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import styles from '@/assets/styles/components.module.scss'
</script>

<style module lang="scss">
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
}

.title {
  color: var(--color-text-primary);
  font-size: 2rem;
  margin-bottom: 1rem;
}
</style>
```

### 条件付きクラス適用

```vue
<template>
  <button
    :class="[styles.btn, isPrimary ? styles.btnPrimary : styles.btnSecondary, $style.customButton]"
  >
    {{ label }}
  </button>
</template>

<script setup lang="ts">
import styles from '@/assets/styles/components.module.scss'

interface Props {
  label: string
  isPrimary?: boolean
}

defineProps<Props>()
</script>

<style module lang="scss">
.customButton {
  min-width: 120px;

  &:hover {
    transform: translateY(-2px);
  }
}
</style>
```
