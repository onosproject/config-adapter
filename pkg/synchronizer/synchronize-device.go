// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package synchronizer

import (
	"github.com/onosproject/config-adapter/pkg/gnmi"
)

// SynchronizeDevice synchronizes a device. Two sets of error state are returned:
//   1) pushFailures -- a count of pushes that failed to the core. Synchronizer should retry again later.
//   2) error -- a fatal error that occurred during synchronization.
func (s *Synchronizer) SynchronizeDevice(allConfig *gnmi.ConfigForest) (int, error) {

	// Forget all current metrics. We'll compute and report them inside the sync loop.
	KpiSliceBitrate.Reset()
	KpiApplicationBitrate.Reset()
	KpiDeviceGroupBitrate.Reset()

	pushFailures := 0

	// TODO: everything

	return pushFailures, nil
}
