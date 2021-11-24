package core

func Boot() {
	initViper()
	initLog()
	InitRPCService()
	initRoute()
}
