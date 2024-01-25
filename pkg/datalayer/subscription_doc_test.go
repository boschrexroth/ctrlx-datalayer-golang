package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

func ExampleSubscription_Subscribe() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
	}
	// Create subscription
	p := datalayer.DefaultSubscriptionProperties()
	s, _ := c.CreateSubscription("mySub", p, onSubscribe)

	//Subscribe to the address
	s.Subscribe("address")
}

func ExampleSubscription_Unsubscribe() {
	// Initialize and destroy the client
	var c *datalayer.Client
	defer datalayer.DeleteClient(c)

	//callback function
	onSubscribe := func(result datalayer.Result, items map[string]datalayer.Variant) {
	}

	// Create subscription
	p := datalayer.DefaultSubscriptionProperties()
	s, _ := c.CreateSubscription("mySub", p, onSubscribe)

	//Subscribe and unsubscribe the address
	s.Subscribe("address")
	s.Unsubscribe("address")

}
