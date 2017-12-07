package dsc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viant/dsc"
	"github.com/viant/dsunit"
)

func TestCreateFromURL(t *testing.T) {
	factory := dsc.NewManagerFactory()
	{
		_, err := factory.CreateFromURL(dsunit.ExpandTestProtocolAsURLIfNeeded("test:///test/file_config3.json"))
		assert.NotNil(t, err)
	}

	{
		_, err := factory.CreateFromURL(dsunit.ExpandTestProtocolAsURLIfNeeded("test:///test/file_config.json"))
		assert.NotNil(t, err)
	}
}

func TestMissingDricer(t *testing.T) {
	factory := dsc.NewManagerFactory()
	{
		_, err := factory.CreateFromURL(dsunit.ExpandTestProtocolAsURLIfNeeded("test:///test/file_config3.json"))
		assert.NotNil(t, err)
	}

}
