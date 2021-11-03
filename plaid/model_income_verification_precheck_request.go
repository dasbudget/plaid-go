/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationPrecheckRequest IncomeVerificationPrecheckRequest defines the request schema for `/income/verification/precheck`
type IncomeVerificationPrecheckRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	User NullableIncomeVerificationPrecheckUser `json:"user,omitempty"`
	Employer NullableIncomeVerificationPrecheckEmployer `json:"employer,omitempty"`
	// The access token associated with the Item data is being requested for.
	TransactionsAccessToken NullableString `json:"transactions_access_token,omitempty"`
	// An array of access tokens corresponding to the Items that will be cross-referenced with the product data. If the `transactions` product was not initialized for the Items during link, it will be initialized after this Link session.
	TransactionsAccessTokens *[]string `json:"transactions_access_tokens,omitempty"`
	UsMilitaryInfo NullableIncomeVerificationPrecheckMilitaryInfo `json:"us_military_info,omitempty"`
}

// NewIncomeVerificationPrecheckRequest instantiates a new IncomeVerificationPrecheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckRequest() *IncomeVerificationPrecheckRequest {
	this := IncomeVerificationPrecheckRequest{}
	return &this
}

// NewIncomeVerificationPrecheckRequestWithDefaults instantiates a new IncomeVerificationPrecheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckRequestWithDefaults() *IncomeVerificationPrecheckRequest {
	this := IncomeVerificationPrecheckRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationPrecheckRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationPrecheckRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckRequest) GetUser() IncomeVerificationPrecheckUser {
	if o == nil || o.User.Get() == nil {
		var ret IncomeVerificationPrecheckUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckRequest) GetUserOk() (*IncomeVerificationPrecheckUser, bool) {
	if o == nil  {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableIncomeVerificationPrecheckUser and assigns it to the User field.
func (o *IncomeVerificationPrecheckRequest) SetUser(v IncomeVerificationPrecheckUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *IncomeVerificationPrecheckRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *IncomeVerificationPrecheckRequest) UnsetUser() {
	o.User.Unset()
}

// GetEmployer returns the Employer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckRequest) GetEmployer() IncomeVerificationPrecheckEmployer {
	if o == nil || o.Employer.Get() == nil {
		var ret IncomeVerificationPrecheckEmployer
		return ret
	}
	return *o.Employer.Get()
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckRequest) GetEmployerOk() (*IncomeVerificationPrecheckEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Employer.Get(), o.Employer.IsSet()
}

// HasEmployer returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasEmployer() bool {
	if o != nil && o.Employer.IsSet() {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given NullableIncomeVerificationPrecheckEmployer and assigns it to the Employer field.
func (o *IncomeVerificationPrecheckRequest) SetEmployer(v IncomeVerificationPrecheckEmployer) {
	o.Employer.Set(&v)
}
// SetEmployerNil sets the value for Employer to be an explicit nil
func (o *IncomeVerificationPrecheckRequest) SetEmployerNil() {
	o.Employer.Set(nil)
}

// UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
func (o *IncomeVerificationPrecheckRequest) UnsetEmployer() {
	o.Employer.Unset()
}

// GetTransactionsAccessToken returns the TransactionsAccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessToken() string {
	if o == nil || o.TransactionsAccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionsAccessToken.Get()
}

// GetTransactionsAccessTokenOk returns a tuple with the TransactionsAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionsAccessToken.Get(), o.TransactionsAccessToken.IsSet()
}

// HasTransactionsAccessToken returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasTransactionsAccessToken() bool {
	if o != nil && o.TransactionsAccessToken.IsSet() {
		return true
	}

	return false
}

// SetTransactionsAccessToken gets a reference to the given NullableString and assigns it to the TransactionsAccessToken field.
func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessToken(v string) {
	o.TransactionsAccessToken.Set(&v)
}
// SetTransactionsAccessTokenNil sets the value for TransactionsAccessToken to be an explicit nil
func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessTokenNil() {
	o.TransactionsAccessToken.Set(nil)
}

