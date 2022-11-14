package cloudcmd

import (
	"context"
	"errors"
	"github.com/TykTechnologies/cloud-sdk/cloud"
	mock "github.com/TykTechnologies/tykctl/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestGetOrg(t *testing.T) {
	testCases := []struct {
		name              string
		mockResponse      cloud.InlineResponse20014
		mockHttpResponse  *http.Response
		mockError         error
		ExpectedError     error
		ExpectedOrgLength int
	}{
		{
			name: "Check status Ok",
			mockResponse: cloud.InlineResponse20014{
				Error_:  "",
				Payload: &cloud.Organisations{Organisations: generateOrgs(4)},
				Status:  "ok",
			},
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedError:     nil,
			mockError:         nil,
			ExpectedOrgLength: 4,
		},
		{
			name: "Check status code not 200",
			mockResponse: cloud.InlineResponse20014{
				Status:  "",
				Error_:  "error is here",
				Payload: &cloud.Organisations{Organisations: nil},
			},
			mockError:         nil,
			mockHttpResponse:  &http.Response{StatusCode: http.StatusForbidden},
			ExpectedError:     ErrorFetchingOrg,
			ExpectedOrgLength: 0,
		},
		{
			name: "Test when cloud returns an error",
			mockResponse: cloud.InlineResponse20014{
				Payload: nil,
				Status:  "ok",
				Error_:  "",
			},
			mockError:         ErrorGenericError,
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedError:     ErrorGenericError,
			ExpectedOrgLength: 0,
		},
		{
			name: "Test nil payload",
			mockResponse: cloud.InlineResponse20014{
				Payload: nil,
				Status:  "ok",
				Error_:  "",
			},
			mockError:         nil,
			mockHttpResponse:  &http.Response{StatusCode: http.StatusOK},
			ExpectedError:     nil,
			ExpectedOrgLength: 0,
		},
		{
			name:              "Test payload status with error",
			ExpectedError:     errors.New("there is an error here"),
			ExpectedOrgLength: 0,
			mockResponse: cloud.InlineResponse20014{
				Payload: nil,
				Status:  "error",
				Error_:  "there is an error here",
			},
			mockError:        nil,
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetOrgs(gomock.Any()).Times(1).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			orgs, err := GetOrgs(context.Background(), m)
			assert.Equal(t, tt.ExpectedError, err)
			if tt.mockResponse.Payload != nil {
				assert.Equal(t, tt.mockResponse.Payload.Organisations, orgs)
			}
			assert.Equal(t, tt.ExpectedOrgLength, len(orgs))
		})
	}

}
func TestGetOrgById(t *testing.T) {
	testCases := []struct {
		name             string
		mockHttpResponse *http.Response
		mockResponse     cloud.InlineResponse2005
		ExpectedError    error
		mockError        error
	}{
		{
			name:             "Check fetch org success",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: cloud.InlineResponse2005{
				Error_:  "",
				Payload: &generateOrgs(1)[0],
				Status:  "ok",
			},
			ExpectedError: nil,
			mockError:     nil,
		},
		{
			name:             "Test Status code 403 ",
			mockHttpResponse: &http.Response{StatusCode: http.StatusForbidden},
			mockResponse: cloud.InlineResponse2005{
				Error_:  "I have an error here",
				Payload: nil,
				Status:  "error",
			},
			ExpectedError: ErrorFetchingOrg,
			mockError:     nil,
		},
		{
			name:             "Test response body status is not sucess",
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockResponse: cloud.InlineResponse2005{
				Error_:  "i have an error here",
				Payload: nil,
				Status:  "error",
			},
			ExpectedError: errors.New("i have an error here"),
			mockError:     nil,
		},
		{
			name:             "Check when cloud returns an error",
			mockResponse:     cloud.InlineResponse2005{},
			mockHttpResponse: &http.Response{StatusCode: http.StatusOK},
			mockError:        ErrorOutPutFormat,
			ExpectedError:    ErrorOutPutFormat,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			m.EXPECT().GetOrgById(gomock.Any(), gomock.Any()).Return(tt.mockResponse, tt.mockHttpResponse, tt.mockError)
			organization, err := GetOrgById(context.Background(), m, "oid")
			assert.Equal(t, organization, tt.mockResponse.Payload)
			assert.Equal(t, tt.ExpectedError, err)
		})

	}
}
func TestGetEntitlements(t *testing.T) {
	type args struct {
		counter map[string]cloud.CounterEntitlement
		key     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test key is present",
			args: args{
				counter: map[string]cloud.CounterEntitlement{
					"present": {
						Allowed:  10,
						Consumed: 2,
						Name:     "",
					},
				},
				key: "present",
			},
			want: "2 of 10",
		},
		{
			name: "Test key is absent",
			args: args{
				counter: map[string]cloud.CounterEntitlement{
					"absent": {
						Allowed:  10,
						Consumed: 2,
						Name:     "",
					},
				},
				key: "absen",
			},
			want: "- of -",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getEntitlements(tt.args.counter, tt.args.key), "getEntitlements(%v, %v)", tt.args.counter, tt.args.key)
		})
	}
}

