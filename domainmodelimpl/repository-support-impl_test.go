package domainmodelimpl

import (
    "testing"
    "github.com/yfuksenk/gocmis/domainmodel"
)

func TestBaseObjectTypes(t *testing.T) {
	var _ domainmodel.BaseObjectTypeIder = (*CMISDocument)(nil)
	var _ domainmodel.BaseObjectTypeIder = (*CMISFolder)(nil)
	var _ domainmodel.BaseObjectTypeIder = (*CMISPolicy)(nil)
	var _ domainmodel.BaseObjectTypeIder = (*CMISRelationship)(nil)
	var _ domainmodel.BaseObjectTypeIder = (*CMISItem)(nil)
}

func TestSupportedPermissions(t *testing.T) {
	var _ domainmodel.SupportedPermissioner = (*BasicPermissions)(nil)
	var _ domainmodel.SupportedPermissioner = (*RepositoryPermissions)(nil)
	var _ domainmodel.SupportedPermissioner = (*BothPermissions)(nil)
}

func TestPermissionropagations(t *testing.T) {
	var _ domainmodel.Permissioner = (*ObjectOnly)(nil)
	var _ domainmodel.Permissioner = (*Propagate)(nil)
	var _ domainmodel.Permissioner = (*RepositoryDetermined)(nil)
}