// UnsetTransactionsAccessToken ensures that no value is present for TransactionsAccessToken, not even an explicit nil
func (o *IncomeVerificationPrecheckRequest) UnsetTransactionsAccessToken() {
	o.TransactionsAccessToken.Unset()
}

// GetTransactionsAccessTokens returns the TransactionsAccessTokens field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokens() []string {
	if o == nil || o.TransactionsAccessTokens == nil {
		var ret []string
		return ret
	}
	return *o.TransactionsAccessTokens
}

// GetTransactionsAccessTokensOk returns a tuple with the TransactionsAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokensOk() (*[]string, bool) {
	if o == nil || o.TransactionsAccessTokens == nil {
		return nil, false
	}
	return o.TransactionsAccessTokens, true
}

// HasTransactionsAccessTokens returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasTransactionsAccessTokens() bool {
	if o != nil && o.TransactionsAccessTokens != nil {
		return true
	}

	return false
}

// SetTransactionsAccessTokens gets a reference to the given []string and assigns it to the TransactionsAccessTokens field.
func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessTokens(v []string) {
	o.TransactionsAccessTokens = &v
}

// GetUsMilitaryInfo returns the UsMilitaryInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckRequest) GetUsMilitaryInfo() IncomeVerificationPrecheckMilitaryInfo {
	if o == nil || o.UsMilitaryInfo.Get() == nil {
		var ret IncomeVerificationPrecheckMilitaryInfo
		return ret
	}
	return *o.UsMilitaryInfo.Get()
}

// GetUsMilitaryInfoOk returns a tuple with the UsMilitaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckRequest) GetUsMilitaryInfoOk() (*IncomeVerificationPrecheckMilitaryInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UsMilitaryInfo.Get(), o.UsMilitaryInfo.IsSet()
}

// HasUsMilitaryInfo returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckRequest) HasUsMilitaryInfo() bool {
	if o != nil && o.UsMilitaryInfo.IsSet() {
		return true
	}

	return false
}

// SetUsMilitaryInfo gets a reference to the given NullableIncomeVerificationPrecheckMilitaryInfo and assigns it to the UsMilitaryInfo field.
func (o *IncomeVerificationPrecheckRequest) SetUsMilitaryInfo(v IncomeVerificationPrecheckMilitaryInfo) {
	o.UsMilitaryInfo.Set(&v)
}
// SetUsMilitaryInfoNil sets the value for UsMilitaryInfo to be an explicit nil
func (o *IncomeVerificationPrecheckRequest) SetUsMilitaryInfoNil() {
	o.UsMilitaryInfo.Set(nil)
}

// UnsetUsMilitaryInfo ensures that no value is present for UsMilitaryInfo, not even an explicit nil
func (o *IncomeVerificationPrecheckRequest) UnsetUsMilitaryInfo() {
	o.UsMilitaryInfo.Unset()
}

func (o IncomeVerificationPrecheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Employer.IsSet() {
		toSerialize["employer"] = o.Employer.Get()
	}
	if o.TransactionsAccessToken.IsSet() {
		toSerialize["transactions_access_token"] = o.TransactionsAccessToken.Get()
	}
	if o.TransactionsAccessTokens != nil {
		toSerialize["transactions_access_tokens"] = o.TransactionsAccessTokens
	}
	if o.UsMilitaryInfo.IsSet() {
		toSerialize["us_military_info"] = o.UsMilitaryInfo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckRequest struct {
	value *IncomeVerificationPrecheckRequest
	isSet bool
}

func (v NullableIncomeVerificationPrecheckRequest) Get() *IncomeVerificationPrecheckRequest {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckRequest) Set(val *IncomeVerificationPrecheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckRequest(val *IncomeVerificationPrecheckRequest) *NullableIncomeVerificationPrecheckRequest {
	return &NullableIncomeVerificationPrecheckRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

