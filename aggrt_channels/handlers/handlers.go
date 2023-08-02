package handlers

import (
	"strconv"
	"sync"
	//"time"

	"example/aggrt_channels/helpers"

	"github.com/gofiber/fiber/v2"
)

func FindAggrt(c *fiber.Ctx) error{

    num1,err:=strconv.Atoi(c.Params("num1"))
    if err != nil {
	  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number format"})
  }


    num2,err:=strconv.Atoi(c.Params("num2"))
    if err != nil {
	  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number format"})
  }


 sumChan := make(chan int,2)
 avgChan := make(chan float64,2)


 var wg sync.WaitGroup
  wg.Add(2)

  go func() {
	  defer wg.Done()
	  helpers.CalculateSum(num1, num2, sumChan)
  }()
  go func() {
	  defer wg.Done()
	  helpers.CalculateAverage(sumChan,avgChan)
  }()

  wg.Wait()


  prod:=num1 * num2

  return c.JSON(fiber.Map{
	"sum":         <-sumChan,
	"average":     <-avgChan,
	"multiplication": prod,
 })
 //close(avgChan)
 //close(sumChan)
 }






  //go helpers.CalculateSum(num1, num2, sumChan)
  //go helpers.CalculateAverage(num1, num2,sumChan,avgChan)



 /*func calculateSum(num1, num2 int, sumChan chan<- int) {
	sum := num1 + num2

	sumChan <- sum
	sumChan <- sum

}

func calculateAverage(num1, num2 int,sumChan chan int,avgChan chan<- float64) {

	tempsum := <-sumChan
	avgval := float64(tempsum) / float64(2)
	avgChan <- avgval
}


var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		helpers.CalculateSum(num1, num2, sumChan)
	}()

	go func() {
		defer wg.Done()
		helpers.CalculateAverage(num1, num2,sumChan,avgChan)
	}()

	wg.Wait()*/






	/*package handlers

import (
	"strconv"
	//"sync"

	"example/aggrt_channels/helpers"

	"github.com/gofiber/fiber/v2"
)

func FindAggrt(c *fiber.Ctx) error{

    num1,err:=strconv.Atoi(c.Params("num1"))
    if err != nil {
	  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number format"})
  }


    num2,err:=strconv.Atoi(c.Params("num2"))
    if err != nil {
	  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number format"})
  }


 sumChan := make(chan int)
 avgChan := make(chan float64)

 go helpers.CalculateSum(num1, num2, sumChan)
 go helpers.CalculateAverage(num1, num2,sumChan,avgChan)
 

  sum:=<-sumChan
  avg:=<-avgChan

  prod:=num1*num2

  return c.JSON(fiber.Map{
	"sum":         sum,
	"average":     avg,
	"multiplication": prod,
 })
 }*/
