package contor

import "net"

type Conn struct {
    UserId int64 `json:"user_id"`
    Con net.Conn
}

func NewConn(userId int64,con net.Conn) *Conn {
    return &Conn{
        UserId: userId,
        Con: con,
    }
}

func (cn *Conn)Do()  {

}

type ConnMange struct {
    G1 map[int64]*Conn
}

func (cm *ConnMange) Add(c *Conn)  {
    cm.G1[c.UserId] = c
}

func (cm *ConnMange) Remove(userId int64)  {
    delete(cm.G1,userId)
}

func (cm *ConnMange) GetConn(userId int64)  *Conn{
    conn := cm.G1[userId]
    return conn
}

func Init() *ConnMange {
    cm := &ConnMange{
        G1:make(map[int64]*Conn),
    }
    return cm
}
