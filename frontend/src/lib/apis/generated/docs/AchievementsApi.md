# AchievementsApi

All URIs are relative to */api*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**achievementsGet**](#achievementsget) | **GET** /achievements | 実績一覧の取得|
|[**tryAchieveIdPost**](#tryachieveidpost) | **POST** /try-achieve/{id} | 確率に応じてイベントが発火するかどうか決める|

# **achievementsGet**
> Array<Achievement> achievementsGet()


### Example

```typescript
import {
    AchievementsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let traqId: string; //特定のユーザーの実績のみを取得 (optional) (default to undefined)

const { status, data } = await apiInstance.achievementsGet(
    traqId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **traqId** | [**string**] | 特定のユーザーの実績のみを取得 | (optional) defaults to undefined|


### Return type

**Array<Achievement>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 実績一覧 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **tryAchieveIdPost**
> TryAchieveIdPost200Response tryAchieveIdPost()


### Example

```typescript
import {
    AchievementsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let id: string; //イベントID (default to undefined)

const { status, data } = await apiInstance.tryAchieveIdPost(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**string**] | イベントID | defaults to undefined|


### Return type

**TryAchieveIdPost200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | イベント処理結果 |  -  |
|**404** | 指定されたIDのイベントが見つからない |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

