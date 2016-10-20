package releaseupgradepath_test

import (
	"bytes"
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands/releaseupgradepath"
	"github.com/pivotal-cf/pivnet-cli/commands/releaseupgradepath/releaseupgradepathfakes"
	"github.com/pivotal-cf/pivnet-cli/errorhandler/errorhandlerfakes"
	"github.com/pivotal-cf/pivnet-cli/printer"
)

var _ = Describe("releaseupgradepath commands", func() {
	var (
		fakePivnetClient *releaseupgradepathfakes.FakePivnetClient
		fakeFilter       *releaseupgradepathfakes.FakeFilter

		fakeErrorHandler *errorhandlerfakes.FakeErrorHandler

		outBuffer bytes.Buffer

		existingReleases []pivnet.Release
		filteredReleases []pivnet.Release

		releaseUpgradePaths []pivnet.ReleaseUpgradePath

		client *releaseupgradepath.ReleaseUpgradePathClient
	)

	BeforeEach(func() {
		fakePivnetClient = &releaseupgradepathfakes.FakePivnetClient{}
		fakeFilter = &releaseupgradepathfakes.FakeFilter{}

		outBuffer = bytes.Buffer{}

		fakeErrorHandler = &errorhandlerfakes.FakeErrorHandler{}

		releaseUpgradePaths = []pivnet.ReleaseUpgradePath{
			{
				Release: pivnet.UpgradePathRelease{
					ID: 1234,
				},
			},
			{
				Release: pivnet.UpgradePathRelease{
					ID: 2345,
				},
			},
		}

		existingReleases = []pivnet.Release{
			{
				ID:      1234,
				Version: "1.2.3",
			},
			{
				ID:      1235,
				Version: "1.2.4",
			},
			{
				ID:      2345,
				Version: "2.3.4",
			},
		}

		filteredReleases = []pivnet.Release{
			existingReleases[0],
			existingReleases[1],
		}

		fakePivnetClient.ReleasesForProductSlugReturns(existingReleases, nil)
		fakePivnetClient.ReleaseForVersionReturns(existingReleases[2], nil)
		fakePivnetClient.ReleaseUpgradePathsReturns(releaseUpgradePaths, nil)
		fakeFilter.ReleasesByVersionReturns(filteredReleases, nil)

		client = releaseupgradepath.NewReleaseUpgradePathClient(
			fakePivnetClient,
			fakeErrorHandler,
			printer.PrintAsJSON,
			&outBuffer,
			printer.NewPrinter(&outBuffer),
			fakeFilter,
		)
	})

	Describe("ReleaseUpgradePaths", func() {
		var (
			productSlug    string
			releaseVersion string
		)

		BeforeEach(func() {
			productSlug = "some product slug"
			releaseVersion = "some release version"
		})

		It("lists all ReleaseUpgradePaths", func() {
			err := client.List(productSlug, releaseVersion)
			Expect(err).NotTo(HaveOccurred())

			var returnedReleaseUpgradePaths []pivnet.ReleaseUpgradePath
			err = json.Unmarshal(outBuffer.Bytes(), &returnedReleaseUpgradePaths)
			Expect(err).NotTo(HaveOccurred())

			Expect(returnedReleaseUpgradePaths).To(Equal(releaseUpgradePaths))
		})

		Context("when there is an error", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releaseUpgradePaths error")
				fakePivnetClient.ReleaseUpgradePathsReturns(nil, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.List(productSlug, releaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error getting release", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releases error")
				fakePivnetClient.ReleaseForVersionReturns(pivnet.Release{}, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.List(productSlug, releaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})
	})

	Describe("AddReleaseUpgradePath", func() {
		var (
			productSlug            string
			releaseVersion         string
			previousReleaseVersion string
		)

		BeforeEach(func() {
			productSlug = "some product slug"
			releaseVersion = "some release version"
			previousReleaseVersion = "previous release version"
		})

		It("adds ReleaseUpgradePath", func() {
			err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
			Expect(err).NotTo(HaveOccurred())

			invokedReleases, invokedFilterVersion := fakeFilter.ReleasesByVersionArgsForCall(0)
			Expect(invokedReleases).To(Equal(existingReleases))
			Expect(invokedFilterVersion).To(Equal(previousReleaseVersion))

			Expect(fakePivnetClient.AddReleaseUpgradePathCallCount()).To(Equal(2))

			invokedProductSlug0, invokedReleaseID0, invokedPreviousReleaseID0 :=
				fakePivnetClient.AddReleaseUpgradePathArgsForCall(0)

			Expect(invokedProductSlug0).To(Equal(productSlug))
			Expect(invokedReleaseID0).To(Equal(2345))
			Expect(invokedPreviousReleaseID0).To(Equal(existingReleases[0].ID))

			invokedProductSlug1, invokedReleaseID1, invokedPreviousReleaseID1 :=
				fakePivnetClient.AddReleaseUpgradePathArgsForCall(1)

			Expect(invokedProductSlug1).To(Equal(productSlug))
			Expect(invokedReleaseID1).To(Equal(2345))
			Expect(invokedPreviousReleaseID1).To(Equal(existingReleases[1].ID))
		})

		Context("when provided version does not match", func() {
			BeforeEach(func() {
				fakeFilter.ReleasesByVersionReturns(nil, nil)
			})

			It("invokes the error handler", func() {
				err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
			})
		})

		Context("when there is an error", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releaseUpgradePaths error")
				fakePivnetClient.AddReleaseUpgradePathReturns(expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error getting release", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releases error")
				fakePivnetClient.ReleaseForVersionReturns(pivnet.Release{}, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error getting previous release", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releases error")
				fakePivnetClient.ReleasesForProductSlugReturns(nil, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error filtering releases", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("filter error")
				fakeFilter.ReleasesByVersionReturns(nil, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Add(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})
	})

	Describe("RemoveReleaseUpgradePath", func() {
		var (
			productSlug            string
			releaseVersion         string
			previousReleaseVersion string
		)

		BeforeEach(func() {
			productSlug = "some product slug"
			releaseVersion = "some release version"
			previousReleaseVersion = "previous release version"
		})

		It("removes ReleaseUpgradePath", func() {
			err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
			Expect(err).NotTo(HaveOccurred())

			invokedReleases, invokedFilterVersion := fakeFilter.ReleasesByVersionArgsForCall(0)
			Expect(invokedReleases).To(Equal(existingReleases))
			Expect(invokedFilterVersion).To(Equal(previousReleaseVersion))

			Expect(fakePivnetClient.RemoveReleaseUpgradePathCallCount()).To(Equal(2))

			invokedProductSlug0, invokedReleaseID0, invokedPreviousReleaseID0 :=
				fakePivnetClient.RemoveReleaseUpgradePathArgsForCall(0)

			Expect(invokedProductSlug0).To(Equal(productSlug))
			Expect(invokedReleaseID0).To(Equal(2345))
			Expect(invokedPreviousReleaseID0).To(Equal(existingReleases[0].ID))

			invokedProductSlug1, invokedReleaseID1, invokedPreviousReleaseID1 :=
				fakePivnetClient.RemoveReleaseUpgradePathArgsForCall(1)

			Expect(invokedProductSlug1).To(Equal(productSlug))
			Expect(invokedReleaseID1).To(Equal(2345))
			Expect(invokedPreviousReleaseID1).To(Equal(existingReleases[1].ID))
		})

		Context("when provided version does not match", func() {
			BeforeEach(func() {
				fakeFilter.ReleasesByVersionReturns(nil, nil)
			})

			It("invokes the error handler", func() {
				err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
			})
		})

		Context("when there is an error", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releaseUpgradePaths error")
				fakePivnetClient.RemoveReleaseUpgradePathReturns(expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error getting release", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releases error")
				fakePivnetClient.ReleaseForVersionReturns(pivnet.Release{}, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error getting previous release", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("releases error")
				fakePivnetClient.ReleasesForProductSlugReturns(nil, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})

		Context("when there is an error filtering releases", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("filter error")
				fakeFilter.ReleasesByVersionReturns(nil, expectedErr)
			})

			It("invokes the error handler", func() {
				err := client.Remove(productSlug, releaseVersion, previousReleaseVersion)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeErrorHandler.HandleErrorCallCount()).To(Equal(1))
				Expect(fakeErrorHandler.HandleErrorArgsForCall(0)).To(Equal(expectedErr))
			})
		})
	})
})
