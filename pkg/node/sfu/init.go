package sfu

import (
	nprotoo "github.com/cloudwebrtc/nats-protoo"
	"github.com/pion/ion/pkg/log"
	"github.com/pion/ion/pkg/proto"
	"github.com/pion/ion/pkg/rtc"
	"github.com/pion/ion/pkg/util"
)

var (
	dc          = "default"
	protoo      *nprotoo.NatsProtoo
	broadcaster *nprotoo.Broadcaster
)

// Init func
func Init(dcID, rpcID, eventID, natsURL string) {
	dc = dcID
	protoo = nprotoo.NewNatsProtoo(natsURL)
	broadcaster = protoo.NewBroadcaster(eventID)
	handleRequest(rpcID)
	checkRTC()
}

// checkRTC send `stream-remove` msg to islb when some pub has been cleaned
func checkRTC() {
	log.Infof("SFU.checkRTC start")
	go func() {
		for mid := range rtc.CleanChannel {
			broadcaster.Say(proto.IslbOnStreamRemove, util.Map("mid", mid))
		}
	}()
}