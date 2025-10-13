# InvitationFailedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**Msg** | **interface{}** |  | 
**Code** | **interface{}** |  | 
**Errors** | Pointer to [**[][]InvitationFailedErrorAllOfErrorsInnerInner**]([]InvitationFailedErrorAllOfErrorsInnerInner.md) | An array of arrays of length 3, where each inner array consists of (a) an email address that was skipped while sending invitations, (b) the corresponding error message, and (c) a boolean which is &#x60;true&#x60; when the email address already uses Zulip and the corresponding user is deactivated in the organization.  | [optional] 
**SentInvitations** | Pointer to **bool** | A boolean specifying whether any invitations were sent.  | [optional] 
**DailyLimitReached** | Pointer to **bool** | A boolean specifying whether the limit on the number of invitations that can be sent in the organization in a day has been reached.  | [optional] 
**LicenseLimitReached** | Pointer to **bool** | A boolean specifying whether the organization have enough unused Zulip licenses to invite specified number of users.  | [optional] 

## Methods

### NewInvitationFailedError

`func NewInvitationFailedError(result interface{}, msg interface{}, code interface{}, ) *InvitationFailedError`

NewInvitationFailedError instantiates a new InvitationFailedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationFailedErrorWithDefaults

`func NewInvitationFailedErrorWithDefaults() *InvitationFailedError`

NewInvitationFailedErrorWithDefaults instantiates a new InvitationFailedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *InvitationFailedError) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *InvitationFailedError) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *InvitationFailedError) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *InvitationFailedError) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *InvitationFailedError) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *InvitationFailedError) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InvitationFailedError) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InvitationFailedError) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *InvitationFailedError) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *InvitationFailedError) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *InvitationFailedError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvitationFailedError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvitationFailedError) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *InvitationFailedError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InvitationFailedError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetErrors

`func (o *InvitationFailedError) GetErrors() [][]InvitationFailedErrorAllOfErrorsInnerInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InvitationFailedError) GetErrorsOk() (*[][]InvitationFailedErrorAllOfErrorsInnerInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InvitationFailedError) SetErrors(v [][]InvitationFailedErrorAllOfErrorsInnerInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InvitationFailedError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSentInvitations

`func (o *InvitationFailedError) GetSentInvitations() bool`

GetSentInvitations returns the SentInvitations field if non-nil, zero value otherwise.

### GetSentInvitationsOk

`func (o *InvitationFailedError) GetSentInvitationsOk() (*bool, bool)`

GetSentInvitationsOk returns a tuple with the SentInvitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentInvitations

`func (o *InvitationFailedError) SetSentInvitations(v bool)`

SetSentInvitations sets SentInvitations field to given value.

### HasSentInvitations

`func (o *InvitationFailedError) HasSentInvitations() bool`

HasSentInvitations returns a boolean if a field has been set.

### GetDailyLimitReached

`func (o *InvitationFailedError) GetDailyLimitReached() bool`

GetDailyLimitReached returns the DailyLimitReached field if non-nil, zero value otherwise.

### GetDailyLimitReachedOk

`func (o *InvitationFailedError) GetDailyLimitReachedOk() (*bool, bool)`

GetDailyLimitReachedOk returns a tuple with the DailyLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitReached

`func (o *InvitationFailedError) SetDailyLimitReached(v bool)`

SetDailyLimitReached sets DailyLimitReached field to given value.

### HasDailyLimitReached

`func (o *InvitationFailedError) HasDailyLimitReached() bool`

HasDailyLimitReached returns a boolean if a field has been set.

### GetLicenseLimitReached

`func (o *InvitationFailedError) GetLicenseLimitReached() bool`

GetLicenseLimitReached returns the LicenseLimitReached field if non-nil, zero value otherwise.

### GetLicenseLimitReachedOk

`func (o *InvitationFailedError) GetLicenseLimitReachedOk() (*bool, bool)`

GetLicenseLimitReachedOk returns a tuple with the LicenseLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLimitReached

`func (o *InvitationFailedError) SetLicenseLimitReached(v bool)`

SetLicenseLimitReached sets LicenseLimitReached field to given value.

### HasLicenseLimitReached

`func (o *InvitationFailedError) HasLicenseLimitReached() bool`

HasLicenseLimitReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


