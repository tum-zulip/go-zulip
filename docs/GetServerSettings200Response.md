# GetServerSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**IgnoredParametersUnsupported** | Pointer to **interface{}** |  | [optional] 
**AuthenticationMethods** | Pointer to [**AuthMethods**](AuthMethods.md) |  | [optional] 
**ExternalAuthenticationMethods** | Pointer to [**[]ExternalAuthMethod**](ExternalAuthMethod.md) | A list of dictionaries describing the available external authentication methods (E.g. Google, GitHub, or SAML) enabled for this organization.  The list is sorted in the order in which these authentication methods should be displayed.  **Changes**: New in Zulip 2.1.0.  | [optional] 
**ZulipFeatureLevel** | Pointer to **int32** | An integer indicating what features are available on the server. The feature level increases monotonically; a value of N means the server supports all API features introduced before feature level N. This is designed to provide a simple way for client apps to decide whether the server supports a given feature or API change. See the [changelog](/api/changelog) for details on what each feature level means.  **Changes**: New in Zulip 3.0 (feature level 1). We recommend using an implied value of 0 for Zulip servers that do not send this field.  | [optional] 
**ZulipVersion** | Pointer to **string** | The server&#39;s version number. This is often a release version number, like &#x60;2.1.7&#x60;. But for a server running a [version from Git][git-release], it will be a Git reference to the commit, like &#x60;5.0-dev-1650-gc3fd37755f&#x60;.  [git-release]: https://zulip.readthedocs.io/en/latest/overview/release-lifecycle.html#git-versions  | [optional] 
**ZulipMergeBase** | Pointer to **string** | The &#x60;git merge-base&#x60; between &#x60;zulip_version&#x60; and official branches in the public [Zulip server and web app repository](https://github.com/zulip/zulip), in the same format as &#x60;zulip_version&#x60;. This will equal &#x60;zulip_version&#x60; if the server is not running a fork of the Zulip server.  This will be &#x60;\&quot;\&quot;&#x60; if unavailable.  **Changes**: New in Zulip 5.0 (feature level 88).  | [optional] 
**PushNotificationsEnabled** | Pointer to **bool** | Whether mobile/push notifications are configured.  | [optional] 
**IsIncompatible** | Pointer to **bool** | Whether the Zulip client that has sent a request to this endpoint is deemed incompatible with the server.  | [optional] 
**EmailAuthEnabled** | Pointer to **bool** | Setting for allowing users authenticate with an email-password combination.  | [optional] 
**RequireEmailFormatUsernames** | Pointer to **bool** | Whether all valid usernames for authentication to this organization will be email addresses. This is important for clients to know whether to do client side validation of email address format in a login prompt.  This value will be false if the server has [LDAP authentication][ldap-auth] enabled with a username and password combination.  [ldap-auth]: https://zulip.readthedocs.io/en/latest/production/authentication-methods.html#ldap-including-active-directory  | [optional] 
**RealmUri** | Pointer to **string** | The organization&#39;s canonical URL. Alias of &#x60;realm_url&#x60;.  **Changes**: Deprecated in Zulip 9.0 (feature level 257). The term \&quot;URI\&quot; is deprecated in [web standards](https://url.spec.whatwg.org/#goals).  | [optional] 
**RealmUrl** | Pointer to **string** | The organization&#39;s canonical URL.  **Changes**: New in Zulip 9.0 (feature level 257), replacing the deprecated &#x60;realm_uri&#x60;.  | [optional] 
**RealmName** | Pointer to **string** | The organization&#39;s name (for display purposes).  | [optional] 
**RealmIcon** | Pointer to **string** | The URL for the organization&#39;s logo formatted as a square image, used for identifying the organization in small locations in the mobile and desktop apps.  | [optional] 
**RealmDescription** | Pointer to **string** | HTML description of the organization, as configured by the [organization profile](/help/create-your-organization-profile).  | [optional] 
**RealmWebPublicAccessEnabled** | Pointer to **bool** | Whether the organization has enabled the creation of [web-public channels](/help/public-access-option) and at least one web-public channel on the server currently exists. Clients that support viewing content in web-public channels without an account can use this to determine whether to offer that feature on the login page for an organization.  **Changes**: New in Zulip 5.0 (feature level 116).  | [optional] 

