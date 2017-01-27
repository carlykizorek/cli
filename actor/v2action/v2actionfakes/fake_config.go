// This file was generated by counterfeiter
package v2actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
)

type FakeConfig struct {
	UnsetOrganizationInformationStub        func()
	unsetOrganizationInformationMutex       sync.RWMutex
	unsetOrganizationInformationArgsForCall []struct{}
	UnsetSpaceInformationStub               func()
	unsetSpaceInformationMutex              sync.RWMutex
	unsetSpaceInformationArgsForCall        []struct{}
	SetTargetInformationStub                func(api string, apiVersion string, auth string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool)
	setTargetInformationMutex               sync.RWMutex
	setTargetInformationArgsForCall         []struct {
		api               string
		apiVersion        string
		auth              string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}
	SetTokenInformationStub        func(accessToken string, refreshToken string, sshOAuthClient string)
	setTokenInformationMutex       sync.RWMutex
	setTokenInformationArgsForCall []struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct{}
	skipSSLValidationReturns     struct {
		result1 bool
	}
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct{}
	targetReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) UnsetOrganizationInformation() {
	fake.unsetOrganizationInformationMutex.Lock()
	fake.unsetOrganizationInformationArgsForCall = append(fake.unsetOrganizationInformationArgsForCall, struct{}{})
	fake.recordInvocation("UnsetOrganizationInformation", []interface{}{})
	fake.unsetOrganizationInformationMutex.Unlock()
	if fake.UnsetOrganizationInformationStub != nil {
		fake.UnsetOrganizationInformationStub()
	}
}

func (fake *FakeConfig) UnsetOrganizationInformationCallCount() int {
	fake.unsetOrganizationInformationMutex.RLock()
	defer fake.unsetOrganizationInformationMutex.RUnlock()
	return len(fake.unsetOrganizationInformationArgsForCall)
}

func (fake *FakeConfig) UnsetSpaceInformation() {
	fake.unsetSpaceInformationMutex.Lock()
	fake.unsetSpaceInformationArgsForCall = append(fake.unsetSpaceInformationArgsForCall, struct{}{})
	fake.recordInvocation("UnsetSpaceInformation", []interface{}{})
	fake.unsetSpaceInformationMutex.Unlock()
	if fake.UnsetSpaceInformationStub != nil {
		fake.UnsetSpaceInformationStub()
	}
}

func (fake *FakeConfig) UnsetSpaceInformationCallCount() int {
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	return len(fake.unsetSpaceInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformation(api string, apiVersion string, auth string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		api               string
		apiVersion        string
		auth              string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}{api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.recordInvocation("SetTargetInformation", []interface{}{api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) (string, string, string, string, string, string, string, bool) {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.setTargetInformationArgsForCall[i].api, fake.setTargetInformationArgsForCall[i].apiVersion, fake.setTargetInformationArgsForCall[i].auth, fake.setTargetInformationArgsForCall[i].minCLIVersion, fake.setTargetInformationArgsForCall[i].doppler, fake.setTargetInformationArgsForCall[i].uaa, fake.setTargetInformationArgsForCall[i].routing, fake.setTargetInformationArgsForCall[i].skipSSLValidation
}

func (fake *FakeConfig) SetTokenInformation(accessToken string, refreshToken string, sshOAuthClient string) {
	fake.setTokenInformationMutex.Lock()
	fake.setTokenInformationArgsForCall = append(fake.setTokenInformationArgsForCall, struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}{accessToken, refreshToken, sshOAuthClient})
	fake.recordInvocation("SetTokenInformation", []interface{}{accessToken, refreshToken, sshOAuthClient})
	fake.setTokenInformationMutex.Unlock()
	if fake.SetTokenInformationStub != nil {
		fake.SetTokenInformationStub(accessToken, refreshToken, sshOAuthClient)
	}
}

func (fake *FakeConfig) SetTokenInformationCallCount() int {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return len(fake.setTokenInformationArgsForCall)
}

func (fake *FakeConfig) SetTokenInformationArgsForCall(i int) (string, string, string) {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return fake.setTokenInformationArgsForCall[i].accessToken, fake.setTokenInformationArgsForCall[i].refreshToken, fake.setTokenInformationArgsForCall[i].sshOAuthClient
}

func (fake *FakeConfig) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	} else {
		return fake.skipSSLValidationReturns.result1
	}
}

func (fake *FakeConfig) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeConfig) SkipSSLValidationReturns(result1 bool) {
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Target() string {
	fake.targetMutex.Lock()
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct{}{})
	fake.recordInvocation("Target", []interface{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	} else {
		return fake.targetReturns.result1
	}
}

func (fake *FakeConfig) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeConfig) TargetReturns(result1 string) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unsetOrganizationInformationMutex.RLock()
	defer fake.unsetOrganizationInformationMutex.RUnlock()
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ v2action.Config = new(FakeConfig)