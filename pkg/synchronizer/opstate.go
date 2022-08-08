// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package synchronizer

import (
	"context"
	"flag"
	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/config-adapter/pkg/gnmi"
)

var (
	kafkaURI   = flag.String("kafka_uri", "", "URI of kafka")
	kafkaTopic = flag.String("kafka_topic", "reference_adapter", "kafka topic to fetch from")
)

// Repeatedly loop, receiving messages from Kafka
func (s *Synchronizer) receiveKafkaLoop(config *gnmi.ConfigForest) {
	log.Info("starting kafka receiver loop")
	for {
		select {
		case stringMsg := <-s.kafkaMsgChannel:
			_ = stringMsg
			// TODO: handle kafka message
		case err := <-s.kafkaErrorChannel:
			log.Warnf("Kafka Error: %v", err)
		}
	}
}

// Start opstate synchronization thread
func (s *Synchronizer) startOpstate(config *gnmi.ConfigForest) {
	s.opstateStarted = true

	if *kafkaURI == "" {
		log.Info("no kafkaURI specified; not starting kafka client")
		return
	}

	log.Info("starting opstate processor on topic %s for URI %s", *kafkaTopic, *kafkaURI)

	go kafkaClient.StartTopicReader(context.Background(),
		s.kafkaMsgChannel,
		s.kafkaErrorChannel,
		[]string{*kafkaURI},
		*kafkaTopic,
		"opstate",
	)

	go s.receiveKafkaLoop(config)
}
