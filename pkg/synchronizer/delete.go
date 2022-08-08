// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package synchronizer

/*
 * Deletes are always handled synchronously. The synchronizer stops and wait for a delete to be
 * completed before it continues. This is to handle the case where we fail while performing the
 * delete. It'll get marked as a FAILED transaction in onos-config.
 *
 * This is in contrasts to configuration updates, which are generally handled asynchronously.
 */

import (
	"errors"
	"github.com/onosproject/config-adapter/pkg/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// HandleDelete synchronously performs a delete
func (s *Synchronizer) HandleDelete(config *gnmi.ConfigForest, path *pb.Path) error {
	if path == nil || len(path.Elem) == 0 {
		return errors.New("Delete of whole enterprise is not currently supported")
	}

	target := path.Target
	if target == "" {
		return errors.New("Refusing to handle delete without target specified")
	}

	rootDeviceInterface, okay := config.Configs[target]
	if !okay {
		log.Infof("Delete on target %s is for an empty tree", target)
		return nil
	}

	_ = rootDeviceInterface

	return nil
}
