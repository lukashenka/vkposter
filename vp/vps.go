package vp

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type VpsProcessSuccess struct {
	Vps     *VkPostSingle
	Message string
	Elapsed time.Duration
}
type VpsProcessErr struct {
	Vps *VkPostSingle
	err error
}

type VpsAbortCh struct {
	GroupId string
}

type VkPostSingle struct {
	From        string
	To          string
	TimeRefresh int
	processCh   chan VpsProcessSuccess
	errorCh     chan VpsProcessErr
	abortCh     chan VpsAbortCh
	logger      *log.Entry
	stop        bool
	mu          *sync.Mutex
	quit        chan struct{}
}

func (vps *VkPostSingle) Start() {
	vps.mu = &sync.Mutex{}
	vps.quit = make(chan struct{}, 0)
	logger := vps.getLogCtx()
	vps.logger = logger
	vps.logger.Infof("Starting job for %s", vps.From)
	vps.startCycle()
}

func (vps *VkPostSingle) startCycle() {
	for {
		select {
		case <-vps.quit:
			vps.logger.Warn("quit")
			return
		case <-time.After(time.Second * time.Duration(vps.TimeRefresh)):
			start := time.Now()
			err := vps.runJob()
			elapsed := time.Since(start)

			if err != nil {
				vps.error(err)
				vps.wait()
				continue
			}
			vps.done(elapsed)

		}
	}

}

func (vps *VkPostSingle) Stop() {
	vps.quit <- struct{}{}
}

func (vps *VkPostSingle) runJob() (err error) {
	randSl := rand.Intn(2) + 1
	vps.logger.Infof("Mock Sleeping %vs", randSl)
	time.Sleep(time.Second * time.Duration(randSl))
	if randSl == 1 {
		return errors.New("mock_error")
	}

	return
}

func (vps *VkPostSingle) done(elapsed time.Duration) {
	vps.processCh <- VpsProcessSuccess{
		Vps:     vps,
		Message: "success!!",
		Elapsed: elapsed,
	}
}
func (vps *VkPostSingle) error(err error) {
	vps.errorCh <- VpsProcessErr{
		Vps: vps,
		err: errors.New("mock_error"),
	}
}
func (vps *VkPostSingle) wait() {
	sleepingTime := time.Duration(vps.TimeRefresh) + time.Second
	vps.logger.Infof("Sleeping time %vs", sleepingTime)
	time.Sleep(sleepingTime)
}

func (vps *VkPostSingle) getLogCtx() *log.Entry {
	ctx := log.WithFields(log.Fields{
		"group_id": vps.From,
	})
	return ctx
}
