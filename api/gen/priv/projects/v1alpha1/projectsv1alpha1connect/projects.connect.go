//
// The projects service is written as a model endpoint that new endpoints should emulate.
// The comments in this file are written such that API documentation can be
// automatically generated from it.
// If you're adding an 'internal' comment, don't place it next to a field, method, or message.

//*
// API to manage projects

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: priv/projects/v1alpha1/projects.proto

package projectsv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "go.jetpack.io/pkg/api/gen/priv/projects/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProjectsServiceName is the fully-qualified name of the ProjectsService service.
	ProjectsServiceName = "priv.projects.v1alpha1.ProjectsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProjectsServiceGetProjectProcedure is the fully-qualified name of the ProjectsService's
	// GetProject RPC.
	ProjectsServiceGetProjectProcedure = "/priv.projects.v1alpha1.ProjectsService/GetProject"
	// ProjectsServiceListProjectsProcedure is the fully-qualified name of the ProjectsService's
	// ListProjects RPC.
	ProjectsServiceListProjectsProcedure = "/priv.projects.v1alpha1.ProjectsService/ListProjects"
	// ProjectsServiceSearchProjectsProcedure is the fully-qualified name of the ProjectsService's
	// SearchProjects RPC.
	ProjectsServiceSearchProjectsProcedure = "/priv.projects.v1alpha1.ProjectsService/SearchProjects"
	// ProjectsServiceCreateProjectProcedure is the fully-qualified name of the ProjectsService's
	// CreateProject RPC.
	ProjectsServiceCreateProjectProcedure = "/priv.projects.v1alpha1.ProjectsService/CreateProject"
	// ProjectsServiceDeleteProjectProcedure is the fully-qualified name of the ProjectsService's
	// DeleteProject RPC.
	ProjectsServiceDeleteProjectProcedure = "/priv.projects.v1alpha1.ProjectsService/DeleteProject"
	// ProjectsServicePatchProjectProcedure is the fully-qualified name of the ProjectsService's
	// PatchProject RPC.
	ProjectsServicePatchProjectProcedure = "/priv.projects.v1alpha1.ProjectsService/PatchProject"
	// ProjectsServiceUpdateProjectProcedure is the fully-qualified name of the ProjectsService's
	// UpdateProject RPC.
	ProjectsServiceUpdateProjectProcedure = "/priv.projects.v1alpha1.ProjectsService/UpdateProject"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	projectsServiceServiceDescriptor              = v1alpha1.File_priv_projects_v1alpha1_projects_proto.Services().ByName("ProjectsService")
	projectsServiceGetProjectMethodDescriptor     = projectsServiceServiceDescriptor.Methods().ByName("GetProject")
	projectsServiceListProjectsMethodDescriptor   = projectsServiceServiceDescriptor.Methods().ByName("ListProjects")
	projectsServiceSearchProjectsMethodDescriptor = projectsServiceServiceDescriptor.Methods().ByName("SearchProjects")
	projectsServiceCreateProjectMethodDescriptor  = projectsServiceServiceDescriptor.Methods().ByName("CreateProject")
	projectsServiceDeleteProjectMethodDescriptor  = projectsServiceServiceDescriptor.Methods().ByName("DeleteProject")
	projectsServicePatchProjectMethodDescriptor   = projectsServiceServiceDescriptor.Methods().ByName("PatchProject")
	projectsServiceUpdateProjectMethodDescriptor  = projectsServiceServiceDescriptor.Methods().ByName("UpdateProject")
)

