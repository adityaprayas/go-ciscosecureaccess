# VirtualApplianceObjectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalIPs** | Pointer to **[]string** | A list of internal IP addresses. | [optional] 
**ExternalIP** | Pointer to **string** | The external IP address. | [optional] 
**HostType** | Pointer to **string** | The type of the host. | [optional] 
**Uptime** | Pointer to **int64** | The uptime in seconds. | [optional] 
**IsDnscryptEnabled** | Pointer to **bool** | Specifies whether DNSCrypt is enabled. | [optional] 
**Version** | Pointer to **string** | Specifies the version of the virtual appliance. | [optional] 
**UpgradeError** | Pointer to **string** | Specifies the upgrade error. | [optional] 
**Domains** | Pointer to **[]string** | The list of domains. | [optional] 
**LastSyncTime** | Pointer to **time.Time** | The date and time (ISO8601 timestamp) of the last sync. | [optional] 

## Methods

### NewVirtualApplianceObjectSettings

`func NewVirtualApplianceObjectSettings() *VirtualApplianceObjectSettings`

NewVirtualApplianceObjectSettings instantiates a new VirtualApplianceObjectSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualApplianceObjectSettingsWithDefaults

`func NewVirtualApplianceObjectSettingsWithDefaults() *VirtualApplianceObjectSettings`

NewVirtualApplianceObjectSettingsWithDefaults instantiates a new VirtualApplianceObjectSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalIPs

`func (o *VirtualApplianceObjectSettings) GetInternalIPs() []string`

GetInternalIPs returns the InternalIPs field if non-nil, zero value otherwise.

### GetInternalIPsOk

`func (o *VirtualApplianceObjectSettings) GetInternalIPsOk() (*[]string, bool)`

GetInternalIPsOk returns a tuple with the InternalIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIPs

`func (o *VirtualApplianceObjectSettings) SetInternalIPs(v []string)`

SetInternalIPs sets InternalIPs field to given value.

### HasInternalIPs

`func (o *VirtualApplianceObjectSettings) HasInternalIPs() bool`

HasInternalIPs returns a boolean if a field has been set.

### GetExternalIP

`func (o *VirtualApplianceObjectSettings) GetExternalIP() string`

GetExternalIP returns the ExternalIP field if non-nil, zero value otherwise.

### GetExternalIPOk

`func (o *VirtualApplianceObjectSettings) GetExternalIPOk() (*string, bool)`

GetExternalIPOk returns a tuple with the ExternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIP

`func (o *VirtualApplianceObjectSettings) SetExternalIP(v string)`

SetExternalIP sets ExternalIP field to given value.

### HasExternalIP

`func (o *VirtualApplianceObjectSettings) HasExternalIP() bool`

HasExternalIP returns a boolean if a field has been set.

### GetHostType

`func (o *VirtualApplianceObjectSettings) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *VirtualApplianceObjectSettings) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *VirtualApplianceObjectSettings) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *VirtualApplianceObjectSettings) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetUptime

`func (o *VirtualApplianceObjectSettings) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VirtualApplianceObjectSettings) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VirtualApplianceObjectSettings) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VirtualApplianceObjectSettings) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetIsDnscryptEnabled

`func (o *VirtualApplianceObjectSettings) GetIsDnscryptEnabled() bool`

GetIsDnscryptEnabled returns the IsDnscryptEnabled field if non-nil, zero value otherwise.

### GetIsDnscryptEnabledOk

`func (o *VirtualApplianceObjectSettings) GetIsDnscryptEnabledOk() (*bool, bool)`

GetIsDnscryptEnabledOk returns a tuple with the IsDnscryptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDnscryptEnabled

`func (o *VirtualApplianceObjectSettings) SetIsDnscryptEnabled(v bool)`

SetIsDnscryptEnabled sets IsDnscryptEnabled field to given value.

### HasIsDnscryptEnabled

`func (o *VirtualApplianceObjectSettings) HasIsDnscryptEnabled() bool`

HasIsDnscryptEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualApplianceObjectSettings) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualApplianceObjectSettings) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualApplianceObjectSettings) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualApplianceObjectSettings) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpgradeError

`func (o *VirtualApplianceObjectSettings) GetUpgradeError() string`

GetUpgradeError returns the UpgradeError field if non-nil, zero value otherwise.

### GetUpgradeErrorOk

`func (o *VirtualApplianceObjectSettings) GetUpgradeErrorOk() (*string, bool)`

GetUpgradeErrorOk returns a tuple with the UpgradeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeError

`func (o *VirtualApplianceObjectSettings) SetUpgradeError(v string)`

SetUpgradeError sets UpgradeError field to given value.

### HasUpgradeError

`func (o *VirtualApplianceObjectSettings) HasUpgradeError() bool`

HasUpgradeError returns a boolean if a field has been set.

### GetDomains

`func (o *VirtualApplianceObjectSettings) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *VirtualApplianceObjectSettings) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *VirtualApplianceObjectSettings) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *VirtualApplianceObjectSettings) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetLastSyncTime

`func (o *VirtualApplianceObjectSettings) GetLastSyncTime() time.Time`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *VirtualApplianceObjectSettings) GetLastSyncTimeOk() (*time.Time, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *VirtualApplianceObjectSettings) SetLastSyncTime(v time.Time)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *VirtualApplianceObjectSettings) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


