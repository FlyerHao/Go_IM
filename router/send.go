package router

import (
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gobwas/ws"
    "github.com/gobwas/ws/wsutil"
    "goIm/contor"
    "goIm/global"
    "goIm/request"
)

type Send struct {

}

func Login(c *gin.Context)  {
        r := c.Request
        w := c.Writer
        conn, _, _, err := ws.UpgradeHTTP(r, w)
        if err != nil {
            // handle error
        }
        go func() {
            defer conn.Close()

            for {
                msg, op, err := wsutil.ReadClientData(conn)
                if err != nil {
                    // handle error
                }
                seq :=request.SendRequest{}
                err = json.Unmarshal(msg, &seq)
                fmt.Println(seq)

                nc := contor.NewConn(seq.UserId,conn)
                global.ConnMan.Add(nc)

                err = wsutil.WriteServerMessage(conn, op, msg)
                if err != nil {
                    // handle error
                }
            }
        }()
}

func LogOut(c *gin.Context)  {
    r := c.Request
    w := c.Writer
    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
        // handle error
    }
    go func() {
        defer conn.Close()

        for {
            msg, op, err := wsutil.ReadClientData(conn)
            if err != nil {
                // handle error
            }
            seq :=request.SendRequest{}
            err = json.Unmarshal(msg, &seq)
            fmt.Println(seq)

            nc := contor.NewConn(seq.UserId,conn)
            global.ConnMan.Add(nc)

            err = wsutil.WriteServerMessage(conn, op, msg)
            if err != nil {
                // handle error
            }
        }
    }()
}

func (s Send) SendMessage(c *gin.Context)  {
    r := c.Request
    w := c.Writer
    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
        // handle error
    }
    go func() {
        defer conn.Close()

        for {
            msg, op, err := wsutil.ReadClientData(conn)
            if err != nil {
                // handle error
            }
            seq :=request.SendRequest{}
            err = json.Unmarshal(msg, &seq)
            fmt.Println(seq)

            nc := contor.NewConn(seq.UserId,conn)
            global.ConnMan.Add(nc)

            err = wsutil.WriteServerMessage(conn, op, msg)
            if err != nil {
                // handle error
            }
        }
    }()
}