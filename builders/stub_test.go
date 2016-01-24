package builders_test

import (
	. "github.com/durmaze/gobank/builders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"sync"
)

var _ = Describe("Stub Builder Tests", func() {

	Describe("When building a Stub with single Response", func() {
		var (
			actualResponse Response
			expectedResponse Response
			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func(){
				headers := make(map[string]string)
				headers["Content-Type"] = "application/json"
				headers["X-Custom-Header"] = "ABC123"

				is := Is{
					StatusCode: 200,
					Headers: headers,
					Body: "{ \"greeting\": \"Hello GoBank\" }",
				}

				expectedResponse = Response{
					Is: is,
				}

				stub := NewStubBuilder().AddResponse(expectedResponse).Build()

				actualResponse = stub.Responses[0]
			})
	  })

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Content-Type header", func() {
			Expect(actualResponse.Is.Headers["Content-Type"]).To(Equal(expectedResponse.Is.Headers["Content-Type"]))
		})

		It("should create a Stub that returns a Response with the correct Custom header", func() {
			Expect(actualResponse.Is.Headers["X-Custom-Header"]).To(Equal(expectedResponse.Is.Headers["X-Custom-Header"]))
		})
		
		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})
	})

})