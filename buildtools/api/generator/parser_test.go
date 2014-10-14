package generator

import "testing"

var PETSTORE_EXPANDED_FILE string = "test/petstore-expanded.json"

func TestParseSwagger(t *testing.T) {
	sw := loadSwagger(t, PETSTORE_EXPANDED_FILE)
	a := newAssert(t)

	a.assertEqualStr("swagger", sw.Swagger, "2.0")
	a.assertEqualStr("host", sw.Host, "petstore.swagger.wordnik.com")
	a.assertEqualStrSl("schemes", sw.Schemes, []string{"http"})
	a.assertEqualStrSl("consumes", sw.Consumes, []string{"application/json"})
	a.assertEqualStr("info.version", sw.Info.Version, "1.0.0")
	a.assertEqualStr("info.title", sw.Info.Title, "Swagger Petstore")
	a.assertEqualStr("info.contact.name", sw.Info.Contact.Name, "Wordnik API Team")
	a.assertEqualStr("info.license.name", sw.Info.License.Name, "MIT")

	a.assertEqualInt("len(paths)", len(sw.Paths.Items), 2)
}

func TestGetPets(t *testing.T) {
	sw := loadSwagger(t, PETSTORE_EXPANDED_FILE)
	a := newAssert(t)

	a.assertNonNil("paths./pets", sw.Paths.Items["/pets"])

	getPets := sw.Paths.Items["/pets"].Get
	a.assertNonNil("paths./pets.get", getPets)
	a.assertEqualStr("paths./pets.get.operationId", getPets.OperationId, "findPets")
	a.assertEqualStrSl("paths./pets.get.produces", getPets.Produces, []string{
		"application/json",
		"application/xml",
		"text/xml",
		"text/html",
	})
	a.assertEqualInt("len(paths./pets.get.parameters)", len(getPets.Parameters), 2)

	tagsParam := getPets.Parameters[0]
	a.assertEqualStr("paths./pets.get.parameters[0].name", tagsParam.Name, "tags")
	a.assertEqualStr("paths./pets.get.parameters[0].in", tagsParam.In, "query")
	a.assertEqualBool("paths./pets.get.parameters[0].required", tagsParam.Required, false)
	a.assertEqualStr("paths./pets.get.parameters[0].type", tagsParam.Type, "array")
	a.assertEqualStr("paths./pets.get.parameters[0].items.type", tagsParam.Items.Type, "string")
	a.assertEqualStr("paths./pets.get.parameters[0].collectionFormat", tagsParam.CollectionFormat, "csv")

	limitParam := getPets.Parameters[1]
	a.assertEqualStr("paths./pets.get.parameters[1].name", limitParam.Name, "limit")
	a.assertEqualStr("paths./pets.get.parameters[1].in", limitParam.In, "query")
	a.assertEqualBool("paths./pets.get.parameters[1].required", limitParam.Required, false)
	a.assertEqualStr("paths./pets.get.parameters[1].type", limitParam.Type, "integer")
	a.assertEqualStr("paths./pets.get.parameters[1].format", limitParam.Format, "int32")

	resp := getPets.Responses.Default
	a.assertEqualStr("paths./pets.get.responses.default.description", resp.Description, "unexpected error")
	a.assertEqualStr("paths./pets.get.responses.default.schema.ref", resp.Schema.Ref, "#/definitions/errorModel")

	resp = getPets.Responses.Status[200]
	a.assertEqualStr("paths./pets.get.responses.200.description", resp.Description, "pet response")
	a.assertEqualStr("paths./pets.get.responses.200.schema.type", resp.Schema.Type, "array")
	a.assertEqualStr("paths./pets.get.responses.200.schema.items.ref", resp.Schema.Items.Ref, "#/definitions/pet")
}

