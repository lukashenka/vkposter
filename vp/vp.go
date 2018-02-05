package vp

import (
	"os"
	"sync"

	"github.com/lukashenka/vkposter/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

type VkPoster struct {
	FromID        []string
	ToId          string
	RefreshTime   int
	abortCh       chan struct{}
	processChan   chan VpsProcessSuccess
	errorsChan    chan VpsProcessErr
	postersWaitGr *sync.WaitGroup
	vps           []VkPostSingle
}

func InitProcessing() *VkPoster {
	c := config.GetConfig()
	vp := &VkPoster{
		FromID:        c.VkGroupFrom,
		ToId:          c.VkGroupTo,
		postersWaitGr: &sync.WaitGroup{},
	}
	return vp

}

func (vp *VkPoster) Start() {
	c := config.GetConfig()
	log.Info("Start processing")

	vp.processChan = make(chan VpsProcessSuccess, len(c.VkGroupTo))
	vp.errorsChan = make(chan VpsProcessErr, len(c.VkGroupTo))

	vp.vps = make([]VkPostSingle, len(c.VkGroupTo))

	for i, vpf := range vp.FromID {
		vp.postersWaitGr.Add(1)
		vps := VkPostSingle{
			From:        vpf,
			To:          vp.ToId,
			TimeRefresh: c.RefreshTimePerGroup,
			processCh:   vp.processChan,
			errorCh:     vp.errorsChan,
		}
		vp.vps[i] = vps

		go func(i int) {
			defer vp.postersWaitGr.Done()
			vp.vps[i].Start()
		}(i)

	}

	go vp.processListen()
}

func (vp *VkPoster) Stop() {

	for _, vps := range vp.vps {
		go func(vps VkPostSingle) {
			vps.Stop()
		}(vps)

	}
	vp.postersWaitGr.Wait()
	vp.postersWaitGr.Wait()
	close(vp.processChan)
	close(vp.errorsChan)
	log.Infof("All jobs done")

}

func (vp *VkPoster) processListen() {
	for {
		select {
		case process := <-vp.processChan:
			{
				process.Vps.logger.Infof("%s elapsed:%vs", process.Message, process.Elapsed)
			}
		case err := <-vp.errorsChan:
			{
				err.Vps.logger.Errorf("%s", err.err)
			}

		}
	}
}
