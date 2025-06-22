/**
 * 引数の確率でtrueを返す関数
 * @param probability 0から1の間の確率値（0 = 0%, 1 = 100%）
 * @returns 指定された確率でtrueを返し、それ以外はfalseを返す
 * @throws probability が0未満または1より大きい場合にエラーをスロー
 *
 * @example
 * // 50%の確率でtrueを返す
 * randomBoolean(0.5)
 *
 * // 10%の確率でtrueを返す
 * randomBoolean(0.1)
 *
 * // 100%の確率でtrueを返す（常にtrue）
 * randomBoolean(1.0)
 *
 * // 0%の確率でtrueを返す（常にfalse）
 * randomBoolean(0.0)
 */
export function randomBoolean(probability: number): boolean {
  if (probability < 0 || probability > 1) {
    throw new Error('Probability must be between 0 and 1')
  }

  return Math.random() < probability
}

/**
 * パーセンテージで確率を指定してtrueを返す関数
 * @param percentage 0から100の間のパーセンテージ値
 * @returns 指定されたパーセンテージの確率でtrueを返し、それ以外はfalseを返す
 * @throws percentage が0未満または100より大きい場合にエラーをスロー
 *
 * @example
 * // 50%の確率でtrueを返す
 * randomBooleanByPercent(50)
 *
 * // 10%の確率でtrueを返す
 * randomBooleanByPercent(10)
 */
export function randomBooleanByPercent(percentage: number): boolean {
  if (percentage < 0 || percentage > 100) {
    throw new Error('Percentage must be between 0 and 100')
  }

  return randomBoolean(percentage / 100)
}

/**
 * 指定された配列からランダムに要素を選択する関数
 * @param array 選択元の配列
 * @returns 配列からランダムに選択された要素
 * @throws 空の配列が渡された場合にエラーをスロー
 *
 * @example
 * randomChoice(['red', 'blue', 'green']) // 'red', 'blue', 'green'のいずれかを返す
 * randomChoice([1, 2, 3, 4, 5]) // 1〜5のいずれかを返す
 */
export function randomChoice<T>(array: T[]): T {
  if (array.length === 0) {
    throw new Error('Array cannot be empty')
  }

  const randomIndex = Math.floor(Math.random() * array.length)
  return array[randomIndex]
}

/**
 * 指定された範囲内でランダムな整数を生成する関数
 * @param min 最小値（含む）
 * @param max 最大値（含む）
 * @returns min以上max以下のランダムな整数
 *
 * @example
 * randomInt(1, 6) // サイコロの目（1〜6）
 * randomInt(0, 100) // 0〜100の整数
 */
export function randomInt(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1)) + min
}
