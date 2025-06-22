# Reply


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 返信ID | [default to undefined]
**author** | **string** | 投稿者のtraqID | [default to undefined]
**content** | **string** | 返信本文 | [default to undefined]
**imageId** | **string** | 添付画像のID | [default to undefined]
**reactions** | [**Reactions**](Reactions.md) |  | [default to undefined]
**createdAt** | **string** | 作成日時 | [default to undefined]

## Example

```typescript
import { Reply } from './api';

const instance: Reply = {
    id,
    author,
    content,
    imageId,
    reactions,
    createdAt,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