## Methods

### NewGetServerSettings200Response

`func NewGetServerSettings200Response(result interface{}, msg interface{}, ) *GetServerSettings200Response`

NewGetServerSettings200Response instantiates a new GetServerSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerSettings200ResponseWithDefaults

`func NewGetServerSettings200ResponseWithDefaults() *GetServerSettings200Response`

NewGetServerSettings200ResponseWithDefaults instantiates a new GetServerSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetServerSettings200Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetServerSettings200Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetServerSettings200Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *GetServerSettings200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetServerSettings200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *GetServerSettings200Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetServerSettings200Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetServerSettings200Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *GetServerSettings200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *GetServerSettings200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetIgnoredParametersUnsupported

`func (o *GetServerSettings200Response) GetIgnoredParametersUnsupported() interface{}`

GetIgnoredParametersUnsupported returns the IgnoredParametersUnsupported field if non-nil, zero value otherwise.

### GetIgnoredParametersUnsupportedOk

`func (o *GetServerSettings200Response) GetIgnoredParametersUnsupportedOk() (*interface{}, bool)`

GetIgnoredParametersUnsupportedOk returns a tuple with the IgnoredParametersUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredParametersUnsupported

`func (o *GetServerSettings200Response) SetIgnoredParametersUnsupported(v interface{})`

SetIgnoredParametersUnsupported sets IgnoredParametersUnsupported field to given value.

### HasIgnoredParametersUnsupported

`func (o *GetServerSettings200Response) HasIgnoredParametersUnsupported() bool`

HasIgnoredParametersUnsupported returns a boolean if a field has been set.

### SetIgnoredParametersUnsupportedNil

`func (o *GetServerSettings200Response) SetIgnoredParametersUnsupportedNil(b bool)`

 SetIgnoredParametersUnsupportedNil sets the value for IgnoredParametersUnsupported to be an explicit nil

### UnsetIgnoredParametersUnsupported
`func (o *GetServerSettings200Response) UnsetIgnoredParametersUnsupported()`

UnsetIgnoredParametersUnsupported ensures that no value is present for IgnoredParametersUnsupported, not even an explicit nil
### GetAuthenticationMethods

`func (o *GetServerSettings200Response) GetAuthenticationMethods() AuthMethods`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *GetServerSettings200Response) GetAuthenticationMethodsOk() (*AuthMethods, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *GetServerSettings200Response) SetAuthenticationMethods(v AuthMethods)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *GetServerSettings200Response) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetExternalAuthenticationMethods

`func (o *GetServerSettings200Response) GetExternalAuthenticationMethods() []ExternalAuthMethod`

GetExternalAuthenticationMethods returns the ExternalAuthenticationMethods field if non-nil, zero value otherwise.

### GetExternalAuthenticationMethodsOk

`func (o *GetServerSettings200Response) GetExternalAuthenticationMethodsOk() (*[]ExternalAuthMethod, bool)`

GetExternalAuthenticationMethodsOk returns a tuple with the ExternalAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuthenticationMethods

`func (o *GetServerSettings200Response) SetExternalAuthenticationMethods(v []ExternalAuthMethod)`

SetExternalAuthenticationMethods sets ExternalAuthenticationMethods field to given value.

### HasExternalAuthenticationMethods

`func (o *GetServerSettings200Response) HasExternalAuthenticationMethods() bool`

HasExternalAuthenticationMethods returns a boolean if a field has been set.

### GetZulipFeatureLevel

`func (o *GetServerSettings200Response) GetZulipFeatureLevel() int32`

GetZulipFeatureLevel returns the ZulipFeatureLevel field if non-nil, zero value otherwise.

### GetZulipFeatureLevelOk

`func (o *GetServerSettings200Response) GetZulipFeatureLevelOk() (*int32, bool)`

GetZulipFeatureLevelOk returns a tuple with the ZulipFeatureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipFeatureLevel

`func (o *GetServerSettings200Response) SetZulipFeatureLevel(v int32)`

SetZulipFeatureLevel sets ZulipFeatureLevel field to given value.

### HasZulipFeatureLevel

`func (o *GetServerSettings200Response) HasZulipFeatureLevel() bool`

