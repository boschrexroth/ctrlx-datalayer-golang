// MIT License
//
// Copyright (c) 2021-2022 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package datalayer

import (
	"encoding/json"

	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
)

// SubscriptionProperties struct
type SubscriptionProperties struct {
	Id                    string
	KeepaliveInterval     uint32
	PublishInterval       uint32
	ErrorInterval         uint32
	SamplingInterval      uint64
	QueueSize             uint32
	QueueBehaviour        fbs.QueueBehaviour
	DeadBandValue         float32
	DataChangeTrigger     fbs.DataChangeTrigger
	BrowseListChange      bool
	MetaDataChange        bool
	CountingSubscriptions bool
	RateLimit             uint32
}

// DefaultSubscriptionProperties returns all default subscription properties.
func DefaultSubscriptionProperties() SubscriptionProperties {
	return SubscriptionProperties{
		Id:                    "",
		KeepaliveInterval:     60000,
		PublishInterval:       1000,
		ErrorInterval:         10000,
		SamplingInterval:      1000000,
		QueueSize:             10,
		QueueBehaviour:        fbs.QueueBehaviourDiscardOldest,
		DeadBandValue:         0,
		DataChangeTrigger:     fbs.DataChangeTriggerStatusValue,
		BrowseListChange:      false,
		MetaDataChange:        false,
		CountingSubscriptions: false,
		RateLimit:             0,
	}
}

// BuildJson constructs the Json file with default subscription properties.
// It returns the encoded JSON.
func (s SubscriptionProperties) BuildJson() ([]byte, error) {
	d := DefaultSubscriptionProperties()
	v := map[string]interface{}{
		"id":                s.Id,
		"keepaliveInterval": s.KeepaliveInterval,
		"publishInterval":   s.PublishInterval,
		"errorInterval":     s.ErrorInterval,
	}
	v["rules"] = make([]interface{}, 0)
	if s.QueueSize != d.QueueSize ||
		s.QueueBehaviour != d.QueueBehaviour {
		p := make(map[string]interface{})
		p["rule_type"] = "Queueing"
		r := make(map[string]interface{})
		p["rule"] = r
		r["queueSize"] = s.QueueSize
		r["behaviour"] = s.QueueBehaviour
		rs := v["rules"].([]interface{})
		rs = append(rs, p)
		v["rules"] = rs
	}
	if s.DeadBandValue != d.DeadBandValue {
		p := make(map[string]interface{})
		p["rule_type"] = "DataChangeFilter"
		r := make(map[string]interface{})
		p["rule"] = r
		r["deadBandValue"] = s.DeadBandValue
		rs := v["rules"].([]interface{})
		rs = append(rs, p)
		v["rules"] = rs
	}
	if s.DataChangeTrigger != d.DataChangeTrigger ||
		s.BrowseListChange != d.BrowseListChange ||
		s.MetaDataChange != d.MetaDataChange {
		p := make(map[string]interface{})
		p["rule_type"] = "ChangeEvents"
		r := make(map[string]interface{})
		p["rule"] = r
		switch s.DataChangeTrigger {
		case fbs.DataChangeTriggerStatus:
			r["valueChange"] = "Status"
		case fbs.DataChangeTriggerStatusValue:
			r["valueChange"] = "StatusValue"
		case fbs.DataChangeTriggerStatusValueTimestamp:
			r["valueChange"] = "StatusValueTimestamp"
		}
		r["browselistChange"] = s.BrowseListChange
		r["metadataChange"] = s.MetaDataChange
		rs := v["rules"].([]interface{})
		rs = append(rs, p)
		v["rules"] = rs
	}
	if s.CountingSubscriptions != d.CountingSubscriptions {
		p := make(map[string]interface{})
		p["rule_type"] = "Counting"
		r := make(map[string]interface{})
		p["rule"] = r
		r["countSubscriptions"] = s.CountingSubscriptions
		rs := v["rules"].([]interface{})
		rs = append(rs, p)
		v["rules"] = rs
	}
	if s.RateLimit != d.RateLimit {
		p := make(map[string]interface{})
		p["rule_type"] = "LosslessRateLimit"
		r := make(map[string]interface{})
		p["rule"] = r
		r["rateLimit"] = s.RateLimit
		rs := v["rules"].([]interface{})
		rs = append(rs, p)
		v["rules"] = rs
	}
	return json.Marshal(v)
}
