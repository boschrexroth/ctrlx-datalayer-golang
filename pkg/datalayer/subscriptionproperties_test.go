// MIT License
//
// Copyright (c) 2021 Bosch Rexroth AG
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
	"testing"

	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/comm/datalayer"
	"github.com/oliveagle/jsonpath"
	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	s := NewSystem("")
	defer DeleteSystem(s)
	want := SubscriptionProperties{
		Id:                    "Test",
		KeepaliveInterval:     10000,
		PublishInterval:       500,
		ErrorInterval:         20000,
		QueueSize:             5,
		QueueBehaviour:        fbs.QueueBehaviourDiscardNewest,
		DataChangeTrigger:     fbs.DataChangeTriggerStatus,
		CountingSubscriptions: true,
	}
	sc := newSubscriptionPropertiesConverter(s)
	j, err := want.BuildJson()
	assert.Nil(t, err)
	b, err := sc.JsonToFlatBuffer(j)
	assert.Nil(t, err)
	c := s.JSONConverter()
	r, jsd := c.GenerateJsonComplex(b, sc.schema, 2)
	assert.Equal(t, Result(0), r)
	bytes := []byte(jsd)
	jsMap := make(map[string]interface{})
	err = json.Unmarshal(bytes[:len(bytes)-1], &jsMap)
	assert.Nil(t, err)
	res, err := jsonpath.JsonPathLookup(jsMap, "$.id")
	assert.Nil(t, err)
	assert.Equal(t, res, want.Id)
	res, err = jsonpath.JsonPathLookup(jsMap, "$.rules")
	assert.Nil(t, err)
	rules := res.([]interface{})
	assert.NotNil(t, rules)
	for _, rule := range rules {
		res, err = jsonpath.JsonPathLookup(rule, "$.rule_type")
		assert.Nil(t, err)
		switch res.(string) {
		case "Queueing":
			res, err = jsonpath.JsonPathLookup(rule, "$.rule.queueSize")
			assert.Nil(t, err)
			assert.Equal(t, 5.0, res)
		}
	}
}

func _TestSubscriptionAPI(t *testing.T) {
	system := NewSystem("")
	system.Start(true)
	client := system.Factory().CreateClient("tcp://boschrexroth:boschrexroth@192.168.178.20:2069")
	sub, result := client.CreateSubscription("test", DefaultSubscriptionProperties(), nil)
	assert.Equal(t, Result(0), result)
	assert.NotNil(t, sub)
}
