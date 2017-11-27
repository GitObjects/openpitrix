// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app.proto

package openpitrix

import regexp "regexp"
import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_App_Id = regexp.MustCompile("^app-[a-zA-Z0-9]{8}$")

func (this *App) Validate() error {
	if this.Id != nil {
		if !_regex_App_Id.MatchString(*(this.Id)) {
			return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^app-[a-zA-Z0-9]{8}$"`, *(this.Id)))
		}
	}
	if this.Created != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.Created))); err != nil {
			return go_proto_validators.FieldError("Created", err)
		}
	}
	if this.LastModified != nil {
		if err := go_proto_validators.CallValidatorIfExists(&(*(this.LastModified))); err != nil {
			return go_proto_validators.FieldError("LastModified", err)
		}
	}
	return nil
}

var _regex_AppId_Id = regexp.MustCompile("^app-[a-zA-Z0-9]{8}$")

func (this *AppId) Validate() error {
	if this.Id != nil {
		if !_regex_AppId_Id.MatchString(*(this.Id)) {
			return go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^app-[a-zA-Z0-9]{8}$"`, *(this.Id)))
		}
	}
	return nil
}
func (this *AppListRequest) Validate() error {
	return nil
}
func (this *AppListResponse) Validate() error {
	for _, item := range this.Items {
		if err := go_proto_validators.CallValidatorIfExists(&(*(item))); err != nil {
			return go_proto_validators.FieldError("Items", err)
		}
	}
	return nil
}
