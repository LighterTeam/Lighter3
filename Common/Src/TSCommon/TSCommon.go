package TSCommon

import (
	"fmt"
	"runtime"
)

func init () {
	fmt.Println("TSCommon init");
	runtime.GOMAXPROCS(runtime.NumCPU())	
}