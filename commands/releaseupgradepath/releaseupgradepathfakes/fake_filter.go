// This file was generated by counterfeiter
package releaseupgradepathfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands/releaseupgradepath"
)

type FakeFilter struct {
	ReleasesByVersionStub        func(releases []go_pivnet.Release, version string) ([]go_pivnet.Release, error)
	releasesByVersionMutex       sync.RWMutex
	releasesByVersionArgsForCall []struct {
		releases []go_pivnet.Release
		version  string
	}
	releasesByVersionReturns struct {
		result1 []go_pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilter) ReleasesByVersion(releases []go_pivnet.Release, version string) ([]go_pivnet.Release, error) {
	var releasesCopy []go_pivnet.Release
	if releases != nil {
		releasesCopy = make([]go_pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.releasesByVersionMutex.Lock()
	fake.releasesByVersionArgsForCall = append(fake.releasesByVersionArgsForCall, struct {
		releases []go_pivnet.Release
		version  string
	}{releasesCopy, version})
	fake.recordInvocation("ReleasesByVersion", []interface{}{releasesCopy, version})
	fake.releasesByVersionMutex.Unlock()
	if fake.ReleasesByVersionStub != nil {
		return fake.ReleasesByVersionStub(releases, version)
	} else {
		return fake.releasesByVersionReturns.result1, fake.releasesByVersionReturns.result2
	}
}

func (fake *FakeFilter) ReleasesByVersionCallCount() int {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return len(fake.releasesByVersionArgsForCall)
}

func (fake *FakeFilter) ReleasesByVersionArgsForCall(i int) ([]go_pivnet.Release, string) {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return fake.releasesByVersionArgsForCall[i].releases, fake.releasesByVersionArgsForCall[i].version
}

func (fake *FakeFilter) ReleasesByVersionReturns(result1 []go_pivnet.Release, result2 error) {
	fake.ReleasesByVersionStub = nil
	fake.releasesByVersionReturns = struct {
		result1 []go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFilter) recordInvocation(key string, args []interface{}) {
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

var _ releaseupgradepath.Filter = new(FakeFilter)
