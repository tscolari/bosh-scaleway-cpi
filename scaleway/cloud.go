package scaleway

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/scaleway/scaleway-cli/pkg/api"
	"github.com/tscolari/bosh-c3pi/cloud"
)

func New(apiClient *api.ScalewayAPI) *Client {
	return &Client{
		api: apiClient,
	}
}

type Client struct {
	api *api.ScalewayAPI
}

func (c *Client) CreateStemcell(imagePath string, cloudProperties cloud.CloudProperties) (string, error) {
	return "", nil
}
func (c *Client) DeleteStemcell(stemcellID string) error {
	return nil
}

func (c *Client) CurrentVmID() string {
	return ""
}
func (c *Client) CreateVm(agentID, stemcellID string, cloudProperties cloud.CloudProperties, networks cloud.Networks, diskLocality string, env cloud.Environment) (string, error) {
	definition := api.ScalewayServerDefinition{
		Name:           "bosh-" + uuid.New(),
		CommercialType: cloudProperties["instance_type"],
		Tags:           []string{"bosh"},
	}

	c.api.PostServer(definition)
	return "", nil
}
func (c *Client) DeleteVm(vmID string) error {
	return c.api.DeleteServer(vmID)
}
func (c *Client) HasVm(vmID string) (bool, error) {
	return true, nil
}
func (c *Client) RebootVm(vmID string) error {
	return c.api.PostServerAction(vmID, "reboot")
}
func (c *Client) SetVmMetadata(vm string, metadata cloud.Metadata) error {
	return nil
}

func (c *Client) CreateDisk(size int, cloudProperties cloud.CloudProperties, vmLocality string) (string, error) {
	return "", nil
}
func (c *Client) GetDisks(vmID string) ([]string, error) {
	return []string{}, nil
}
func (c *Client) HasDisk(diskID string) (bool, error) {
	return true, nil
}
func (c *Client) DeleteDisk(diskID string) error {
	return nil
}
func (c *Client) AttachDisk(vmID, diskID string) error {
	return nil
}
func (c *Client) DetachDisk(vmID, diskID string) error {
	return nil
}

func (c *Client) SnapshotDisk(diskID string, metadata cloud.Metadata) (string, error) {
	return "", nil
}
func (c *Client) DeleteSnapshot(snapshotID string) error {
	return nil
}