HasZulipFeatureLevel returns a boolean if a field has been set.

### GetZulipVersion

`func (o *GetServerSettings200Response) GetZulipVersion() string`

GetZulipVersion returns the ZulipVersion field if non-nil, zero value otherwise.

### GetZulipVersionOk

`func (o *GetServerSettings200Response) GetZulipVersionOk() (*string, bool)`

GetZulipVersionOk returns a tuple with the ZulipVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipVersion

`func (o *GetServerSettings200Response) SetZulipVersion(v string)`

SetZulipVersion sets ZulipVersion field to given value.

### HasZulipVersion

`func (o *GetServerSettings200Response) HasZulipVersion() bool`

HasZulipVersion returns a boolean if a field has been set.

### GetZulipMergeBase

`func (o *GetServerSettings200Response) GetZulipMergeBase() string`

GetZulipMergeBase returns the ZulipMergeBase field if non-nil, zero value otherwise.

### GetZulipMergeBaseOk

`func (o *GetServerSettings200Response) GetZulipMergeBaseOk() (*string, bool)`

GetZulipMergeBaseOk returns a tuple with the ZulipMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZulipMergeBase

`func (o *GetServerSettings200Response) SetZulipMergeBase(v string)`

SetZulipMergeBase sets ZulipMergeBase field to given value.

### HasZulipMergeBase

`func (o *GetServerSettings200Response) HasZulipMergeBase() bool`

HasZulipMergeBase returns a boolean if a field has been set.

### GetPushNotificationsEnabled

`func (o *GetServerSettings200Response) GetPushNotificationsEnabled() bool`

GetPushNotificationsEnabled returns the PushNotificationsEnabled field if non-nil, zero value otherwise.

### GetPushNotificationsEnabledOk

`func (o *GetServerSettings200Response) GetPushNotificationsEnabledOk() (*bool, bool)`

GetPushNotificationsEnabledOk returns a tuple with the PushNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotificationsEnabled

`func (o *GetServerSettings200Response) SetPushNotificationsEnabled(v bool)`

SetPushNotificationsEnabled sets PushNotificationsEnabled field to given value.

### HasPushNotificationsEnabled

`func (o *GetServerSettings200Response) HasPushNotificationsEnabled() bool`

HasPushNotificationsEnabled returns a boolean if a field has been set.

### GetIsIncompatible

`func (o *GetServerSettings200Response) GetIsIncompatible() bool`

GetIsIncompatible returns the IsIncompatible field if non-nil, zero value otherwise.

### GetIsIncompatibleOk

`func (o *GetServerSettings200Response) GetIsIncompatibleOk() (*bool, bool)`

GetIsIncompatibleOk returns a tuple with the IsIncompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncompatible

`func (o *GetServerSettings200Response) SetIsIncompatible(v bool)`

SetIsIncompatible sets IsIncompatible field to given value.

### HasIsIncompatible

`func (o *GetServerSettings200Response) HasIsIncompatible() bool`

HasIsIncompatible returns a boolean if a field has been set.

### GetEmailAuthEnabled

`func (o *GetServerSettings200Response) GetEmailAuthEnabled() bool`

GetEmailAuthEnabled returns the EmailAuthEnabled field if non-nil, zero value otherwise.

### GetEmailAuthEnabledOk

`func (o *GetServerSettings200Response) GetEmailAuthEnabledOk() (*bool, bool)`

GetEmailAuthEnabledOk returns a tuple with the EmailAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAuthEnabled

`func (o *GetServerSettings200Response) SetEmailAuthEnabled(v bool)`

SetEmailAuthEnabled sets EmailAuthEnabled field to given value.

### HasEmailAuthEnabled

`func (o *GetServerSettings200Response) HasEmailAuthEnabled() bool`

HasEmailAuthEnabled returns a boolean if a field has been set.

### GetRequireEmailFormatUsernames

`func (o *GetServerSettings200Response) GetRequireEmailFormatUsernames() bool`

GetRequireEmailFormatUsernames returns the RequireEmailFormatUsernames field if non-nil, zero value otherwise.

### GetRequireEmailFormatUsernamesOk

`func (o *GetServerSettings200Response) GetRequireEmailFormatUsernamesOk() (*bool, bool)`

