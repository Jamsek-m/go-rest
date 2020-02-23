package rest

type EntityList struct {
	Count    int
	Entities []interface{}
}

func NewEntityList(entities []interface{}, count int) EntityList {
	return EntityList{
		Count:    count,
		Entities: entities,
	}
}
