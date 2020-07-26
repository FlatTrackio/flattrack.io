package test_test

import (
	"net/http"
	"log"
	"encoding/json"
	"bytes"
	"os"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/flattrack/flattrack.io/src/backend/common"
	"gitlab.com/flattrack/flattrack.io/src/backend/interested"
	"gitlab.com/flattrack/flattrack.io/src/backend/migrations"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
	"gitlab.com/flattrack/flattrack.io/src/backend/database"
)

var _ = Describe("API interested tests", func() {
	cwd, _ := os.Getwd()
	os.Setenv("APP_DB_MIGRATIONS_PATH", fmt.Sprintf("%v/../../../migrations", cwd))

	db, err := database.DB(
		common.GetDBusername(),
		common.GetDBpassword(),
		common.GetDBhost(),
		common.GetDBdatabase(),
		common.GetDBsslmode(),
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	BeforeSuite(func() {
		err = migrations.Reset(db)
		Expect(err).To(BeNil(), "failed to reset migrations")
		err = migrations.Migrate(db)
		Expect(err).To(BeNil(), "failed to migrate")
	})

	AfterSuite(func() {
		err = migrations.Reset(db)
		Expect(err).To(BeNil(), "failed to reset migrations")
		err = migrations.Migrate(db)
		Expect(err).To(BeNil(), "failed to migrate")
	})

	AfterEach(func() {
		interested.ResetAllEntries(db)
	})

        It("should allow valid email addresses", func() {
                By("preparing emails")

		emails := []types.InterestedSpec{
			{
				Email: "email@example.com",
			},
			{
				Email: "email.1@example.com",
			},
			{
				Email: "email.1@example.coop",
			},
			{
				Email: "emailaaaaa@exampla.e.coop",
			},
		}

		By("sending emails individually")
		for _, email := range emails {
			emailBytes, err := json.Marshal(email)
			Expect(err).To(BeNil(), "failed to marshal to JSON")

			resp, err := httpRequestWithHeader("POST", "http://localhost:8080/api/interested", emailBytes, "")
			Expect(err).To(BeNil(), "API request should not return error")
			Expect(resp.StatusCode).To(Equal(200), "Server MUST response with a 200 value status code")
		}
        })

        It("should not allow non email address strings", func() {
                By("preparing strings")
		emails := []types.InterestedSpec{
			{
				Email: "userexample.com",
			},
			{
				Email: "userexamplecom",
			},
			{
				Email: "",
			},
		}

		By("sending strings individually")
		for _, email := range emails {
			emailBytes, err := json.Marshal(email)
			Expect(err).To(BeNil(), "failed to marshal to JSON")

			resp, err := httpRequestWithHeader("POST", "http://localhost:8080/api/interested", emailBytes, "")
			Expect(err).To(BeNil(), "API request should not return error")
			Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")
		}

        })

        It("should not allow email address strings with length above 70 characters", func() {
		email := types.InterestedSpec{
			Email: "lqDhxqzxymVTmmsZxFUaIMEqYDfOQmkhY5D8TDG6qrwZpLhAKaVU4Wbb5GTLKSMt8nE4AuHU@example.com",
		}
		emailBytes, err := json.Marshal(email)
		Expect(err).To(BeNil(), "failed to marshal to JSON")

		resp, err := httpRequestWithHeader("POST", "http://localhost:8080/api/interested", emailBytes, "")
		Expect(err).To(BeNil(), "API request should not return error")
		Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")
        })

        It("should not allow empty strings", func() {
		email := types.InterestedSpec{
			Email: "",
		}
		emailBytes, err := json.Marshal(email)
		Expect(err).To(BeNil(), "failed to marshal to JSON")

		resp, err := httpRequestWithHeader("POST", "http://localhost:8080/api/interested", emailBytes, "")
		Expect(err).To(BeNil(), "API request should not return error")
		Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")
        })
})

func httpRequestWithHeader(verb string, url string, data []byte, jwt string) (resp *http.Response, err error) {
	req, err := http.NewRequest(verb, url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err = client.Do(req)
	return resp, err
}
