package engineio

import  (
	"net/http"
)


const engineioPath = "/engine.io/default"

var handleHttpRequest = func (w http.ResponseWriter, r *http.Request){

}

func Attach(){
	http.HandleFunc(engineioPath, handleHttpRequest)
}

//func Attach(mux *ServeMux){
// 	mux.HandlerFunc(engineioPath, handleHttpRequest)
//}