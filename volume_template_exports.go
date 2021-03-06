// Code generated by gen_exports.go DO NOT EDIT

package metalcloud

//VolumeTemplateGet returns the specified volume template
func (c *Client) VolumeTemplateGet(volumeTemplateID int) (*VolumeTemplate, error) {
	return c.volumeTemplateGet(volumeTemplateID)
}

//VolumeTemplateGetByLabel returns the specified volume template
func (c *Client) VolumeTemplateGetByLabel(volumeTemplateLabel string) (*VolumeTemplate, error) {
	return c.volumeTemplateGet(volumeTemplateLabel)
}

//VolumeTemplateCreateFromDrive creates a private volume template from a drive
func (c *Client) VolumeTemplateCreateFromDrive(driveID int, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error) {
	return c.volumeTemplateCreateFromDrive(driveID,objVolumeTemplate)
}

//VolumeTemplateCreateFromDriveByLabel creates a private volume template from a drive
func (c *Client) VolumeTemplateCreateFromDriveByLabel(driveLabel string, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error) {
	return c.volumeTemplateCreateFromDrive(driveLabel,objVolumeTemplate)
}
