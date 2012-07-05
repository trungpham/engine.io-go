package engineio

import (

)

type Socket struct{

}

func (s *Socket) Send(string){

}

func (s *Socket) Close(){

}

func (s *Socket) OnMessage(func(data []byte)){


}

func (s *Socket) OnClose(func()){

}