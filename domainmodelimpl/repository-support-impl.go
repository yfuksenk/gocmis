package domainmodelimpl

//
//Author: yfuksenko
//

type CMISDocument struct{}
func (CMISDocument) BaseObjectTypeId() string{
	return "cmis:document"
}

type CMISFolder struct{}
func (CMISFolder) BaseObjectTypeId() string{
	return "cmis:folder"
}

type CMISPolicy struct{}
func (CMISPolicy) BaseObjectTypeId() string{
	return "cmis:policy"
}

type CMISRelationship struct{}
func (CMISRelationship) BaseObjectTypeId() string{
	return "cmis:relationship"
}

type CMISItem struct{}
func (CMISItem) BaseObjectTypeId() string{
	return "cmis:item"
}

type BasicPermissions struct{}
func (BasicPermissions) SupportedPermission() string{
	return "basic"
}

type RepositoryPermissions struct{}
func (RepositoryPermissions) SupportedPermission() string{
	return "repository"
}

type BothPermissions struct{}
func (BothPermissions) SupportedPermission() string{
	return "both"
}

type ObjectOnly struct{}
func (ObjectOnly) Permission() string{
	return "objectonly"
}

type Propagate struct{}
func (Propagate) Permission() string{
	return "propagate"
}

type RepositoryDetermined struct{}
func (RepositoryDetermined) Permission() string{
	return "repositorydetermined"
}