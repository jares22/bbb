package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"github.com/jares22/bbb/entity"
)

func TestNegative (t*testing.T){

	g:=NewGomegaWithT(t)

	t.Run("positivee",func(t*testing.T){
		customer := entity.Customer{
			Name: "John Doe",
			Age: "aaa",
			CustomerID: "CU12345679",
		}
		
		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("age must be an integer"))
	})
}