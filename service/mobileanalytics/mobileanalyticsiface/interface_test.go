// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalyticsiface_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*mobileanalyticsiface.MobileAnalyticsAPI)(nil), mobileanalytics.New(nil))
}
