// This file was generated by counterfeiter
package productfilefakes

import (
	"io"
	"os"
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands/productfile"
)

type FakePivnetClient struct {
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	ProductFilesStub        func(productSlug string) ([]go_pivnet.ProductFile, error)
	productFilesMutex       sync.RWMutex
	productFilesArgsForCall []struct {
		productSlug string
	}
	productFilesReturns struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}
	ProductFilesForReleaseStub        func(productSlug string, releaseID int) ([]go_pivnet.ProductFile, error)
	productFilesForReleaseMutex       sync.RWMutex
	productFilesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	productFilesForReleaseReturns struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}
	ProductFileStub        func(productSlug string, productFileID int) (go_pivnet.ProductFile, error)
	productFileMutex       sync.RWMutex
	productFileArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	productFileReturns struct {
		result1 go_pivnet.ProductFile
		result2 error
	}
	ProductFileForReleaseStub        func(productSlug string, releaseID int, productFileID int) (go_pivnet.ProductFile, error)
	productFileForReleaseMutex       sync.RWMutex
	productFileForReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	productFileForReleaseReturns struct {
		result1 go_pivnet.ProductFile
		result2 error
	}
	CreateProductFileStub        func(config go_pivnet.CreateProductFileConfig) (go_pivnet.ProductFile, error)
	createProductFileMutex       sync.RWMutex
	createProductFileArgsForCall []struct {
		config go_pivnet.CreateProductFileConfig
	}
	createProductFileReturns struct {
		result1 go_pivnet.ProductFile
		result2 error
	}
	UpdateProductFileStub        func(productSlug string, productFile go_pivnet.ProductFile) (go_pivnet.ProductFile, error)
	updateProductFileMutex       sync.RWMutex
	updateProductFileArgsForCall []struct {
		productSlug string
		productFile go_pivnet.ProductFile
	}
	updateProductFileReturns struct {
		result1 go_pivnet.ProductFile
		result2 error
	}
	AddProductFileToReleaseStub        func(productSlug string, releaseID int, productFileID int) error
	addProductFileToReleaseMutex       sync.RWMutex
	addProductFileToReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	addProductFileToReleaseReturns struct {
		result1 error
	}
	RemoveProductFileFromReleaseStub        func(productSlug string, releaseID int, productFileID int) error
	removeProductFileFromReleaseMutex       sync.RWMutex
	removeProductFileFromReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	removeProductFileFromReleaseReturns struct {
		result1 error
	}
	AddProductFileToFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	addProductFileToFileGroupMutex       sync.RWMutex
	addProductFileToFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	addProductFileToFileGroupReturns struct {
		result1 error
	}
	RemoveProductFileFromFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	removeProductFileFromFileGroupMutex       sync.RWMutex
	removeProductFileFromFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	removeProductFileFromFileGroupReturns struct {
		result1 error
	}
	DeleteProductFileStub        func(productSlug string, releaseID int) (go_pivnet.ProductFile, error)
	deleteProductFileMutex       sync.RWMutex
	deleteProductFileArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	deleteProductFileReturns struct {
		result1 go_pivnet.ProductFile
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	DownloadProductFileStub        func(location *os.File, productSlug string, releaseID int, productFileID int, progressWriter io.Writer) error
	downloadProductFileMutex       sync.RWMutex
	downloadProductFileArgsForCall []struct {
		location       *os.File
		productSlug    string
		releaseID      int
		productFileID  int
		progressWriter io.Writer
	}
	downloadProductFileReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 go_pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFiles(productSlug string) ([]go_pivnet.ProductFile, error) {
	fake.productFilesMutex.Lock()
	fake.productFilesArgsForCall = append(fake.productFilesArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("ProductFiles", []interface{}{productSlug})
	fake.productFilesMutex.Unlock()
	if fake.ProductFilesStub != nil {
		return fake.ProductFilesStub(productSlug)
	} else {
		return fake.productFilesReturns.result1, fake.productFilesReturns.result2
	}
}

func (fake *FakePivnetClient) ProductFilesCallCount() int {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	return len(fake.productFilesArgsForCall)
}

func (fake *FakePivnetClient) ProductFilesArgsForCall(i int) string {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	return fake.productFilesArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) ProductFilesReturns(result1 []go_pivnet.ProductFile, result2 error) {
	fake.ProductFilesStub = nil
	fake.productFilesReturns = struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesForRelease(productSlug string, releaseID int) ([]go_pivnet.ProductFile, error) {
	fake.productFilesForReleaseMutex.Lock()
	fake.productFilesForReleaseArgsForCall = append(fake.productFilesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ProductFilesForRelease", []interface{}{productSlug, releaseID})
	fake.productFilesForReleaseMutex.Unlock()
	if fake.ProductFilesForReleaseStub != nil {
		return fake.ProductFilesForReleaseStub(productSlug, releaseID)
	} else {
		return fake.productFilesForReleaseReturns.result1, fake.productFilesForReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) ProductFilesForReleaseCallCount() int {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return len(fake.productFilesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFilesForReleaseArgsForCall(i int) (string, int) {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return fake.productFilesForReleaseArgsForCall[i].productSlug, fake.productFilesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ProductFilesForReleaseReturns(result1 []go_pivnet.ProductFile, result2 error) {
	fake.ProductFilesForReleaseStub = nil
	fake.productFilesForReleaseReturns = struct {
		result1 []go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFile(productSlug string, productFileID int) (go_pivnet.ProductFile, error) {
	fake.productFileMutex.Lock()
	fake.productFileArgsForCall = append(fake.productFileArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("ProductFile", []interface{}{productSlug, productFileID})
	fake.productFileMutex.Unlock()
	if fake.ProductFileStub != nil {
		return fake.ProductFileStub(productSlug, productFileID)
	} else {
		return fake.productFileReturns.result1, fake.productFileReturns.result2
	}
}

func (fake *FakePivnetClient) ProductFileCallCount() int {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	return len(fake.productFileArgsForCall)
}

func (fake *FakePivnetClient) ProductFileArgsForCall(i int) (string, int) {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	return fake.productFileArgsForCall[i].productSlug, fake.productFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) ProductFileReturns(result1 go_pivnet.ProductFile, result2 error) {
	fake.ProductFileStub = nil
	fake.productFileReturns = struct {
		result1 go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileForRelease(productSlug string, releaseID int, productFileID int) (go_pivnet.ProductFile, error) {
	fake.productFileForReleaseMutex.Lock()
	fake.productFileForReleaseArgsForCall = append(fake.productFileForReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("ProductFileForRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.productFileForReleaseMutex.Unlock()
	if fake.ProductFileForReleaseStub != nil {
		return fake.ProductFileForReleaseStub(productSlug, releaseID, productFileID)
	} else {
		return fake.productFileForReleaseReturns.result1, fake.productFileForReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) ProductFileForReleaseCallCount() int {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return len(fake.productFileForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFileForReleaseArgsForCall(i int) (string, int, int) {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return fake.productFileForReleaseArgsForCall[i].productSlug, fake.productFileForReleaseArgsForCall[i].releaseID, fake.productFileForReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) ProductFileForReleaseReturns(result1 go_pivnet.ProductFile, result2 error) {
	fake.ProductFileForReleaseStub = nil
	fake.productFileForReleaseReturns = struct {
		result1 go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateProductFile(config go_pivnet.CreateProductFileConfig) (go_pivnet.ProductFile, error) {
	fake.createProductFileMutex.Lock()
	fake.createProductFileArgsForCall = append(fake.createProductFileArgsForCall, struct {
		config go_pivnet.CreateProductFileConfig
	}{config})
	fake.recordInvocation("CreateProductFile", []interface{}{config})
	fake.createProductFileMutex.Unlock()
	if fake.CreateProductFileStub != nil {
		return fake.CreateProductFileStub(config)
	} else {
		return fake.createProductFileReturns.result1, fake.createProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) CreateProductFileCallCount() int {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return len(fake.createProductFileArgsForCall)
}

func (fake *FakePivnetClient) CreateProductFileArgsForCall(i int) go_pivnet.CreateProductFileConfig {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return fake.createProductFileArgsForCall[i].config
}

func (fake *FakePivnetClient) CreateProductFileReturns(result1 go_pivnet.ProductFile, result2 error) {
	fake.CreateProductFileStub = nil
	fake.createProductFileReturns = struct {
		result1 go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateProductFile(productSlug string, productFile go_pivnet.ProductFile) (go_pivnet.ProductFile, error) {
	fake.updateProductFileMutex.Lock()
	fake.updateProductFileArgsForCall = append(fake.updateProductFileArgsForCall, struct {
		productSlug string
		productFile go_pivnet.ProductFile
	}{productSlug, productFile})
	fake.recordInvocation("UpdateProductFile", []interface{}{productSlug, productFile})
	fake.updateProductFileMutex.Unlock()
	if fake.UpdateProductFileStub != nil {
		return fake.UpdateProductFileStub(productSlug, productFile)
	} else {
		return fake.updateProductFileReturns.result1, fake.updateProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) UpdateProductFileCallCount() int {
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	return len(fake.updateProductFileArgsForCall)
}

func (fake *FakePivnetClient) UpdateProductFileArgsForCall(i int) (string, go_pivnet.ProductFile) {
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	return fake.updateProductFileArgsForCall[i].productSlug, fake.updateProductFileArgsForCall[i].productFile
}

func (fake *FakePivnetClient) UpdateProductFileReturns(result1 go_pivnet.ProductFile, result2 error) {
	fake.UpdateProductFileStub = nil
	fake.updateProductFileReturns = struct {
		result1 go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddProductFileToRelease(productSlug string, releaseID int, productFileID int) error {
	fake.addProductFileToReleaseMutex.Lock()
	fake.addProductFileToReleaseArgsForCall = append(fake.addProductFileToReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("AddProductFileToRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.addProductFileToReleaseMutex.Unlock()
	if fake.AddProductFileToReleaseStub != nil {
		return fake.AddProductFileToReleaseStub(productSlug, releaseID, productFileID)
	} else {
		return fake.addProductFileToReleaseReturns.result1
	}
}

func (fake *FakePivnetClient) AddProductFileToReleaseCallCount() int {
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	return len(fake.addProductFileToReleaseArgsForCall)
}

func (fake *FakePivnetClient) AddProductFileToReleaseArgsForCall(i int) (string, int, int) {
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	return fake.addProductFileToReleaseArgsForCall[i].productSlug, fake.addProductFileToReleaseArgsForCall[i].releaseID, fake.addProductFileToReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) AddProductFileToReleaseReturns(result1 error) {
	fake.AddProductFileToReleaseStub = nil
	fake.addProductFileToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromRelease(productSlug string, releaseID int, productFileID int) error {
	fake.removeProductFileFromReleaseMutex.Lock()
	fake.removeProductFileFromReleaseArgsForCall = append(fake.removeProductFileFromReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("RemoveProductFileFromRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.removeProductFileFromReleaseMutex.Unlock()
	if fake.RemoveProductFileFromReleaseStub != nil {
		return fake.RemoveProductFileFromReleaseStub(productSlug, releaseID, productFileID)
	} else {
		return fake.removeProductFileFromReleaseReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseCallCount() int {
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	return len(fake.removeProductFileFromReleaseArgsForCall)
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseArgsForCall(i int) (string, int, int) {
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	return fake.removeProductFileFromReleaseArgsForCall[i].productSlug, fake.removeProductFileFromReleaseArgsForCall[i].releaseID, fake.removeProductFileFromReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseReturns(result1 error) {
	fake.RemoveProductFileFromReleaseStub = nil
	fake.removeProductFileFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddProductFileToFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.addProductFileToFileGroupMutex.Lock()
	fake.addProductFileToFileGroupArgsForCall = append(fake.addProductFileToFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("AddProductFileToFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.addProductFileToFileGroupMutex.Unlock()
	if fake.AddProductFileToFileGroupStub != nil {
		return fake.AddProductFileToFileGroupStub(productSlug, fileGroupID, productFileID)
	} else {
		return fake.addProductFileToFileGroupReturns.result1
	}
}

func (fake *FakePivnetClient) AddProductFileToFileGroupCallCount() int {
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	return len(fake.addProductFileToFileGroupArgsForCall)
}

func (fake *FakePivnetClient) AddProductFileToFileGroupArgsForCall(i int) (string, int, int) {
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	return fake.addProductFileToFileGroupArgsForCall[i].productSlug, fake.addProductFileToFileGroupArgsForCall[i].fileGroupID, fake.addProductFileToFileGroupArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) AddProductFileToFileGroupReturns(result1 error) {
	fake.AddProductFileToFileGroupStub = nil
	fake.addProductFileToFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.removeProductFileFromFileGroupMutex.Lock()
	fake.removeProductFileFromFileGroupArgsForCall = append(fake.removeProductFileFromFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("RemoveProductFileFromFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.removeProductFileFromFileGroupMutex.Unlock()
	if fake.RemoveProductFileFromFileGroupStub != nil {
		return fake.RemoveProductFileFromFileGroupStub(productSlug, fileGroupID, productFileID)
	} else {
		return fake.removeProductFileFromFileGroupReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupCallCount() int {
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	return len(fake.removeProductFileFromFileGroupArgsForCall)
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupArgsForCall(i int) (string, int, int) {
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	return fake.removeProductFileFromFileGroupArgsForCall[i].productSlug, fake.removeProductFileFromFileGroupArgsForCall[i].fileGroupID, fake.removeProductFileFromFileGroupArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupReturns(result1 error) {
	fake.RemoveProductFileFromFileGroupStub = nil
	fake.removeProductFileFromFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteProductFile(productSlug string, releaseID int) (go_pivnet.ProductFile, error) {
	fake.deleteProductFileMutex.Lock()
	fake.deleteProductFileArgsForCall = append(fake.deleteProductFileArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("DeleteProductFile", []interface{}{productSlug, releaseID})
	fake.deleteProductFileMutex.Unlock()
	if fake.DeleteProductFileStub != nil {
		return fake.DeleteProductFileStub(productSlug, releaseID)
	} else {
		return fake.deleteProductFileReturns.result1, fake.deleteProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) DeleteProductFileCallCount() int {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return len(fake.deleteProductFileArgsForCall)
}

func (fake *FakePivnetClient) DeleteProductFileArgsForCall(i int) (string, int) {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return fake.deleteProductFileArgsForCall[i].productSlug, fake.deleteProductFileArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) DeleteProductFileReturns(result1 go_pivnet.ProductFile, result2 error) {
	fake.DeleteProductFileStub = nil
	fake.deleteProductFileReturns = struct {
		result1 go_pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	} else {
		return fake.acceptEULAReturns.result1
	}
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DownloadProductFile(location *os.File, productSlug string, releaseID int, productFileID int, progressWriter io.Writer) error {
	fake.downloadProductFileMutex.Lock()
	fake.downloadProductFileArgsForCall = append(fake.downloadProductFileArgsForCall, struct {
		location       *os.File
		productSlug    string
		releaseID      int
		productFileID  int
		progressWriter io.Writer
	}{location, productSlug, releaseID, productFileID, progressWriter})
	fake.recordInvocation("DownloadProductFile", []interface{}{location, productSlug, releaseID, productFileID, progressWriter})
	fake.downloadProductFileMutex.Unlock()
	if fake.DownloadProductFileStub != nil {
		return fake.DownloadProductFileStub(location, productSlug, releaseID, productFileID, progressWriter)
	} else {
		return fake.downloadProductFileReturns.result1
	}
}

func (fake *FakePivnetClient) DownloadProductFileCallCount() int {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return len(fake.downloadProductFileArgsForCall)
}

func (fake *FakePivnetClient) DownloadProductFileArgsForCall(i int) (*os.File, string, int, int, io.Writer) {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return fake.downloadProductFileArgsForCall[i].location, fake.downloadProductFileArgsForCall[i].productSlug, fake.downloadProductFileArgsForCall[i].releaseID, fake.downloadProductFileArgsForCall[i].productFileID, fake.downloadProductFileArgsForCall[i].progressWriter
}

func (fake *FakePivnetClient) DownloadProductFileReturns(result1 error) {
	fake.DownloadProductFileStub = nil
	fake.downloadProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ productfile.PivnetClient = new(FakePivnetClient)
