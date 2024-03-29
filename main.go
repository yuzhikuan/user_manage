package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/e421083458/golang_common/lib"
	"github.com/yuzhikuan/user_manage/public"
	"github.com/yuzhikuan/user_manage/router"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql"})
	defer lib.Destroy()

	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
