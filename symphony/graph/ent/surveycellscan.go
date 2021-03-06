// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/surveycellscan"
)

// SurveyCellScan is the model entity for the SurveyCellScan schema.
type SurveyCellScan struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// NetworkType holds the value of the "network_type" field.
	NetworkType string `json:"network_type,omitempty"`
	// SignalStrength holds the value of the "signal_strength" field.
	SignalStrength int `json:"signal_strength,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// BaseStationID holds the value of the "base_station_id" field.
	BaseStationID string `json:"base_station_id,omitempty"`
	// NetworkID holds the value of the "network_id" field.
	NetworkID string `json:"network_id,omitempty"`
	// SystemID holds the value of the "system_id" field.
	SystemID string `json:"system_id,omitempty"`
	// CellID holds the value of the "cell_id" field.
	CellID string `json:"cell_id,omitempty"`
	// LocationAreaCode holds the value of the "location_area_code" field.
	LocationAreaCode string `json:"location_area_code,omitempty"`
	// MobileCountryCode holds the value of the "mobile_country_code" field.
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	// MobileNetworkCode holds the value of the "mobile_network_code" field.
	MobileNetworkCode string `json:"mobile_network_code,omitempty"`
	// PrimaryScramblingCode holds the value of the "primary_scrambling_code" field.
	PrimaryScramblingCode string `json:"primary_scrambling_code,omitempty"`
	// Operator holds the value of the "operator" field.
	Operator string `json:"operator,omitempty"`
	// Arfcn holds the value of the "arfcn" field.
	Arfcn int `json:"arfcn,omitempty"`
	// PhysicalCellID holds the value of the "physical_cell_id" field.
	PhysicalCellID string `json:"physical_cell_id,omitempty"`
	// TrackingAreaCode holds the value of the "tracking_area_code" field.
	TrackingAreaCode string `json:"tracking_area_code,omitempty"`
	// TimingAdvance holds the value of the "timing_advance" field.
	TimingAdvance int `json:"timing_advance,omitempty"`
	// Earfcn holds the value of the "earfcn" field.
	Earfcn int `json:"earfcn,omitempty"`
	// Uarfcn holds the value of the "uarfcn" field.
	Uarfcn int `json:"uarfcn,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SurveyCellScan) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullTime{},
		&sql.NullTime{},
		&sql.NullString{},
		&sql.NullInt64{},
		&sql.NullTime{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullInt64{},
		&sql.NullString{},
		&sql.NullString{},
		&sql.NullInt64{},
		&sql.NullInt64{},
		&sql.NullInt64{},
		&sql.NullFloat64{},
		&sql.NullFloat64{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SurveyCellScan fields.
func (scs *SurveyCellScan) assignValues(values ...interface{}) error {
	if m, n := len(values), len(surveycellscan.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	scs.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		scs.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		scs.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field network_type", values[2])
	} else if value.Valid {
		scs.NetworkType = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field signal_strength", values[3])
	} else if value.Valid {
		scs.SignalStrength = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field timestamp", values[4])
	} else if value.Valid {
		scs.Timestamp = value.Time
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field base_station_id", values[5])
	} else if value.Valid {
		scs.BaseStationID = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field network_id", values[6])
	} else if value.Valid {
		scs.NetworkID = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field system_id", values[7])
	} else if value.Valid {
		scs.SystemID = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cell_id", values[8])
	} else if value.Valid {
		scs.CellID = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field location_area_code", values[9])
	} else if value.Valid {
		scs.LocationAreaCode = value.String
	}
	if value, ok := values[10].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mobile_country_code", values[10])
	} else if value.Valid {
		scs.MobileCountryCode = value.String
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mobile_network_code", values[11])
	} else if value.Valid {
		scs.MobileNetworkCode = value.String
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field primary_scrambling_code", values[12])
	} else if value.Valid {
		scs.PrimaryScramblingCode = value.String
	}
	if value, ok := values[13].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field operator", values[13])
	} else if value.Valid {
		scs.Operator = value.String
	}
	if value, ok := values[14].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field arfcn", values[14])
	} else if value.Valid {
		scs.Arfcn = int(value.Int64)
	}
	if value, ok := values[15].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field physical_cell_id", values[15])
	} else if value.Valid {
		scs.PhysicalCellID = value.String
	}
	if value, ok := values[16].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tracking_area_code", values[16])
	} else if value.Valid {
		scs.TrackingAreaCode = value.String
	}
	if value, ok := values[17].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field timing_advance", values[17])
	} else if value.Valid {
		scs.TimingAdvance = int(value.Int64)
	}
	if value, ok := values[18].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field earfcn", values[18])
	} else if value.Valid {
		scs.Earfcn = int(value.Int64)
	}
	if value, ok := values[19].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field uarfcn", values[19])
	} else if value.Valid {
		scs.Uarfcn = int(value.Int64)
	}
	if value, ok := values[20].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field latitude", values[20])
	} else if value.Valid {
		scs.Latitude = value.Float64
	}
	if value, ok := values[21].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field longitude", values[21])
	} else if value.Valid {
		scs.Longitude = value.Float64
	}
	return nil
}

