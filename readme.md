#### ListenAndServer() is a Blocking call. It makes sense to block because otherwise the program would run to end and the server would be closed. If you need it not to block, run it in a go routine & use channels for communication and SIGINT signals for closing server etc...

##### 	server := &http.Server{
		Addr:    ":8080",
		Handler: http.Handler, // we can also use chi.NewRouter()
	}
