package custonlog

import (
	"time"

	"github.com/fatih/color"
)

//InfoLog is the funcion for log
func InfoLog(Processo, Message string) {

	c := color.New(color.FgWhite)

	c.Printf("[%v ] - INFO \t - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), Processo, Message)

}

//WarnLog is the funcion for log
func WarnLog(Processo, Message string) {

	c := color.New(color.FgWhite)
	c.Printf("[%v ] - WARNING \t - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), Processo, Message)

}

//ErrorLog is the funcion for log
func ErrorLog(Processo, Message string) {

	c := color.New(color.FgRed)
	c.Printf("[%v ] - ERROR \t - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), Processo, Message)

}

//DebugLog is the funcion for log
func DebugLog(Processo, Message string) {

	c := color.New(color.FgRed)
	c.Printf("[%v ] - DEBUG \t - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), Processo, Message)

}
