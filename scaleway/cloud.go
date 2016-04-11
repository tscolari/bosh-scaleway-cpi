package scaleway

import (
	"fmt"

	"github.com/satori/go.uuid"
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
		Name:           fmt.Sprintf("bosh-%s", uuid.NewV4()),
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
	server, err := c.api.GetServer(vmID)
	if err != nil {
		return false, err
	}

	if server != nil {
		return true, nil
	}

	return false, nil
}

func (c *Client) RebootVm(vmID string) error {
	return c.api.PostServerAction(vmID, "reboot")
}

func (c *Client) SetVmMetadata(vm string, metadata cloud.Metadata) error {
	return nil
}

func (c *Client) CreateDisk(size int, cloudProperties cloud.CloudProperties, vmLocality string) (string, error) {
	volumeDefinition := api.ScalewayVolumeDefinition{
		Name:         fmt.Sprintf("bosh-%s", uuid.NewV4()),
		Size:         uint64(size),
		Type:         "l_ssd",
		Organization: c.api.Organization,
	}

	return c.api.PostVolume(volumeDefinition)
}

func (c *Client) GetDisks(vmID string) ([]string, error) {
	volumes, err := c.api.GetVolumes()
	if err != nil {
		return nil, err
	}

	volumeIDs := []string{}
	for _, volume := range *volumes {
		volumeIDs = append(volumeIDs, volume.Identifier)
	}

	return volumeIDs, nil
}

func (c *Client) HasDisk(diskID string) (bool, error) {
	volume, err := c.api.GetVolume(diskID)
	if err != nil {
		return false, err
	}

	if volume != nil {
		return true, nil
	}

	return false, nil
}

func (c *Client) DeleteDisk(diskID string) error {
	return c.api.DeleteVolume(diskID)
}

func (c *Client) AttachDisk(vmID, diskID string) error {
	return nil
}

func (c *Client) DetachDisk(vmID, diskID string) error {
	return nil
}

func (c *Client) SnapshotDisk(diskID string, metadata cloud.Metadata) (string, error) {
	return c.api.PostSnapshot(diskID, fmt.Sprintf("bosh-%s", uuid.NewV4()))
}

func (c *Client) DeleteSnapshot(snapshotID string) error {
	return c.api.DeleteSnapshot(snapshotID)
}
