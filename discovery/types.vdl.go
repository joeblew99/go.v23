// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package discovery

import (
	// VDL system imports
	"v.io/v23/vdl"
)

// Service represents service information for service discovery.
type Service struct {
	// The universal unique identifier of a service instance.
	// If this is not specified, a random 128 bit (16 byte) UUID will be used.
	InstanceUuid []byte
	// Optional name of the service instance.
	InstanceName string
	// The interface that the service implements.
	// E.g., 'v.io/v23/services/vtrace.Store'.
	InterfaceName string
	// The service attributes.
	// E.g., {'resolution': '1024x768'}.
	Attrs Attributes
	// The addresses (vanadium object names) that the service is served on.
	// E.g., '/host:port/a/b/c', '/ns.dev.v.io:8101/blah/blah'.
	Addrs []string
}

func (Service) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Service"`
}) {
}

// Attributes represents service attributes as a key/value pair.
//
// The key must be US-ASCII printable characters, excluding the '=' character
// and should not start with '_' character.
type Attributes map[string]string

func (Attributes) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attributes"`
}) {
}

// Found represents a service that is discovered by scan.
type Found struct {
	Service Service
}

func (Found) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Found"`
}) {
}

// Lost represents a service that is lost during scan.
type Lost struct {
	InstanceUuid []byte
}

func (Lost) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Lost"`
}) {
}

type (
	// Update represents any single field of the Update union type.
	//
	// Update represents a discovery update.
	Update interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Update union type.
		__VDLReflect(__UpdateReflect)
	}
	// UpdateFound represents field Found of the Update union type.
	UpdateFound struct{ Value Found }
	// UpdateLost represents field Lost of the Update union type.
	UpdateLost struct{ Value Lost }
	// __UpdateReflect describes the Update union type.
	__UpdateReflect struct {
		Name  string `vdl:"v.io/v23/discovery.Update"`
		Type  Update
		Union struct {
			Found UpdateFound
			Lost  UpdateLost
		}
	}
)

func (x UpdateFound) Index() int                   { return 0 }
func (x UpdateFound) Interface() interface{}       { return x.Value }
func (x UpdateFound) Name() string                 { return "Found" }
func (x UpdateFound) __VDLReflect(__UpdateReflect) {}

func (x UpdateLost) Index() int                   { return 1 }
func (x UpdateLost) Interface() interface{}       { return x.Value }
func (x UpdateLost) Name() string                 { return "Lost" }
func (x UpdateLost) __VDLReflect(__UpdateReflect) {}

func init() {
	vdl.Register((*Service)(nil))
	vdl.Register((*Attributes)(nil))
	vdl.Register((*Found)(nil))
	vdl.Register((*Lost)(nil))
	vdl.Register((*Update)(nil))
}
