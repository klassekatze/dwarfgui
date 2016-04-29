package main

import (
	"log"
	"sync"

	"github.com/BenLubar/arm_ok/dfhack"
	"github.com/BenLubar/arm_ok/dfhack/dfproto"
	"github.com/golang/protobuf/proto"
)

//cherry picked this bit from github.com/BenLubar/arm_ok/armok_web/proxy.go
//unfortunately, that package wouldn't import or I'd have just imported the whole hog and called it from there.

var remoteOnce sync.Once
var Remote *dfhack.Conn
var RemoteVersion *dfproto.StringMessage
var RemoteDFVersion *dfproto.StringMessage

func remote() {
	var err error
	Remote, err = dfhack.Connect()
	if err != nil {
		log.Panicln("remote connect:", err)
	}

	version, _, err := Remote.GetVersion()
	if err != nil {
		log.Panicln("remote GetVersion:", err)
	}
	RemoteVersion = &dfproto.StringMessage{Value: proto.String(version)}

	version, _, err = Remote.GetDFVersion()
	if err != nil {
		log.Panicln("remote GetDFVersion:", err)
	}
	RemoteDFVersion = &dfproto.StringMessage{Value: proto.String(version)}

	_, err = Remote.ResetMapHashes()
	if err != nil {
		log.Panicln("remote ResetMapHashes:", err)
	}
}
