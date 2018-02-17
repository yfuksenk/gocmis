package domainmodel

//
//Author: yfuksenko
//

type RepositoryInfo interface{
	RepositoryId() string
	RepositoryName() string
	RepositoryDescription() string
	VendorName() string
	ProductName() string
	ProductVersion() string
	RootFolderId() string
	Capabilities() []Capability
	LatestChangeLogToken() string
	CmisVersionSupported() string
	ThinClientURI() string
	ChangesIncomplete() bool
	ChangesOnType() []BaseObjectTypeIder
	SupportedPermissions() SupportedPermissioner
	Propagation() Propagationer
	Permission() []Permissioner
	Mapping() []PermissionMapping
	PrincipalAnonymous() string
	PrincipalAnyone() string
	ExtendedFeatures() []RepositoryFeature
} 

type Capability interface{
	
}

type BaseObjectTypeIder interface{
	BaseObjectTypeId() string
}

type SupportedPermissioner interface{
	 SupportedPermission() string
}

type Propagationer interface{
	Propagation() string
}

type Permissioner interface{
	Permission() string
}

type PermissionMapping interface{
	Key() string
	Permissions() []string 
}

type RepositoryFeature interface{
	Id() string
	Name() string
	Description() string
	Version() string
} 
