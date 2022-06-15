// code generated by "./internal/cmd/gen/generic", DON'T EDIT IT.

package generic

import (
	"database/sql/driver"

	"github.com/bangumi/server/internal/model"
)

func SubjectIDToValuerSlice(a []model.SubjectID) []driver.Valuer {
	b := make([]driver.Valuer, len(a))
	for i := range a {
		b[i] = a[i]
	}

	return b
}

func PersonIDToValuerSlice(a []model.PersonID) []driver.Valuer {
	b := make([]driver.Valuer, len(a))
	for i := range a {
		b[i] = a[i]
	}

	return b
}

func UserIDToValuerSlice(a []model.UserID) []driver.Valuer {
	b := make([]driver.Valuer, len(a))
	for i := range a {
		b[i] = a[i]
	}

	return b
}

func CharacterIDToValuerSlice(a []model.CharacterID) []driver.Valuer {
	b := make([]driver.Valuer, len(a))
	for i := range a {
		b[i] = a[i]
	}

	return b
}

func EpisodeIDToValuerSlice(a []model.EpisodeID) []driver.Valuer {
	b := make([]driver.Valuer, len(a))
	for i := range a {
		b[i] = a[i]
	}

	return b
}
