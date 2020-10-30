/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package daemons

import (
	"context"
	"sync"
	"time"

	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/model"
	"github.com/IBAX-io/go-ibax/packages/transaction"

	log "github.com/sirupsen/logrus"
)

var mutex = sync.Mutex{}
	// poll a base with period
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			if model.DBConn != nil && CheckDB() {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// CheckDB check if installation complete or not
func CheckDB() bool {
	install := &model.Install{}

	err := install.Get()
	if err != nil {
		log.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("getting install")
	}

	if install.Progress == model.ProgressComplete {
		return true
	}

	return false
}

// DBLock locks daemons
func DBLock() {
	mutex.Lock()
}

// DBUnlock unlocks database
func DBUnlock() {
	transaction.CleanCache()
	mutex.Unlock()
}
