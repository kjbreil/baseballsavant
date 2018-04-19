package baseballsavant

import (
	"testing"
)

func Test_reflectTest(t *testing.T) {
	var f Filter
	// ty := reflect.TypeOf(f)
	// l := ty.NumField()

	f.All = true
	something := "Something"
	f.HfSit = &something
	// va := reflect.ValueOf(f)

	// test := va.FieldByName("HfPR").Interface()

	// log.Println(test)

	// tagMap := make(map[string]string)

	// for i := 0; i < l; i++ {
	// 	// .Tag.Get("filter")
	// 	name := ty.Field(i).Tag.Get("filter")
	// 	oname := ty.Field(i).Name
	// 	tagMap[name] = ty.Field(i).Name
	// 	log.Println(name, va.FieldByName(oname))
	// }
	// // tag, _ := ty.FieldByName("All")
	// // log.Println(tag.Tag.Get("column"))
	// log.Println(tagMap)

	// fmt.Println(f.All)
	// fmt.Println(va.FieldByName("All"))

	f.URL()

	t.Fail()

}
