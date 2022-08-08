// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"github.com/onosproject/config-adapter/pkg/gnmi"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

type target struct {
	*gnmi.Server
	Model       *gnmi.Model
	UpdateChann chan *pb.Update
}

var log = logging.GetLogger("target")