// ProjectsServiceClient is a client for the priv.projects.v1alpha1.ProjectsService service.
type ProjectsServiceClient interface {
	// Get a project
	//
	// Retrieves the details of an existing project identified by its unique
	// project id.
	GetProject(context.Context, *connect.Request[v1alpha1.GetProjectRequest]) (*connect.Response[v1alpha1.GetProjectResponse], error)
	// List the projects in an organization
	//
	// Lists the projects belonging to the given organization. The projects are
	// sorted by creation date, with the most recently created projects appearing
	// first.
	ListProjects(context.Context, *connect.Request[v1alpha1.ListProjectsRequest]) (*connect.Response[v1alpha1.ListProjectsResponse], error)
	// Search for projects in an organization
	//
	// Searches for products previously created in the given organization.
	// Don't use search in read-after-write flows where strict consistency is
	// necessary.
	SearchProjects(context.Context, *connect.Request[v1alpha1.SearchProjectsRequest]) (*connect.Response[v1alpha1.SearchProjectsResponse], error)
	// Create a new project
	//
	// Creates a new project in the specified organization. The authenticated user
	// must be a member of the organization.
	CreateProject(context.Context, *connect.Request[v1alpha1.CreateProjectRequest]) (*connect.Response[v1alpha1.CreateProjectResponse], error)
	// Delete a project
	//
	// Deletes the project specified by the given id.
	DeleteProject(context.Context, *connect.Request[v1alpha1.DeleteProjectRequest]) (*connect.Response[v1alpha1.DeleteProjectResponse], error)
	// Patch a project
	//
	// Patches the specified project with the provided fields. Any fields that
	// are not provided, will be left unchanged.
	PatchProject(context.Context, *connect.Request[v1alpha1.PatchProjectRequest]) (*connect.Response[v1alpha1.PatchProjectResponse], error)
	// Update a project
	//
	// Updates the specified project by setting the values of the provided fields.
	// All fields will be updates. If you'd like to partially update some fields,
	// use Patch instead.
	UpdateProject(context.Context, *connect.Request[v1alpha1.UpdateProjectRequest]) (*connect.Response[v1alpha1.UpdateProjectResponse], error)
}

