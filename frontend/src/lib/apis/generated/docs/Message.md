# Message


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | メッセージID | [default to undefined]
**author** | **string** | 投稿者のtraqID | [default to undefined]
**content** | **string** | メッセージ本文 | [default to undefined]
**imageId** | **string** | 添付画像のID | [default to undefined]
**reactions** | [**Reactions**](Reactions.md) |  | [default to undefined]
**replyCount** | **number** | 返信数 | [default to undefined]
**createdAt** | **string** | 作成日時 | [default to undefined]

## Example

```typescript
import { Message } from './api';

const instance: Message = {
    id,
    author,
    content,
    imageId,
    reactions,
    replyCount,
    createdAt,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
