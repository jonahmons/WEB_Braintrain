package model

type UpperCategory struct {
	Name       string
	Id         int
	Categories []Category
}

type Category struct {
	Id   int
	Name string
}

var Categories = []UpperCategory{
	UpperCategory{
		Name: "Naturwissenschaften",
		Id:   1,
		Categories: []Category{
			Category{Id: 1, Name: "Biologie"},
			Category{Id: 2, Name: "Chemie"},
			Category{Id: 3, Name: "Elektotechnik"},
			Category{Id: 4, Name: "Informatik"},
			Category{Id: 5, Name: "Mathematik"},
			Category{Id: 6, Name: "Medizin"},
			Category{Id: 7, Name: "Naturkunde"},
			Category{Id: 8, Name: "Physik"},
			Category{Id: 9, Name: "Sonstiges"},
		},
	},
	UpperCategory{
		Name: "Sprachen",
		Id:   2,
		Categories: []Category{
			Category{Id: 10, Name: "Chinesisch"},
			Category{Id: 11, Name: "Chemie"},
			Category{Id: 12, Name: "Deutsch"},
			Category{Id: 13, Name: "Englisch"},
			Category{Id: 14, Name: "Französisch"},
			Category{Id: 15, Name: "Griechisch"},
			Category{Id: 16, Name: "Italienisch"},
			Category{Id: 17, Name: "Latein"},
			Category{Id: 18, Name: "Russisch"},
			Category{Id: 19, Name: "Sonstiges"},
		},
	},
	UpperCategory{
		Name: "Gesellschaft",
		Id:   3,
		Categories: []Category{
			Category{Id: 20, Name: "Ethik"},
			Category{Id: 21, Name: "Geschichte"},
			Category{Id: 22, Name: "Literatur"},
			Category{Id: 23, Name: "Musik"},
			Category{Id: 24, Name: "Politik"},
			Category{Id: 25, Name: "Recht"},
			Category{Id: 26, Name: "Soziales"},
			Category{Id: 27, Name: "Sport"},
			Category{Id: 28, Name: "Verkehrskunde"},
			Category{Id: 29, Name: "Sonstiges"},
		},
	},
	UpperCategory{
		Name: "Wirtschaft",
		Id:   4,
		Categories: []Category{
			Category{Id: 30, Name: "BWL"},
			Category{Id: 31, Name: "Finanzen"},
			Category{Id: 32, Name: "Landwirtschaft"},
			Category{Id: 33, Name: "Marketing"},
			Category{Id: 34, Name: "VWL"},
			Category{Id: 35, Name: "Sonstiges"},
		},
	},
	UpperCategory{
		Name: "Geisteswissenschaften",
		Id:   5,
		Categories: []Category{
			Category{Id: 36, Name: "Kriminologie"},
			Category{Id: 37, Name: "Philosophie"},
			Category{Id: 38, Name: "Psychologie"},
			Category{Id: 39, Name: "Pädagogik"},
			Category{Id: 40, Name: "Theologie"},
			Category{Id: 41, Name: "Sonstiges"},
		},
	},
}

//GetUpperCategoryByCatID funktion
func GetUpperCategoryByCatID(id int) (int, string) {
	for _, uc := range Categories {
		for _, c := range uc.Categories {
			if c.Id == id {
				return uc.Id, uc.Name
			}
		}
	}
	return 0, ""
}

func GetCategoryByID(id int) string {
	for _, uc := range Categories {
		for _, c := range uc.Categories {
			if c.Id == id {
				return c.Name
			}
		}
	}
	return ""
}
