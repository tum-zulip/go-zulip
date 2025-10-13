# SendInvites400Response

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

### NewSendInvites400Response

`func NewSendInvites400Response(result interface{}, msg interface{}, code interface{}, ) *SendInvites400Response`

NewSendInvites400Response instantiates a new SendInvites400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInvites400ResponseWithDefaults

`func NewSendInvites400ResponseWithDefaults() *SendInvites400Response`

NewSendInvites400ResponseWithDefaults instantiates a new SendInvites400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SendInvites400Response) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SendInvites400Response) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SendInvites400Response) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *SendInvites400Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *SendInvites400Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetMsg

`func (o *SendInvites400Response) GetMsg() interface{}`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SendInvites400Response) GetMsgOk() (*interface{}, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SendInvites400Response) SetMsg(v interface{})`

SetMsg sets Msg field to given value.


### SetMsgNil

`func (o *SendInvites400Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *SendInvites400Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetCode

`func (o *SendInvites400Response) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SendInvites400Response) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SendInvites400Response) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *SendInvites400Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SendInvites400Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetErrors

`func (o *SendInvites400Response) GetErrors() [][]InvitationFailedErrorAllOfErrorsInnerInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SendInvites400Response) GetErrorsOk() (*[][]InvitationFailedErrorAllOfErrorsInnerInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SendInvites400Response) SetErrors(v [][]InvitationFailedErrorAllOfErrorsInnerInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SendInvites400Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSentInvitations

`func (o *SendInvites400Response) GetSentInvitations() bool`

GetSentInvitations returns the SentInvitations field if non-nil, zero value otherwise.

### GetSentInvitationsOk

`func (o *SendInvites400Response) GetSentInvitationsOk() (*bool, bool)`

GetSentInvitationsOk returns a tuple with the SentInvitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentInvitations

`func (o *SendInvites400Response) SetSentInvitations(v bool)`

SetSentInvitations sets SentInvitations field to given value.

### HasSentInvitations

`func (o *SendInvites400Response) HasSentInvitations() bool`

HasSentInvitations returns a boolean if a field has been set.

### GetDailyLimitReached

`func (o *SendInvites400Response) GetDailyLimitReached() bool`

GetDailyLimitReached returns the DailyLimitReached field if non-nil, zero value otherwise.

### GetDailyLimitReachedOk

`func (o *SendInvites400Response) GetDailyLimitReachedOk() (*bool, bool)`

GetDailyLimitReachedOk returns a tuple with the DailyLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitReached

`func (o *SendInvites400Response) SetDailyLimitReached(v bool)`

SetDailyLimitReached sets DailyLimitReached field to given value.

### HasDailyLimitReached

`func (o *SendInvites400Response) HasDailyLimitReached() bool`

HasDailyLimitReached returns a boolean if a field has been set.

### GetLicenseLimitReached

`func (o *SendInvites400Response) GetLicenseLimitReached() bool`

GetLicenseLimitReached returns the LicenseLimitReached field if non-nil, zero value otherwise.

### GetLicenseLimitReachedOk

`func (o *SendInvites400Response) GetLicenseLimitReachedOk() (*bool, bool)`

GetLicenseLimitReachedOk returns a tuple with the LicenseLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLimitReached

`func (o *SendInvites400Response) SetLicenseLimitReached(v bool)`

SetLicenseLimitReached sets LicenseLimitReached field to given value.

### HasLicenseLimitReached

`func (o *SendInvites400Response) HasLicenseLimitReached() bool`

HasLicenseLimitReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


