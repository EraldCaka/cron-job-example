package customer_jobs

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func (c *CustomerJob) InsertCustomerDataInsideExcelFile() *excelize.File {
	f := excelize.NewFile()
	index, err := f.NewSheet("Customers-1")
	if err != nil {
		log.Fatal(err)
	}

	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F81BD"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	dataStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Color: "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	headers := []string{"Name", "Money", "Age"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col)
		f.SetCellValue("Customers-1", cell, header)
		f.SetCellStyle("Customers-1", cell, cell, headerStyle)
	}

	for i, customer := range c.Customers {
		f.SetCellValue("Customers-1", fmt.Sprintf("A%v", i+2), customer.Name)
		f.SetCellValue("Customers-1", fmt.Sprintf("B%v", i+2), customer.Money)
		f.SetCellValue("Customers-1", fmt.Sprintf("C%v", i+2), customer.Age)
	}

	f.SetCellStyle("Customers-1", "A2", fmt.Sprintf("C%v", len(c.Customers)+1), dataStyle)

	f.SetActiveSheet(index)

	if err := f.SaveAs("CustomersData.xlsx"); err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Data written to CustomersData.xlsx successfully")
	return f
}
