package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// Close
// 实现GobCodec结构体的接口
func (g *GobCodec) Close() error {
	//TODO implement me
	return g.conn.Close()
}

// ReadHeader
// 将header解码到本地
func (g *GobCodec) ReadHeader(header *Header) error {
	return g.dec.Decode(header)
}

// ReadBody
// 将body解码到本地
func (g *GobCodec) ReadBody(body interface{}) error {
	return g.dec.Decode(body)
}

func (g *GobCodec) Write(header *Header, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

var _ Codec = (*GobCodec)(nil)

// NewGodCodec
// 初始化GobCodec
func NewGodCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn,
		buf,
		gob.NewDecoder(conn),
		gob.NewEncoder(buf),
	}
}