// QuerySurveyQuestion queries the survey_question edge of the SurveyCellScan.
func (scs *SurveyCellScan) QuerySurveyQuestion() *SurveyQuestionQuery {
	return (&SurveyCellScanClient{scs.config}).QuerySurveyQuestion(scs)
}

// QueryLocation queries the location edge of the SurveyCellScan.
func (scs *SurveyCellScan) QueryLocation() *LocationQuery {
	return (&SurveyCellScanClient{scs.config}).QueryLocation(scs)
}

// Update returns a builder for updating this SurveyCellScan.
// Note that, you need to call SurveyCellScan.Unwrap() before calling this method, if this SurveyCellScan
// was returned from a transaction, and the transaction was committed or rolled back.
func (scs *SurveyCellScan) Update() *SurveyCellScanUpdateOne {
	return (&SurveyCellScanClient{scs.config}).UpdateOne(scs)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (scs *SurveyCellScan) Unwrap() *SurveyCellScan {
	tx, ok := scs.config.driver.(*txDriver)
	if !ok {
		panic("ent: SurveyCellScan is not a transactional entity")
	}
	scs.config.driver = tx.drv
	return scs
}

// String implements the fmt.Stringer.
func (scs *SurveyCellScan) String() string {
	var builder strings.Builder
	builder.WriteString("SurveyCellScan(")
	builder.WriteString(fmt.Sprintf("id=%v", scs.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(scs.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(scs.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", network_type=")
	builder.WriteString(scs.NetworkType)
	builder.WriteString(", signal_strength=")
	builder.WriteString(fmt.Sprintf("%v", scs.SignalStrength))
	builder.WriteString(", timestamp=")
	builder.WriteString(scs.Timestamp.Format(time.ANSIC))
	builder.WriteString(", base_station_id=")
	builder.WriteString(scs.BaseStationID)
	builder.WriteString(", network_id=")
	builder.WriteString(scs.NetworkID)
	builder.WriteString(", system_id=")
	builder.WriteString(scs.SystemID)
	builder.WriteString(", cell_id=")
	builder.WriteString(scs.CellID)
	builder.WriteString(", location_area_code=")
	builder.WriteString(scs.LocationAreaCode)
	builder.WriteString(", mobile_country_code=")
	builder.WriteString(scs.MobileCountryCode)
	builder.WriteString(", mobile_network_code=")
	builder.WriteString(scs.MobileNetworkCode)
	builder.WriteString(", primary_scrambling_code=")
	builder.WriteString(scs.PrimaryScramblingCode)
	builder.WriteString(", operator=")
	builder.WriteString(scs.Operator)
	builder.WriteString(", arfcn=")
	builder.WriteString(fmt.Sprintf("%v", scs.Arfcn))
	builder.WriteString(", physical_cell_id=")
	builder.WriteString(scs.PhysicalCellID)
	builder.WriteString(", tracking_area_code=")
	builder.WriteString(scs.TrackingAreaCode)
	builder.WriteString(", timing_advance=")
	builder.WriteString(fmt.Sprintf("%v", scs.TimingAdvance))
	builder.WriteString(", earfcn=")
	builder.WriteString(fmt.Sprintf("%v", scs.Earfcn))
	builder.WriteString(", uarfcn=")
	builder.WriteString(fmt.Sprintf("%v", scs.Uarfcn))
	builder.WriteString(", latitude=")
	builder.WriteString(fmt.Sprintf("%v", scs.Latitude))
	builder.WriteString(", longitude=")
	builder.WriteString(fmt.Sprintf("%v", scs.Longitude))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (scs *SurveyCellScan) id() int {
	id, _ := strconv.Atoi(scs.ID)
	return id
}

// SurveyCellScans is a parsable slice of SurveyCellScan.
type SurveyCellScans []*SurveyCellScan

func (scs SurveyCellScans) config(cfg config) {
	for _i := range scs {
		scs[_i].config = cfg
	}
}
