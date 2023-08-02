package helpers

//import "time"

func CalculateSum(num1, num2 int, sumChan chan<- int) {
	sum := num1 + num2

	sumChan <- sum
	sumChan <- sum

}

func CalculateAverage(sumChan <-chan int,avgChan chan<- float64) {

	sum := <-sumChan
	avgval := float64(sum) / 2.0
	avgChan <- avgval
	avgChan <- avgval
	//close(avgChan)


}