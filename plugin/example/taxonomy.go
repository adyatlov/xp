package example

import "github.com/adyatlov/xp/data"

var TCompany = &data.ObjectType{
	Name:        "company",
	PluralName:  "companies",
	Description: "Commercial business",
	Properties: []*data.PropertyType{
		PEstablished,
		PCompanyForm,
		PIncome,
		PExpenses,
	},
}

var TDivision = &data.ObjectType{
	Name:        "division",
	PluralName:  "divisions",
	Description: "business unit",
	Properties: []*data.PropertyType{
		PEstablished,
		PIncome,
		PExpenses,
	},
}

var TEmployee = &data.ObjectType{
	Name:        "employee",
	PluralName:  "employees",
	Description: "person employed for wages or salary",
	Properties: []*data.PropertyType{
		PPosition,
		PFirstName,
		PLastName,
		PBirthDay,
		PFirstDay,
		PIncome,
		PEmail,
	},
}

var PEstablished = &data.PropertyType{
	Name:        "established",
	Type:        data.PVTTimestamp,
	Description: "day of foundation",
}

var PCompanyForm = &data.PropertyType{
	Name:        "company_form",
	Type:        data.PVTType,
	Description: "type of business",
}

var PIncome = &data.PropertyType{
	Name:        "income",
	Type:        data.PVTInteger,
	Description: "money received yearly",
}

var PExpenses = &data.PropertyType{
	Name:        "expenses",
	Type:        data.PVTInteger,
	Description: "money spent yearly",
}

var PBirthDay = &data.PropertyType{
	Name:        "birth_day",
	Type:        data.PVTTimestamp,
	Description: "day of birth",
}

var PFirstDay = &data.PropertyType{
	Name:        "first_day",
	Type:        data.PVTTimestamp,
	Description: "employment begins",
}

var PPosition = &data.PropertyType{
	Name:        "position",
	Type:        data.PVTType,
	Description: "job title",
}

var PEmail = &data.PropertyType{
	Name:        "email",
	Type:        data.PVTString,
	Description: "e-mail address",
}

var PFirstName = &data.PropertyType{
	Name:        "first_name",
	Type:        data.PVTString,
	Description: "first name",
}

var PLastName = &data.PropertyType{
	Name:        "last_name",
	Type:        data.PVTString,
	Description: "last name",
}
