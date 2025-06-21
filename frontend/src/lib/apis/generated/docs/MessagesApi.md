# MessagesApi

All URIs are relative to */api*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**messagesGet**](#messagesget) | **GET** /messages | メッセージ一覧の取得|
|[**messagesIdGet**](#messagesidget) | **GET** /messages/{id} | 指定されたIDのメッセージの詳細を取得|
|[**messagesPost**](#messagespost) | **POST** /messages | メッセージの投稿|

# **messagesGet**
> Array<Message> messagesGet()


### Example

```typescript
import {
    MessagesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let limit: number; //取得する件数の上限 (optional) (default to 20)
let offset: number; //取得開始位置 (optional) (default to 0)
let traqId: string; //特定のユーザーのメッセージのみを取得 (optional) (default to undefined)

const { status, data } = await apiInstance.messagesGet(
    limit,
    offset,
    traqId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | [**number**] | 取得する件数の上限 | (optional) defaults to 20|
| **offset** | [**number**] | 取得開始位置 | (optional) defaults to 0|
| **traqId** | [**string**] | 特定のユーザーのメッセージのみを取得 | (optional) defaults to undefined|


### Return type

**Array<Message>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | メッセージ一覧 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **messagesIdGet**
> MessageDetail messagesIdGet()


### Example

```typescript
import {
    MessagesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let id: string; //メッセージID (default to undefined)

const { status, data } = await apiInstance.messagesIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | メッセージID | defaults to undefined|


### Return type

**MessageDetail**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | メッセージの詳細 |  -  |
|**404** | 指定されたIDのメッセージが見つからない |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **messagesPost**
> MessageDetail messagesPost()


### Example

```typescript
import {
    MessagesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new MessagesApi(configuration);

let message: string; //メッセージ本文 (default to undefined)
let repliesTo: string; //返信先のメッセージID (optional) (default to undefined)
let image: File; //添付画像 (optional) (default to undefined)

const { status, data } = await apiInstance.messagesPost(
    message,
    repliesTo,
    image
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **message** | [**string**] | メッセージ本文 | defaults to undefined|
| **repliesTo** | [**string**] | 返信先のメッセージID | (optional) defaults to undefined|
| **image** | [**File**] | 添付画像 | (optional) defaults to undefined|


### Return type

**MessageDetail**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | メッセージが正常に作成された |  -  |
|**400** | リクエストが不正（メッセージ本文が空、または返信先のメッセージが見つからない） |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

