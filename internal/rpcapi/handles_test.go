// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package rpcapi

import (
	"testing"

	"github.com/hashicorp/go-slug/sourcebundle"

	"github.com/hashicorp/terraform/internal/stacks/stackconfig"
	"github.com/hashicorp/terraform/internal/stacks/stackstate"
)

func TestNewHandleWithDependency(t *testing.T) {
	t.Run("valid parent", func(t *testing.T) {
		handles := newHandleTable()
		bundleHnd := handles.NewSourceBundle((*sourcebundle.Bundle)(nil))

		configHnd, err := handles.NewStackConfig((*stackconfig.Config)(nil), bundleHnd)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if configHnd.IsNil() {
			t.Fatal("expected a non-zero handle")
		}
	})

	t.Run("missing parent", func(t *testing.T) {
		handles := newHandleTable()

		_, err := handles.NewStackConfig((*stackconfig.Config)(nil), handle[*sourcebundle.Bundle](42))
		if err != newHandleErrorNoParent {
			t.Fatalf("wrong error\ngot:  %v\nwant: %v", err, newHandleErrorNoParent)
		}
	})

	t.Run("parent of the wrong type", func(t *testing.T) {
		handles := newHandleTable()
		stateHnd := handles.NewStackState((*stackstate.State)(nil))

		_, err := handles.NewStackConfig((*stackconfig.Config)(nil), handle[*sourcebundle.Bundle](stateHnd))
		if err != newHandleErrorWrongParentType {
			t.Fatalf("wrong error\ngot:  %v\nwant: %v", err, newHandleErrorWrongParentType)
		}
	})
}
