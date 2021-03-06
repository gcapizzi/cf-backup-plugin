package fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
)

type FakeCliConnection struct {
	CliCommandWithoutTerminalOutputStub        func(args ...string) ([]string, error)
	cliCommandWithoutTerminalOutputMutex       sync.RWMutex
	cliCommandWithoutTerminalOutputArgsForCall []struct {
		args []string
	}
	cliCommandWithoutTerminalOutputReturns struct {
		result1 []string
		result2 error
	}
	cliCommandWithoutTerminalOutputReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	CliCommandStub        func(args ...string) ([]string, error)
	cliCommandMutex       sync.RWMutex
	cliCommandArgsForCall []struct {
		args []string
	}
	cliCommandReturns struct {
		result1 []string
		result2 error
	}
	cliCommandReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetCurrentOrgStub        func() (plugin_models.Organization, error)
	getCurrentOrgMutex       sync.RWMutex
	getCurrentOrgArgsForCall []struct{}
	getCurrentOrgReturns     struct {
		result1 plugin_models.Organization
		result2 error
	}
	getCurrentOrgReturnsOnCall map[int]struct {
		result1 plugin_models.Organization
		result2 error
	}
	GetCurrentSpaceStub        func() (plugin_models.Space, error)
	getCurrentSpaceMutex       sync.RWMutex
	getCurrentSpaceArgsForCall []struct{}
	getCurrentSpaceReturns     struct {
		result1 plugin_models.Space
		result2 error
	}
	getCurrentSpaceReturnsOnCall map[int]struct {
		result1 plugin_models.Space
		result2 error
	}
	UsernameStub        func() (string, error)
	usernameMutex       sync.RWMutex
	usernameArgsForCall []struct{}
	usernameReturns     struct {
		result1 string
		result2 error
	}
	usernameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UserGuidStub        func() (string, error)
	userGuidMutex       sync.RWMutex
	userGuidArgsForCall []struct{}
	userGuidReturns     struct {
		result1 string
		result2 error
	}
	userGuidReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UserEmailStub        func() (string, error)
	userEmailMutex       sync.RWMutex
	userEmailArgsForCall []struct{}
	userEmailReturns     struct {
		result1 string
		result2 error
	}
	userEmailReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	IsLoggedInStub        func() (bool, error)
	isLoggedInMutex       sync.RWMutex
	isLoggedInArgsForCall []struct{}
	isLoggedInReturns     struct {
		result1 bool
		result2 error
	}
	isLoggedInReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	IsSSLDisabledStub        func() (bool, error)
	isSSLDisabledMutex       sync.RWMutex
	isSSLDisabledArgsForCall []struct{}
	isSSLDisabledReturns     struct {
		result1 bool
		result2 error
	}
	isSSLDisabledReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	HasOrganizationStub        func() (bool, error)
	hasOrganizationMutex       sync.RWMutex
	hasOrganizationArgsForCall []struct{}
	hasOrganizationReturns     struct {
		result1 bool
		result2 error
	}
	hasOrganizationReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	HasSpaceStub        func() (bool, error)
	hasSpaceMutex       sync.RWMutex
	hasSpaceArgsForCall []struct{}
	hasSpaceReturns     struct {
		result1 bool
		result2 error
	}
	hasSpaceReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ApiEndpointStub        func() (string, error)
	apiEndpointMutex       sync.RWMutex
	apiEndpointArgsForCall []struct{}
	apiEndpointReturns     struct {
		result1 string
		result2 error
	}
	apiEndpointReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ApiVersionStub        func() (string, error)
	apiVersionMutex       sync.RWMutex
	apiVersionArgsForCall []struct{}
	apiVersionReturns     struct {
		result1 string
		result2 error
	}
	apiVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HasAPIEndpointStub        func() (bool, error)
	hasAPIEndpointMutex       sync.RWMutex
	hasAPIEndpointArgsForCall []struct{}
	hasAPIEndpointReturns     struct {
		result1 bool
		result2 error
	}
	hasAPIEndpointReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	LoggregatorEndpointStub        func() (string, error)
	loggregatorEndpointMutex       sync.RWMutex
	loggregatorEndpointArgsForCall []struct{}
	loggregatorEndpointReturns     struct {
		result1 string
		result2 error
	}
	loggregatorEndpointReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DopplerEndpointStub        func() (string, error)
	dopplerEndpointMutex       sync.RWMutex
	dopplerEndpointArgsForCall []struct{}
	dopplerEndpointReturns     struct {
		result1 string
		result2 error
	}
	dopplerEndpointReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	AccessTokenStub        func() (string, error)
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
		result2 error
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetAppStub        func(string) (plugin_models.GetAppModel, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		arg1 string
	}
	getAppReturns struct {
		result1 plugin_models.GetAppModel
		result2 error
	}
	getAppReturnsOnCall map[int]struct {
		result1 plugin_models.GetAppModel
		result2 error
	}
	GetAppsStub        func() ([]plugin_models.GetAppsModel, error)
	getAppsMutex       sync.RWMutex
	getAppsArgsForCall []struct{}
	getAppsReturns     struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}
	getAppsReturnsOnCall map[int]struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}
	GetOrgsStub        func() ([]plugin_models.GetOrgs_Model, error)
	getOrgsMutex       sync.RWMutex
	getOrgsArgsForCall []struct{}
	getOrgsReturns     struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}
	getOrgsReturnsOnCall map[int]struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}
	GetSpacesStub        func() ([]plugin_models.GetSpaces_Model, error)
	getSpacesMutex       sync.RWMutex
	getSpacesArgsForCall []struct{}
	getSpacesReturns     struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}
	getSpacesReturnsOnCall map[int]struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}
	GetOrgUsersStub        func(string, ...string) ([]plugin_models.GetOrgUsers_Model, error)
	getOrgUsersMutex       sync.RWMutex
	getOrgUsersArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getOrgUsersReturns struct {
		result1 []plugin_models.GetOrgUsers_Model
		result2 error
	}
	getOrgUsersReturnsOnCall map[int]struct {
		result1 []plugin_models.GetOrgUsers_Model
		result2 error
	}
	GetSpaceUsersStub        func(string, string) ([]plugin_models.GetSpaceUsers_Model, error)
	getSpaceUsersMutex       sync.RWMutex
	getSpaceUsersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceUsersReturns struct {
		result1 []plugin_models.GetSpaceUsers_Model
		result2 error
	}
	getSpaceUsersReturnsOnCall map[int]struct {
		result1 []plugin_models.GetSpaceUsers_Model
		result2 error
	}
	GetServicesStub        func() ([]plugin_models.GetServices_Model, error)
	getServicesMutex       sync.RWMutex
	getServicesArgsForCall []struct{}
	getServicesReturns     struct {
		result1 []plugin_models.GetServices_Model
		result2 error
	}
	getServicesReturnsOnCall map[int]struct {
		result1 []plugin_models.GetServices_Model
		result2 error
	}
	GetServiceStub        func(string) (plugin_models.GetService_Model, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		arg1 string
	}
	getServiceReturns struct {
		result1 plugin_models.GetService_Model
		result2 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 plugin_models.GetService_Model
		result2 error
	}
	GetOrgStub        func(string) (plugin_models.GetOrg_Model, error)
	getOrgMutex       sync.RWMutex
	getOrgArgsForCall []struct {
		arg1 string
	}
	getOrgReturns struct {
		result1 plugin_models.GetOrg_Model
		result2 error
	}
	getOrgReturnsOnCall map[int]struct {
		result1 plugin_models.GetOrg_Model
		result2 error
	}
	GetSpaceStub        func(string) (plugin_models.GetSpace_Model, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		arg1 string
	}
	getSpaceReturns struct {
		result1 plugin_models.GetSpace_Model
		result2 error
	}
	getSpaceReturnsOnCall map[int]struct {
		result1 plugin_models.GetSpace_Model
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutput(args ...string) ([]string, error) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	ret, specificReturn := fake.cliCommandWithoutTerminalOutputReturnsOnCall[len(fake.cliCommandWithoutTerminalOutputArgsForCall)]
	fake.cliCommandWithoutTerminalOutputArgsForCall = append(fake.cliCommandWithoutTerminalOutputArgsForCall, struct {
		args []string
	}{args})
	fake.recordInvocation("CliCommandWithoutTerminalOutput", []interface{}{args})
	fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	if fake.CliCommandWithoutTerminalOutputStub != nil {
		return fake.CliCommandWithoutTerminalOutputStub(args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.cliCommandWithoutTerminalOutputReturns.result1, fake.cliCommandWithoutTerminalOutputReturns.result2
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputCallCount() int {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	return len(fake.cliCommandWithoutTerminalOutputArgsForCall)
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputArgsForCall(i int) []string {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	return fake.cliCommandWithoutTerminalOutputArgsForCall[i].args
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputReturns(result1 []string, result2 error) {
	fake.CliCommandWithoutTerminalOutputStub = nil
	fake.cliCommandWithoutTerminalOutputReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputReturnsOnCall(i int, result1 []string, result2 error) {
	fake.CliCommandWithoutTerminalOutputStub = nil
	if fake.cliCommandWithoutTerminalOutputReturnsOnCall == nil {
		fake.cliCommandWithoutTerminalOutputReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.cliCommandWithoutTerminalOutputReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) CliCommand(args ...string) ([]string, error) {
	fake.cliCommandMutex.Lock()
	ret, specificReturn := fake.cliCommandReturnsOnCall[len(fake.cliCommandArgsForCall)]
	fake.cliCommandArgsForCall = append(fake.cliCommandArgsForCall, struct {
		args []string
	}{args})
	fake.recordInvocation("CliCommand", []interface{}{args})
	fake.cliCommandMutex.Unlock()
	if fake.CliCommandStub != nil {
		return fake.CliCommandStub(args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.cliCommandReturns.result1, fake.cliCommandReturns.result2
}

func (fake *FakeCliConnection) CliCommandCallCount() int {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return len(fake.cliCommandArgsForCall)
}

func (fake *FakeCliConnection) CliCommandArgsForCall(i int) []string {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return fake.cliCommandArgsForCall[i].args
}

func (fake *FakeCliConnection) CliCommandReturns(result1 []string, result2 error) {
	fake.CliCommandStub = nil
	fake.cliCommandReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) CliCommandReturnsOnCall(i int, result1 []string, result2 error) {
	fake.CliCommandStub = nil
	if fake.cliCommandReturnsOnCall == nil {
		fake.cliCommandReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.cliCommandReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentOrg() (plugin_models.Organization, error) {
	fake.getCurrentOrgMutex.Lock()
	ret, specificReturn := fake.getCurrentOrgReturnsOnCall[len(fake.getCurrentOrgArgsForCall)]
	fake.getCurrentOrgArgsForCall = append(fake.getCurrentOrgArgsForCall, struct{}{})
	fake.recordInvocation("GetCurrentOrg", []interface{}{})
	fake.getCurrentOrgMutex.Unlock()
	if fake.GetCurrentOrgStub != nil {
		return fake.GetCurrentOrgStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCurrentOrgReturns.result1, fake.getCurrentOrgReturns.result2
}

func (fake *FakeCliConnection) GetCurrentOrgCallCount() int {
	fake.getCurrentOrgMutex.RLock()
	defer fake.getCurrentOrgMutex.RUnlock()
	return len(fake.getCurrentOrgArgsForCall)
}

func (fake *FakeCliConnection) GetCurrentOrgReturns(result1 plugin_models.Organization, result2 error) {
	fake.GetCurrentOrgStub = nil
	fake.getCurrentOrgReturns = struct {
		result1 plugin_models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentOrgReturnsOnCall(i int, result1 plugin_models.Organization, result2 error) {
	fake.GetCurrentOrgStub = nil
	if fake.getCurrentOrgReturnsOnCall == nil {
		fake.getCurrentOrgReturnsOnCall = make(map[int]struct {
			result1 plugin_models.Organization
			result2 error
		})
	}
	fake.getCurrentOrgReturnsOnCall[i] = struct {
		result1 plugin_models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentSpace() (plugin_models.Space, error) {
	fake.getCurrentSpaceMutex.Lock()
	ret, specificReturn := fake.getCurrentSpaceReturnsOnCall[len(fake.getCurrentSpaceArgsForCall)]
	fake.getCurrentSpaceArgsForCall = append(fake.getCurrentSpaceArgsForCall, struct{}{})
	fake.recordInvocation("GetCurrentSpace", []interface{}{})
	fake.getCurrentSpaceMutex.Unlock()
	if fake.GetCurrentSpaceStub != nil {
		return fake.GetCurrentSpaceStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCurrentSpaceReturns.result1, fake.getCurrentSpaceReturns.result2
}

func (fake *FakeCliConnection) GetCurrentSpaceCallCount() int {
	fake.getCurrentSpaceMutex.RLock()
	defer fake.getCurrentSpaceMutex.RUnlock()
	return len(fake.getCurrentSpaceArgsForCall)
}

func (fake *FakeCliConnection) GetCurrentSpaceReturns(result1 plugin_models.Space, result2 error) {
	fake.GetCurrentSpaceStub = nil
	fake.getCurrentSpaceReturns = struct {
		result1 plugin_models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentSpaceReturnsOnCall(i int, result1 plugin_models.Space, result2 error) {
	fake.GetCurrentSpaceStub = nil
	if fake.getCurrentSpaceReturnsOnCall == nil {
		fake.getCurrentSpaceReturnsOnCall = make(map[int]struct {
			result1 plugin_models.Space
			result2 error
		})
	}
	fake.getCurrentSpaceReturnsOnCall[i] = struct {
		result1 plugin_models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) Username() (string, error) {
	fake.usernameMutex.Lock()
	ret, specificReturn := fake.usernameReturnsOnCall[len(fake.usernameArgsForCall)]
	fake.usernameArgsForCall = append(fake.usernameArgsForCall, struct{}{})
	fake.recordInvocation("Username", []interface{}{})
	fake.usernameMutex.Unlock()
	if fake.UsernameStub != nil {
		return fake.UsernameStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.usernameReturns.result1, fake.usernameReturns.result2
}

func (fake *FakeCliConnection) UsernameCallCount() int {
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	return len(fake.usernameArgsForCall)
}

func (fake *FakeCliConnection) UsernameReturns(result1 string, result2 error) {
	fake.UsernameStub = nil
	fake.usernameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UsernameReturnsOnCall(i int, result1 string, result2 error) {
	fake.UsernameStub = nil
	if fake.usernameReturnsOnCall == nil {
		fake.usernameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.usernameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserGuid() (string, error) {
	fake.userGuidMutex.Lock()
	ret, specificReturn := fake.userGuidReturnsOnCall[len(fake.userGuidArgsForCall)]
	fake.userGuidArgsForCall = append(fake.userGuidArgsForCall, struct{}{})
	fake.recordInvocation("UserGuid", []interface{}{})
	fake.userGuidMutex.Unlock()
	if fake.UserGuidStub != nil {
		return fake.UserGuidStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.userGuidReturns.result1, fake.userGuidReturns.result2
}

func (fake *FakeCliConnection) UserGuidCallCount() int {
	fake.userGuidMutex.RLock()
	defer fake.userGuidMutex.RUnlock()
	return len(fake.userGuidArgsForCall)
}

func (fake *FakeCliConnection) UserGuidReturns(result1 string, result2 error) {
	fake.UserGuidStub = nil
	fake.userGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserGuidReturnsOnCall(i int, result1 string, result2 error) {
	fake.UserGuidStub = nil
	if fake.userGuidReturnsOnCall == nil {
		fake.userGuidReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.userGuidReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserEmail() (string, error) {
	fake.userEmailMutex.Lock()
	ret, specificReturn := fake.userEmailReturnsOnCall[len(fake.userEmailArgsForCall)]
	fake.userEmailArgsForCall = append(fake.userEmailArgsForCall, struct{}{})
	fake.recordInvocation("UserEmail", []interface{}{})
	fake.userEmailMutex.Unlock()
	if fake.UserEmailStub != nil {
		return fake.UserEmailStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.userEmailReturns.result1, fake.userEmailReturns.result2
}

func (fake *FakeCliConnection) UserEmailCallCount() int {
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	return len(fake.userEmailArgsForCall)
}

func (fake *FakeCliConnection) UserEmailReturns(result1 string, result2 error) {
	fake.UserEmailStub = nil
	fake.userEmailReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserEmailReturnsOnCall(i int, result1 string, result2 error) {
	fake.UserEmailStub = nil
	if fake.userEmailReturnsOnCall == nil {
		fake.userEmailReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.userEmailReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsLoggedIn() (bool, error) {
	fake.isLoggedInMutex.Lock()
	ret, specificReturn := fake.isLoggedInReturnsOnCall[len(fake.isLoggedInArgsForCall)]
	fake.isLoggedInArgsForCall = append(fake.isLoggedInArgsForCall, struct{}{})
	fake.recordInvocation("IsLoggedIn", []interface{}{})
	fake.isLoggedInMutex.Unlock()
	if fake.IsLoggedInStub != nil {
		return fake.IsLoggedInStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isLoggedInReturns.result1, fake.isLoggedInReturns.result2
}

func (fake *FakeCliConnection) IsLoggedInCallCount() int {
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	return len(fake.isLoggedInArgsForCall)
}

func (fake *FakeCliConnection) IsLoggedInReturns(result1 bool, result2 error) {
	fake.IsLoggedInStub = nil
	fake.isLoggedInReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsLoggedInReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsLoggedInStub = nil
	if fake.isLoggedInReturnsOnCall == nil {
		fake.isLoggedInReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isLoggedInReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsSSLDisabled() (bool, error) {
	fake.isSSLDisabledMutex.Lock()
	ret, specificReturn := fake.isSSLDisabledReturnsOnCall[len(fake.isSSLDisabledArgsForCall)]
	fake.isSSLDisabledArgsForCall = append(fake.isSSLDisabledArgsForCall, struct{}{})
	fake.recordInvocation("IsSSLDisabled", []interface{}{})
	fake.isSSLDisabledMutex.Unlock()
	if fake.IsSSLDisabledStub != nil {
		return fake.IsSSLDisabledStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isSSLDisabledReturns.result1, fake.isSSLDisabledReturns.result2
}

func (fake *FakeCliConnection) IsSSLDisabledCallCount() int {
	fake.isSSLDisabledMutex.RLock()
	defer fake.isSSLDisabledMutex.RUnlock()
	return len(fake.isSSLDisabledArgsForCall)
}

func (fake *FakeCliConnection) IsSSLDisabledReturns(result1 bool, result2 error) {
	fake.IsSSLDisabledStub = nil
	fake.isSSLDisabledReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsSSLDisabledReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsSSLDisabledStub = nil
	if fake.isSSLDisabledReturnsOnCall == nil {
		fake.isSSLDisabledReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isSSLDisabledReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasOrganization() (bool, error) {
	fake.hasOrganizationMutex.Lock()
	ret, specificReturn := fake.hasOrganizationReturnsOnCall[len(fake.hasOrganizationArgsForCall)]
	fake.hasOrganizationArgsForCall = append(fake.hasOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("HasOrganization", []interface{}{})
	fake.hasOrganizationMutex.Unlock()
	if fake.HasOrganizationStub != nil {
		return fake.HasOrganizationStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasOrganizationReturns.result1, fake.hasOrganizationReturns.result2
}

func (fake *FakeCliConnection) HasOrganizationCallCount() int {
	fake.hasOrganizationMutex.RLock()
	defer fake.hasOrganizationMutex.RUnlock()
	return len(fake.hasOrganizationArgsForCall)
}

func (fake *FakeCliConnection) HasOrganizationReturns(result1 bool, result2 error) {
	fake.HasOrganizationStub = nil
	fake.hasOrganizationReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasOrganizationReturnsOnCall(i int, result1 bool, result2 error) {
	fake.HasOrganizationStub = nil
	if fake.hasOrganizationReturnsOnCall == nil {
		fake.hasOrganizationReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasOrganizationReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasSpace() (bool, error) {
	fake.hasSpaceMutex.Lock()
	ret, specificReturn := fake.hasSpaceReturnsOnCall[len(fake.hasSpaceArgsForCall)]
	fake.hasSpaceArgsForCall = append(fake.hasSpaceArgsForCall, struct{}{})
	fake.recordInvocation("HasSpace", []interface{}{})
	fake.hasSpaceMutex.Unlock()
	if fake.HasSpaceStub != nil {
		return fake.HasSpaceStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasSpaceReturns.result1, fake.hasSpaceReturns.result2
}

func (fake *FakeCliConnection) HasSpaceCallCount() int {
	fake.hasSpaceMutex.RLock()
	defer fake.hasSpaceMutex.RUnlock()
	return len(fake.hasSpaceArgsForCall)
}

func (fake *FakeCliConnection) HasSpaceReturns(result1 bool, result2 error) {
	fake.HasSpaceStub = nil
	fake.hasSpaceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasSpaceReturnsOnCall(i int, result1 bool, result2 error) {
	fake.HasSpaceStub = nil
	if fake.hasSpaceReturnsOnCall == nil {
		fake.hasSpaceReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasSpaceReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiEndpoint() (string, error) {
	fake.apiEndpointMutex.Lock()
	ret, specificReturn := fake.apiEndpointReturnsOnCall[len(fake.apiEndpointArgsForCall)]
	fake.apiEndpointArgsForCall = append(fake.apiEndpointArgsForCall, struct{}{})
	fake.recordInvocation("ApiEndpoint", []interface{}{})
	fake.apiEndpointMutex.Unlock()
	if fake.ApiEndpointStub != nil {
		return fake.ApiEndpointStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.apiEndpointReturns.result1, fake.apiEndpointReturns.result2
}

func (fake *FakeCliConnection) ApiEndpointCallCount() int {
	fake.apiEndpointMutex.RLock()
	defer fake.apiEndpointMutex.RUnlock()
	return len(fake.apiEndpointArgsForCall)
}

func (fake *FakeCliConnection) ApiEndpointReturns(result1 string, result2 error) {
	fake.ApiEndpointStub = nil
	fake.apiEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiEndpointReturnsOnCall(i int, result1 string, result2 error) {
	fake.ApiEndpointStub = nil
	if fake.apiEndpointReturnsOnCall == nil {
		fake.apiEndpointReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.apiEndpointReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiVersion() (string, error) {
	fake.apiVersionMutex.Lock()
	ret, specificReturn := fake.apiVersionReturnsOnCall[len(fake.apiVersionArgsForCall)]
	fake.apiVersionArgsForCall = append(fake.apiVersionArgsForCall, struct{}{})
	fake.recordInvocation("ApiVersion", []interface{}{})
	fake.apiVersionMutex.Unlock()
	if fake.ApiVersionStub != nil {
		return fake.ApiVersionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.apiVersionReturns.result1, fake.apiVersionReturns.result2
}

func (fake *FakeCliConnection) ApiVersionCallCount() int {
	fake.apiVersionMutex.RLock()
	defer fake.apiVersionMutex.RUnlock()
	return len(fake.apiVersionArgsForCall)
}

func (fake *FakeCliConnection) ApiVersionReturns(result1 string, result2 error) {
	fake.ApiVersionStub = nil
	fake.apiVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.ApiVersionStub = nil
	if fake.apiVersionReturnsOnCall == nil {
		fake.apiVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.apiVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasAPIEndpoint() (bool, error) {
	fake.hasAPIEndpointMutex.Lock()
	ret, specificReturn := fake.hasAPIEndpointReturnsOnCall[len(fake.hasAPIEndpointArgsForCall)]
	fake.hasAPIEndpointArgsForCall = append(fake.hasAPIEndpointArgsForCall, struct{}{})
	fake.recordInvocation("HasAPIEndpoint", []interface{}{})
	fake.hasAPIEndpointMutex.Unlock()
	if fake.HasAPIEndpointStub != nil {
		return fake.HasAPIEndpointStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasAPIEndpointReturns.result1, fake.hasAPIEndpointReturns.result2
}

func (fake *FakeCliConnection) HasAPIEndpointCallCount() int {
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	return len(fake.hasAPIEndpointArgsForCall)
}

func (fake *FakeCliConnection) HasAPIEndpointReturns(result1 bool, result2 error) {
	fake.HasAPIEndpointStub = nil
	fake.hasAPIEndpointReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasAPIEndpointReturnsOnCall(i int, result1 bool, result2 error) {
	fake.HasAPIEndpointStub = nil
	if fake.hasAPIEndpointReturnsOnCall == nil {
		fake.hasAPIEndpointReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasAPIEndpointReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) LoggregatorEndpoint() (string, error) {
	fake.loggregatorEndpointMutex.Lock()
	ret, specificReturn := fake.loggregatorEndpointReturnsOnCall[len(fake.loggregatorEndpointArgsForCall)]
	fake.loggregatorEndpointArgsForCall = append(fake.loggregatorEndpointArgsForCall, struct{}{})
	fake.recordInvocation("LoggregatorEndpoint", []interface{}{})
	fake.loggregatorEndpointMutex.Unlock()
	if fake.LoggregatorEndpointStub != nil {
		return fake.LoggregatorEndpointStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loggregatorEndpointReturns.result1, fake.loggregatorEndpointReturns.result2
}

func (fake *FakeCliConnection) LoggregatorEndpointCallCount() int {
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	return len(fake.loggregatorEndpointArgsForCall)
}

func (fake *FakeCliConnection) LoggregatorEndpointReturns(result1 string, result2 error) {
	fake.LoggregatorEndpointStub = nil
	fake.loggregatorEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) LoggregatorEndpointReturnsOnCall(i int, result1 string, result2 error) {
	fake.LoggregatorEndpointStub = nil
	if fake.loggregatorEndpointReturnsOnCall == nil {
		fake.loggregatorEndpointReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.loggregatorEndpointReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) DopplerEndpoint() (string, error) {
	fake.dopplerEndpointMutex.Lock()
	ret, specificReturn := fake.dopplerEndpointReturnsOnCall[len(fake.dopplerEndpointArgsForCall)]
	fake.dopplerEndpointArgsForCall = append(fake.dopplerEndpointArgsForCall, struct{}{})
	fake.recordInvocation("DopplerEndpoint", []interface{}{})
	fake.dopplerEndpointMutex.Unlock()
	if fake.DopplerEndpointStub != nil {
		return fake.DopplerEndpointStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.dopplerEndpointReturns.result1, fake.dopplerEndpointReturns.result2
}

func (fake *FakeCliConnection) DopplerEndpointCallCount() int {
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	return len(fake.dopplerEndpointArgsForCall)
}

func (fake *FakeCliConnection) DopplerEndpointReturns(result1 string, result2 error) {
	fake.DopplerEndpointStub = nil
	fake.dopplerEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) DopplerEndpointReturnsOnCall(i int, result1 string, result2 error) {
	fake.DopplerEndpointStub = nil
	if fake.dopplerEndpointReturnsOnCall == nil {
		fake.dopplerEndpointReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.dopplerEndpointReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) AccessToken() (string, error) {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.accessTokenReturns.result1, fake.accessTokenReturns.result2
}

func (fake *FakeCliConnection) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeCliConnection) AccessTokenReturns(result1 string, result2 error) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) AccessTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetApp(arg1 string) (plugin_models.GetAppModel, error) {
	fake.getAppMutex.Lock()
	ret, specificReturn := fake.getAppReturnsOnCall[len(fake.getAppArgsForCall)]
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetApp", []interface{}{arg1})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppReturns.result1, fake.getAppReturns.result2
}

func (fake *FakeCliConnection) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeCliConnection) GetAppArgsForCall(i int) string {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetAppReturns(result1 plugin_models.GetAppModel, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 plugin_models.GetAppModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetAppReturnsOnCall(i int, result1 plugin_models.GetAppModel, result2 error) {
	fake.GetAppStub = nil
	if fake.getAppReturnsOnCall == nil {
		fake.getAppReturnsOnCall = make(map[int]struct {
			result1 plugin_models.GetAppModel
			result2 error
		})
	}
	fake.getAppReturnsOnCall[i] = struct {
		result1 plugin_models.GetAppModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetApps() ([]plugin_models.GetAppsModel, error) {
	fake.getAppsMutex.Lock()
	ret, specificReturn := fake.getAppsReturnsOnCall[len(fake.getAppsArgsForCall)]
	fake.getAppsArgsForCall = append(fake.getAppsArgsForCall, struct{}{})
	fake.recordInvocation("GetApps", []interface{}{})
	fake.getAppsMutex.Unlock()
	if fake.GetAppsStub != nil {
		return fake.GetAppsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppsReturns.result1, fake.getAppsReturns.result2
}

func (fake *FakeCliConnection) GetAppsCallCount() int {
	fake.getAppsMutex.RLock()
	defer fake.getAppsMutex.RUnlock()
	return len(fake.getAppsArgsForCall)
}

func (fake *FakeCliConnection) GetAppsReturns(result1 []plugin_models.GetAppsModel, result2 error) {
	fake.GetAppsStub = nil
	fake.getAppsReturns = struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetAppsReturnsOnCall(i int, result1 []plugin_models.GetAppsModel, result2 error) {
	fake.GetAppsStub = nil
	if fake.getAppsReturnsOnCall == nil {
		fake.getAppsReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetAppsModel
			result2 error
		})
	}
	fake.getAppsReturnsOnCall[i] = struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgs() ([]plugin_models.GetOrgs_Model, error) {
	fake.getOrgsMutex.Lock()
	ret, specificReturn := fake.getOrgsReturnsOnCall[len(fake.getOrgsArgsForCall)]
	fake.getOrgsArgsForCall = append(fake.getOrgsArgsForCall, struct{}{})
	fake.recordInvocation("GetOrgs", []interface{}{})
	fake.getOrgsMutex.Unlock()
	if fake.GetOrgsStub != nil {
		return fake.GetOrgsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrgsReturns.result1, fake.getOrgsReturns.result2
}

func (fake *FakeCliConnection) GetOrgsCallCount() int {
	fake.getOrgsMutex.RLock()
	defer fake.getOrgsMutex.RUnlock()
	return len(fake.getOrgsArgsForCall)
}

func (fake *FakeCliConnection) GetOrgsReturns(result1 []plugin_models.GetOrgs_Model, result2 error) {
	fake.GetOrgsStub = nil
	fake.getOrgsReturns = struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgsReturnsOnCall(i int, result1 []plugin_models.GetOrgs_Model, result2 error) {
	fake.GetOrgsStub = nil
	if fake.getOrgsReturnsOnCall == nil {
		fake.getOrgsReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetOrgs_Model
			result2 error
		})
	}
	fake.getOrgsReturnsOnCall[i] = struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaces() ([]plugin_models.GetSpaces_Model, error) {
	fake.getSpacesMutex.Lock()
	ret, specificReturn := fake.getSpacesReturnsOnCall[len(fake.getSpacesArgsForCall)]
	fake.getSpacesArgsForCall = append(fake.getSpacesArgsForCall, struct{}{})
	fake.recordInvocation("GetSpaces", []interface{}{})
	fake.getSpacesMutex.Unlock()
	if fake.GetSpacesStub != nil {
		return fake.GetSpacesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpacesReturns.result1, fake.getSpacesReturns.result2
}

func (fake *FakeCliConnection) GetSpacesCallCount() int {
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	return len(fake.getSpacesArgsForCall)
}

func (fake *FakeCliConnection) GetSpacesReturns(result1 []plugin_models.GetSpaces_Model, result2 error) {
	fake.GetSpacesStub = nil
	fake.getSpacesReturns = struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpacesReturnsOnCall(i int, result1 []plugin_models.GetSpaces_Model, result2 error) {
	fake.GetSpacesStub = nil
	if fake.getSpacesReturnsOnCall == nil {
		fake.getSpacesReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetSpaces_Model
			result2 error
		})
	}
	fake.getSpacesReturnsOnCall[i] = struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgUsers(arg1 string, arg2 ...string) ([]plugin_models.GetOrgUsers_Model, error) {
	fake.getOrgUsersMutex.Lock()
	ret, specificReturn := fake.getOrgUsersReturnsOnCall[len(fake.getOrgUsersArgsForCall)]
	fake.getOrgUsersArgsForCall = append(fake.getOrgUsersArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("GetOrgUsers", []interface{}{arg1, arg2})
	fake.getOrgUsersMutex.Unlock()
	if fake.GetOrgUsersStub != nil {
		return fake.GetOrgUsersStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrgUsersReturns.result1, fake.getOrgUsersReturns.result2
}

func (fake *FakeCliConnection) GetOrgUsersCallCount() int {
	fake.getOrgUsersMutex.RLock()
	defer fake.getOrgUsersMutex.RUnlock()
	return len(fake.getOrgUsersArgsForCall)
}

func (fake *FakeCliConnection) GetOrgUsersArgsForCall(i int) (string, []string) {
	fake.getOrgUsersMutex.RLock()
	defer fake.getOrgUsersMutex.RUnlock()
	return fake.getOrgUsersArgsForCall[i].arg1, fake.getOrgUsersArgsForCall[i].arg2
}

func (fake *FakeCliConnection) GetOrgUsersReturns(result1 []plugin_models.GetOrgUsers_Model, result2 error) {
	fake.GetOrgUsersStub = nil
	fake.getOrgUsersReturns = struct {
		result1 []plugin_models.GetOrgUsers_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgUsersReturnsOnCall(i int, result1 []plugin_models.GetOrgUsers_Model, result2 error) {
	fake.GetOrgUsersStub = nil
	if fake.getOrgUsersReturnsOnCall == nil {
		fake.getOrgUsersReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetOrgUsers_Model
			result2 error
		})
	}
	fake.getOrgUsersReturnsOnCall[i] = struct {
		result1 []plugin_models.GetOrgUsers_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaceUsers(arg1 string, arg2 string) ([]plugin_models.GetSpaceUsers_Model, error) {
	fake.getSpaceUsersMutex.Lock()
	ret, specificReturn := fake.getSpaceUsersReturnsOnCall[len(fake.getSpaceUsersArgsForCall)]
	fake.getSpaceUsersArgsForCall = append(fake.getSpaceUsersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceUsers", []interface{}{arg1, arg2})
	fake.getSpaceUsersMutex.Unlock()
	if fake.GetSpaceUsersStub != nil {
		return fake.GetSpaceUsersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceUsersReturns.result1, fake.getSpaceUsersReturns.result2
}

func (fake *FakeCliConnection) GetSpaceUsersCallCount() int {
	fake.getSpaceUsersMutex.RLock()
	defer fake.getSpaceUsersMutex.RUnlock()
	return len(fake.getSpaceUsersArgsForCall)
}

func (fake *FakeCliConnection) GetSpaceUsersArgsForCall(i int) (string, string) {
	fake.getSpaceUsersMutex.RLock()
	defer fake.getSpaceUsersMutex.RUnlock()
	return fake.getSpaceUsersArgsForCall[i].arg1, fake.getSpaceUsersArgsForCall[i].arg2
}

func (fake *FakeCliConnection) GetSpaceUsersReturns(result1 []plugin_models.GetSpaceUsers_Model, result2 error) {
	fake.GetSpaceUsersStub = nil
	fake.getSpaceUsersReturns = struct {
		result1 []plugin_models.GetSpaceUsers_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaceUsersReturnsOnCall(i int, result1 []plugin_models.GetSpaceUsers_Model, result2 error) {
	fake.GetSpaceUsersStub = nil
	if fake.getSpaceUsersReturnsOnCall == nil {
		fake.getSpaceUsersReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetSpaceUsers_Model
			result2 error
		})
	}
	fake.getSpaceUsersReturnsOnCall[i] = struct {
		result1 []plugin_models.GetSpaceUsers_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetServices() ([]plugin_models.GetServices_Model, error) {
	fake.getServicesMutex.Lock()
	ret, specificReturn := fake.getServicesReturnsOnCall[len(fake.getServicesArgsForCall)]
	fake.getServicesArgsForCall = append(fake.getServicesArgsForCall, struct{}{})
	fake.recordInvocation("GetServices", []interface{}{})
	fake.getServicesMutex.Unlock()
	if fake.GetServicesStub != nil {
		return fake.GetServicesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServicesReturns.result1, fake.getServicesReturns.result2
}

func (fake *FakeCliConnection) GetServicesCallCount() int {
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	return len(fake.getServicesArgsForCall)
}

func (fake *FakeCliConnection) GetServicesReturns(result1 []plugin_models.GetServices_Model, result2 error) {
	fake.GetServicesStub = nil
	fake.getServicesReturns = struct {
		result1 []plugin_models.GetServices_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetServicesReturnsOnCall(i int, result1 []plugin_models.GetServices_Model, result2 error) {
	fake.GetServicesStub = nil
	if fake.getServicesReturnsOnCall == nil {
		fake.getServicesReturnsOnCall = make(map[int]struct {
			result1 []plugin_models.GetServices_Model
			result2 error
		})
	}
	fake.getServicesReturnsOnCall[i] = struct {
		result1 []plugin_models.GetServices_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetService(arg1 string) (plugin_models.GetService_Model, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetService", []interface{}{arg1})
	fake.getServiceMutex.Unlock()
	if fake.GetServiceStub != nil {
		return fake.GetServiceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceReturns.result1, fake.getServiceReturns.result2
}

func (fake *FakeCliConnection) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *FakeCliConnection) GetServiceArgsForCall(i int) string {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return fake.getServiceArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetServiceReturns(result1 plugin_models.GetService_Model, result2 error) {
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 plugin_models.GetService_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetServiceReturnsOnCall(i int, result1 plugin_models.GetService_Model, result2 error) {
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 plugin_models.GetService_Model
			result2 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 plugin_models.GetService_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrg(arg1 string) (plugin_models.GetOrg_Model, error) {
	fake.getOrgMutex.Lock()
	ret, specificReturn := fake.getOrgReturnsOnCall[len(fake.getOrgArgsForCall)]
	fake.getOrgArgsForCall = append(fake.getOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrg", []interface{}{arg1})
	fake.getOrgMutex.Unlock()
	if fake.GetOrgStub != nil {
		return fake.GetOrgStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOrgReturns.result1, fake.getOrgReturns.result2
}

func (fake *FakeCliConnection) GetOrgCallCount() int {
	fake.getOrgMutex.RLock()
	defer fake.getOrgMutex.RUnlock()
	return len(fake.getOrgArgsForCall)
}

func (fake *FakeCliConnection) GetOrgArgsForCall(i int) string {
	fake.getOrgMutex.RLock()
	defer fake.getOrgMutex.RUnlock()
	return fake.getOrgArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetOrgReturns(result1 plugin_models.GetOrg_Model, result2 error) {
	fake.GetOrgStub = nil
	fake.getOrgReturns = struct {
		result1 plugin_models.GetOrg_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgReturnsOnCall(i int, result1 plugin_models.GetOrg_Model, result2 error) {
	fake.GetOrgStub = nil
	if fake.getOrgReturnsOnCall == nil {
		fake.getOrgReturnsOnCall = make(map[int]struct {
			result1 plugin_models.GetOrg_Model
			result2 error
		})
	}
	fake.getOrgReturnsOnCall[i] = struct {
		result1 plugin_models.GetOrg_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpace(arg1 string) (plugin_models.GetSpace_Model, error) {
	fake.getSpaceMutex.Lock()
	ret, specificReturn := fake.getSpaceReturnsOnCall[len(fake.getSpaceArgsForCall)]
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSpace", []interface{}{arg1})
	fake.getSpaceMutex.Unlock()
	if fake.GetSpaceStub != nil {
		return fake.GetSpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceReturns.result1, fake.getSpaceReturns.result2
}

func (fake *FakeCliConnection) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *FakeCliConnection) GetSpaceArgsForCall(i int) string {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return fake.getSpaceArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetSpaceReturns(result1 plugin_models.GetSpace_Model, result2 error) {
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 plugin_models.GetSpace_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaceReturnsOnCall(i int, result1 plugin_models.GetSpace_Model, result2 error) {
	fake.GetSpaceStub = nil
	if fake.getSpaceReturnsOnCall == nil {
		fake.getSpaceReturnsOnCall = make(map[int]struct {
			result1 plugin_models.GetSpace_Model
			result2 error
		})
	}
	fake.getSpaceReturnsOnCall[i] = struct {
		result1 plugin_models.GetSpace_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	fake.getCurrentOrgMutex.RLock()
	defer fake.getCurrentOrgMutex.RUnlock()
	fake.getCurrentSpaceMutex.RLock()
	defer fake.getCurrentSpaceMutex.RUnlock()
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	fake.userGuidMutex.RLock()
	defer fake.userGuidMutex.RUnlock()
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	fake.isSSLDisabledMutex.RLock()
	defer fake.isSSLDisabledMutex.RUnlock()
	fake.hasOrganizationMutex.RLock()
	defer fake.hasOrganizationMutex.RUnlock()
	fake.hasSpaceMutex.RLock()
	defer fake.hasSpaceMutex.RUnlock()
	fake.apiEndpointMutex.RLock()
	defer fake.apiEndpointMutex.RUnlock()
	fake.apiVersionMutex.RLock()
	defer fake.apiVersionMutex.RUnlock()
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	fake.getAppsMutex.RLock()
	defer fake.getAppsMutex.RUnlock()
	fake.getOrgsMutex.RLock()
	defer fake.getOrgsMutex.RUnlock()
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	fake.getOrgUsersMutex.RLock()
	defer fake.getOrgUsersMutex.RUnlock()
	fake.getSpaceUsersMutex.RLock()
	defer fake.getSpaceUsersMutex.RUnlock()
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	fake.getOrgMutex.RLock()
	defer fake.getOrgMutex.RUnlock()
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCliConnection) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ plugin.CliConnection = new(FakeCliConnection)
