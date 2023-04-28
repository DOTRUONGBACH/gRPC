package main

import (
	"ent-grpc-example/ent/proto/entpb"
	"testing"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:         "bach",
		EmailAddress: "vuabongdapro@gmail.com",
	}
	if user.GetName() != "cc" {
		t.Fatal("expected user name to be bach")
	}
	if user.GetEmailAddress() != "vuabongdapro@gmail.com" {
		t.Fatal("expected email address to be vuabongdapro@gmail.com ")
	}
}
