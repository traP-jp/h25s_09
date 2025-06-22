# AchievementsApi

All URIs are relative to */api*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**achievementsGet**](#achievementsget) | **GET** /achievements | 実績一覧の取得|
|[**meAchievementsPost**](#meachievementspost) | **POST** /me-achievements | 実績の作成|

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

# **meAchievementsPost**
> Achievement meAchievementsPost(meAchievementsPostRequest)


### Example

```typescript
import {
    AchievementsApi,
    Configuration,
    MeAchievementsPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AchievementsApi(configuration);

let meAchievementsPostRequest: MeAchievementsPostRequest; //

const { status, data } = await apiInstance.meAchievementsPost(
    meAchievementsPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **meAchievementsPostRequest** | **MeAchievementsPostRequest**|  | |


### Return type

**Achievement**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | 実績が正常に作成された |  -  |
|**400** | リクエストが不正 |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

