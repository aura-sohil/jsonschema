package jsonschema_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aura-sohil/jsonschema"
)

type SampleUser struct {
	ID          int            `json:"id"`
	Name        string         `json:"name" jsonschema:"title=the name,description=The name of a friend,example=joe,example=lucy,default=alex"`
	Friends     []int          `json:"friends,omitempty" jsonschema_description:"The list of IDs, omitted when empty"`
	Tags        map[string]any `json:"tags,omitempty" jsonschema_extras:"a=b,foo=bar,foo=bar1"`
	BirthDate   time.Time      `json:"birth_date,omitempty" jsonschema:"oneof_required=date"`
	YearOfBirth string         `json:"year_of_birth,omitempty" jsonschema:"oneof_required=year"`
	Metadata    any            `json:"metadata,omitempty" jsonschema:"oneof_type=string;array"`
	FavColor    string         `json:"fav_color,omitempty" jsonschema:"enum=red,enum=green,enum=blue"`
}

func ExampleReflect() {
	s := jsonschema.Reflect(&SampleUser{})
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "$id": "https://github.com/aura-sohil/jsonschema_test/sample-user",
	//   "$ref": "#/$defs/SampleUser",
	//   "$defs": {
	//     "SampleUser": {
	//       "oneOf": [
	//         {
	//           "required": [
	//             "birth_date"
	//           ],
	//           "title": "date"
	//         },
	//         {
	//           "required": [
	//             "year_of_birth"
	//           ],
	//           "title": "year"
	//         }
	//       ],
	//       "properties": {
	//         "id": {
	//           "type": "integer"
	//         },
	//         "name": {
	//           "type": "string",
	//           "title": "the name",
	//           "description": "The name of a friend",
	//           "default": "alex",
	//           "examples": [
	//             "joe",
	//             "lucy"
	//           ]
	//         },
	//         "friends": {
	//           "oneOf": [
	//             {
	//               "items": {
	//                 "type": "integer"
	//               },
	//               "type": "array",
	//               "description": "The list of IDs, omitted when empty"
	//             },
	//             {
	//               "type": "null"
	//             }
	//           ]
	//         },
	//         "tags": {
	//           "oneOf": [
	//             {
	//               "type": "object",
	//               "a": "b",
	//               "foo": [
	//                 "bar",
	//                 "bar1"
	//               ]
	//             },
	//             {
	//               "type": "null"
	//             }
	//           ]
	//         },
	//         "birth_date": {
	//           "oneOf": [
	//             {
	//               "type": "string",
	//               "format": "date-time"
	//             },
	//             {
	//               "type": "null"
	//             }
	//           ]
	//         },
	//         "year_of_birth": {
	//           "oneOf": [
	//             {
	//               "type": "string"
	//             },
	//             {
	//               "type": "null"
	//             }
	//           ]
	//         },
	//         "metadata": {
	//           "oneOf": [
	//             {
	//               "type": "string"
	//             },
	//             {
	//               "type": "array"
	//             }
	//           ]
	//         },
	//         "fav_color": {
	//           "oneOf": [
	//             {
	//               "type": "string",
	//               "enum": [
	//                 "red",
	//                 "green",
	//                 "blue"
	//               ]
	//             },
	//             {
	//               "type": "null"
	//             }
	//           ]
	//         }
	//       },
	//       "additionalProperties": false,
	//       "type": "object",
	//       "required": [
	//         "id",
	//         "name"
	//       ]
	//     }
	//   }
	// }
}
