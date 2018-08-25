package consts

import (
	"errors"
	"sync"

	pb "github.com/lon9/ww/proto"
)

var kinds = map[pb.Kind]string{
	pb.Kind_CITIZEN:  "Citizen",
	pb.Kind_WAREWOLF: "Warewolf",
	pb.Kind_TELLER:   "Fortune teller",
	pb.Kind_KNIGHT:   "Knight",
}

var mKinds = new(sync.Mutex)

var camps = map[pb.Camp]string{
	pb.Camp_GOOD: "Good",
	pb.Camp_EVIL: "Evil",
}

var mCamps = new(sync.Mutex)

// GetKind returns string for the Kind
func GetKind(c pb.Kind) (string, error) {
	mKinds.Lock()
	defer mKinds.Unlock()
	if v, ok := kinds[c]; ok {
		return v, nil
	}
	return "", errors.New("The kind is not defined")
}

// GetCamp returns string for the camp
func GetCamp(c pb.Camp) (string, error) {
	mCamps.Lock()
	defer mCamps.Unlock()
	if v, ok := camps[c]; ok {
		return v, nil
	}
	return "", errors.New("The camp is not defined.")
}
