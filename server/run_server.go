package server

func RunServer(){
	s := Server{
		Host: "localhost",
		Port: "8000",
	}
	
	s.Start()
}