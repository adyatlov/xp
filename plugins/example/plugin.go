package example

import (
	"net/url"
	"strconv"

	"github.com/adyatlov/xp/data/mem"

	"github.com/brianvoe/gofakeit/v5"

	"github.com/adyatlov/xp/data"
)

func init() {
	p := &data.Plugin{
		Name: "Example Plugin",
		Description: "The example plugin demonstrates the possibilities of XP using a fake organization structure. " +
			"Plugin URL should look like example.com/?minEmployee=10&maxEmployee=100&nDivision=11. " +
			"minEmployee and maxEmployee set min and max amount of employees per division, " +
			"and nDivision is the number of divisions in the organization.",
		Open:       open,
		Compatible: compatible,
	}
	data.RegisterPlugin(p)
}

func open(urlStr string) (data.Dataset, error) {
	minEmployee, maxEmployee, nDivision, err := parseURL(urlStr)
	if err != nil {
		return nil, err
	}
	return generateDataset(minEmployee, maxEmployee, nDivision), nil
}

func compatible(urlStr string) (bool, error) {
	if _, _, _, err := parseURL(urlStr); err != nil {
		return false, nil
	}
	return true, nil
}

// example.com/?minEmployee=10&maxEmployee=100&nDivision=11
func parseURL(urlStr string) (int, int, int, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return 0, 0, 0, err
	}
	minEmployee, err := strconv.Atoi(u.Query().Get("minEmployee"))
	if err != nil {
		return 0, 0, 0, err
	}
	maxEmployee, err := strconv.Atoi(u.Query().Get("maxEmployee"))
	if err != nil {
		return 0, 0, 0, err
	}
	nDivision, err := strconv.Atoi(u.Query().Get("nDivision"))
	if err != nil {
		return 0, 0, 0, err
	}
	return minEmployee, maxEmployee, nDivision, nil
}

func generateDataset(minEmployee, maxEmployee, nDivision int) *mem.Dataset {
	dataset, company := mem.NewDataset(data.DatasetId(gofakeit.UUID()),
		TCompany,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(gofakeit.Company()),
	)
	company.AddProperty(PEstablished.Name, gofakeit.Date())
	company.AddProperty(PCompanyForm.Name, gofakeit.CompanySuffix())
	company.AddProperty(PIncome.Name, gofakeit.Number(5e6, 10e6))
	company.AddProperty(PExpenses.Name, gofakeit.Number(1e6, 5e6))
	for i := 0; i < nDivision; i++ {
		division := generateDivision(dataset, minEmployee, maxEmployee)
		company.AddChild(division)
	}
	return dataset
}

func generateDivision(d *mem.Dataset, minEmployee, maxEmployee int) *mem.Object {
	division := d.NewObject(TDivision,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(gofakeit.Vegetable()))
	division.AddProperty(PEstablished.Name, gofakeit.Date())
	division.AddProperty(PIncome.Name, gofakeit.Number(1e6, 5e6))
	division.AddProperty(PExpenses.Name, gofakeit.Number(1e5, 1e6))
	for i := 0; i < gofakeit.Number(minEmployee, maxEmployee); i++ {
		division.AddChild(generateEmployee(d))
	}
	return division
}

func generateEmployee(d *mem.Dataset) *mem.Object {
	person := gofakeit.Person()
	employee := d.NewObject(TEmployee,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(person.FirstName+" "+person.LastName))
	employee.AddProperty(PPosition.Name, person.Job.Title)
	employee.AddProperty(PFirstName.Name, person.FirstName)
	employee.AddProperty(PLastName.Name, person.LastName)
	employee.AddProperty(PBirthDay.Name, gofakeit.Date())
	employee.AddProperty(PFirstDay.Name, gofakeit.Date())
	employee.AddProperty(PIncome.Name, gofakeit.Number(4e4, 1e5))
	employee.AddProperty(PEmail.Name, person.Contact.Email)
	return employee
}
