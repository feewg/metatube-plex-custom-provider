package main

import "testing"

func TestProviderDocumentUsesRegistrationRelativeFeaturePaths(t *testing.T) {
	document := ProviderDocument(Settings{
		ProviderIdentifier: "tv.plex.agents.custom.metatube.movie",
		ProviderTitle:      "MetaTube Movie Provider",
		AuthPath:           "_metatube",
		AuthToken:          "secret",
	})

	features := document.MediaProvider.Feature
	if features[0].Key != metadataPath {
		t.Fatalf("unexpected metadata feature path: %q", features[0].Key)
	}
	if features[1].Key != matchPath {
		t.Fatalf("unexpected match feature path: %q", features[1].Key)
	}
}

func TestRequestPathStripsAuthenticationPrefix(t *testing.T) {
	settings := Settings{AuthPath: "_metatube", AuthToken: "secret"}

	got, ok := requestPath("/_metatube/secret/library/metadata/matches", settings)
	if !ok {
		t.Fatal("expected authenticated path to be accepted")
	}
	if got != matchPath {
		t.Fatalf("unexpected request path: %q", got)
	}
}
