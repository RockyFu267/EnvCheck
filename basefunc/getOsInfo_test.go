package basefunc

import (
	"fmt"
	"testing"
)

func Test_GetOsInfo(t *testing.T) {

	res := GetOsInfo()
	fmt.Println(res)

}