// NewProjectsServiceClient constructs a client for the priv.projects.v1alpha1.ProjectsService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProjectsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &projectsServiceClient{
		getProject: connect.NewClient[v1alpha1.GetProjectRequest, v1alpha1.GetProjectResponse](
			httpClient,
			baseURL+ProjectsServiceGetProjectProcedure,
			connect.WithSchema(projectsServiceGetProjectMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listProjects: connect.NewClient[v1alpha1.ListProjectsRequest, v1alpha1.ListProjectsResponse](
			httpClient,
			baseURL+ProjectsServiceListProjectsProcedure,
			connect.WithSchema(projectsServiceListProjectsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		searchProjects: connect.NewClient[v1alpha1.SearchProjectsRequest, v1alpha1.SearchProjectsResponse](
			httpClient,
			baseURL+ProjectsServiceSearchProjectsProcedure,
			connect.WithSchema(projectsServiceSearchProjectsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createProject: connect.NewClient[v1alpha1.CreateProjectRequest, v1alpha1.CreateProjectResponse](
			httpClient,
			baseURL+ProjectsServiceCreateProjectProcedure,
			connect.WithSchema(projectsServiceCreateProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteProject: connect.NewClient[v1alpha1.DeleteProjectRequest, v1alpha1.DeleteProjectResponse](
			httpClient,
			baseURL+ProjectsServiceDeleteProjectProcedure,
			connect.WithSchema(projectsServiceDeleteProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		patchProject: connect.NewClient[v1alpha1.PatchProjectRequest, v1alpha1.PatchProjectResponse](
			httpClient,
			baseURL+ProjectsServicePatchProjectProcedure,
			connect.WithSchema(projectsServicePatchProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateProject: connect.NewClient[v1alpha1.UpdateProjectRequest, v1alpha1.UpdateProjectResponse](
			httpClient,
			baseURL+ProjectsServiceUpdateProjectProcedure,
			connect.WithSchema(projectsServiceUpdateProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// projectsServiceClient implements ProjectsServiceClient.
type projectsServiceClient struct {
	getProject     *connect.Client[v1alpha1.GetProjectRequest, v1alpha1.GetProjectResponse]
	listProjects   *connect.Client[v1alpha1.ListProjectsRequest, v1alpha1.ListProjectsResponse]
	searchProjects *connect.Client[v1alpha1.SearchProjectsRequest, v1alpha1.SearchProjectsResponse]
	createProject  *connect.Client[v1alpha1.CreateProjectRequest, v1alpha1.CreateProjectResponse]
	deleteProject  *connect.Client[v1alpha1.DeleteProjectRequest, v1alpha1.DeleteProjectResponse]
	patchProject   *connect.Client[v1alpha1.PatchProjectRequest, v1alpha1.PatchProjectResponse]
	updateProject  *connect.Client[v1alpha1.UpdateProjectRequest, v1alpha1.UpdateProjectResponse]
}

// GetProject calls priv.projects.v1alpha1.ProjectsService.GetProject.
func (c *projectsServiceClient) GetProject(ctx context.Context, req *connect.Request[v1alpha1.GetProjectRequest]) (*connect.Response[v1alpha1.GetProjectResponse], error) {
	return c.getProject.CallUnary(ctx, req)
}

// ListProjects calls priv.projects.v1alpha1.ProjectsService.ListProjects.
func (c *projectsServiceClient) ListProjects(ctx context.Context, req *connect.Request[v1alpha1.ListProjectsRequest]) (*connect.Response[v1alpha1.ListProjectsResponse], error) {
	return c.listProjects.CallUnary(ctx, req)
}

// SearchProjects calls priv.projects.v1alpha1.ProjectsService.SearchProjects.
func (c *projectsServiceClient) SearchProjects(ctx context.Context, req *connect.Request[v1alpha1.SearchProjectsRequest]) (*connect.Response[v1alpha1.SearchProjectsResponse], error) {
	return c.searchProjects.CallUnary(ctx, req)
}

// CreateProject calls priv.projects.v1alpha1.ProjectsService.CreateProject.
func (c *projectsServiceClient) CreateProject(ctx context.Context, req *connect.Request[v1alpha1.CreateProjectRequest]) (*connect.Response[v1alpha1.CreateProjectResponse], error) {
	return c.createProject.CallUnary(ctx, req)
}

// DeleteProject calls priv.projects.v1alpha1.ProjectsService.DeleteProject.
func (c *projectsServiceClient) DeleteProject(ctx context.Context, req *connect.Request[v1alpha1.DeleteProjectRequest]) (*connect.Response[v1alpha1.DeleteProjectResponse], error) {
	return c.deleteProject.CallUnary(ctx, req)
}

// PatchProject calls priv.projects.v1alpha1.ProjectsService.PatchProject.
func (c *projectsServiceClient) PatchProject(ctx context.Context, req *connect.Request[v1alpha1.PatchProjectRequest]) (*connect.Response[v1alpha1.PatchProjectResponse], error) {
	return c.patchProject.CallUnary(ctx, req)
}

// UpdateProject calls priv.projects.v1alpha1.ProjectsService.UpdateProject.
func (c *projectsServiceClient) UpdateProject(ctx context.Context, req *connect.Request[v1alpha1.UpdateProjectRequest]) (*connect.Response[v1alpha1.UpdateProjectResponse], error) {
	return c.updateProject.CallUnary(ctx, req)
}

// ProjectsServiceHandler is an implementation of the priv.projects.v1alpha1.ProjectsService
// service.
type ProjectsServiceHandler interface {
	// Get a project
	//
	// Retrieves the details of an existing project identified by its unique
	// project id.
	GetProject(context.Context, *connect.Request[v1alpha1.GetProjectRequest]) (*connect.Response[v1alpha1.GetProjectResponse], error)
	// List the projects in an organization
	//
	// Lists the projects belonging to the given organization. The projects are
	// sorted by creation date, with the most recently created projects appearing
	// first.
	ListProjects(context.Context, *connect.Request[v1alpha1.ListProjectsRequest]) (*connect.Response[v1alpha1.ListProjectsResponse], error)
	// Search for projects in an organization
	//
	// Searches for products previously created in the given organization.
	// Don't use search in read-after-write flows where strict consistency is
	// necessary.
	SearchProjects(context.Context, *connect.Request[v1alpha1.SearchProjectsRequest]) (*connect.Response[v1alpha1.SearchProjectsResponse], error)
	// Create a new project
	//
	// Creates a new project in the specified organization. The authenticated user
	// must be a member of the organization.
	CreateProject(context.Context, *connect.Request[v1alpha1.CreateProjectRequest]) (*connect.Response[v1alpha1.CreateProjectResponse], error)
	// Delete a project
	//
	// Deletes the project specified by the given id.
	DeleteProject(context.Context, *connect.Request[v1alpha1.DeleteProjectRequest]) (*connect.Response[v1alpha1.DeleteProjectResponse], error)
	// Patch a project
	//
	// Patches the specified project with the provided fields. Any fields that
	// are not provided, will be left unchanged.
	PatchProject(context.Context, *connect.Request[v1alpha1.PatchProjectRequest]) (*connect.Response[v1alpha1.PatchProjectResponse], error)
	// Update a project
	//
	// Updates the specified project by setting the values of the provided fields.
	// All fields will be updates. If you'd like to partially update some fields,
	// use Patch instead.
	UpdateProject(context.Context, *connect.Request[v1alpha1.UpdateProjectRequest]) (*connect.Response[v1alpha1.UpdateProjectResponse], error)
}

// NewProjectsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectsServiceHandler(svc ProjectsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	projectsServiceGetProjectHandler := connect.NewUnaryHandler(
		ProjectsServiceGetProjectProcedure,
		svc.GetProject,
		connect.WithSchema(projectsServiceGetProjectMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	projectsServiceListProjectsHandler := connect.NewUnaryHandler(
		ProjectsServiceListProjectsProcedure,
		svc.ListProjects,
		connect.WithSchema(projectsServiceListProjectsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	projectsServiceSearchProjectsHandler := connect.NewUnaryHandler(
		ProjectsServiceSearchProjectsProcedure,
		svc.SearchProjects,
		connect.WithSchema(projectsServiceSearchProjectsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	projectsServiceCreateProjectHandler := connect.NewUnaryHandler(
		ProjectsServiceCreateProjectProcedure,
		svc.CreateProject,
		connect.WithSchema(projectsServiceCreateProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectsServiceDeleteProjectHandler := connect.NewUnaryHandler(
		ProjectsServiceDeleteProjectProcedure,
		svc.DeleteProject,
		connect.WithSchema(projectsServiceDeleteProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectsServicePatchProjectHandler := connect.NewUnaryHandler(
		ProjectsServicePatchProjectProcedure,
		svc.PatchProject,
		connect.WithSchema(projectsServicePatchProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	projectsServiceUpdateProjectHandler := connect.NewUnaryHandler(
		ProjectsServiceUpdateProjectProcedure,
		svc.UpdateProject,
		connect.WithSchema(projectsServiceUpdateProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/priv.projects.v1alpha1.ProjectsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProjectsServiceGetProjectProcedure:
			projectsServiceGetProjectHandler.ServeHTTP(w, r)
		case ProjectsServiceListProjectsProcedure:
			projectsServiceListProjectsHandler.ServeHTTP(w, r)
		case ProjectsServiceSearchProjectsProcedure:
			projectsServiceSearchProjectsHandler.ServeHTTP(w, r)
		case ProjectsServiceCreateProjectProcedure:
			projectsServiceCreateProjectHandler.ServeHTTP(w, r)
		case ProjectsServiceDeleteProjectProcedure:
			projectsServiceDeleteProjectHandler.ServeHTTP(w, r)
		case ProjectsServicePatchProjectProcedure:
			projectsServicePatchProjectHandler.ServeHTTP(w, r)
		case ProjectsServiceUpdateProjectProcedure:
			projectsServiceUpdateProjectHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProjectsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectsServiceHandler struct{}

func (UnimplementedProjectsServiceHandler) GetProject(context.Context, *connect.Request[v1alpha1.GetProjectRequest]) (*connect.Response[v1alpha1.GetProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.GetProject is not implemented"))
}

func (UnimplementedProjectsServiceHandler) ListProjects(context.Context, *connect.Request[v1alpha1.ListProjectsRequest]) (*connect.Response[v1alpha1.ListProjectsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.ListProjects is not implemented"))
}

func (UnimplementedProjectsServiceHandler) SearchProjects(context.Context, *connect.Request[v1alpha1.SearchProjectsRequest]) (*connect.Response[v1alpha1.SearchProjectsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.SearchProjects is not implemented"))
}

func (UnimplementedProjectsServiceHandler) CreateProject(context.Context, *connect.Request[v1alpha1.CreateProjectRequest]) (*connect.Response[v1alpha1.CreateProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.CreateProject is not implemented"))
}

func (UnimplementedProjectsServiceHandler) DeleteProject(context.Context, *connect.Request[v1alpha1.DeleteProjectRequest]) (*connect.Response[v1alpha1.DeleteProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.DeleteProject is not implemented"))
}

func (UnimplementedProjectsServiceHandler) PatchProject(context.Context, *connect.Request[v1alpha1.PatchProjectRequest]) (*connect.Response[v1alpha1.PatchProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.PatchProject is not implemented"))
}

func (UnimplementedProjectsServiceHandler) UpdateProject(context.Context, *connect.Request[v1alpha1.UpdateProjectRequest]) (*connect.Response[v1alpha1.UpdateProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("priv.projects.v1alpha1.ProjectsService.UpdateProject is not implemented"))
}
