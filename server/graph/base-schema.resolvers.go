package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/opendatahub-io/model-registry/model/graph"
)

// Type is the resolver for the type field.
func (r *artifactResolver) Type(ctx context.Context, obj *graph.Artifact, filter *graph.InstanceFilter) (*graph.ArtifactType, error) {
	id := "1"
	return &graph.ArtifactType{ID: &id, Name: "TestType"}, nil
}

// Artifact returns ArtifactResolver implementation.
func (r *Resolver) Artifact() ArtifactResolver { return &artifactResolver{r} }

type artifactResolver struct{ *Resolver }