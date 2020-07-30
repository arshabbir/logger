package logger

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	//"github.com/arshabbir/logger/logger"
)

func BenchmarkLogger(b *testing.B) {

	log := NewLogger(".\\logs", 2*1024, false)

	for i := 0; i <= b.N; i++ {
		logMesg := fmt.Sprintf("%d,,%s", rand.Intn(10), "this isfhhgdgdhgd ddg dgdhgdg dgd gdghdgdg dgd gdghdkghd gkdg dg dgdhgd gdg dgdhgkdhgdk gd gdghdgkdghdkg dg dgdkgdhgkdhkdhdgd   a ERROR message")
		log.LogString(ERROR, http.StatusOK, logMesg)
		//time.Sleep(time.Millisecond * 1)
	}
}
