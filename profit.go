package main

import "fmt" 


func main(){
   //var revenue1,revenue2,revenue3,revenue4 float64
   //var expenses1,expenses2,expenses3,expenses4 float64
   //var taxRate float64

   var numProducts=4
    revenues := make([]float64,numProducts)
    expenses :=make([]float64,numProducts)
   var taxRate float64
   
   for i:=0; i<numProducts; i++{
    fmt.Printf("Revenue of Product %d: ", i+1 )
    fmt.Scan(&revenues[i])

    fmt.Printf("Expenses of Product %v: ",i+1)
    fmt.Scan(&expenses[i])

   }
     
   
   //fmt.Print("Revenue of the first product: ")
   //fmt.Scan(&revenue1)

   //fmt.Print("Revenue of the second product: ")
   //fmt.Scan(&revenue2)

   //fmt.Print("Revenue of the third product: ")
   //fmt.Scan(&revenue3)

   //fmt.Print("Revenue of the fourth product: ")
   //fmt.Scan(&revenue4)

   //fmt.Print("Expenses of the first product: ")
   //fmt.Scan(&expenses1)

   //fmt.Print("Expenses of the second product: ")
   //fmt.Scan(&expenses2)

   //fmt.Print("Expenses of the third product: ")
   //fmt.Scan(&expenses3)

   //fmt.Print("Expenses of the    fourth product: ")
   //fmt.Scan(&expenses4)

   fmt.Print("Tax Rate: ")
   fmt.Scan(&taxRate)

   //totalRevenue := revenue1+revenue2+revenue3+revenue4
   //totalExpenses := expenses1+expenses2+expenses3+expenses4
   var totalRevenue,totalExpenses float64
   for i:=0; i<numProducts; i++{
    totalRevenue +=revenues[i]
    totalExpenses +=expenses[i]
   }
   
   earningBT :=totalRevenue- totalExpenses
   profit := earningBT*(1-taxRate/100) //earningsAT
   profitMargin := (profit/totalRevenue)*100

   var investmentIntoProdutct float64
   if profitMargin>=20{
    investmentIntoProdutct = 0.1*profit
   } else{
    investmentIntoProdutct = 0
   }

  fmt.Println("Total Revenue: ",totalRevenue)
  fmt.Println("Total Expenses: ",totalExpenses)
  fmt.Println("Earning Before the Tax: ",earningBT)
  fmt.Println("Profit: ",profit)
  fmt.Println("Profit Margin: ",profitMargin, "%")
  fmt.Println("Investment: ",investmentIntoProdutct)

}
