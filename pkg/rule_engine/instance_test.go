package rule_engine

import (
	"io/ioutil"
	"testing"
)

func TestNewRuleChainInstance(t *testing.T) {
	buf, err := ioutil.ReadFile("./manifest/manifest_sample.json")
	if err != nil {
		t.Error(err)
	}
	_, errs := NewRuleChainInstance(buf)
	if len(errs) > 0 {
		t.Error(errs[0])
	}
}