func TestCreateOrgHeaderAndRows(t *testing.T) {
	type args struct {
		organizations []cloud.Organisation
	}
	tests := []struct {
		name    string
		args    args
		headers []string
		rows    [][]string
	}{
		{
			name: "Test Create Org Row Success",
			args: args{
				organizations: []cloud.Organisation{{
					Name: "test organisation-8e392d",
					UID:  "2c05d09d-f8cc-4333-9a3a-4531298e392d",
					Entitlements: &cloud.Entitlements{
						Counters: map[string]cloud.CounterEntitlement{
							"MaxAnalytics": {
								Allowed:  100,
								Consumed: 45,
								Name:     "",
							},
							"MaxDashboardCount": {
								Allowed:  29,
								Consumed: 13,
								Name:     "",
							},
							"MaxGatewayCount": {
								Allowed:  7,
								Consumed: 4,
								Name:     "",
							},
							"MaxLoadoutCount": {
								Allowed:  51,
								Consumed: 27,
								Name:     "",
							},
							"MaxTeamCount": {
								Allowed:  200,
								Consumed: 47,
								Name:     "",
							},
						},
					},
				}},
			},
			headers: []string{"Name", "ID", "Teams", "Environments", "Control planes", "Edge"},
			rows:    [][]string{{"test organisation-8e392d", "2c05d09d-f8cc-4333-9a3a-4531298e392d", "47 of 200", "27 of 51", "13 of 29", "4 of 7"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnedHeaders, returnedRows := CreateOrgHeaderAndRows(tt.args.organizations)
			assert.Equalf(t, tt.headers, returnedHeaders, "CreateOrgHeaderAndRows(%v)", tt.args.organizations)
			assert.Equalf(t, tt.rows, returnedRows, "CreateOrgHeaderAndRows(%v)", tt.args.organizations)
		})
	}
}

func generateOrgs(size int) []cloud.Organisation {
	var organizations []cloud.Organisation
	for i := 0; i < size; i++ {
		organizations = append(organizations, cloud.Organisation{
			AccountID: strconv.Itoa(i),
			UID:       strconv.Itoa(i),
		})
	}

	return organizations

}

func TestFetchAndPrintOrgById(t *testing.T) {
	type args struct {
		output        string
		ExpectedError error
	}
	tests := []struct {
		args args
		name string
	}{
		{
			name: "Test correct format is passed",
			args: args{
				output:        "yml",
				ExpectedError: ErrorOutPutFormat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			err := FetchAndPrintOrgById(context.Background(), m, tt.args.output, "test")
			assert.Equal(t, tt.args.ExpectedError, err)
		})
	}
}
func TestFetchAndPrintOrganizations(t *testing.T) {
	type args struct {
		ExpectedError error
		output        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test correct format is passed",
			args: args{
				ExpectedError: ErrorOutPutFormat,
				output:        "yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock.NewMockCloudClient(ctrl)
			err := FetchAndPrintOrganizations(context.Background(), m, tt.args.output)
			assert.Equal(t, tt.args.ExpectedError, err)

		})
	}
}
