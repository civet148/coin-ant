package api

type ManagerApi interface {
	//run service loop
	Run() (err error)

	//close service
	Close()
}
