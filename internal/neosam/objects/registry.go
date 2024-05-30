// /******************************************************************************
//  * Date   : May 2024
//  * Brief  : Object registry implementation using a singleton design.
//  *
//  * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
//  * SPDX-License-Identifier: MIT
//  *****************************************************************************/

// package neon_samurai

// import (
// 	"fmt"
// 	"reflect"
// 	"sync"
// )

// /* ------------------------------- Types ----------------------------------- */

// type objectRegistry struct {
// 	interfaces map[objectID]objectInterface
// 	sync.Mutex
// }

// /* ------------------------------- Constants ------------------------------- */
// /* ------------------------------- Variables ------------------------------- */

// var (
// 	singleton *objectRegistry
// 	once sync.Once
// )

// var idMap = map[objectID]objectInterface{
// 	objectSystemConfig: reflect.TypeOf((*config)(nil)).Elem(),
// }

// /* ------------------------------- Functions ------------------------------- */

// func SetupRegistry() *objectRegistry {
// 	once.Do(func() {
// 		singleton = &objectRegistry{interfaces: idMap}
// 	})
// 	return singleton
// }

// func (r *objectRegistry) register(id objectID, t objectInterface) {
// 	r.Lock()
// 	defer r.Unlock()
// 	r.interfaces[id] = t
// }

// func (r *objectRegistry) getInterface(id objectID) (objectInterface, error) {
// 	r.Lock()
// 	defer r.Unlock()
// 	if inf, exists := r.interfaces[id]; !exists {
// 		return nil, fmt.Errorf("interface not registered for this object type yet")
// 	} else {
// 		return inf, nil
// 	}
// }
