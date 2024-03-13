package ext2mime

import (
	"sync"
)

type Mapping struct {
	mutex sync.RWMutex
	exttomime map[string]string
}

func (receiver *Mapping) Get(fileExtension string) (mimeType string, found bool) {
	if nil == receiver {
		return "", false
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	var exttomime map[string]string = receiver.exttomime
	if len(exttomime) < 1 {
		return "", false
	}

	mimeType, found = exttomime[fileExtension]
	return
}


func (receiver *Mapping) Set(fileExtension string, mimeType string) error {
	if nil == receiver {
		 return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.exttomime {
		receiver.exttomime = map[string]string{}
	}

	receiver.exttomime[fileExtension] = mimeType
	return nil
}
