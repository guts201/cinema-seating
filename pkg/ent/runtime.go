// Code generated by ent, DO NOT EDIT.

package ent

import (
	cinema "cinema/api"
	entcinema "cinema/pkg/ent/cinema"
	"cinema/pkg/ent/movie"
	"cinema/pkg/ent/screening"
	"cinema/pkg/ent/seatreservation"
	"cinema/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entcinemaMixin := schema.Cinema{}.Mixin()
	entcinemaMixinFields0 := entcinemaMixin[0].Fields()
	_ = entcinemaMixinFields0
	entcinemaFields := schema.Cinema{}.Fields()
	_ = entcinemaFields
	// entcinemaDescCreatedAt is the schema descriptor for created_at field.
	entcinemaDescCreatedAt := entcinemaMixinFields0[1].Descriptor()
	// entcinema.DefaultCreatedAt holds the default value on creation for the created_at field.
	entcinema.DefaultCreatedAt = entcinemaDescCreatedAt.Default.(func() time.Time)
	// entcinemaDescUpdatedAt is the schema descriptor for updated_at field.
	entcinemaDescUpdatedAt := entcinemaMixinFields0[2].Descriptor()
	// entcinema.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entcinema.DefaultUpdatedAt = entcinemaDescUpdatedAt.Default.(func() time.Time)
	// entcinema.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entcinema.UpdateDefaultUpdatedAt = entcinemaDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entcinemaDescNumRow is the schema descriptor for num_row field.
	entcinemaDescNumRow := entcinemaFields[0].Descriptor()
	// entcinema.NumRowValidator is a validator for the "num_row" field. It is called by the builders before save.
	entcinema.NumRowValidator = entcinemaDescNumRow.Validators[0].(func(uint32) error)
	// entcinemaDescNumColumn is the schema descriptor for num_column field.
	entcinemaDescNumColumn := entcinemaFields[1].Descriptor()
	// entcinema.NumColumnValidator is a validator for the "num_column" field. It is called by the builders before save.
	entcinema.NumColumnValidator = entcinemaDescNumColumn.Validators[0].(func(uint32) error)
	// entcinemaDescName is the schema descriptor for name field.
	entcinemaDescName := entcinemaFields[2].Descriptor()
	// entcinema.NameValidator is a validator for the "name" field. It is called by the builders before save.
	entcinema.NameValidator = entcinemaDescName.Validators[0].(func(string) error)
	// entcinemaDescMinDistance is the schema descriptor for min_distance field.
	entcinemaDescMinDistance := entcinemaFields[3].Descriptor()
	// entcinema.MinDistanceValidator is a validator for the "min_distance" field. It is called by the builders before save.
	entcinema.MinDistanceValidator = entcinemaDescMinDistance.Validators[0].(func(uint32) error)
	movieMixin := schema.Movie{}.Mixin()
	movieMixinFields0 := movieMixin[0].Fields()
	_ = movieMixinFields0
	movieFields := schema.Movie{}.Fields()
	_ = movieFields
	// movieDescCreatedAt is the schema descriptor for created_at field.
	movieDescCreatedAt := movieMixinFields0[1].Descriptor()
	// movie.DefaultCreatedAt holds the default value on creation for the created_at field.
	movie.DefaultCreatedAt = movieDescCreatedAt.Default.(func() time.Time)
	// movieDescUpdatedAt is the schema descriptor for updated_at field.
	movieDescUpdatedAt := movieMixinFields0[2].Descriptor()
	// movie.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	movie.DefaultUpdatedAt = movieDescUpdatedAt.Default.(func() time.Time)
	// movie.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	movie.UpdateDefaultUpdatedAt = movieDescUpdatedAt.UpdateDefault.(func() time.Time)
	// movieDescTitle is the schema descriptor for title field.
	movieDescTitle := movieFields[0].Descriptor()
	// movie.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	movie.TitleValidator = movieDescTitle.Validators[0].(func(string) error)
	screeningMixin := schema.Screening{}.Mixin()
	screeningMixinFields0 := screeningMixin[0].Fields()
	_ = screeningMixinFields0
	screeningFields := schema.Screening{}.Fields()
	_ = screeningFields
	// screeningDescCreatedAt is the schema descriptor for created_at field.
	screeningDescCreatedAt := screeningMixinFields0[1].Descriptor()
	// screening.DefaultCreatedAt holds the default value on creation for the created_at field.
	screening.DefaultCreatedAt = screeningDescCreatedAt.Default.(func() time.Time)
	// screeningDescUpdatedAt is the schema descriptor for updated_at field.
	screeningDescUpdatedAt := screeningMixinFields0[2].Descriptor()
	// screening.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	screening.DefaultUpdatedAt = screeningDescUpdatedAt.Default.(func() time.Time)
	// screening.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	screening.UpdateDefaultUpdatedAt = screeningDescUpdatedAt.UpdateDefault.(func() time.Time)
	// screeningDescTitle is the schema descriptor for title field.
	screeningDescTitle := screeningFields[0].Descriptor()
	// screening.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	screening.TitleValidator = screeningDescTitle.Validators[0].(func(string) error)
	seatreservationMixin := schema.SeatReservation{}.Mixin()
	seatreservationMixinFields0 := seatreservationMixin[0].Fields()
	_ = seatreservationMixinFields0
	seatreservationFields := schema.SeatReservation{}.Fields()
	_ = seatreservationFields
	// seatreservationDescCreatedAt is the schema descriptor for created_at field.
	seatreservationDescCreatedAt := seatreservationMixinFields0[1].Descriptor()
	// seatreservation.DefaultCreatedAt holds the default value on creation for the created_at field.
	seatreservation.DefaultCreatedAt = seatreservationDescCreatedAt.Default.(func() time.Time)
	// seatreservationDescUpdatedAt is the schema descriptor for updated_at field.
	seatreservationDescUpdatedAt := seatreservationMixinFields0[2].Descriptor()
	// seatreservation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	seatreservation.DefaultUpdatedAt = seatreservationDescUpdatedAt.Default.(func() time.Time)
	// seatreservation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	seatreservation.UpdateDefaultUpdatedAt = seatreservationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// seatreservationDescReservedAt is the schema descriptor for reserved_at field.
	seatreservationDescReservedAt := seatreservationFields[0].Descriptor()
	// seatreservation.DefaultReservedAt holds the default value on creation for the reserved_at field.
	seatreservation.DefaultReservedAt = seatreservationDescReservedAt.Default.(func() time.Time)
	// seatreservationDescStatus is the schema descriptor for status field.
	seatreservationDescStatus := seatreservationFields[2].Descriptor()
	// seatreservation.DefaultStatus holds the default value on creation for the status field.
	seatreservation.DefaultStatus = cinema.SeatReservationStatus(seatreservationDescStatus.Default.(int32))
	// seatreservationDescRowNum is the schema descriptor for row_num field.
	seatreservationDescRowNum := seatreservationFields[3].Descriptor()
	// seatreservation.RowNumValidator is a validator for the "row_num" field. It is called by the builders before save.
	seatreservation.RowNumValidator = seatreservationDescRowNum.Validators[0].(func(uint32) error)
	// seatreservationDescColumnNum is the schema descriptor for column_num field.
	seatreservationDescColumnNum := seatreservationFields[4].Descriptor()
	// seatreservation.ColumnNumValidator is a validator for the "column_num" field. It is called by the builders before save.
	seatreservation.ColumnNumValidator = seatreservationDescColumnNum.Validators[0].(func(uint32) error)
}
