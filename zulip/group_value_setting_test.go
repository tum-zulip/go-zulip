package zulip_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tum-zulip/go-zulip/zulip"
)

func TestGroupSettingValueUnmarshalJSON(t *testing.T) {
	var v zulip.GroupSettingValue

	err := json.Unmarshal([]byte(`1`), &v)
	require.NoError(t, err)
	assert.Equal(t, v.GroupId, ptrTo(int64(1)))

	err = json.Unmarshal([]byte(`{"direct_members":[1,2,3],"direct_subgroups":[4,5,6]}`), &v)
	require.NoError(t, err)
	require.NotNil(t, v.ComplexGroupSettingValue)
	assert.Equal(t, []int64{1, 2, 3}, v.ComplexGroupSettingValue.DirectMembers)
	assert.Equal(t, []int64{4, 5, 6}, v.ComplexGroupSettingValue.DirectSubgroups)

	err = json.Unmarshal([]byte(`{"direct_members":[],"direct_subgroups":[]}`), &v)
	require.NoError(t, err)
	require.NotNil(t, v.ComplexGroupSettingValue)
	assert.Equal(t, []int64{}, v.ComplexGroupSettingValue.DirectMembers)
	assert.Equal(t, []int64{}, v.ComplexGroupSettingValue.DirectSubgroups)
}
