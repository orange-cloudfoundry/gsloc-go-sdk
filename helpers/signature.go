package helpers

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
	"google.golang.org/protobuf/proto"
)

func MessageSignature(m proto.Message) (string, error) {
	// make fingerprint from marshalling
	// this is not the best way to do it as deterministic marshalling
	//is not guaranteed in all langages but has we only plan to
	// work with go for signature generation it's ok.
	// if schema change it will produce a different signature but we what it that way anyway.
	b, err := proto.MarshalOptions{
		AllowPartial:  true,
		Deterministic: true,
		UseCachedSize: true,
	}.Marshal(m)
	if err != nil {
		return "", err
	}
	h := xxhash.New()
	h.Write(b)
	return fmt.Sprint(h.Sum64()), nil
}
