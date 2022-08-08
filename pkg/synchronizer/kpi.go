// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package synchronizer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// KpiSynchronizationTotal is the count of Synchronizations
	KpiSynchronizationTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "synchronization_total",
		Help: "The total number of synchronizations",
	},
		[]string{"enterprise"},
	)

	// KpiSynchronizationDuration is a histogram of duration of synchronizations
	KpiSynchronizationDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "synchronization_duration",
		Help: "The duration of synchronizations",
	},
		[]string{"enterprise"},
	)

	// Note: Prometheus best practice is to track "total requests" and "total failures" as
	// separate metrics, rather than "total successes" and "total failures", or a
	// combined metric with a "status" label.
	//
	// For that reason, we first ResourceTotal, which is our total requests. Then we track
	// FailedTotal, which represents the errors.

	// KpiSynchronizationResourceTotal is the total number of resources that were
	// attempted to be synchronized.
	KpiSynchronizationResourceTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "synchronization_resource_total",
		Help: "The total number of resources synchronized",
	},
		[]string{"enterprise", "kind"},
	)

	// KpiSynchronizationFailedTotal is a count of failed synchronizations
	KpiSynchronizationFailedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "synchronization_failed_total",
		Help: "The total number of failed synchronizations",
	},
		[]string{"enterprise", "kind", "destination"},
	)

	// KpiSliceBitrate is the Configured MBR for slices
	KpiSliceBitrate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "slice_bitrate",
		Help: "per-Slice configured maximum bitrate",
	},
		[]string{"enterprise", "site", "slice", "direction"},
	)

	// KpiApplicationBitrate is the Configured MBR for Applications
	KpiApplicationBitrate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "application_bitrate",
		Help: "per-Device per-Application configured maximim bitrate",
	},
		[]string{"enterprise", "site", "slice", "application", "endpoint", "direction"},
	)

	// KpiDeviceGroupBitrate is the Configured MBR for Device Groups
	KpiDeviceGroupBitrate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_group_bitrate",
		Help: "per-Device configured maximum bitrate",
	},
		[]string{"enterprise", "site", "devicegroup", "direction"},
	)
)
