package inframodel

type IisSiteApplicationMemberPool struct {
	ParentSite        IisSite
	ChildApplications []IisApplication
	// TODO: Implement InitPools
}
