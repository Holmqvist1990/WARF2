// Package job provides the interface
// implementations of Job to be used
// in JobService.
package job

import (
	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

const (
	finished   = true
	unfinished = false
)

// Job declares the common interface
// for jobs, in order to be used within
// the job service.
type Job interface {
	NeedsToBeRemoved(*m.Map) bool
	PerformWork(*m.Map, []*dwarf.Dwarf) bool
	Finish(*m.Map, *room.Service)
	Priority() int

	GetWorker() *dwarf.Dwarf
	SetWorker(*dwarf.Dwarf)
	GetDestinations() []int
	String() string
}
