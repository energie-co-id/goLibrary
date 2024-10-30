// Copyright 2024 openGemini Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openGeminiLib

/*
	The example code use the dot import, but the user should choose the package import method according to their own needs
*/

import (
	// "context"
	"fmt"

	// "math/rand"
	// "time"

	"github.com/libgox/addr"
	"github.com/openGemini/opengemini-client-go/opengemini"
)

func CreateClient(host string, port int) (opengemini.Client, error) {
	// create an openGemini client
	config := &opengemini.Config{
		Addresses: []addr.Address{{
			Host: host,
			Port: port,
		}},
	}
	client, err := opengemini.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return client, err
	}
	return client, nil
}

func Write(client opengemini.Client, database string, measurement string) {
	// use point write method
	point := &opengemini.Point{}
	point.Measurement = measurement
	point.AddTag("Weather", "foggy")
	point.AddField("Humidity", 87)
	point.AddField("Temperature", 25)
	err := client.WritePoint(database, point, func(err error) {
		if err != nil {
			fmt.Printf("write point failed for %s", err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
}
func RawQuery(client opengemini.Client, database string, query string) (*opengemini.QueryResult, error) {
	// do a query
	q := opengemini.Query{
		Database: database,
		Command:  query,
	}
	res, err := client.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
