// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalMeasurementMethod(t *testing.T) {
	testData := MeasurementMethod{
		Event: true,
		Volum: false,
		Durat: true,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{5}, buf)
}

func TestUnmarshalMeasurementMethod(t *testing.T) {
	buf := []byte{5}
	var testData MeasurementMethod
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := MeasurementMethod{
		Event: true,
		Volum: false,
		Durat: true,
	}
	assert.Equal(t, expectData, testData)
}
