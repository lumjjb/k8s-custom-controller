package crd_controller

import (
	trust_schema "cit_custom_controller/crd_schema/cit_trust_schema"
	api "k8s.io/client-go/pkg/api/v1"
	"testing"
)

func TestGetPLCrdDef(t *testing.T) {
	expecPlCrd := CrdDefinition{
		Plural:   "platformcrds",
		Singular: "platformcrd",
		Group:    "cit.intel.com",
		Kind:     "PlatformCrd",
	}
	recvPlCrd := GetPLCrdDef()
	if expecPlCrd != recvPlCrd {
		t.Fatalf("Changes found in PL CRD Definition ")
		t.Fatalf("Expected :%v however Received: %v ", expecPlCrd, recvPlCrd)
	}
	t.Logf("Test GetPLCrd Def success")
}

func TestGetPlObjLabel(t *testing.T) {
	trust_obj := trust_schema.HostList{
		Hostname:             "Node123",
		Trusted:              "true",
		TrustTagExpiry:       "12-23-45T123.91.12",
		TrustTagSignedReport: "495270d6242e2c67e24e22bad49dgdah",
	}
	node := &api.Node{}
	recvlabel, recannotate := GetPlObjLabel(trust_obj, node)
	if _, ok := recvlabel["trusted"]; ok {
		t.Logf("Found in PL label Trusted field")
	} else {
		t.Fatalf("Could not get label trusted from PL Report")
	}
	if _, ok := recvlabel["TrustTagExpiry"]; ok {
		t.Logf("Found in PL label TrustTagExpiry field")
	} else {
		t.Fatalf("Could not get label TrustTagExpiry from PL Report")
	}
	if _, ok := recannotate["TrustTagSignedReport"]; ok {
		t.Logf("Found in PL annotation TrustTagSignedReport ")
	} else {
		t.Fatalf("Could not get annotation TrustTagSignedReport from PL Report")
	}
	t.Logf("Test getPlObjLabel success")
}
