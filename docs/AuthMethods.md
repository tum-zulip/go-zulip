# AuthMethods

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

### NewAuthMethods

`func NewAuthMethods() *AuthMethods`

NewAuthMethods instantiates a new AuthMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodsWithDefaults

`func NewAuthMethodsWithDefaults() *AuthMethods`

NewAuthMethodsWithDefaults instantiates a new AuthMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AuthMethods) GetPassword() bool`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthMethods) GetPasswordOk() (*bool, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthMethods) SetPassword(v bool)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthMethods) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDev

`func (o *AuthMethods) GetDev() bool`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *AuthMethods) GetDevOk() (*bool, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *AuthMethods) SetDev(v bool)`

SetDev sets Dev field to given value.

### HasDev

`func (o *AuthMethods) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetEmail

`func (o *AuthMethods) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthMethods) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthMethods) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthMethods) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLdap

`func (o *AuthMethods) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AuthMethods) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AuthMethods) SetLdap(v bool)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *AuthMethods) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetRemoteuser

`func (o *AuthMethods) GetRemoteuser() bool`

GetRemoteuser returns the Remoteuser field if non-nil, zero value otherwise.

### GetRemoteuserOk

`func (o *AuthMethods) GetRemoteuserOk() (*bool, bool)`

GetRemoteuserOk returns a tuple with the Remoteuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteuser

`func (o *AuthMethods) SetRemoteuser(v bool)`

SetRemoteuser sets Remoteuser field to given value.

### HasRemoteuser

`func (o *AuthMethods) HasRemoteuser() bool`

HasRemoteuser returns a boolean if a field has been set.

### GetGithub

`func (o *AuthMethods) GetGithub() bool`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *AuthMethods) GetGithubOk() (*bool, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *AuthMethods) SetGithub(v bool)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *AuthMethods) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetAzuread

`func (o *AuthMethods) GetAzuread() bool`

GetAzuread returns the Azuread field if non-nil, zero value otherwise.

### GetAzureadOk

`func (o *AuthMethods) GetAzureadOk() (*bool, bool)`

GetAzureadOk returns a tuple with the Azuread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuread

`func (o *AuthMethods) SetAzuread(v bool)`

SetAzuread sets Azuread field to given value.

### HasAzuread

`func (o *AuthMethods) HasAzuread() bool`

HasAzuread returns a boolean if a field has been set.

### GetGitlab

`func (o *AuthMethods) GetGitlab() bool`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *AuthMethods) GetGitlabOk() (*bool, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *AuthMethods) SetGitlab(v bool)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *AuthMethods) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetApple

`func (o *AuthMethods) GetApple() bool`

GetApple returns the Apple field if non-nil, zero value otherwise.

### GetAppleOk

`func (o *AuthMethods) GetAppleOk() (*bool, bool)`

GetAppleOk returns a tuple with the Apple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApple

`func (o *AuthMethods) SetApple(v bool)`

SetApple sets Apple field to given value.

### HasApple

`func (o *AuthMethods) HasApple() bool`

HasApple returns a boolean if a field has been set.

### GetGoogle

`func (o *AuthMethods) GetGoogle() bool`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *AuthMethods) GetGoogleOk() (*bool, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *AuthMethods) SetGoogle(v bool)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *AuthMethods) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetSaml

`func (o *AuthMethods) GetSaml() bool`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *AuthMethods) GetSamlOk() (*bool, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *AuthMethods) SetSaml(v bool)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *AuthMethods) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOpenidConnect

`func (o *AuthMethods) GetOpenidConnect() bool`

GetOpenidConnect returns the OpenidConnect field if non-nil, zero value otherwise.

### GetOpenidConnectOk

`func (o *AuthMethods) GetOpenidConnectOk() (*bool, bool)`

GetOpenidConnectOk returns a tuple with the OpenidConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidConnect

`func (o *AuthMethods) SetOpenidConnect(v bool)`

SetOpenidConnect sets OpenidConnect field to given value.

### HasOpenidConnect

`func (o *AuthMethods) HasOpenidConnect() bool`

HasOpenidConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


