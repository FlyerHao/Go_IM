package global

import "goIm/contor"

var ConnMan *contor.ConnMange

func init()  {
    ConnMan = contor.Init()
}