package plugin

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
	DefaultProperties: []*data.PropertyType{
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
	DefaultProperties: []*data.PropertyType{
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
	DefaultProperties: []*data.PropertyType{
		PPosition,
		PFirstDay,
		PIncome,
	},
}

var PEstablished = &data.PropertyType{
	Name:        "established",
	ValueType:   data.MVTTimestamp,
	Description: "day of foundation",
}

var PCompanyForm = &data.PropertyType{
	Name:        "company_form",
	ValueType:   data.MVTType,
	Description: "type of business",
}

var PIncome = &data.PropertyType{
	Name:        "income",
	ValueType:   data.MVTInteger,
	Description: "money received yearly",
}

var PExpenses = &data.PropertyType{
	Name:        "expenses",
	ValueType:   data.MVTReal,
	Description: "money spent yearly",
}

var PBirthDay = &data.PropertyType{
	Name:        "birth_day",
	ValueType:   data.MVTTimestamp,
	Description: "day of birth",
}

var PFirstDay = &data.PropertyType{
	Name:        "first_day",
	ValueType:   data.MVTTimestamp,
	Description: "employment begins",
}

var PPosition = &data.PropertyType{
	Name:        "position",
	ValueType:   data.MVTType,
	Description: "job title",
}

var PEmail = &data.PropertyType{
	Name:        "email",
	ValueType:   data.MVTString,
	Description: "e-mail address",
}

var PFirstName = &data.PropertyType{
	Name:        "first_name",
	ValueType:   data.MVTString,
	Description: "first name",
}

var PLastName = &data.PropertyType{
	Name:        "last_name",
	ValueType:   data.MVTString,
	Description: "last name",
}
