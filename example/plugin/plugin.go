package plugin

import (
	"net/url"
	"strconv"

	"github.com/adyatlov/xp/data"
	"github.com/brianvoe/gofakeit/v5"
)

func init() {
	plugin := NewPlugin("example", "example company plugin", open)
	data.RegisterPlugin(plugin)
}

func open(urlStr string) (data.Dataset, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	minEmployee, err := strconv.Atoi(u.Query().Get("minEmployee"))
	if err != nil {
		return nil, err
	}
	maxEmployee, err := strconv.Atoi(u.Query().Get("maxEmployee"))
	if err != nil {
		return nil, err
	}
	nDivision, err := strconv.Atoi(u.Query().Get("nDivision"))
	if err != nil {
		return nil, err
	}
	return generateDataset(minEmployee, maxEmployee, nDivision), nil
}

func generateDataset(minEmployee, maxEmployee, nDivision int) *Dataset {
	company := NewObject(TCompany,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(gofakeit.Company()))
	d := NewDataset(data.DatasetId(gofakeit.UUID()), company)
	company.AddProperty(NewProperty(PEstablished, gofakeit.Date()))
	company.AddProperty(NewProperty(PCompanyForm, gofakeit.CompanySuffix()))
	company.AddProperty(NewProperty(PIncome, gofakeit.Number(5e6, 10e6)))
	company.AddProperty(NewProperty(PExpenses, gofakeit.Number(1e6, 5e6)))
	for i := 0; i < nDivision; i++ {
		division := generateDivision(d, minEmployee, maxEmployee)
		company.AddChild(division)
		groups, err := division.Children(TEmployee.Name)
		if err != nil {
			panic("division should have at employees")
		}
		for _, employee := range groups[0].Objects() {
			company.AddChild(employee)
		}
	}
	return d
}

func generateDivision(d *Dataset, minEmployee, maxEmployee int) *Object {
	division := d.NewObject(TDivision,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(gofakeit.Vegetable()))
	division.AddProperty(NewProperty(PEstablished, gofakeit.Date()))
	division.AddProperty(NewProperty(PIncome, gofakeit.Number(1e6, 5e6)))
	division.AddProperty(NewProperty(PExpenses, gofakeit.Number(1e5, 1e6)))
	for i := 0; i < gofakeit.Number(minEmployee, maxEmployee); i++ {
		division.AddChild(generateEmployee(d))
	}
	return division
}

func generateEmployee(d *Dataset) *Object {
	person := gofakeit.Person()
	employee := d.NewObject(TEmployee,
		data.ObjectId(gofakeit.UUID()),
		data.ObjectName(person.FirstName+" "+person.LastName))
	employee.AddProperty(NewProperty(PPosition, person.Job.Title))
	employee.AddProperty(NewProperty(PFirstName, person.FirstName))
	employee.AddProperty(NewProperty(PLastName, person.LastName))
	employee.AddProperty(NewProperty(PBirthDay, gofakeit.Date()))
	employee.AddProperty(NewProperty(PFirstDay, gofakeit.Date()))
	employee.AddProperty(NewProperty(PIncome, gofakeit.Number(4e4, 1e5)))
	employee.AddProperty(NewProperty(PEmail, person.Contact.Email))
	return employee
}
