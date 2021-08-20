package armstatus

import (
	"encoding/json"
	"io/ioutil"
)

type AStatus struct {
	State string
}

func WriteStatus(astat AStatus) {
	file, _ := json.MarshalIndent(astat, "", " ")
	_ = ioutil.WriteFile("/app/state/armstate.json", file, 0644)
}

//func ReadStatus() {


//}
