// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package extra25519

import (
	"bytes"
	"crypto/rand"
	"testing"

	"crypto/ed25519"

	"golang.org/x/crypto/curve25519"
)

func TestCurve25519Conversion(t *testing.T) {
	public, private, _ := ed25519.GenerateKey(rand.Reader)

	var curve25519Public, curve25519Public2, curve25519Private [32]byte
	var pk [64]byte
	var pub [32]byte
	copy(pk[:], private)
	copy(pub[:], public)
	PrivateKeyToCurve25519(&curve25519Private, &pk)
	curve25519.ScalarBaseMult(&curve25519Public, &curve25519Private)

	if !PublicKeyToCurve25519(&curve25519Public2, &pub) {
		t.Fatalf("PublicKeyToCurve25519 failed")
	}

	if !bytes.Equal(curve25519Public[:], curve25519Public2[:]) {
		t.Errorf("Values didn't match: curve25519 produced %x, conversion produced %x", curve25519Public[:], curve25519Public2[:])
	}
}

func TestElligator(t *testing.T) {
	var publicKey, publicKey2, publicKey3, representative, privateKey [32]byte

	for i := 0; i < 1000; i++ {
		if _, err := rand.Reader.Read(privateKey[:]); err != nil {
			t.Fatal(err.Error())
		}

		if !ScalarBaseMult(&publicKey, &representative, &privateKey) {
			continue
		}
		RepresentativeToPublicKey(&publicKey2, &representative)
		if !bytes.Equal(publicKey[:], publicKey2[:]) {
			t.Fatal("The resulting public key doesn't match the initial one.")
		}

		curve25519.ScalarBaseMult(&publicKey3, &privateKey)
		if !bytes.Equal(publicKey[:], publicKey3[:]) {
			t.Fatal("The public key doesn't match the value that curve25519 produced.")
		}
	}
}

func BenchmarkKeyGeneration(b *testing.B) {
	var publicKey, representative, privateKey [32]byte

	// Find the private key that results in a point that's in the image of the map.
	for {
		if _, err := rand.Reader.Read(privateKey[:]); err != nil {
			b.Fatal(err.Error())
		}
		if ScalarBaseMult(&publicKey, &representative, &privateKey) {
			break
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ScalarBaseMult(&publicKey, &representative, &privateKey)
	}
}

func BenchmarkMap(b *testing.B) {
	var publicKey, representative [32]byte
	if _, err := rand.Reader.Read(representative[:]); err != nil {
		b.Fatal(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RepresentativeToPublicKey(&publicKey, &representative)
	}
}
