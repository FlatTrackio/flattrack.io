package test_test

import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
        "os"
        "fmt"
        "net/url"
        "net/http"
        "gitlab.com/flattrack/flattrack.io/src/backend/common"
)

var _ = Describe("API interested tests", func() {
        currentWorkingDirectory, _ := os.Getwd()
        deploymentConfigStoreLocation := fmt.Sprintf("%v/%v/%v/%v/%v/%v", currentWorkingDirectory, "..", "..", "..", "deployment", "interested.json")
        BeforeEach(func() {
                common.DeinitJSONstore(deploymentConfigStoreLocation)
                common.InitJSONstore(deploymentConfigStoreLocation)
        })
        AfterEach(func() {
                common.DeinitJSONstore(deploymentConfigStoreLocation)
        })

        It("should allow email addresses to be written", func() {
                By("Not failing the request")

                // set up request
                user1formEmail := "user@example.com"
                user1form := url.Values{}
                user1form.Set("email", user1formEmail)

                // send the request request
                resp, err := http.PostForm("http://localhost:8080/api/interested", user1form)
                Expect(err).To(BeNil(), "API request should not return error")

                By("Finding posted value in deployment")

                // get the deployment interested.json's contents
                emailStore := common.ReadJSONstore(deploymentConfigStoreLocation)
                _, found := common.Find(emailStore.Emails, user1formEmail)
                Expect(found).To(BeTrue(), "Saved deployment file MUST container the just posted email")
                Expect(resp.StatusCode).To(Equal(200), "Server MUST response with a 200 value status code")
        })

        It("should not allow non email address strings", func() {
                user1formEmail := "userexamplecom"
                user1form := url.Values{}
                user1form.Set("email", user1formEmail)

                resp, err := http.PostForm("http://localhost:8080/api/interested", user1form)
                Expect(err).To(BeNil(), "API request should return error")

                By("Not finding posted value in deployment")
                emailStore := common.ReadJSONstore(deploymentConfigStoreLocation)
                _, found := common.Find(emailStore.Emails, user1formEmail)
                Expect(found).To(BeFalse(), "Saved deployment file MUST NOT container the just posted email")
                Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")

        })

        It("should not allow email address strings with length above 70 characters", func() {
                user1formEmail := "lqDhxqzxymVTmmsZxFUaIMEqYDfOQmkhY5D8TDG6qrwZpLhAKaVU4Wbb5GTLKSMt8nE4AuHU@example.com"
                user1form := url.Values{}
                user1form.Set("email", user1formEmail)

                resp, err := http.PostForm("http://localhost:8080/api/interested", user1form)
                Expect(err).To(BeNil(), "API request should return error")

                By("Not finding posted value in deployment")
                emailStore := common.ReadJSONstore(deploymentConfigStoreLocation)
                _, found := common.Find(emailStore.Emails, user1formEmail)
                Expect(found).To(BeFalse(), "Saved deployment file MUST NOT container the just posted email")
                Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")
        })

        It("should not allow empty strings", func() {
                user1formEmail := ""
                user1form := url.Values{}
                user1form.Set("email", user1formEmail)

                resp, err := http.PostForm("http://localhost:8080/api/interested", user1form)
                Expect(err).To(BeNil(), "API request should return error")

                By("Not finding posted value in deployment")
                emailStore := common.ReadJSONstore(deploymentConfigStoreLocation)
                _, found := common.Find(emailStore.Emails, user1formEmail)
                Expect(found).To(BeFalse(), "Saved deployment file MUST NOT container the just posted email")
                Expect(resp.StatusCode).To(Equal(400), "Server MUST response with a 400 value status code")
        })

        It("should not allow duplicate email address strings to be saved", func() {
                By("Not failing the request")
                for i := 0; i < 2; i++ {
                        // set up request
                        user1formEmail := "user@example.com"
                        user1form := url.Values{}
                        user1form.Set("email", user1formEmail)

                        // send the request request
                        resp, err := http.PostForm("http://localhost:8080/api/interested", user1form)
                        Expect(err).To(BeNil(), "API request should not return error")

                        By("Finding posted value in deployment")

                        // get the deployment interested.json's contents
                        emailStore := common.ReadJSONstore(deploymentConfigStoreLocation)
                        _, found := common.Find(emailStore.Emails, user1formEmail)
                        Expect(found).To(BeTrue(), "Saved deployment file MUST container the just posted email")
                        Expect(resp.StatusCode).To(Equal(200), "Server MUST response with a 200 value status code")
                }
        })
})