GetRequireEmailFormatUsernamesOk returns a tuple with the RequireEmailFormatUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEmailFormatUsernames

`func (o *GetServerSettings200Response) SetRequireEmailFormatUsernames(v bool)`

SetRequireEmailFormatUsernames sets RequireEmailFormatUsernames field to given value.

### HasRequireEmailFormatUsernames

`func (o *GetServerSettings200Response) HasRequireEmailFormatUsernames() bool`

HasRequireEmailFormatUsernames returns a boolean if a field has been set.

### GetRealmUri

`func (o *GetServerSettings200Response) GetRealmUri() string`

GetRealmUri returns the RealmUri field if non-nil, zero value otherwise.

### GetRealmUriOk

`func (o *GetServerSettings200Response) GetRealmUriOk() (*string, bool)`

GetRealmUriOk returns a tuple with the RealmUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUri

`func (o *GetServerSettings200Response) SetRealmUri(v string)`

SetRealmUri sets RealmUri field to given value.

### HasRealmUri

`func (o *GetServerSettings200Response) HasRealmUri() bool`

HasRealmUri returns a boolean if a field has been set.

### GetRealmUrl

`func (o *GetServerSettings200Response) GetRealmUrl() string`

GetRealmUrl returns the RealmUrl field if non-nil, zero value otherwise.

### GetRealmUrlOk

`func (o *GetServerSettings200Response) GetRealmUrlOk() (*string, bool)`

GetRealmUrlOk returns a tuple with the RealmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmUrl

`func (o *GetServerSettings200Response) SetRealmUrl(v string)`

SetRealmUrl sets RealmUrl field to given value.

### HasRealmUrl

`func (o *GetServerSettings200Response) HasRealmUrl() bool`

HasRealmUrl returns a boolean if a field has been set.

### GetRealmName

`func (o *GetServerSettings200Response) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *GetServerSettings200Response) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *GetServerSettings200Response) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *GetServerSettings200Response) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetRealmIcon

`func (o *GetServerSettings200Response) GetRealmIcon() string`

GetRealmIcon returns the RealmIcon field if non-nil, zero value otherwise.

### GetRealmIconOk

`func (o *GetServerSettings200Response) GetRealmIconOk() (*string, bool)`

GetRealmIconOk returns a tuple with the RealmIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmIcon

`func (o *GetServerSettings200Response) SetRealmIcon(v string)`

SetRealmIcon sets RealmIcon field to given value.

### HasRealmIcon

`func (o *GetServerSettings200Response) HasRealmIcon() bool`

HasRealmIcon returns a boolean if a field has been set.

### GetRealmDescription

`func (o *GetServerSettings200Response) GetRealmDescription() string`

GetRealmDescription returns the RealmDescription field if non-nil, zero value otherwise.

### GetRealmDescriptionOk

`func (o *GetServerSettings200Response) GetRealmDescriptionOk() (*string, bool)`

GetRealmDescriptionOk returns a tuple with the RealmDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmDescription

`func (o *GetServerSettings200Response) SetRealmDescription(v string)`

SetRealmDescription sets RealmDescription field to given value.

### HasRealmDescription

`func (o *GetServerSettings200Response) HasRealmDescription() bool`

HasRealmDescription returns a boolean if a field has been set.

### GetRealmWebPublicAccessEnabled

`func (o *GetServerSettings200Response) GetRealmWebPublicAccessEnabled() bool`

GetRealmWebPublicAccessEnabled returns the RealmWebPublicAccessEnabled field if non-nil, zero value otherwise.

### GetRealmWebPublicAccessEnabledOk

`func (o *GetServerSettings200Response) GetRealmWebPublicAccessEnabledOk() (*bool, bool)`

GetRealmWebPublicAccessEnabledOk returns a tuple with the RealmWebPublicAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmWebPublicAccessEnabled

`func (o *GetServerSettings200Response) SetRealmWebPublicAccessEnabled(v bool)`

SetRealmWebPublicAccessEnabled sets RealmWebPublicAccessEnabled field to given value.

### HasRealmWebPublicAccessEnabled

`func (o *GetServerSettings200Response) HasRealmWebPublicAccessEnabled() bool`

HasRealmWebPublicAccessEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


