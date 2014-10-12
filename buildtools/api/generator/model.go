package generator

import (
	"encoding/json"
	"strconv"
)

type Swagger struct {
	Swagger      float32              `json:"swagger"`
	Info         Info                 `json:"info"`
	ExternalDocs ExternalDocs         `json:"externalDocs"`
	Host         string               `json:"host"`
	BasePath     string               `json:"basePath"`
	Schemes      []string             `json:"schemes"`
	Consumes     []string             `json:"consumes"`
	Produces     []string             `json:"produces"`
	Paths        Paths                `json:"paths"`
	Definitions  map[string]Schema    `json:"definitions"`
	Parameters   map[string]Parameter `json:"parameters"`
	Responses    map[string]Response  `json:"responses"`
	Tags         Tags                 `json:"tags"`
}

type RefType struct {
	Ref string `json:"$ref"`
}

type Info struct {
	Version        string  `json:"version"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
}

type Contact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Paths struct {
	Items      map[string]PathItem
	Extensions map[string]interface{}
}

func (p *Paths) UnmarshalJSON(b []byte) (err error) {
	i := make(map[string]PathItem)
	if err = json.Unmarshal(b, &i); err != nil {
		return
	}
	p.Items = i
	return
}

type PathItem struct {
	RefType
	Get        Operation   `json:"get"`
	Put        Operation   `json:"put"`
	Post       Operation   `json:"post"`
	Delete     Operation   `json:"delete"`
	Options    Operation   `json:"options"`
	Head       Operation   `json:"head"`
	Patch      Operation   `json:"patch"`
	Parameters []Parameter `json:"parameters"`
}

type Operation struct {
	Tags         []string     `json:"tags"`
	Summary      string       `json:"summary"`
	Description  string       `json:"Description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
	OperationId  string       `json:"operationId"`
	Consumes     []string     `json:"consumes"`
	Produces     []string     `json:"produces"`
	Parameters   []Parameter  `json:"parameters"`
	Responses    Responses    `json:"responses"`
	Schemes      []string     `json:"schemes"`
	Deprecated   bool         `json:"deprecated"`
}

type Responses struct {
	Default    Response
	Status     map[int]Response
	Extensions map[string]interface{}
}

func (r *Responses) UnmarshalJSON(b []byte) (err error) {
	res := make(map[string]Response)
	if err = json.Unmarshal(b, &res); err == nil {
		r.Status = make(map[int]Response)
		for k, v := range res {
			if k == "default" {
				r.Default = v
			} else if stCode, stCodeErr := strconv.Atoi(k); stCodeErr == nil {
				r.Status[stCode] = v
			}
		}
	}

	return
}

type Response struct {
	RefType
	Description string                 `json:"Description"`
	Schema      Schema                 `json:"schema"`
	Headers     map[string]ItemsType   `json:"headers"`
	Examples    map[string]interface{} `json:"examples"`
}

type Parameter struct {
	ItemsType
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Schema      Schema `json:"schema"`
	Extensions  map[string]interface{}
}

type Schema struct {
	ItemsType
	Required   []string            `json:"required"`
	Properties map[string]Property `json:"properties"`
}

type Property struct {
}

type ItemsType struct {
	RefType
	Type             string        `json:"type"`
	Format           string        `json:"format"`
	Items            *ItemsType    `json:"items"`
	CollectionFormat string        `json:"collectionFormat"`
	Default          string        `json:"default"`
	Maximum          float64       `json:"maximum"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          float64       `json:"minimum"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int64         `json:"maxLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         int64         `json:"maxItems"`
	MinItems         int64         `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum"`
	multipleOf       float64       `json:"multipleOf"`
}

type Tags struct {
	Tag        Tag
	Extensions map[string]interface{}
}

type Tag struct {
	Name         string       `json:"name"`
	Description  string       `json:"Description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
}

type ExternalDocs struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}
