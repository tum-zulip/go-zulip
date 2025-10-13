# GetServerSettings200ResponseAllOfAuthenticationMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **bool** | Whether the user can authenticate using password.  | [optional] 
**Dev** | Pointer to **bool** | Whether the user can authenticate using development API key.  | [optional] 
**Email** | Pointer to **bool** | Whether the user can authenticate using email.  | [optional] 
**Ldap** | Pointer to **bool** | Whether the user can authenticate using LDAP.  | [optional] 
**Remoteuser** | Pointer to **bool** | Whether the user can authenticate using REMOTE_USER.  | [optional] 
**Github** | Pointer to **bool** | Whether the user can authenticate using their GitHub account.  | [optional] 
**Azuread** | Pointer to **bool** | Whether the user can authenticate using their Microsoft Entra ID account.  | [optional] 
**Gitlab** | Pointer to **bool** | Whether the user can authenticate using their GitLab account.  **Changes**: New in Zulip 3.0 (feature level 1).  | [optional] 
**Apple** | Pointer to **bool** | Whether the user can authenticate using their Apple account.  | [optional] 
**Google** | Pointer to **bool** | Whether the user can authenticate using their Google account.  | [optional] 
**Saml** | Pointer to **bool** | Whether the user can authenticate using SAML.  | [optional] 
**OpenidConnect** | Pointer to **bool** | Whether the user can authenticate using OpenID Connect.  | [optional] 

## Methods

### NewGetServerSettings200ResponseAllOfAuthenticationMethods

`func NewGetServerSettings200ResponseAllOfAuthenticationMethods() *GetServerSettings200ResponseAllOfAuthenticationMethods`

NewGetServerSettings200ResponseAllOfAuthenticationMethods instantiates a new GetServerSettings200ResponseAllOfAuthenticationMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerSettings200ResponseAllOfAuthenticationMethodsWithDefaults

`func NewGetServerSettings200ResponseAllOfAuthenticationMethodsWithDefaults() *GetServerSettings200ResponseAllOfAuthenticationMethods`

NewGetServerSettings200ResponseAllOfAuthenticationMethodsWithDefaults instantiates a new GetServerSettings200ResponseAllOfAuthenticationMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetPassword() bool`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetPasswordOk() (*bool, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetPassword(v bool)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDev

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetDev() bool`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetDevOk() (*bool, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetDev(v bool)`

SetDev sets Dev field to given value.

### HasDev

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetEmail

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLdap

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetRemoteuser

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetRemoteuser() bool`

GetRemoteuser returns the Remoteuser field if non-nil, zero value otherwise.

### GetRemoteuserOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetRemoteuserOk() (*bool, bool)`

GetRemoteuserOk returns a tuple with the Remoteuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteuser

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetRemoteuser(v bool)`

SetRemoteuser sets Remoteuser field to given value.

### HasRemoteuser

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasRemoteuser() bool`

HasRemoteuser returns a boolean if a field has been set.

### GetGithub

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGithub() bool`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGithubOk() (*bool, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetGithub(v bool)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetAzuread

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetAzuread() bool`

GetAzuread returns the Azuread field if non-nil, zero value otherwise.

### GetAzureadOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetAzureadOk() (*bool, bool)`

GetAzureadOk returns a tuple with the Azuread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuread

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetAzuread(v bool)`

SetAzuread sets Azuread field to given value.

### HasAzuread

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasAzuread() bool`

HasAzuread returns a boolean if a field has been set.

### GetGitlab

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGitlab() bool`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGitlabOk() (*bool, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetGitlab(v bool)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetApple

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetApple() bool`

GetApple returns the Apple field if non-nil, zero value otherwise.

### GetAppleOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetAppleOk() (*bool, bool)`

GetAppleOk returns a tuple with the Apple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApple

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetApple(v bool)`

SetApple sets Apple field to given value.

### HasApple

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasApple() bool`

HasApple returns a boolean if a field has been set.

### GetGoogle

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGoogle() bool`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetGoogleOk() (*bool, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetGoogle(v bool)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetSaml

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetSaml() bool`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetSamlOk() (*bool, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetSaml(v bool)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOpenidConnect

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetOpenidConnect() bool`

GetOpenidConnect returns the OpenidConnect field if non-nil, zero value otherwise.

### GetOpenidConnectOk

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) GetOpenidConnectOk() (*bool, bool)`

GetOpenidConnectOk returns a tuple with the OpenidConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidConnect

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) SetOpenidConnect(v bool)`

SetOpenidConnect sets OpenidConnect field to given value.

### HasOpenidConnect

`func (o *GetServerSettings200ResponseAllOfAuthenticationMethods) HasOpenidConnect() bool`

HasOpenidConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


