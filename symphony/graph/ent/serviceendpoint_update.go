// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/equipmentport"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/service"
	"github.com/facebookincubator/symphony/graph/ent/serviceendpoint"
)

// ServiceEndpointUpdate is the builder for updating ServiceEndpoint entities.
type ServiceEndpointUpdate struct {
	config

	update_time    *time.Time
	role           *string
	port           map[string]struct{}
	service        map[string]struct{}
	clearedPort    bool
	clearedService bool
	predicates     []predicate.ServiceEndpoint
}

// Where adds a new predicate for the builder.
func (seu *ServiceEndpointUpdate) Where(ps ...predicate.ServiceEndpoint) *ServiceEndpointUpdate {
	seu.predicates = append(seu.predicates, ps...)
	return seu
}

// SetRole sets the role field.
func (seu *ServiceEndpointUpdate) SetRole(s string) *ServiceEndpointUpdate {
	seu.role = &s
	return seu
}

// SetPortID sets the port edge to EquipmentPort by id.
func (seu *ServiceEndpointUpdate) SetPortID(id string) *ServiceEndpointUpdate {
	if seu.port == nil {
		seu.port = make(map[string]struct{})
	}
	seu.port[id] = struct{}{}
	return seu
}

// SetNillablePortID sets the port edge to EquipmentPort by id if the given value is not nil.
func (seu *ServiceEndpointUpdate) SetNillablePortID(id *string) *ServiceEndpointUpdate {
	if id != nil {
		seu = seu.SetPortID(*id)
	}
	return seu
}

// SetPort sets the port edge to EquipmentPort.
func (seu *ServiceEndpointUpdate) SetPort(e *EquipmentPort) *ServiceEndpointUpdate {
	return seu.SetPortID(e.ID)
}

// SetServiceID sets the service edge to Service by id.
func (seu *ServiceEndpointUpdate) SetServiceID(id string) *ServiceEndpointUpdate {
	if seu.service == nil {
		seu.service = make(map[string]struct{})
	}
	seu.service[id] = struct{}{}
	return seu
}

// SetNillableServiceID sets the service edge to Service by id if the given value is not nil.
func (seu *ServiceEndpointUpdate) SetNillableServiceID(id *string) *ServiceEndpointUpdate {
	if id != nil {
		seu = seu.SetServiceID(*id)
	}
	return seu
}

// SetService sets the service edge to Service.
func (seu *ServiceEndpointUpdate) SetService(s *Service) *ServiceEndpointUpdate {
	return seu.SetServiceID(s.ID)
}

// ClearPort clears the port edge to EquipmentPort.
func (seu *ServiceEndpointUpdate) ClearPort() *ServiceEndpointUpdate {
	seu.clearedPort = true
	return seu
}

