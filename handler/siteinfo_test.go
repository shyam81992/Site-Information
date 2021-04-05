package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/shyam81992/Site-Information/controllers"
	scrap "github.com/shyam81992/Site-Information/scrapper"
	scrapmock "github.com/shyam81992/Site-Information/scrapper/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCity(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockScrapper := scrapmock.NewMockIScrapper(mockCtrl)
	//mockRow := db.NewMockIRow(mockCtrl)

	siteInfoController := &controllers.SiteInfoCtl{Scrap: mockScrapper}

	if os.Getenv("INTEGRATION_TESTING") == "true" {

		siteInfoController = &controllers.SiteInfoCtl{Scrap: &scrap.Scrapper{}}
	}

	NewHandler(&Config{
		R:                  router,
		SiteInfoController: siteInfoController,
	})

	var testCases = []struct {
		name          string
		input         gin.H
		buildStubs    func(mockDb *scrapmock.MockIScrapper)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Get Page Info Success",
			input: gin.H{
				"url": "https://www.google.com",
			},
			buildStubs: func(mockScrapper *scrapmock.MockIScrapper) {
				mockScrapper.EXPECT().GetPageInfo(gomock.Any()).Return(gin.H{
					"Page Title":         "pagetitle",
					"Headings":           gin.H{"h1": 10},
					"Internal Links":     10,
					"External Links":     10,
					"Inaccessible Links": 10,
					"loginform":          true,
				})

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			//.After(callfirst)
			// create a request body with invalid fields
			reqBody, err := json.Marshal(test.input)
			assert.NoError(t, err)

			if os.Getenv("INTEGRATION_TESTING") != "true" {
				test.buildStubs(mockScrapper)
			}

			request, err := http.NewRequest(http.MethodPost, "/static/siteinfo", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)

			test.checkResponse(t, rr)

		})
	}

}
