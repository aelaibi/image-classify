//go:generate go run github.com/99designs/gqlgen
package image_classify

import (
	"context"
	"fmt"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	projects []*Project
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Project() ProjectResolver {
	return &projectResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProject(ctx context.Context, input NewProject) (*Project, error) {
	project := &Project{
		Name:   input.Name,
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.projects = append(r.projects, project)
	return project, nil
}

type projectResolver struct{ *Resolver }

func (r *projectResolver) User(ctx context.Context, obj *Project) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Projects(ctx context.Context) ([]*Project, error) {
	return r.projects, nil
}