func TestPostPets(t *testing.T) {
	sw := loadSwagger(t, PETSTORE_EXPANDED_FILE)
	a := newAssert(t)

	a.assertNonNil("paths./pets", sw.Paths.Items["/pets"])

	postPets := sw.Paths.Items["/pets"].Post
	a.assertNonNil("paths./pets.post", postPets)
	a.assertEqualStr("paths./pets.post.description", postPets.Description, "post description")
	a.assertEqualStr("paths./pets.post.operationId", postPets.OperationId, "addPet")
	a.assertEqualStrSl("paths./pets.post.produces", postPets.Produces, []string{
		"application/json",
	})
	a.assertEqualInt("len(paths./pets.post.parameters)", len(postPets.Parameters), 1)

	petParam := postPets.Parameters[0]
	a.assertEqualStr("paths./pets.post.parameters[0].name", petParam.Name, "pet")
	a.assertEqualStr("paths./pets.post.parameters[0].in", petParam.In, "body")
	a.assertEqualBool("paths./pets.post.parameters[0].required", petParam.Required, true)
	a.assertEqualStr("paths./pets.post.parameters[0].items.type", petParam.Schema.Ref, "#/definitions/newPet")
}

func TestDefinitions(t *testing.T) {
	sw := loadSwagger(t, PETSTORE_EXPANDED_FILE)
	a := newAssert(t)

	def := sw.Definitions["pet"]
	a.assertNonNil("definitions.pet", def)

	a.assertEqualStrSl("definitions.pet.required", def.Required, []string{
		"id",
		"name",
	})
	a.assertEqualInt("len(definitions.pet.properties)", len(def.Properties), 3)
	a.assertEqualStr("definitions.pet.properties.id.type", def.Properties["id"].Type, "integer")
	a.assertEqualStr("definitions.pet.properties.id.format", def.Properties["id"].Format, "int64")
	a.assertEqualStr("definitions.pet.properties.name.type", def.Properties["name"].Type, "string")
	a.assertEqualStr("definitions.pet.properties.tag.type", def.Properties["tag"].Type, "string")

	def = sw.Definitions["newPet"]
	a.assertNonNil("definitions.newPet.allOf", def.AllOf)
	a.assertEqualInt("len(definitions.newPet.allOf)", len(def.AllOf), 2)
	a.assertEqualStr("definitions.newPet.allOf[0].ref", def.AllOf[0].Ref, "pet")

	newPet := def.AllOf[1]
	a.assertEqualStrSl("definitions.newPet.allOf[1].required", newPet.Required, []string{
		"name",
	})
	a.assertEqualInt("len(definitions.newPet.allOf[1].properties)", len(newPet.Properties), 1)
	a.assertEqualStr("definitions.newPet.allOf[1].properties.id.type", newPet.Properties["id"].Type, "number")
	a.assertEqualStr("definitions.newPet.allOf[1].properties.id.format", newPet.Properties["id"].Format, "float64")
}

func loadSwagger(t *testing.T, filePath string) *Swagger {
	var err error
	var sw *Swagger
	if sw, err = ParseSwagger(filePath); err != nil {
		t.Errorf("Error parsing api description: %s", err.Error())
	}
	return sw
}

type assert struct {
	*testing.T
}

func newAssert(t *testing.T) *assert {
	return &assert{t}
}

func (t *assert) assertEqualFl32(info string, actual, expected float32) {
	if actual != expected {
		t.Errorf("%s: expected:%f actual:%f", info, expected, actual)
	}
}

func (t *assert) assertEqualInt(info string, actual, expected int) {
	if actual != expected {
		t.Errorf("%s: expected:%d actual:%d", info, expected, actual)
	}
}

func (t *assert) assertEqualBool(info string, actual, expected bool) {
	if actual != expected {
		t.Errorf("%s: expected:%v :%v", info, expected, actual)
	}
}

func (t *assert) assertEqualStr(info, actual, expected string) {
	if actual != expected {
		t.Errorf("%s: expected:%s actual:%s", info, expected, actual)
	}
}

func (t *assert) assertNonNil(info string, actual interface{}) {
	if actual == nil {
		t.Errorf("%s: expected non nil actual:%f", info, actual)
	}
}

func (t *assert) assertEqualStrSl(info string, actual, expected []string) {
	if len(actual) != len(expected) {
		t.Errorf("%s: expected:%v actual:%v", info, expected, actual)
		return
	}

	for i, s := range expected {
		if actual[i] != s {
			t.Errorf("%s: expected:%v actual:%v", info, expected, actual)
			return
		}
	}
}
