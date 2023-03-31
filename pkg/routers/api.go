/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package routers

import (
	"github.com/corelayer/nsregistry/pkg/controllers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ApiRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/customers", listCustomers)
	r.Post("/customers", createCustomer)

	return r
}

func listCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := controllers.ListCustomers()
	if err != nil {
		w.WriteHeader(503)
		return
	}
	w.Write(c)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Customer created"))
}
