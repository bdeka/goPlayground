package main

// import the necessary packages
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	// define the structures to hold our data
	type Product struct {
		XMLName   xml.Name `xml:"product"`
		Name      string   `xml:"name"`
		Quantity  string   `xml:"quantity"`
		UnitPrice string   `xml:"unitPrice"`
	}
	type Transaction struct {
		XMLName     xml.Name `xml:"transaction"`
		Type        string   `xml:"type,attr"`
		TotalAmount string   `xml:"totalAmount"`
		Product     Product  `xml:"product"`
	}

	type SalesData struct {
		XMLName      xml.Name      `xml:"salesData"`
		Transactions []Transaction `xml:"transaction"`
	}

	// Open the xmlFile
	xmlFile, err := os.Open("salesData.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened salesData.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our transactions array
	var transactions SalesData
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &transactions)

	// we iterate through transaction within our transactions array and
	// print out some of the details of the tansactions
	var nofTransactions int = len(transactions.Transactions)

	fmt.Println("No of transactions found: " + strconv.Itoa(nofTransactions))

	for i := 0; i < nofTransactions; i++ {
		fmt.Println("Transaction Type: " + transactions.Transactions[i].Type)
		fmt.Println("Total Amount: " + transactions.Transactions[i].TotalAmount)
		fmt.Println("Product Name: " + transactions.Transactions[i].Product.Name)
	}

}
