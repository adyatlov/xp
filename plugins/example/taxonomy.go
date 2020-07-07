package example

import "github.com/adyatlov/xp/data"

var TCompany = &data.ObjectType{
	Name:        "company",
	PluralName:  "companies",
	Description: "Commercial business",
	PropertyTypes: []*data.PropertyType{
		PEstablished,
		PCompanyForm,
		PIncome,
		PExpenses,
	},
	ChildTypes: []*data.ObjectType{
		TDivision,
		TEmployee,
	},
}

var TDivision = &data.ObjectType{
	Name:        "division",
	PluralName:  "divisions",
	Description: "business unit",
	PropertyTypes: []*data.PropertyType{
		PEstablished,
		PIncome,
		PExpenses,
	},
	ChildTypes: []*data.ObjectType{
		TEmployee,
	},
}

var TEmployee = &data.ObjectType{
	Name:        "employee",
	PluralName:  "employees",
	Description: "person employed for wages or salary",
	PropertyTypes: []*data.PropertyType{
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
	ValueType:   data.PVTTimestamp,
	Description: "day of foundation",
}

var PCompanyForm = &data.PropertyType{
	Name:        "company form",
	ValueType:   data.PVTType,
	Description: "type of business",
}

var PIncome = &data.PropertyType{
	Name:        "income",
	ValueType:   data.PVTInteger,
	Description: "money received yearly",
}

var PExpenses = &data.PropertyType{
	Name:        "expenses",
	ValueType:   data.PVTInteger,
	Description: "money spent yearly",
}

var PBirthDay = &data.PropertyType{
	Name:        "birth day",
	ValueType:   data.PVTTimestamp,
	Description: "day of birth",
}

var PFirstDay = &data.PropertyType{
	Name:        "first day",
	ValueType:   data.PVTTimestamp,
	Description: "employment begins",
}

var PPosition = &data.PropertyType{
	Name:        "position",
	ValueType:   data.PVTType,
	Description: "job title",
}

var PEmail = &data.PropertyType{
	Name:        "email",
	ValueType:   data.PVTString,
	Description: "e-mail address",
}

var PFirstName = &data.PropertyType{
	Name:        "first name",
	ValueType:   data.PVTString,
	Description: "first name",
}

var PLastName = &data.PropertyType{
	Name:        "last name",
	ValueType:   data.PVTString,
	Description: "last name",
}
