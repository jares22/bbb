package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"github.com/jares22/bbb/entity"
)

func TestPositive (t*testing.T){

	g:=NewGomegaWithT(t)

	t.Run("positive",func(t*testing.T){
		customer := entity.Customer{
			Name: "John Doe",
			Age: "30",
			CustomerID: "CU12345679",
		}
		
		ok,err := govalidator.ValidateStruct(customer)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}