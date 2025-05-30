package certificate

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/oracle/oci-go-sdk/v65/certificates"
	"github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
	"github.com/oracle/oci-go-sdk/v65/common"
	ociclient "github.com/oracle/oci-native-ingress-controller/pkg/oci/client"
)

const (
	ErrorListingCaBundle = "error listing Ca Bundles"
	errorMsg             = "no cert found"
	namespace            = "test"
	errorImportCert      = "errorImportCert"
)

func setup() *CertificatesClient {
	certClient := GetCertClient()
	certManageClient := GetCertManageClient()
	return New(certManageClient, certClient)
}

func TestNew(t *testing.T) {
	RegisterTestingT(t)
	client := setup()
	Expect(client).Should(Not(BeNil()))
}

func TestCertificatesClient_Cache(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.CreateCertificateRequest{
		CreateCertificateDetails: certificatesmanagement.CreateCertificateDetails{
			Name:          common.String("certificate-name"),
			CompartmentId: common.String("compartment-id"),
		},
		OpcRequestId:    nil,
		OpcRetryToken:   nil,
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.CreateCertificate(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

}

func TestCertificatesClient_CreateCertificate(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.CreateCertificateRequest{
		CreateCertificateDetails: certificatesmanagement.CreateCertificateDetails{
			Name:          common.String("certificate-name"),
			CompartmentId: common.String("compartment-id"),
		},
		OpcRequestId:    nil,
		OpcRetryToken:   nil,
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.CreateCertificate(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

}

func TestCertificatesClient_CreateCaBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.CreateCaBundleRequest{
		CreateCaBundleDetails: certificatesmanagement.CreateCaBundleDetails{
			Name:          common.String("cabundle-name"),
			CompartmentId: common.String("compartment-id"),
		},
		OpcRequestId:    nil,
		OpcRetryToken:   nil,
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.CreateCaBundle(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

}
func TestCertificatesClient_GetCertificate(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.GetCertificateRequest{
		CertificateId:   common.String("id"),
		OpcRequestId:    nil,
		RequestMetadata: common.RequestMetadata{},
	}
	cert, etag, err := client.GetCertificate(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(etag).Should(Equal("etag"))
	Expect(cert).Should(Not(BeNil()))

}
func TestCertificatesClient_ListCertificates(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.ListCertificatesRequest{
		CertificateId:   common.String("id"),
		OpcRequestId:    nil,
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.ListCertificates(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

}

func TestCertificatesClient_UpdateCertificate(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.UpdateCertificateRequest{
		CertificateId:   common.String("id"),
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.UpdateCertificate(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

	request.CertificateId = common.String("error")
	cert, _, err = client.UpdateCertificate(context.TODO(), request)
	Expect(err).ShouldNot(BeNil())
	Expect(cert).Should(BeNil())
}

func TestCertificatesClient_ListCertificateVersions(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.ListCertificateVersionsRequest{
		CertificateId: common.String("id"),
		SortOrder:     certificatesmanagement.ListCertificateVersionsSortOrderDesc,
	}
	certVersionSummary, _, err := client.ListCertificateVersions(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(certVersionSummary).ShouldNot(BeNil())
}

func TestCertificatesClient_ScheduleCertificateVersionDeletion(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.ScheduleCertificateVersionDeletionRequest{
		CertificateId:            common.String("id"),
		CertificateVersionNumber: common.Int64(3),
	}
	cert, _, err := client.ScheduleCertificateVersionDeletion(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).ShouldNot(BeNil())
}

func TestCertificatesClient_waitForActiveCertificate(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	certificateId := "id"
	cert, _, err := client.waitForActiveCertificate(context.TODO(), certificateId)
	Expect(err).Should(BeNil())
	Expect(cert).ShouldNot(BeNil())
}

func TestCertificatesClient_GetCaBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.GetCaBundleRequest{
		CaBundleId:      common.String("id"),
		OpcRequestId:    nil,
		RequestMetadata: common.RequestMetadata{},
	}
	caBundle, etag, err := client.GetCaBundle(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(etag).Should(Equal("etag"))
	Expect(caBundle).Should(Not(BeNil()))

}
func TestCertificatesClient_GetCertificateBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificates.GetCertificateBundleRequest{
		CertificateId:          common.String("id"),
		OpcRequestId:           nil,
		VersionNumber:          nil,
		CertificateVersionName: nil,
		Stage:                  "",
		CertificateBundleType:  "",
		RequestMetadata:        common.RequestMetadata{},
	}
	caBundle, err := client.GetCertificateBundle(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(caBundle).Should(Not(BeNil()))

}

func TestCertificatesClient_ListCaBundles(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.ListCaBundlesRequest{
		Name:           common.String("name"),
		CompartmentId:  common.String("compartmentId"),
		LifecycleState: certificatesmanagement.ListCaBundlesLifecycleStateActive,
	}
	caBundle, _, err := client.ListCaBundles(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(caBundle).Should(Not(BeNil()))

	request = certificatesmanagement.ListCaBundlesRequest{
		Name:           common.String("name"),
		CompartmentId:  common.String("compartmentId"),
		LifecycleState: certificatesmanagement.ListCaBundlesLifecycleStateDeleted,
	}
	caBundle, _, err = client.ListCaBundles(context.TODO(), request)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(Equal(ErrorListingCaBundle))

}

func TestCertificatesClient_UpdateCaBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	request := certificatesmanagement.UpdateCaBundleRequest{
		CaBundleId:      common.String("id"),
		RequestMetadata: common.RequestMetadata{},
	}
	cert, _, err := client.UpdateCaBundle(context.TODO(), request)
	Expect(err).Should(BeNil())
	Expect(cert).Should(Not(BeNil()))

	request.CaBundleId = common.String("error")
	cert, _, err = client.UpdateCaBundle(context.TODO(), request)
	Expect(err).ShouldNot(BeNil())
	Expect(cert).Should(BeNil())
}

func TestScheduleCertificateDeletion(t *testing.T) {
	RegisterTestingT(t)
	client := setup()
	id := "id"
	request := certificatesmanagement.ScheduleCertificateDeletionRequest{
		CertificateId: &id,
	}
	err := client.ScheduleCertificateDeletion(context.TODO(), request)
	Expect(err).Should(BeNil())

	id = "error"
	request = certificatesmanagement.ScheduleCertificateDeletionRequest{
		CertificateId:                      &id,
		ScheduleCertificateDeletionDetails: certificatesmanagement.ScheduleCertificateDeletionDetails{},
		OpcRequestId:                       nil,
		IfMatch:                            nil,
		RequestMetadata:                    common.RequestMetadata{},
	}
	err = client.ScheduleCertificateDeletion(context.TODO(), request)
	Expect(err).Should(Not(BeNil()))
}

func TestDeleteCaBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()
	id := "id"
	request := getDeleteCaBundleRequest(id)
	res, err := client.DeleteCaBundle(context.TODO(), request)

	Expect(err).Should(BeNil())
	Expect(res.Status).Should(Equal("200"))

	request = getDeleteCaBundleRequest("error")
	res, err = client.DeleteCaBundle(context.TODO(), request)
	Expect(err).Should(Not(BeNil()))
}

func TestCertificatesClient_ListAssociations(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	listAssociationsRequest := certificatesmanagement.ListAssociationsRequest{
		CompartmentId: common.String("compartmentId"),
	}
	_, err := client.ListAssociations(context.Background(), listAssociationsRequest)
	Expect(err).To(BeNil())

	listAssociationsRequest.CompartmentId = common.String("error")
	_, err = client.ListAssociations(context.Background(), listAssociationsRequest)
	Expect(err).ToNot(BeNil())
}

func TestCertificatesClient_waitForActiveCaBundle(t *testing.T) {
	RegisterTestingT(t)
	client := setup()

	certificateId := "id"
	caBundle, _, err := client.waitForActiveCaBundle(context.TODO(), certificateId)
	Expect(err).Should(BeNil())
	Expect(caBundle).ShouldNot(BeNil())
}

func getDeleteCaBundleRequest(id string) certificatesmanagement.DeleteCaBundleRequest {
	request := certificatesmanagement.DeleteCaBundleRequest{
		CaBundleId:      &id,
		OpcRequestId:    &id,
		IfMatch:         nil,
		RequestMetadata: common.RequestMetadata{},
	}
	return request
}

func GetCertManageClient() ociclient.CertificateManagementInterface {
	return &MockCertificateManagerClient{}
}

type MockCertificateManagerClient struct {
}

func (m MockCertificateManagerClient) CreateCertificate(ctx context.Context, request certificatesmanagement.CreateCertificateRequest) (certificatesmanagement.CreateCertificateResponse, error) {
	id := "id"
	etag := "etag"
	return certificatesmanagement.CreateCertificateResponse{
		RawResponse: nil,
		Certificate: certificatesmanagement.Certificate{
			Id: &id,
		},
		Etag:         &etag,
		OpcRequestId: &id,
	}, nil
}

func (m MockCertificateManagerClient) GetCertificate(ctx context.Context, request certificatesmanagement.GetCertificateRequest) (certificatesmanagement.GetCertificateResponse, error) {

	if *request.CertificateId == "error" {
		return certificatesmanagement.GetCertificateResponse{}, errors.New(errorMsg)
	}
	id := "id"
	name := "cert"
	authorityId := "authId"
	etag := "etag"
	var confType certificatesmanagement.CertificateConfigTypeEnum
	if *request.CertificateId == errorImportCert {
		name = "error"
		confType = certificatesmanagement.CertificateConfigTypeImported
	} else {
		confType, _ = certificatesmanagement.GetMappingCertificateConfigTypeEnum(*request.CertificateId)
	}
	var number int64
	number = 234
	certVersionSummary := certificatesmanagement.CertificateVersionSummary{
		VersionNumber: &number,
	}
	return certificatesmanagement.GetCertificateResponse{
		RawResponse: nil,
		Certificate: certificatesmanagement.Certificate{
			Id:                           &id,
			Name:                         &name,
			ConfigType:                   confType,
			IssuerCertificateAuthorityId: &authorityId,
			CurrentVersion:               &certVersionSummary,
			LifecycleState:               certificatesmanagement.CertificateLifecycleStateActive,
		},
		Etag:         &etag,
		OpcRequestId: nil,
	}, nil
}

func (m MockCertificateManagerClient) ListCertificates(ctx context.Context, request certificatesmanagement.ListCertificatesRequest) (certificatesmanagement.ListCertificatesResponse, error) {
	id := "id"
	return certificatesmanagement.ListCertificatesResponse{
		RawResponse:           nil,
		CertificateCollection: certificatesmanagement.CertificateCollection{},
		OpcRequestId:          &id,
		OpcNextPage:           &id,
	}, nil
}

func (m MockCertificateManagerClient) UpdateCertificate(ctx context.Context, request certificatesmanagement.UpdateCertificateRequest) (certificatesmanagement.UpdateCertificateResponse, error) {
	if *request.CertificateId == "error" {
		return certificatesmanagement.UpdateCertificateResponse{}, errors.New("cannot find certificate")
	}
	return certificatesmanagement.UpdateCertificateResponse{}, nil
}

func (m MockCertificateManagerClient) ScheduleCertificateDeletion(ctx context.Context, request certificatesmanagement.ScheduleCertificateDeletionRequest) (certificatesmanagement.ScheduleCertificateDeletionResponse, error) {
	var err error
	if *request.CertificateId == "error" {
		err = errors.New("cert error deletion")
	}
	return certificatesmanagement.ScheduleCertificateDeletionResponse{}, err
}

func (m MockCertificateManagerClient) ListCertificateVersions(ctx context.Context, request certificatesmanagement.ListCertificateVersionsRequest) (certificatesmanagement.ListCertificateVersionsResponse, error) {
	return certificatesmanagement.ListCertificateVersionsResponse{}, nil
}

func (m MockCertificateManagerClient) ScheduleCertificateVersionDeletion(ctx context.Context, request certificatesmanagement.ScheduleCertificateVersionDeletionRequest) (certificatesmanagement.ScheduleCertificateVersionDeletionResponse, error) {
	return certificatesmanagement.ScheduleCertificateVersionDeletionResponse{}, nil
}

func (m MockCertificateManagerClient) CreateCaBundle(ctx context.Context, request certificatesmanagement.CreateCaBundleRequest) (certificatesmanagement.CreateCaBundleResponse, error) {
	id := "id"
	etag := "etag"
	return certificatesmanagement.CreateCaBundleResponse{
		RawResponse: nil,
		CaBundle: certificatesmanagement.CaBundle{
			Id: &id,
		},
		Etag:         &etag,
		OpcRequestId: nil,
	}, nil
}

func (m MockCertificateManagerClient) GetCaBundle(ctx context.Context, request certificatesmanagement.GetCaBundleRequest) (certificatesmanagement.GetCaBundleResponse, error) {
	id := "id"
	name := "cabundle"
	etag := "etag"
	return certificatesmanagement.GetCaBundleResponse{
		RawResponse: nil,
		CaBundle: certificatesmanagement.CaBundle{
			Id:             &id,
			Name:           &name,
			LifecycleState: certificatesmanagement.CaBundleLifecycleStateActive,
		},
		OpcRequestId: &id,
		Etag:         &etag,
	}, nil
}

func (m MockCertificateManagerClient) ListCaBundles(ctx context.Context, request certificatesmanagement.ListCaBundlesRequest) (certificatesmanagement.ListCaBundlesResponse, error) {

	if request.LifecycleState == certificatesmanagement.ListCaBundlesLifecycleStateDeleted {
		err := errors.New(ErrorListingCaBundle)
		return certificatesmanagement.ListCaBundlesResponse{}, err
	}

	if *request.Name == "error" {
		return certificatesmanagement.ListCaBundlesResponse{}, nil
	}

	var items []certificatesmanagement.CaBundleSummary
	name := "ic-oci-config"
	id := "id"
	item := certificatesmanagement.CaBundleSummary{
		Id:   &id,
		Name: &name,
	}
	items = append(items, item)

	return certificatesmanagement.ListCaBundlesResponse{
		RawResponse: nil,
		CaBundleCollection: certificatesmanagement.CaBundleCollection{
			Items: items,
		},
		OpcRequestId: nil,
		OpcNextPage:  nil,
	}, nil
}

func (m MockCertificateManagerClient) UpdateCaBundle(ctx context.Context, request certificatesmanagement.UpdateCaBundleRequest) (certificatesmanagement.UpdateCaBundleResponse, error) {
	if *request.CaBundleId == "error" {
		return certificatesmanagement.UpdateCaBundleResponse{}, errors.New("cannot find ca bundle")
	}
	return certificatesmanagement.UpdateCaBundleResponse{}, nil
}

func (m MockCertificateManagerClient) DeleteCaBundle(ctx context.Context, request certificatesmanagement.DeleteCaBundleRequest) (certificatesmanagement.DeleteCaBundleResponse, error) {
	res := http.Response{
		Status: "200",
	}
	var err error
	if *request.CaBundleId == "error" {
		err = errors.New("error deleting cabundle")
	}
	return certificatesmanagement.DeleteCaBundleResponse{
		RawResponse:  &res,
		OpcRequestId: nil,
	}, err
}

func (m MockCertificateManagerClient) ListAssociations(ctx context.Context, request certificatesmanagement.ListAssociationsRequest) (certificatesmanagement.ListAssociationsResponse, error) {
	if *request.CompartmentId == "error" {
		return certificatesmanagement.ListAssociationsResponse{}, errors.New("error listing associations")
	}
	return certificatesmanagement.ListAssociationsResponse{}, nil
}

func GetCertClient() ociclient.CertificateInterface {
	return &MockCertificateClient{}
}

type MockCertificateClient struct {
}

func (m MockCertificateClient) SetCertCache(cert *certificatesmanagement.Certificate) {

}

func (m MockCertificateClient) GetFromCertCache(certId string) *ociclient.CertCacheObj {
	cert := certificatesmanagement.Certificate{}
	var now time.Time
	if certId == "id" {
		now = time.Now()
	} else {
		now = time.Now()
		now.Add(time.Minute * 15)
	}
	return &ociclient.CertCacheObj{
		Cert: &cert,
		Age:  now,
	}
}

func (m MockCertificateClient) SetCaBundleCache(caBundle *certificatesmanagement.CaBundle) {

}

func (m MockCertificateClient) GetFromCaBundleCache(id string) *ociclient.CaBundleCacheObj {
	return nil
}

func (m MockCertificateClient) CreateCertificate(ctx context.Context, req certificatesmanagement.CreateCertificateRequest) (*certificatesmanagement.Certificate, error) {
	return &certificatesmanagement.Certificate{}, nil
}

func (m MockCertificateClient) CreateCaBundle(ctx context.Context, req certificatesmanagement.CreateCaBundleRequest) (*certificatesmanagement.CaBundle, error) {
	return &certificatesmanagement.CaBundle{}, nil
}

func (m MockCertificateClient) GetCertificate(ctx context.Context, req certificatesmanagement.GetCertificateRequest) (*certificatesmanagement.Certificate, error) {
	id := "id"
	return &certificatesmanagement.Certificate{
		Id: &id,
	}, nil
}

func (m MockCertificateClient) ListCertificates(ctx context.Context, req certificatesmanagement.ListCertificatesRequest) (*certificatesmanagement.CertificateCollection, *string, error) {
	return &certificatesmanagement.CertificateCollection{}, nil, nil
}

func (m MockCertificateClient) ScheduleCertificateDeletion(ctx context.Context, req certificatesmanagement.ScheduleCertificateDeletionRequest) error {
	return nil
}

func (m MockCertificateClient) GetCaBundle(ctx context.Context, req certificatesmanagement.GetCaBundleRequest) (*certificatesmanagement.CaBundle, error) {
	return &certificatesmanagement.CaBundle{}, nil
}

func (m MockCertificateClient) ListCaBundles(ctx context.Context, req certificatesmanagement.ListCaBundlesRequest) (*certificatesmanagement.CaBundleCollection, error) {
	return &certificatesmanagement.CaBundleCollection{}, nil
}

func (m MockCertificateClient) DeleteCaBundle(ctx context.Context, req certificatesmanagement.DeleteCaBundleRequest) (*http.Response, error) {
	return &http.Response{}, nil
}

func (m MockCertificateClient) GetCertificateBundle(ctx context.Context, request certificates.GetCertificateBundleRequest) (certificates.GetCertificateBundleResponse, error) {
	return certificates.GetCertificateBundleResponse{
		RawResponse:       nil,
		CertificateBundle: getMockBundle(),
		Etag:              nil,
		OpcRequestId:      nil,
	}, nil
}

func getMockBundle() certificates.CertificateBundle {
	return &MockCertificateBundle{}
}

type MockCertificateBundle struct {
}

func (m MockCertificateBundle) GetCertificateId() *string {
	return nil
}

func (m MockCertificateBundle) GetCertificateName() *string {
	return nil
}

func (m MockCertificateBundle) GetVersionNumber() *int64 {
	return nil
}

func (m MockCertificateBundle) GetSerialNumber() *string {
	return nil
}

func (m MockCertificateBundle) GetTimeCreated() *common.SDKTime {
	return nil
}

func (m MockCertificateBundle) GetValidity() *certificates.Validity {
	return nil
}

func (m MockCertificateBundle) GetStages() []certificates.VersionStageEnum {
	return nil
}

func (m MockCertificateBundle) GetCertificatePem() *string {
	return nil
}

func (m MockCertificateBundle) GetCertChainPem() *string {
	data := "chain"
	return &data
}

func (m MockCertificateBundle) GetVersionName() *string {
	return nil
}

func (m MockCertificateBundle) GetRevocationStatus() *certificates.RevocationStatus {
	return nil
}
