package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func strPtr(s string) *string { return &s }

// --- roleToState ---

func TestRoleToState_NilIsNull(t *testing.T) {
	got := roleToState(nil)
	if !got.IsNull() {
		t.Errorf("expected null for nil input, got %q", got.ValueString())
	}
}

func TestRoleToState_EmptyStringIsNull(t *testing.T) {
	got := roleToState(strPtr(""))
	if !got.IsNull() {
		t.Errorf("expected null for empty string, got %q", got.ValueString())
	}
}

func TestRoleToState_ValidRolePreserved(t *testing.T) {
	for _, role := range []string{"admin", "write", "plan", "read", "custom"} {
		got := roleToState(strPtr(role))
		if got.IsNull() {
			t.Errorf("role %q: expected non-null", role)
		}
		if got.ValueString() != role {
			t.Errorf("role %q: got %q", role, got.ValueString())
		}
	}
}

// --- resolveJobFlag ---

func TestResolveJobFlag_ExplicitTrueOverridesFalseInherit(t *testing.T) {
	got := resolveJobFlag(types.BoolValue(true), types.BoolValue(false))
	if !got {
		t.Error("expected true: explicit true should override false inherit")
	}
}

func TestResolveJobFlag_ExplicitFalseOverridesTrueInherit(t *testing.T) {
	got := resolveJobFlag(types.BoolValue(false), types.BoolValue(true))
	if got {
		t.Error("expected false: explicit false should override true inherit")
	}
}

func TestResolveJobFlag_NullExplicitInheritsTrue(t *testing.T) {
	got := resolveJobFlag(types.BoolNull(), types.BoolValue(true))
	if !got {
		t.Error("expected true: null explicit should inherit true from manage_job")
	}
}

func TestResolveJobFlag_NullExplicitInheritsFalse(t *testing.T) {
	got := resolveJobFlag(types.BoolNull(), types.BoolValue(false))
	if got {
		t.Error("expected false: null explicit should inherit false from manage_job")
	}
}

func TestResolveJobFlag_UnknownExplicitInheritsTrue(t *testing.T) {
	got := resolveJobFlag(types.BoolUnknown(), types.BoolValue(true))
	if !got {
		t.Error("expected true: unknown explicit should inherit from manage_job")
	}
}

func TestResolveJobFlag_ExplicitFalseWithFalseInherit(t *testing.T) {
	got := resolveJobFlag(types.BoolValue(false), types.BoolValue(false))
	if got {
		t.Error("expected false")
	}
}
