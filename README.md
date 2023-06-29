# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [collector.proto](#collector-proto)
    - [Aws](#collector-Aws)
    - [Azure](#collector-Azure)
    - [HostID](#collector-HostID)
    - [K8s](#collector-K8s)
    - [MessageRequest](#collector-MessageRequest)
    - [MessageResult](#collector-MessageResult)
    - [OboeSetting](#collector-OboeSetting)
    - [OboeSetting.ArgumentsEntry](#collector-OboeSetting-ArgumentsEntry)
    - [PingRequest](#collector-PingRequest)
    - [SettingsRequest](#collector-SettingsRequest)
    - [SettingsResult](#collector-SettingsResult)
  
    - [EncodingType](#collector-EncodingType)
    - [HostType](#collector-HostType)
    - [OboeSettingType](#collector-OboeSettingType)
    - [ResultCode](#collector-ResultCode)
  
    - [TraceCollector](#collector-TraceCollector)
  
- [Scalar Value Types](#scalar-value-types)



<a name="collector-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## collector.proto
This file defines the RPC between APM libraries and the collector


<a name="collector-Aws"></a>

### Aws
Represents AWS metadata from [Instance Metadata Service IMDS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cloudProvider | [string](#string) |  | always `aws` |
| cloudPlatform | [string](#string) |  | always `aws_ec2` |
| cloudAccountId | [string](#string) |  | `accountId` in IMDS metadata |
| cloudRegion | [string](#string) |  | `region` in IMDS metadata |
| cloudAvailabilityZone | [string](#string) |  | `availabilityZone` in IMDS metadata |
| hostId | [string](#string) |  | `instanceId` in IMDS metadata |
| hostImageId | [string](#string) |  | `imageId` in IMDS metadata |
| hostName | [string](#string) |  | hostname from either [gethostname() in Linux](https://man7.org/linux/man-pages/man2/gethostname.2.html) or [gethostname() in Windows](https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-gethostname) |
| hostType | [string](#string) |  | `instanceType` in IMDS metadata |






<a name="collector-Azure"></a>

### Azure
Represents Azure metadata from [Instance Metadata Service IMDS](https://learn.microsoft.com/en-us/azure/virtual-machines/instance-metadata-service?tabs=linux)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cloudProvider | [string](#string) |  | always `azure` |
| cloudPlatform | [string](#string) |  | always `azure_vm` |
| cloudRegion | [string](#string) |  | `location` in IMDS metadata |
| cloudAccountId | [string](#string) |  | `subscriptionId` in IMDS metadata |
| hostId | [string](#string) |  | `vmId` in IMDS metadata |
| hostName | [string](#string) |  | `name` in IMDS metadata |
| azureVmName | [string](#string) |  | `name` in IMDS metadata |
| azureVmSize | [string](#string) |  | `vmSize` in IMDS metadata |
| azureVmScaleSetName | [string](#string) |  | `vmScaleSetName` in IMDS metadata |
| azureResourceGroupName | [string](#string) |  | `resourceGroupName` in IMDS metadata |






<a name="collector-HostID"></a>

### HostID
Represents the host metadata needed to infer entity and make correlations from trace telemetry.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hostname | [string](#string) |  | hostname from either [gethostname() in Linux](https://man7.org/linux/man-pages/man2/gethostname.2.html) or [gethostname() in Windows](https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-gethostname). Java agent will not refresh hostname for now to avoid spawning excessive processes. |
| ip_addresses | [string](#string) | repeated | obsolete |
| uuid | [string](#string) |  | [swo only] A random (version 4) UUID generated on application instance startup, analogous to the [OTel Resource attribute](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/README.md#service) `service.instance.id`, e.g. `51fcbc02-670e-454f-84d1-124a500a2646`, `f620e874-ff3d-4f33-825d-cc4b12b2d005` |
| pid | [int32](#int32) |  | process id from either [getpid()](https://man7.org/linux/man-pages/man2/getpid.2.html) or [GetCurrentProcessId()](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid) |
| ec2InstanceID | [string](#string) |  | `instanceId` in AWS EC2 IMDS metadata |
| ec2AvailabilityZone | [string](#string) |  | `availabilityZone` in AWS EC2 IMDS metadata |
| dockerContainerID | [string](#string) |  | container id from `/proc/self/cgroup` for cgroups v1 |
| macAddresses | [string](#string) | repeated | mac addresses for physical interfaces, skipping point-to-point and interfaces w/o IP address from either [getifaddrs(...)](https://man7.org/linux/man-pages/man3/getifaddrs.3.html) or [GetAdaptersAddresses(...)](https://learn.microsoft.com/en-us/windows/win32/api/iphlpapi/nf-iphlpapi-getadaptersaddresses) |
| herokuDynoID | [string](#string) |  | [heroku dyno id](https://devcenter.heroku.com/articles/dynos#local-environment-variables) from environment variable `DYNO` |
| azAppServiceInstanceID | [string](#string) |  | Azure App Service `WEBSITE_INSTANCE_ID` which is the [unique ID of the current VM instance](https://learn.microsoft.com/en-us/azure/app-service/reference-app-settings?tabs=kudu%2Cdotnet#scaling) |
| hostType | [HostType](#collector-HostType) |  | host type struct |
| uamsClientID | [string](#string) |  | [swo only] [uamsclientid](https://swicloud.atlassian.net/wiki/spaces/arch/pages/2963917281/FAS&#43;-&#43;Universal&#43;AMS&#43;Client&#43;-&#43;Unique&#43;Identification) exposed to the APM library for [Service-to-Host correlation](https://swicloud.atlassian.net/wiki/spaces/NIT/pages/2858778658/FAS&#43;Topology&#43;Map&#43;Data&#43;-&#43;UAMS&#43;APM&#43;on-prem#UAMS-Client-Id-management). It is read from `/opt/solarwinds/uamsclient/var/uamsclientid` or `C:\ProgramData\SolarWinds\UAMSClient\uamsclientid`. Windows path may be different depending on the setup. If not found, it is retrieved from `http://127.0.0.1:2113/info/uamsclient` uamsclient_id property |
| awsMetadata | [Aws](#collector-Aws) |  | [swo only] aws ec2 metadata from IMDS |
| azureMetadata | [Azure](#collector-Azure) |  | [swo only] azure metadata from IMDS |
| k8sMetadata | [K8s](#collector-K8s) |  | [swo only] k8s metadata |






<a name="collector-K8s"></a>

### K8s
Represents k8s metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | equivalent to [k8s.namespace.name](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/k8s.md#namespace). The content from `/run/secrets/kubernetes.io/serviceaccount/namespace` |
| podName | [string](#string) |  | equivalent to [k8s.pod.name](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/k8s.md#pod). hostname from `gethostname()` function call |
| podUid | [string](#string) |  | equivalent to [k8s.pod.uid](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/k8s.md#pod). Parsed from `/proc/self/mountinfo` using best effort |
| containerId | [string](#string) |  | id of a running container inside the pod. This field is not set for `cgroup2fs` system. Parsed from `/proc/self/cgroup` using best effort |






<a name="collector-MessageRequest"></a>

### MessageRequest
Represents the message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_key | [string](#string) |  | the Service Key provided by the customer to authenticate and identify the tenant and service the message is destined for. It is a string composed of two parts, an API token and a descriptive service name, separated by a colon `:`. Example: `qwertyuiop1234567:my_cool_service`. |
| messages | [bytes](#bytes) | repeated | bson messages |
| encoding | [EncodingType](#collector-EncodingType) |  | always `EncodingType::BSON` |
| identity | [HostID](#collector-HostID) |  | host id |






<a name="collector-MessageResult"></a>

### MessageResult
Represents the message results


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [ResultCode](#collector-ResultCode) |  | result code from the collector |
| arg | [string](#string) |  | obsolete |
| warning | [string](#string) |  | user-facing warning message. The APM library attempts to squelch repeated warnings, so care should be taken to ensure that warning messages are consistent across all RPCs. |






<a name="collector-OboeSetting"></a>

### OboeSetting
Represents oboe setting message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [OboeSettingType](#collector-OboeSettingType) |  | oboe setting type struct, always DEFAULT_SAMPLE_RATE |
| flags | [bytes](#bytes) |  | flags where { OK=0x0, INVALID=0x1, OVERRIDE=0x2, SAMPLE_START=0x4, SAMPLE_THROUGH=0x8, SAMPLE_THROUGH_ALWAYS=0x10, TRIGGERED_TRACE=0x20 }. e.g. 54 means OK or OVERRIDE or SAMPLE_START or SAMPLE_THROUGH_ALWAYS or TRIGGERED_TRACE |
| timestamp | [int64](#int64) |  | Epoch timestamp |
| value | [int64](#int64) |  | Sampling rate, 1000000 means 100% |
| layer | [bytes](#bytes) |  | layer name, not set since type is always DEFAULT_SAMPLE_RATE |
| arguments | [OboeSetting.ArgumentsEntry](#collector-OboeSetting-ArgumentsEntry) | repeated | key-value pairs. Keys can be [`BucketCapacity`, `BucketRate`, `TriggerRelaxedBucketCapacity`, `TriggerRelaxedBucketRate`, `TriggerStrictBucketCapacity`, `TriggerStrictBucketRate`, `SignatureKey`] |
| ttl | [int64](#int64) |  | time to live for this setting struct |






<a name="collector-OboeSetting-ArgumentsEntry"></a>

### OboeSetting.ArgumentsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="collector-PingRequest"></a>

### PingRequest
Represents the ping request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_key | [string](#string) |  | the Service Key provided by the customer to authenticate and identify the tenant and service the message is destined for. It is a string composed of two parts, an API token and a descriptive service name, separated by a colon `:`. Example: `qwertyuiop1234567:my_cool_service`. |






<a name="collector-SettingsRequest"></a>

### SettingsRequest
Represents the settings request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_key | [string](#string) |  | the Service Key provided by the customer to authenticate and identify the tenant and service the message is destined for. It is a string composed of two parts, an API token and a descriptive service name, separated by a colon `:`. Example: `qwertyuiop1234567:my_cool_service`. |
| identity | [HostID](#collector-HostID) |  | host id, only the `hostname` field needs to be set, all the other fields should be left empty. |
| clientVersion | [string](#string) |  | always `2` |






<a name="collector-SettingsResult"></a>

### SettingsResult
Represents the settings result


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [ResultCode](#collector-ResultCode) |  | result code from the collector |
| arg | [string](#string) |  | obsolete |
| settings | [OboeSetting](#collector-OboeSetting) | repeated | sampling settings |
| warning | [string](#string) |  | user-facing warning message. The APM library attempts to squelch repeated warnings, so care should be taken to ensure that warning messages are consistent across all RPCs. |





 


<a name="collector-EncodingType"></a>

### EncodingType
Represents the encoding type of messages

| Name | Number | Description |
| ---- | ------ | ----------- |
| BSON | 0 | binary JSON |
| PROTOBUF | 1 | obsolete |



<a name="collector-HostType"></a>

### HostType
Represents the host type the APM library is running in.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PERSISTENT | 0 | persistent host type |
| AWS_LAMBDA | 1 | [ao-only] AWS Lambda function |



<a name="collector-OboeSettingType"></a>

### OboeSettingType
Represents oboe setting type

| Name | Number | Description |
| ---- | ------ | ----------- |
| DEFAULT_SAMPLE_RATE | 0 | DEFAULT_SAMPLE_RATE |
| LAYER_SAMPLE_RATE | 1 | obsolete |
| LAYER_APP_SAMPLE_RATE | 2 | obsolete |
| LAYER_HTTPHOST_SAMPLE_RATE | 3 | obsolete |
| CONFIG_STRING | 4 | obsolete |
| CONFIG_INT | 5 | obsolete |



<a name="collector-ResultCode"></a>

### ResultCode
Represents the result code from collector

| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | means OK |
| TRY_LATER | 1 | APM library will retry the request later |
| INVALID_API_KEY | 2 | obsolete, removed handling in this [PR](https://github.com/librato/oboe/pull/720) |
| LIMIT_EXCEEDED | 3 | APM library will retry the request later |
| REDIRECT | 4 | obsolete, removed handling in this [PR](https://github.com/librato/oboe/pull/720) |


 

 


<a name="collector-TraceCollector"></a>

### TraceCollector
Represents the trace collector service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| postEvents | [MessageRequest](#collector-MessageRequest) | [MessageResult](#collector-MessageResult) | post events (traces) to collector. |
| postMetrics | [MessageRequest](#collector-MessageRequest) | [MessageResult](#collector-MessageResult) | post metrics (internal heartbeats, request counters, summary, runtime or custom metrics) to collector |
| postStatus | [MessageRequest](#collector-MessageRequest) | [MessageResult](#collector-MessageResult) | post [__Init](https://github.com/librato/trace/blob/master/docs/specs/KV/init.md) message to collector. May be used by APM library to validate api_key. |
| getSettings | [SettingsRequest](#collector-SettingsRequest) | [SettingsResult](#collector-SettingsResult) | get sampling and other settings for this connection. Note the SettingsRequest requirement for HostID fields. May be used by APM library to validate api_key. |
| ping | [PingRequest](#collector-PingRequest) | [MessageResult](#collector-MessageResult) | ping is used for keep-alive purpose. The APM library is expected to ping the collector if the connection has been idled for 20 seconds (by default). Take note that keep-alive should only be performed if the connection was previously healthy - last API call gave a response |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

