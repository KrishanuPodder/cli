package v7pushaction

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/types"
)

type PushPlan struct {
	Application v7action.Application
	SpaceGUID   string
	OrgGUID     string
	Overrides   FlagOverrides
	Manifest    []byte

	Archive                bool
	ApplicationNeedsUpdate bool
	BitsPath               string
	AllResources           []sharedaction.Resource
	MatchedResources       []sharedaction.Resource
	UnmatchedResources     []sharedaction.Resource
}

type FlagOverrides struct {
	Buildpacks          []string
	Stack               string
	Disk                types.NullUint64
	DockerImage         string
	DockerPassword      string
	DockerUsername      string
	HealthCheckEndpoint string
	HealthCheckTimeout  int64
	HealthCheckType     constant.HealthCheckType
	Instances           types.NullInt
	Memory              types.NullUint64
	NoStart             bool
	ProvidedAppPath     string
	SkipRouteCreation   bool
	StartCommand        types.FilteredString
}

func (state PushPlan) String() string {
	return fmt.Sprintf(
		"Application: %#v - Space GUID: %s, Org GUID: %s, Archive: %t, Bits Path: %s",
		state.Application,
		state.SpaceGUID,
		state.OrgGUID,
		state.Archive,
		state.BitsPath,
	)
}
