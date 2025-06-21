# MessageDetail


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | メッセージID | [default to undefined]
**author** | **string** | 投稿者のtraqID | [default to undefined]
**content** | **string** | メッセージ本文 | [default to undefined]
**imageId** | **string** | 添付画像のID | [default to undefined]
**reactions** | [**Reactions**](Reactions.md) |  | [default to undefined]
**replies** | [**Array&lt;Reply&gt;**](Reply.md) | 返信一覧 | [default to undefined]
**createdAt** | **string** | 作成日時 | [default to undefined]

## Example

```typescript
import { MessageDetail } from './api';

const instance: MessageDetail = {
    id,
    author,
    content,
    imageId,
    reactions,
    replies,
    createdAt,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
