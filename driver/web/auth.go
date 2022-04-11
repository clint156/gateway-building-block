/*
 *   Copyright (c) 2020 Board of Trustees of the University of Illinois.
 *   All rights reserved.

 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at

 *   http://www.apache.org/licenses/LICENSE-2.0

 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rokwire/core-auth-library-go/authservice"
	"github.com/rokwire/core-auth-library-go/tokenauth"
	"github.com/rokwire/logging-library-go/logs"
)

type cacheUser struct {
	lastUsage time.Time
}

//TokenAuth used to encapsualte the tokenauth type from the core auth library
type TokenAuth struct {
	tokenAuth *tokenauth.TokenAuth
}

func (auth TokenAuth) check(r *http.Request) error {
	_, err := auth.tokenAuth.CheckRequestTokens(r)
	if err != nil {
		log.Printf("auth -> coreAuthCheck: FAILED to validate token: %s", err.Error())
		return err
	}
	return nil
}

func printDeletedAccountIDs(accountIDs []string) error {
	log.Printf("Deleted account IDs: %v\n", accountIDs)
	return nil
}

//NewTokenAuth creats a token auth instance
func NewTokenAuth(serviceHost string, coreHost string) *TokenAuth {
	serviceID := "gateway"
	config := authservice.RemoteAuthDataLoaderConfig{
		AuthServicesHost: coreHost,
		ServiceToken:     "sampleotken",

		DeletedAccountsCallback: printDeletedAccountIDs,
	}

	logger := logs.NewLogger(serviceID, nil)
	dataLoader, err := authservice.NewRemoteAuthDataLoader(config, nil, logger)
	authHost := fmt.Sprintf("%s/bbs/service-regs", coreHost)
	fmt.Println(authHost)
	hostArray := make([]string, 1)
	hostArray[0] = authHost

	//serviceLoader := authservice.NewRemoteServiceRegLoader(hostArray)

	authService, err := authservice.NewAuthService(serviceID, serviceHost, dataLoader)
	var tokenAuth *tokenauth.TokenAuth
	if err == nil {
		tokenAuth, err = tokenauth.NewTokenAuth(true, authService, nil, nil)
		if err != nil {
			log.Printf("auth -> newAuth: FAILED to init token auth: %s", err.Error())
		}
	} else {
		log.Printf("auth -> newAuth: FAILED to init auth service: %s", err.Error())
	}

	auth := TokenAuth{tokenAuth: tokenAuth}
	return &auth
}
