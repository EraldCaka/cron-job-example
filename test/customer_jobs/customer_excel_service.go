package customer_jobs

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
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
	err = f.DeleteSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fileName := "CustomersData.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Data written to CustomersData.xlsx successfully")
	sendEmail(fileName)

	return f
}
func sendEmail(filename string) {
	m := gomail.NewMessage()

	m.SetHeader("From", "name@example.com")
	m.SetHeader("To", "name@exmple.com")
	m.SetHeader("Subject", "Customer Data")
	m.SetBody("text/plain", "Please find the attached customer data.")
	m.Attach(filename)

	d := gomail.NewDialer("smtp.example.com", 587, "name@example.com", "secret password")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent successfully")
}
