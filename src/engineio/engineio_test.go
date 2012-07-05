package engineio

import (
    "testing"
    "net/http"
)

func TestAttach(t *testing.T) {

	intercepted := false
	
	handleHttpRequest = func (w http.ResponseWriter, r *http.Request){
		intercepted = true
	}
	Attach()
	go func(){
		http.ListenAndServe(":12345", nil)
	}()
	http.Get("http://localhost:12345/engine.io/default")
	
	if !intercepted {
		t.Errorf("Calling Attach should intercept request to /engine.io/default")
	}
}

