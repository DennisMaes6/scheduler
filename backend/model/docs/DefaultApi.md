# \DefaultApi

All URIs are relative to *http://localhost:8080/backend*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbScheduleGet**](DefaultApi.md#DbScheduleGet) | **Get** /db-schedule | Returns the schedule as found in the db.
[**FileScheduleGet**](DefaultApi.md#FileScheduleGet) | **Get** /file-schedule | Returns the schedule as found in a file.
[**InstanceDataGetGet**](DefaultApi.md#InstanceDataGetGet) | **Get** /instance-data/get | Returns the current instance data.
[**InstanceDataSetPost**](DefaultApi.md#InstanceDataSetPost) | **Post** /instance-data/set | Sets the insatnce data in the backend.
[**ModelParametersGetGet**](DefaultApi.md#ModelParametersGetGet) | **Get** /model-parameters/get | Returns the current model parameters.
[**ModelParametersSetPost**](DefaultApi.md#ModelParametersSetPost) | **Post** /model-parameters/set | Sets the model paramters in the backend.
[**ScheduleFileSetPost**](DefaultApi.md#ScheduleFileSetPost) | **Post** /schedule-file/set | Sets the schedule DB file in the backend.
[**ScheduleGenerateGet**](DefaultApi.md#ScheduleGenerateGet) | **Get** /schedule-generate | Returns a schedule generated with Java.
[**ScheduleGet**](DefaultApi.md#ScheduleGet) | **Get** /schedule | Returns a schedule generated with MiniZinc.



## DbScheduleGet

> Schedule DbScheduleGet(ctx).Execute()

Returns the schedule as found in the db.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DbScheduleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DbScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DbScheduleGet`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DbScheduleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDbScheduleGetRequest struct via the builder pattern


### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileScheduleGet

> Schedule FileScheduleGet(ctx).Execute()

Returns the schedule as found in a file.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FileScheduleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FileScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileScheduleGet`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FileScheduleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFileScheduleGetRequest struct via the builder pattern


### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceDataGetGet

> InstanceData InstanceDataGetGet(ctx).Execute()

Returns the current instance data.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.InstanceDataGetGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InstanceDataGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstanceDataGetGet`: InstanceData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InstanceDataGetGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceDataGetGetRequest struct via the builder pattern


### Return type

[**InstanceData**](InstanceData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceDataSetPost

> InstanceDataSetPost(ctx).InstanceData(instanceData).Execute()

Sets the insatnce data in the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    instanceData := *openapiclient.NewInstanceData([]openapiclient.Assistant{*openapiclient.NewAssistant(int32(123), "Janssen J", openapiclient.AssistantType("JA"), []int32{int32(123)})}, []openapiclient.Day{*openapiclient.NewDay(int32(123), *openapiclient.NewDayDate(int32(123), int32(123), int32(123)), false)}) // InstanceData | The instance data to be set.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.InstanceDataSetPost(context.Background()).InstanceData(instanceData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InstanceDataSetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstanceDataSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceData** | [**InstanceData**](InstanceData.md) | The instance data to be set. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelParametersGetGet

> ModelParameters ModelParametersGetGet(ctx).Execute()

Returns the current model parameters.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ModelParametersGetGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModelParametersGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModelParametersGetGet`: ModelParameters
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ModelParametersGetGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiModelParametersGetGetRequest struct via the builder pattern


### Return type

[**ModelParameters**](ModelParameters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelParametersSetPost

> ModelParametersSetPost(ctx).ModelParameters(modelParameters).Execute()

Sets the model paramters in the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    modelParameters := *openapiclient.NewModelParameters() // ModelParameters | The model parameters to be set.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ModelParametersSetPost(context.Background()).ModelParameters(modelParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ModelParametersSetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModelParametersSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelParameters** | [**ModelParameters**](ModelParameters.md) | The model parameters to be set. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleFileSetPost

> ScheduleFileSetPost(ctx).DBFile(dBFile).Execute()

Sets the schedule DB file in the backend.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dBFile := *openapiclient.NewDBFile("Filename_example") // DBFile | The db file to be set.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ScheduleFileSetPost(context.Background()).DBFile(dBFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ScheduleFileSetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleFileSetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dBFile** | [**DBFile**](DBFile.md) | The db file to be set. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleGenerateGet

> Schedule ScheduleGenerateGet(ctx).Execute()

Returns a schedule generated with Java.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ScheduleGenerateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ScheduleGenerateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleGenerateGet`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ScheduleGenerateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleGenerateGetRequest struct via the builder pattern


### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleGet

> Schedule ScheduleGet(ctx).Execute()

Returns a schedule generated with MiniZinc.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ScheduleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleGet`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ScheduleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleGetRequest struct via the builder pattern


### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

