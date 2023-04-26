package client

import (
	"net"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/4/23
//***************************************************

type robot net.TCPConn

//func GetConnection(address string) (conn *robot) {
//	add, err := net.ResolveTCPAddr("tcp", address)
//	t, err := net.DialTCP("tcp", nil, add)
//	if err != nil {
//		log.Warn("client error: %v", err)
//		return
//	}
//	conn = (*robot)(t)
//	buf := bufio.NewReader(conn)
//	util.Go(func() {
//		for {
//			//buf.Peek()
//		}
//	})
//	return
//}
//
//func (r *robot) SendMessage(moduleName, reqMethodName string, p protoiface.MessageV1) {
//	var (
//		data, _ = proto.Marshal(p)
//	)
//	reqName := []byte(fmt.Sprintf("%v.%v", moduleName, reqMethodName))
//	tmp := make([]byte, 0, 6+len(data)+len(reqName))
//	buff := bytes.NewBuffer(tmp)
//	buff.WriteByte(250)
//	buff.WriteByte(byte(rpc.Logic))
//	reqLenByte := make([]byte, 2)
//	binary.BigEndian.PutUint16(reqLenByte, uint16(len(reqName)))
//	buff.Write(reqLenByte)
//	reqSizeLenByte := make([]byte, 2)
//	binary.BigEndian.PutUint16(reqSizeLenByte, uint16(len(data)))
//	buff.Write(reqSizeLenByte)
//	buff.Write(reqName)
//	buff.Write(data)
//	cnt, er := r.Write(buff.Bytes())
//	if er != nil {
//		log.Warn("write len %v module=%v,reqMethod=%v error : %v ", cnt, moduleName, reqMethodName, er)
//		return
//	}
//	log.Info("send logic message : module=%v,reqMethod=%v", moduleName, reqMethodName)
//}
