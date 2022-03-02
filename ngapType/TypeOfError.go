// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

package ngapType

import "github.com/free5gc/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

const (
	TypeOfErrorPresentNotUnderstood aper.Enumerated = 0
	TypeOfErrorPresentMissing       aper.Enumerated = 1
)

type TypeOfError struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
