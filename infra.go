//nolint:revive
package nsxt

// The following file has been autogenerated. Please avoid any changes!
import (
	"errors"

	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	client0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt"
	model0 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	client1 "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/orgs/projects"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)


type InfraClientContext utl.ClientContext

func NewInfraClient(sessionContext utl.SessionContext, connector vapiProtocolClient_.Connector) *InfraClientContext {
	var client interface{}

	switch sessionContext.ClientType {

	case utl.Local:
		client = client0.NewInfraClient(connector)

	case utl.Multitenancy:
		client = client1.NewInfraClient(connector)

	default:
		return nil
	}
	return &InfraClientContext{Client: client, ClientType: sessionContext.ClientType, ProjectID: sessionContext.ProjectID, VPCID: sessionContext.VPCID}
}

func (c InfraClientContext) Get(basePathParam *string, filterParam *string, typeFilterParam *string) (model0.Infra, error) {
	var obj model0.Infra
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.InfraClient)
		obj, err = client.Get(basePathParam, filterParam, typeFilterParam)
		if err != nil {
			return obj, err
		}

	case utl.Multitenancy:
		client := c.Client.(client1.InfraClient)
		obj, err = client.Get(utl.DefaultOrgID, c.ProjectID, basePathParam, filterParam, typeFilterParam)
		if err != nil {
			return obj, err
		}

	default:
		return obj, errors.New("invalid infrastructure for model")
	}
	return obj, err
}

func (c InfraClientContext) Patch(infraParam model0.Infra, enforceRevisionCheckParam *bool) error {
	var err error

	switch c.ClientType {

	case utl.Local:
		client := c.Client.(client0.InfraClient)
		err = client.Patch(infraParam, enforceRevisionCheckParam)

	case utl.Multitenancy:
		client := c.Client.(client1.InfraClient)
		err = client.Patch(utl.DefaultOrgID, c.ProjectID, infraParam, enforceRevisionCheckParam)

	default:
		err = errors.New("invalid infrastructure for model")
	}
	return err
}