/*
 * Model Registry REST API
 *
 * REST API for Model Registry to create and manage ML model metadata
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/opendatahub-io/model-registry/internal/apiutils"
	"github.com/opendatahub-io/model-registry/internal/converter"
	"github.com/opendatahub-io/model-registry/internal/converter/generated"
	"github.com/opendatahub-io/model-registry/pkg/api"
	model "github.com/opendatahub-io/model-registry/pkg/openapi"
)

// ModelRegistryServiceAPIService is a service that implements the logic for the ModelRegistryServiceAPIServicer
// This service should implement the business logic for every endpoint for the ModelRegistryServiceAPI s.coreApi.
// Include any external packages or services that will be required by this service.
type ModelRegistryServiceAPIService struct {
	coreApi   api.ModelRegistryApi
	converter converter.OpenAPIConverter
}

// NewModelRegistryServiceAPIService creates a default api service
func NewModelRegistryServiceAPIService(coreApi api.ModelRegistryApi) ModelRegistryServiceAPIServicer {
	return &ModelRegistryServiceAPIService{
		coreApi:   coreApi,
		converter: &generated.OpenAPIConverterImpl{},
	}
}

// CreateEnvironmentInferenceService - Create a InferenceService in ServingEnvironment
func (s *ModelRegistryServiceAPIService) CreateEnvironmentInferenceService(ctx context.Context, servingenvironmentId string, inferenceServiceCreate model.InferenceServiceCreate) (ImplResponse, error) {
	inferenceServiceCreate.ServingEnvironmentId = servingenvironmentId
	return s.CreateInferenceService(ctx, inferenceServiceCreate)
	// TODO: return Response(404, Error{}), nil
}

// CreateInferenceService - Create a InferenceService
func (s *ModelRegistryServiceAPIService) CreateInferenceService(ctx context.Context, inferenceServiceCreate model.InferenceServiceCreate) (ImplResponse, error) {
	entity, err := s.converter.ConvertInferenceServiceCreate(&inferenceServiceCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertInferenceService(entity)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
}

// CreateInferenceServiceServe - Create a ServeModel action in a InferenceService
func (s *ModelRegistryServiceAPIService) CreateInferenceServiceServe(ctx context.Context, inferenceserviceId string, serveModelCreate model.ServeModelCreate) (ImplResponse, error) {
	entity, err := s.converter.ConvertServeModelCreate(&serveModelCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertServeModel(entity, &inferenceserviceId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// CreateModelArtifact - Create a ModelArtifact
func (s *ModelRegistryServiceAPIService) CreateModelArtifact(ctx context.Context, modelArtifactCreate model.ModelArtifactCreate) (ImplResponse, error) {
	entity, err := s.converter.ConvertModelArtifactCreate(&modelArtifactCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertModelArtifact(entity, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
}

// CreateModelVersion - Create a ModelVersion
func (s *ModelRegistryServiceAPIService) CreateModelVersion(ctx context.Context, modelVersionCreate model.ModelVersionCreate) (ImplResponse, error) {
	modelVersion, err := s.converter.ConvertModelVersionCreate(&modelVersionCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertModelVersion(modelVersion, &modelVersionCreate.RegisteredModelID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
}

// CreateModelVersionArtifact - Create an Artifact in a ModelVersion
func (s *ModelRegistryServiceAPIService) CreateModelVersionArtifact(ctx context.Context, modelversionId string, artifact model.Artifact) (ImplResponse, error) {
	if artifact.ModelArtifact == nil {
		return Response(http.StatusNotImplemented, nil), errors.New("unsupported artifactType")
	}
	result, err := s.coreApi.UpsertModelArtifact(artifact.ModelArtifact, &modelversionId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	artifactResult := model.Artifact{
		ModelArtifact: result,
	}
	return Response(201, artifactResult), nil
	// TODO return Response(200, Artifact{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// CreateRegisteredModel - Create a RegisteredModel
func (s *ModelRegistryServiceAPIService) CreateRegisteredModel(ctx context.Context, registeredModelCreate model.RegisteredModelCreate) (ImplResponse, error) {
	registeredModel, err := s.converter.ConvertRegisteredModelCreate(&registeredModelCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertRegisteredModel(registeredModel)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
}

// CreateRegisteredModelVersion - Create a ModelVersion in RegisteredModel
func (s *ModelRegistryServiceAPIService) CreateRegisteredModelVersion(ctx context.Context, registeredmodelId string, modelVersion model.ModelVersion) (ImplResponse, error) {
	result, err := s.coreApi.UpsertModelVersion(&modelVersion, &registeredmodelId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// CreateServingEnvironment - Create a ServingEnvironment
func (s *ModelRegistryServiceAPIService) CreateServingEnvironment(ctx context.Context, servingEnvironmentCreate model.ServingEnvironmentCreate) (ImplResponse, error) {
	entity, err := s.converter.ConvertServingEnvironmentCreate(&servingEnvironmentCreate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}

	result, err := s.coreApi.UpsertServingEnvironment(entity)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(201, result), nil
	// TODO: return Response(400, Error{}), nil
	// TODO: return Response(401, Error{}), nil
}

// FindInferenceService - Get an InferenceServices that matches search parameters.
func (s *ModelRegistryServiceAPIService) FindInferenceService(ctx context.Context, name string, externalID string) (ImplResponse, error) {
	result, err := s.coreApi.GetInferenceServiceByParams(&name, nil, &externalID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return esponse(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// FindModelArtifact - Get a ModelArtifact that matches search parameters.
func (s *ModelRegistryServiceAPIService) FindModelArtifact(ctx context.Context, name string, externalID string, parentResourceID string) (ImplResponse, error) {
	result, err := s.coreApi.GetModelArtifactByParams(&name, &externalID, &parentResourceID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return esponse(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// FindModelVersion - Get a ModelVersion that matches search parameters.
func (s *ModelRegistryServiceAPIService) FindModelVersion(ctx context.Context, name string, externalID string, registeredModelID string) (ImplResponse, error) {
	result, err := s.coreApi.GetModelVersionByParams(&name, &externalID, &registeredModelID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return esponse(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// FindRegisteredModel - Get a RegisteredModel that matches search parameters.
func (s *ModelRegistryServiceAPIService) FindRegisteredModel(ctx context.Context, name string, externalID string) (ImplResponse, error) {
	result, err := s.coreApi.GetRegisteredModelByParams(&name, &externalID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return esponse(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// FindServingEnvironment - Find ServingEnvironment
func (s *ModelRegistryServiceAPIService) FindServingEnvironment(ctx context.Context, name string, externalID string) (ImplResponse, error) {
	result, err := s.coreApi.GetServingEnvironmentByParams(&name, &externalID)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetEnvironmentInferenceServices - List All ServingEnvironment&#39;s InferenceServices
func (s *ModelRegistryServiceAPIService) GetEnvironmentInferenceServices(ctx context.Context, servingenvironmentId string, name string, externalID string, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetInferenceServices(listOpts, &servingenvironmentId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetInferenceService - Get a InferenceService
func (s *ModelRegistryServiceAPIService) GetInferenceService(ctx context.Context, inferenceserviceId string) (ImplResponse, error) {
	result, err := s.coreApi.GetInferenceServiceById(inferenceserviceId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetInferenceServiceModel - Get InferenceService&#39;s RegisteredModel
func (s *ModelRegistryServiceAPIService) GetInferenceServiceModel(ctx context.Context, inferenceserviceId string) (ImplResponse, error) {
	result, err := s.coreApi.GetRegisteredModelByInferenceService(inferenceserviceId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetInferenceServiceServes - List All InferenceService&#39;s ServeModel actions
func (s *ModelRegistryServiceAPIService) GetInferenceServiceServes(ctx context.Context, inferenceserviceId string, name string, externalID string, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetServeModels(listOpts, &inferenceserviceId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetInferenceServiceVersion - Get InferenceService&#39;s ModelVersion
func (s *ModelRegistryServiceAPIService) GetInferenceServiceVersion(ctx context.Context, inferenceserviceId string) (ImplResponse, error) {
	result, err := s.coreApi.GetModelVersionByInferenceService(inferenceserviceId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetInferenceServices - List All InferenceServices
func (s *ModelRegistryServiceAPIService) GetInferenceServices(ctx context.Context, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetInferenceServices(listOpts, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetModelArtifact - Get a ModelArtifact
func (s *ModelRegistryServiceAPIService) GetModelArtifact(ctx context.Context, modelartifactId string) (ImplResponse, error) {
	result, err := s.coreApi.GetModelArtifactById(modelartifactId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetModelArtifacts - List All ModelArtifacts
func (s *ModelRegistryServiceAPIService) GetModelArtifacts(ctx context.Context, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetModelArtifacts(listOpts, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetModelVersion - Get a ModelVersion
func (s *ModelRegistryServiceAPIService) GetModelVersion(ctx context.Context, modelversionId string) (ImplResponse, error) {
	result, err := s.coreApi.GetModelVersionById(modelversionId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetModelVersionArtifacts - List All ModelVersion&#39;s artifacts
func (s *ModelRegistryServiceAPIService) GetModelVersionArtifacts(ctx context.Context, modelversionId string, name string, externalID string, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	// TODO name unused
	// TODO externalID unused
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetModelArtifacts(listOpts, &modelversionId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetModelVersions - List All ModelVersions
func (s *ModelRegistryServiceAPIService) GetModelVersions(ctx context.Context, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetModelVersions(listOpts, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetRegisteredModel - Get a RegisteredModel
func (s *ModelRegistryServiceAPIService) GetRegisteredModel(ctx context.Context, registeredmodelId string) (ImplResponse, error) {
	result, err := s.coreApi.GetRegisteredModelById(registeredmodelId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetRegisteredModelVersions - List All RegisteredModel&#39;s ModelVersions
func (s *ModelRegistryServiceAPIService) GetRegisteredModelVersions(ctx context.Context, registeredmodelId string, name string, externalID string, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	// TODO name unused
	// TODO externalID unused
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetModelVersions(listOpts, &registeredmodelId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetRegisteredModels - List All RegisteredModels
func (s *ModelRegistryServiceAPIService) GetRegisteredModels(ctx context.Context, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetRegisteredModels(listOpts)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// GetServingEnvironment - Get a ServingEnvironment
func (s *ModelRegistryServiceAPIService) GetServingEnvironment(ctx context.Context, servingenvironmentId string) (ImplResponse, error) {
	result, err := s.coreApi.GetServingEnvironmentById(servingenvironmentId)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO: return Response(401, Error{}), nil
	// TODO: return Response(404, Error{}), nil
}

// GetServingEnvironments - List All ServingEnvironments
func (s *ModelRegistryServiceAPIService) GetServingEnvironments(ctx context.Context, pageSize string, orderBy model.OrderByField, sortOrder model.SortOrder, nextPageToken string) (ImplResponse, error) {
	listOpts, err := apiutils.BuildListOption(pageSize, orderBy, sortOrder, nextPageToken)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	result, err := s.coreApi.GetServingEnvironments(listOpts)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(401, Error{}), nil
}

// UpdateInferenceService - Update a InferenceService
func (s *ModelRegistryServiceAPIService) UpdateInferenceService(ctx context.Context, inferenceserviceId string, inferenceServiceUpdate model.InferenceServiceUpdate) (ImplResponse, error) {
	entity, err := s.converter.ConvertInferenceServiceUpdate(&inferenceServiceUpdate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	entity.Id = &inferenceserviceId
	result, err := s.coreApi.UpsertInferenceService(entity)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// UpdateModelArtifact - Update a ModelArtifact
func (s *ModelRegistryServiceAPIService) UpdateModelArtifact(ctx context.Context, modelartifactId string, modelArtifactUpdate model.ModelArtifactUpdate) (ImplResponse, error) {
	modelArtifact, err := s.converter.ConvertModelArtifactUpdate(&modelArtifactUpdate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	modelArtifact.Id = &modelartifactId
	result, err := s.coreApi.UpsertModelArtifact(modelArtifact, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// UpdateModelVersion - Update a ModelVersion
func (s *ModelRegistryServiceAPIService) UpdateModelVersion(ctx context.Context, modelversionId string, modelVersion model.ModelVersion) (ImplResponse, error) {
	// TODO: this API is getting model.ModelVersion instead of model.ModelVersionUpdate.
	// c, err := s.converter.ConvertModelVersionUpdate(&modelVersion)
	// if err != nil {
	// 	return Response(500, model.Error{Message: err.Error()}), nil
	// }
	// modelVersion.Id = &modelversionId
	result, err := s.coreApi.UpsertModelVersion(&modelVersion, nil)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// UpdateRegisteredModel - Update a RegisteredModel
func (s *ModelRegistryServiceAPIService) UpdateRegisteredModel(ctx context.Context, registeredmodelId string, registeredModelUpdate model.RegisteredModelUpdate) (ImplResponse, error) {
	registeredModel, err := s.converter.ConvertRegisteredModelUpdate(&registeredModelUpdate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	registeredModel.Id = &registeredmodelId
	result, err := s.coreApi.UpsertRegisteredModel(registeredModel)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}

// UpdateServingEnvironment - Update a ServingEnvironment
func (s *ModelRegistryServiceAPIService) UpdateServingEnvironment(ctx context.Context, servingenvironmentId string, servingEnvironmentUpdate model.ServingEnvironmentUpdate) (ImplResponse, error) {
	entity, err := s.converter.ConvertServingEnvironmentUpdate(&servingEnvironmentUpdate)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	entity.Id = &servingenvironmentId
	result, err := s.coreApi.UpsertServingEnvironment(entity)
	if err != nil {
		return Response(500, model.Error{Message: err.Error()}), nil
	}
	return Response(200, result), nil
	// TODO return Response(400, Error{}), nil
	// TODO return Response(401, Error{}), nil
	// TODO return Response(404, Error{}), nil
}