// ClearService clears the service edge to Service.
func (seu *ServiceEndpointUpdate) ClearService() *ServiceEndpointUpdate {
	seu.clearedService = true
	return seu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (seu *ServiceEndpointUpdate) Save(ctx context.Context) (int, error) {
	if seu.update_time == nil {
		v := serviceendpoint.UpdateDefaultUpdateTime()
		seu.update_time = &v
	}
	if len(seu.port) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"port\"")
	}
	if len(seu.service) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"service\"")
	}
	return seu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (seu *ServiceEndpointUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *ServiceEndpointUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *ServiceEndpointUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seu *ServiceEndpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceendpoint.Table,
			Columns: serviceendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: serviceendpoint.FieldID,
			},
		},
	}
	if ps := seu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := seu.update_time; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: serviceendpoint.FieldUpdateTime,
		})
	}
	if value := seu.role; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: serviceendpoint.FieldRole,
		})
	}
	if seu.clearedPort {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: equipmentport.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := seu.port; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: equipmentport.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if seu.clearedService {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: service.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := seu.service; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: service.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ServiceEndpointUpdateOne is the builder for updating a single ServiceEndpoint entity.
type ServiceEndpointUpdateOne struct {
	config
	id string

	update_time    *time.Time
	role           *string
	port           map[string]struct{}
	service        map[string]struct{}
	clearedPort    bool
	clearedService bool
}

// SetRole sets the role field.
func (seuo *ServiceEndpointUpdateOne) SetRole(s string) *ServiceEndpointUpdateOne {
	seuo.role = &s
	return seuo
}

// SetPortID sets the port edge to EquipmentPort by id.
func (seuo *ServiceEndpointUpdateOne) SetPortID(id string) *ServiceEndpointUpdateOne {
	if seuo.port == nil {
		seuo.port = make(map[string]struct{})
	}
	seuo.port[id] = struct{}{}
	return seuo
}

// SetNillablePortID sets the port edge to EquipmentPort by id if the given value is not nil.
func (seuo *ServiceEndpointUpdateOne) SetNillablePortID(id *string) *ServiceEndpointUpdateOne {
	if id != nil {
		seuo = seuo.SetPortID(*id)
	}
	return seuo
}

// SetPort sets the port edge to EquipmentPort.
func (seuo *ServiceEndpointUpdateOne) SetPort(e *EquipmentPort) *ServiceEndpointUpdateOne {
	return seuo.SetPortID(e.ID)
}

// SetServiceID sets the service edge to Service by id.
func (seuo *ServiceEndpointUpdateOne) SetServiceID(id string) *ServiceEndpointUpdateOne {
	if seuo.service == nil {
		seuo.service = make(map[string]struct{})
	}
	seuo.service[id] = struct{}{}
	return seuo
}

// SetNillableServiceID sets the service edge to Service by id if the given value is not nil.
func (seuo *ServiceEndpointUpdateOne) SetNillableServiceID(id *string) *ServiceEndpointUpdateOne {
	if id != nil {
		seuo = seuo.SetServiceID(*id)
	}
	return seuo
}

// SetService sets the service edge to Service.
func (seuo *ServiceEndpointUpdateOne) SetService(s *Service) *ServiceEndpointUpdateOne {
	return seuo.SetServiceID(s.ID)
}

// ClearPort clears the port edge to EquipmentPort.
func (seuo *ServiceEndpointUpdateOne) ClearPort() *ServiceEndpointUpdateOne {
	seuo.clearedPort = true
	return seuo
}

// ClearService clears the service edge to Service.
func (seuo *ServiceEndpointUpdateOne) ClearService() *ServiceEndpointUpdateOne {
	seuo.clearedService = true
	return seuo
}

// Save executes the query and returns the updated entity.
func (seuo *ServiceEndpointUpdateOne) Save(ctx context.Context) (*ServiceEndpoint, error) {
	if seuo.update_time == nil {
		v := serviceendpoint.UpdateDefaultUpdateTime()
		seuo.update_time = &v
	}
	if len(seuo.port) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"port\"")
	}
	if len(seuo.service) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"service\"")
	}
	return seuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *ServiceEndpointUpdateOne) SaveX(ctx context.Context) *ServiceEndpoint {
	se, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return se
}

// Exec executes the query on the entity.
func (seuo *ServiceEndpointUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *ServiceEndpointUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seuo *ServiceEndpointUpdateOne) sqlSave(ctx context.Context) (se *ServiceEndpoint, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceendpoint.Table,
			Columns: serviceendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  seuo.id,
				Type:   field.TypeString,
				Column: serviceendpoint.FieldID,
			},
		},
	}
	if value := seuo.update_time; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: serviceendpoint.FieldUpdateTime,
		})
	}
	if value := seuo.role; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: serviceendpoint.FieldRole,
		})
	}
	if seuo.clearedPort {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: equipmentport.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := seuo.port; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: equipmentport.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if seuo.clearedService {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: service.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := seuo.service; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: service.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	se = &ServiceEndpoint{config: seuo.config}
	spec.Assign = se.assignValues
	spec.ScanValues = se.scanValues()
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return se, nil
}
