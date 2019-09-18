package manager_test

import (
	"testing"

	function "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	componentsMocks "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents/mocks"
	"github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/manager"
	"github.com/kyma-project/kyma/components/application-gateway/pkg/apperrors"
	subscriptions "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma-project.io/v1alpha1"
	serviceBindingUsages "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	bindings "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	serviceInstance "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/stretchr/testify/assert"
)

func TestCreateSubscription(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Subscription{}
		subscriptionBody := &subscriptions.Subscription{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, nil)
		component.On("Prepare", "githubRepo", "epo-lambda").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateSubscription(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Subscription{}
		subscriptionBody := &subscriptions.Subscription{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, apperrors.Internal("error"))
		component.On("Prepare", "githubRepo", "epo-lambda").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateSubscription(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateServiceBindingUsages(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.BindingUsage{}
		bindingUsageBody := &serviceBindingUsages.ServiceBindingUsage{}
		component.On("Create", bindingUsageBody).Return(bindingUsageBody, nil)
		component.On("Prepare", "githubRepo", "GITHUB_", "epo-lambda").Return(bindingUsageBody)
		component.On("Prepare", "slackWorkspace", "", "epo-lambda").Return(bindingUsageBody)
		component.On("Prepare", "azureServiceName", "", "epo-lambda").Return(bindingUsageBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindingUsages(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.BindingUsage{}
		bindingUsageBody := &serviceBindingUsages.ServiceBindingUsage{}
		component.On("Create", bindingUsageBody).Return(bindingUsageBody, apperrors.Internal("error"))
		component.On("Prepare", "githubRepo", "GITHUB_", "epo-lambda").Return(bindingUsageBody)
		component.On("Prepare", "slackWorkspace", "", "epo-lambda").Return(bindingUsageBody)
		component.On("Prepare", "azureServiceName", "", "epo-lambda").Return(bindingUsageBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindingUsages(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateServiceBindings(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Binding{}
		bindingBody := &bindings.ServiceBinding{}
		component.On("Create", bindingBody).Return(bindingBody, nil)
		component.On("Prepare", "githubRepo", "epo-lambda").Return(bindingBody)
		component.On("Prepare", "slackWorkspace", "epo-lambda").Return(bindingBody)
		component.On("Prepare", "azureServiceName", "epo-lambda").Return(bindingBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindings(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Binding{}
		bindingBody := &bindings.ServiceBinding{}
		component.On("Create", bindingBody).Return(bindingBody, apperrors.Internal("error"))
		component.On("Prepare", "githubRepo", "epo-lambda").Return(bindingBody)
		component.On("Prepare", "slackWorkspace", "epo-lambda").Return(bindingBody)
		component.On("Prepare", "azureServiceName", "epo-lambda").Return(bindingBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateServiceBindings(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateFunction(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.Function{}
		subscriptionBody := &function.Function{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, nil)
		component.On("Prepare", "githubRepo", "epo-lambda").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateFunction(component)

		//then
		assert.NoError(t, err)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.Function{}
		subscriptionBody := &function.Function{}
		component.On("Create", subscriptionBody).Return(subscriptionBody, apperrors.Internal("error"))
		component.On("Prepare", "githubRepo", "epo-lambda").Return(subscriptionBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		//when
		err := testedManager.CreateFunction(component)

		//then
		assert.Error(t, err)
	})
}

func TestCreateServiceInstances(t *testing.T) {
	t.Run("should return nil when everything is fine", func(t *testing.T) {
		//given
		component := &componentsMocks.ServiceInstance{}
		serviceInstanceBody := &serviceInstance.ServiceInstance{}
		raw := runtime.RawExtension{}
		unmarshalerr := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
		component.On("Create", serviceInstanceBody).Return(serviceInstanceBody, nil)
		component.On("Prepare", "azureServiceName", "azureServiceName", "standard-s0", &raw).Return(serviceInstanceBody)
		component.On("Prepare", "githubRepo", "githubRepo-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		component.On("Prepare", "slackWorkspace", "slackWorkspace-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		serviceInstanceList := serviceInstance.ServiceClassList{
			Items: []serviceInstance.ServiceClass{
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "githubRepo-12345",
						},
					},
				}, serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "slackWorkspace-12345",
						},
					},
				},
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "azureServiceName",
						},
					},
				}},
		}
		//when
		err := testedManager.CreateServiceInstances(component, &serviceInstanceList)

		//then
		assert.NoError(t, err)
		assert.NoError(t, unmarshalerr)
	})

	t.Run("should return error when Create method break up", func(t *testing.T) {
		//given
		component := &componentsMocks.ServiceInstance{}
		serviceInstanceBody := &serviceInstance.ServiceInstance{}
		raw := runtime.RawExtension{}
		unmarshalerr := raw.UnmarshalJSON([]byte(`{"location": "westeurope","resourceGroup": "flying-seals-tmp"}`))
		component.On("Create", serviceInstanceBody).Return(serviceInstanceBody, apperrors.Internal("error"))
		component.On("Prepare", "azureServiceName", "azureServiceName", "standard-s0", &raw).Return(serviceInstanceBody)
		component.On("Prepare", "githubRepo", "githubRepo-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		component.On("Prepare", "slackWorkspace", "slackWorkspace-12345", "default", (*runtime.RawExtension)(nil)).Return(serviceInstanceBody)
		testedManager := manager.NewManager("namespace", "githubRepo", "slackWorkspace", "azureServiceName")
		serviceInstanceList := serviceInstance.ServiceClassList{
			Items: []serviceInstance.ServiceClass{
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "githubRepo-12345",
						},
					},
				}, serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "slackWorkspace-12345",
						},
					},
				},
				serviceInstance.ServiceClass{
					Spec: serviceInstance.ServiceClassSpec{
						CommonServiceClassSpec: serviceInstance.CommonServiceClassSpec{
							ExternalName: "azureServiceName",
						},
					},
				}},
		}
		//when
		err := testedManager.CreateServiceInstances(component, &serviceInstanceList)

		//then
		assert.Error(t, err)
		assert.NoError(t, unmarshalerr)
	})
}