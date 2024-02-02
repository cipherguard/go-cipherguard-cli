package group

import (
	"context"
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/khulnasoft/go-cipherguard/api"
)

// Environments for CEl
var celEnvOptions = []cel.EnvOption{
	cel.Variable("ID", cel.StringType),
	cel.Variable("Name", cel.StringType),
	cel.Variable("CreatedTimestamp", cel.TimestampType),
	cel.Variable("ModifiedTimestamp", cel.TimestampType),
}

// Filters the slice groups by invoke CEL program for each group
func filterGroups(groups *[]api.Group, celCmd string, ctx context.Context) ([]api.Group, error) {
	if celCmd == "" {
		return *groups, nil
	}

	program, err := util.InitCELProgram(celCmd, celEnvOptions...)
	if err != nil {
		return nil, err
	}

	filteredGroups := []api.Group{}
	for _, group := range *groups {
		val, _, err := (*program).ContextEval(ctx, map[string]any{
			"ID":                group.ID,
			"Name":              group.Name,
			"CreatedTimestamp":  group.Created.Time,
			"ModifiedTimestamp": group.Modified.Time,
		})

		if err != nil {
			return nil, err
		}

		if val.Value() == true {
			filteredGroups = append(filteredGroups, group)
		}
	}

	if len(filteredGroups) == 0 {
		return nil, fmt.Errorf("No such groups found with filter %v!", celCmd)
	}

	return filteredGroups, nil
}
