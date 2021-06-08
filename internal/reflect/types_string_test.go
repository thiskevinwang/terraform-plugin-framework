package reflect

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type testStringType struct {
	setErr error
}

var _ attr.Type = testStringType{}

func (t testStringType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.String
}

func (t testStringType) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	if !in.Type().Is(tftypes.String) {
		return nil, fmt.Errorf("unexpected type %s", tftypes.String)
	}
	result := &testStringValue{
		setErr: t.setErr,
	}
	if !in.IsKnown() {
		result.Unknown = true
		return result, nil
	}
	if in.IsNull() {
		result.Null = true
		return result, nil
	}
	err := in.As(&result.Value)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t testStringType) Equal(other attr.Type) bool {
	_, ok := other.(testStringType)
	return ok
}

type testStringValue struct {
	Unknown bool
	Null    bool
	Value   string
	setErr  error
}

var _ attr.Value = &testStringValue{}

func (t *testStringValue) ToTerraformValue(_ context.Context) (interface{}, error) {
	if t.Unknown {
		return tftypes.UnknownValue, nil
	}
	if t.Null {
		return nil, nil
	}
	return t.Value, nil
}

func (t *testStringValue) SetTerraformValue(_ context.Context, in tftypes.Value) error {
	if t.setErr != nil {
		return t.setErr
	}
	t.Unknown = false
	t.Null = false
	t.Value = ""
	if !in.Type().Is(tftypes.String) {
		return fmt.Errorf("unexpected type %s", tftypes.String)
	}
	if !in.IsKnown() {
		t.Unknown = true
		return nil
	}
	if !in.IsNull() {
		t.Null = true
		return nil
	}
	return in.As(&t.Value)
}

func (t *testStringValue) Equal(other attr.Value) bool {
	if t == nil && other == nil {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	o, ok := other.(*testStringValue)
	if !ok {
		return false
	}
	if t.Unknown != o.Unknown {
		return false
	}
	if t.Null != o.Null {
		return false
	}
	if t.Value != o.Value {
		return false
	}
	return true
}