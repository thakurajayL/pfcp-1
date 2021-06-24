// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pfcpType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type QERID struct {
	QERID uint32
}

func (q *QERID) MarshalBinary() ([]byte, error) {
	var buf = &bytes.Buffer{}

	if err := binary.Write(buf, binary.BigEndian, &q.QERID); err != nil {
		return nil, fmt.Errorf("marshal QERID fail: " + err.Error())
	}

	return buf.Bytes(), nil
}

func (q *QERID) UnmarshalBinary(data []byte) error {
	var buf = bytes.NewBuffer(data)

	if err := binary.Read(buf, binary.BigEndian, &q.QERID); err != nil {
		return fmt.Errorf("unmarshal QERID fail: " + err.Error())
	}

	return nil
}
