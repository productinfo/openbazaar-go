package core_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/OpenBazaar/jsonpb"
	"github.com/OpenBazaar/openbazaar-go/pb"
	"github.com/OpenBazaar/openbazaar-go/test"
)

func TestOpenBazaarSignedListings_GetSignedListingFromPath(t *testing.T) {

	absPath, _ := filepath.Abs("../test/contracts/signed_listings_1.json")

	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Error(err)
	}

	sl := new(pb.SignedListing)
	err = jsonpb.UnmarshalString(string(file), sl)
	if err != nil {
		t.Error(err)
	}

	node, err := test.NewNode()
	if err != nil {
		t.Error(err)
	}

}
