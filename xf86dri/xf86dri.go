// Package xf86dri is the X client API for the XFree86-DRI extension.
package xf86dri

/*
	This file was generated by xf86dri.xml on May 11 2012 11:57:19pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the XFree86-DRI extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 11, "XFree86-DRI").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XFree86-DRI could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["XFree86-DRI"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["XFree86-DRI"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XFree86-DRI"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["XFree86-DRI"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XFree86-DRI"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

type DrmClipRect struct {
	X1 int16
	Y1 int16
	X2 int16
	X3 int16
}

// DrmClipRectRead reads a byte slice into a DrmClipRect value.
func DrmClipRectRead(buf []byte, v *DrmClipRect) int {
	b := 0

	v.X1 = int16(xgb.Get16(buf[b:]))
	b += 2

	v.Y1 = int16(xgb.Get16(buf[b:]))
	b += 2

	v.X2 = int16(xgb.Get16(buf[b:]))
	b += 2

	v.X3 = int16(xgb.Get16(buf[b:]))
	b += 2

	return b
}

// DrmClipRectReadList reads a byte slice into a list of DrmClipRect values.
func DrmClipRectReadList(buf []byte, dest []DrmClipRect) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = DrmClipRect{}
		b += DrmClipRectRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a DrmClipRect value to a byte slice.
func (v DrmClipRect) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put16(buf[b:], uint16(v.X1))
	b += 2

	xgb.Put16(buf[b:], uint16(v.Y1))
	b += 2

	xgb.Put16(buf[b:], uint16(v.X2))
	b += 2

	xgb.Put16(buf[b:], uint16(v.X3))
	b += 2

	return buf
}

// DrmClipRectListBytes writes a list of %s(MISSING) values to a byte slice.
func DrmClipRectListBytes(buf []byte, list []DrmClipRect) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn) QueryVersionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn) QueryVersionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DriMajorVersion uint16
	DriMinorVersion uint16
	DriMinorPatch   uint32
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.DriMajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.DriMinorVersion = xgb.Get16(buf[b:])
	b += 2

	v.DriMinorPatch = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryDirectRenderingCapableCookie is a cookie used only for QueryDirectRenderingCapable requests.
type QueryDirectRenderingCapableCookie struct {
	*xgb.Cookie
}

// QueryDirectRenderingCapable sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryDirectRenderingCapableCookie.Reply()
func QueryDirectRenderingCapable(c *xgb.Conn, Screen uint32) QueryDirectRenderingCapableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'QueryDirectRenderingCapable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryDirectRenderingCapableRequest(c, Screen), cookie)
	return QueryDirectRenderingCapableCookie{cookie}
}

// QueryDirectRenderingCapableUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryDirectRenderingCapableUnchecked(c *xgb.Conn, Screen uint32) QueryDirectRenderingCapableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'QueryDirectRenderingCapable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryDirectRenderingCapableRequest(c, Screen), cookie)
	return QueryDirectRenderingCapableCookie{cookie}
}

// QueryDirectRenderingCapableReply represents the data returned from a QueryDirectRenderingCapable request.
type QueryDirectRenderingCapableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	IsCapable bool
}

// Reply blocks and returns the reply data for a QueryDirectRenderingCapable request.
func (cook QueryDirectRenderingCapableCookie) Reply() (*QueryDirectRenderingCapableReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryDirectRenderingCapableReply(buf), nil
}

// queryDirectRenderingCapableReply reads a byte slice into a QueryDirectRenderingCapableReply value.
func queryDirectRenderingCapableReply(buf []byte) *QueryDirectRenderingCapableReply {
	v := new(QueryDirectRenderingCapableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.IsCapable = true
	} else {
		v.IsCapable = false
	}
	b += 1

	return v
}

// Write request to wire for QueryDirectRenderingCapable
// queryDirectRenderingCapableRequest writes a QueryDirectRenderingCapable request to a byte slice.
func queryDirectRenderingCapableRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// OpenConnectionCookie is a cookie used only for OpenConnection requests.
type OpenConnectionCookie struct {
	*xgb.Cookie
}

// OpenConnection sends a checked request.
// If an error occurs, it will be returned with the reply by calling OpenConnectionCookie.Reply()
func OpenConnection(c *xgb.Conn, Screen uint32) OpenConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'OpenConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(openConnectionRequest(c, Screen), cookie)
	return OpenConnectionCookie{cookie}
}

// OpenConnectionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func OpenConnectionUnchecked(c *xgb.Conn, Screen uint32) OpenConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'OpenConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(openConnectionRequest(c, Screen), cookie)
	return OpenConnectionCookie{cookie}
}

// OpenConnectionReply represents the data returned from a OpenConnection request.
type OpenConnectionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	SareaHandleLow  uint32
	SareaHandleHigh uint32
	BusIdLen        uint32
	// padding: 12 bytes
	BusId string // size: xgb.Pad((int(BusIdLen) * 1))
}

// Reply blocks and returns the reply data for a OpenConnection request.
func (cook OpenConnectionCookie) Reply() (*OpenConnectionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return openConnectionReply(buf), nil
}

// openConnectionReply reads a byte slice into a OpenConnectionReply value.
func openConnectionReply(buf []byte) *OpenConnectionReply {
	v := new(OpenConnectionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.SareaHandleLow = xgb.Get32(buf[b:])
	b += 4

	v.SareaHandleHigh = xgb.Get32(buf[b:])
	b += 4

	v.BusIdLen = xgb.Get32(buf[b:])
	b += 4

	b += 12 // padding

	{
		byteString := make([]byte, v.BusIdLen)
		copy(byteString[:v.BusIdLen], buf[b:])
		v.BusId = string(byteString)
		b += xgb.Pad(int(v.BusIdLen))
	}

	return v
}

// Write request to wire for OpenConnection
// openConnectionRequest writes a OpenConnection request to a byte slice.
func openConnectionRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// CloseConnectionCookie is a cookie used only for CloseConnection requests.
type CloseConnectionCookie struct {
	*xgb.Cookie
}

// CloseConnection sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CloseConnection(c *xgb.Conn, Screen uint32) CloseConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CloseConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(closeConnectionRequest(c, Screen), cookie)
	return CloseConnectionCookie{cookie}
}

// CloseConnectionChecked sends a checked request.
// If an error occurs, it can be retrieved using CloseConnectionCookie.Check()
func CloseConnectionChecked(c *xgb.Conn, Screen uint32) CloseConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CloseConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(closeConnectionRequest(c, Screen), cookie)
	return CloseConnectionCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CloseConnectionCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for CloseConnection
// closeConnectionRequest writes a CloseConnection request to a byte slice.
func closeConnectionRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// GetClientDriverNameCookie is a cookie used only for GetClientDriverName requests.
type GetClientDriverNameCookie struct {
	*xgb.Cookie
}

// GetClientDriverName sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetClientDriverNameCookie.Reply()
func GetClientDriverName(c *xgb.Conn, Screen uint32) GetClientDriverNameCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetClientDriverName' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getClientDriverNameRequest(c, Screen), cookie)
	return GetClientDriverNameCookie{cookie}
}

// GetClientDriverNameUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetClientDriverNameUnchecked(c *xgb.Conn, Screen uint32) GetClientDriverNameCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetClientDriverName' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getClientDriverNameRequest(c, Screen), cookie)
	return GetClientDriverNameCookie{cookie}
}

// GetClientDriverNameReply represents the data returned from a GetClientDriverName request.
type GetClientDriverNameReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ClientDriverMajorVersion uint32
	ClientDriverMinorVersion uint32
	ClientDriverPatchVersion uint32
	ClientDriverNameLen      uint32
	// padding: 8 bytes
	ClientDriverName string // size: xgb.Pad((int(ClientDriverNameLen) * 1))
}

// Reply blocks and returns the reply data for a GetClientDriverName request.
func (cook GetClientDriverNameCookie) Reply() (*GetClientDriverNameReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getClientDriverNameReply(buf), nil
}

// getClientDriverNameReply reads a byte slice into a GetClientDriverNameReply value.
func getClientDriverNameReply(buf []byte) *GetClientDriverNameReply {
	v := new(GetClientDriverNameReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ClientDriverMajorVersion = xgb.Get32(buf[b:])
	b += 4

	v.ClientDriverMinorVersion = xgb.Get32(buf[b:])
	b += 4

	v.ClientDriverPatchVersion = xgb.Get32(buf[b:])
	b += 4

	v.ClientDriverNameLen = xgb.Get32(buf[b:])
	b += 4

	b += 8 // padding

	{
		byteString := make([]byte, v.ClientDriverNameLen)
		copy(byteString[:v.ClientDriverNameLen], buf[b:])
		v.ClientDriverName = string(byteString)
		b += xgb.Pad(int(v.ClientDriverNameLen))
	}

	return v
}

// Write request to wire for GetClientDriverName
// getClientDriverNameRequest writes a GetClientDriverName request to a byte slice.
func getClientDriverNameRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// CreateContextCookie is a cookie used only for CreateContext requests.
type CreateContextCookie struct {
	*xgb.Cookie
}

// CreateContext sends a checked request.
// If an error occurs, it will be returned with the reply by calling CreateContextCookie.Reply()
func CreateContext(c *xgb.Conn, Screen uint32, Visual uint32, Context uint32) CreateContextCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CreateContext' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(createContextRequest(c, Screen, Visual, Context), cookie)
	return CreateContextCookie{cookie}
}

// CreateContextUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateContextUnchecked(c *xgb.Conn, Screen uint32, Visual uint32, Context uint32) CreateContextCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CreateContext' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(createContextRequest(c, Screen, Visual, Context), cookie)
	return CreateContextCookie{cookie}
}

// CreateContextReply represents the data returned from a CreateContext request.
type CreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	HwContext uint32
}

// Reply blocks and returns the reply data for a CreateContext request.
func (cook CreateContextCookie) Reply() (*CreateContextReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return createContextReply(buf), nil
}

// createContextReply reads a byte slice into a CreateContextReply value.
func createContextReply(buf []byte) *CreateContextReply {
	v := new(CreateContextReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.HwContext = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(c *xgb.Conn, Screen uint32, Visual uint32, Context uint32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Visual)
	b += 4

	xgb.Put32(buf[b:], Context)
	b += 4

	return buf
}

// DestroyContextCookie is a cookie used only for DestroyContext requests.
type DestroyContextCookie struct {
	*xgb.Cookie
}

// DestroyContext sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DestroyContext(c *xgb.Conn, Screen uint32, Context uint32) DestroyContextCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'DestroyContext' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyContextRequest(c, Screen, Context), cookie)
	return DestroyContextCookie{cookie}
}

// DestroyContextChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroyContextCookie.Check()
func DestroyContextChecked(c *xgb.Conn, Screen uint32, Context uint32) DestroyContextCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'DestroyContext' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyContextRequest(c, Screen, Context), cookie)
	return DestroyContextCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroyContextCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroyContext
// destroyContextRequest writes a DestroyContext request to a byte slice.
func destroyContextRequest(c *xgb.Conn, Screen uint32, Context uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Context)
	b += 4

	return buf
}

// CreateDrawableCookie is a cookie used only for CreateDrawable requests.
type CreateDrawableCookie struct {
	*xgb.Cookie
}

// CreateDrawable sends a checked request.
// If an error occurs, it will be returned with the reply by calling CreateDrawableCookie.Reply()
func CreateDrawable(c *xgb.Conn, Screen uint32, Drawable uint32) CreateDrawableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CreateDrawable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(createDrawableRequest(c, Screen, Drawable), cookie)
	return CreateDrawableCookie{cookie}
}

// CreateDrawableUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateDrawableUnchecked(c *xgb.Conn, Screen uint32, Drawable uint32) CreateDrawableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'CreateDrawable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(createDrawableRequest(c, Screen, Drawable), cookie)
	return CreateDrawableCookie{cookie}
}

// CreateDrawableReply represents the data returned from a CreateDrawable request.
type CreateDrawableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	HwDrawableHandle uint32
}

// Reply blocks and returns the reply data for a CreateDrawable request.
func (cook CreateDrawableCookie) Reply() (*CreateDrawableReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return createDrawableReply(buf), nil
}

// createDrawableReply reads a byte slice into a CreateDrawableReply value.
func createDrawableReply(buf []byte) *CreateDrawableReply {
	v := new(CreateDrawableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.HwDrawableHandle = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for CreateDrawable
// createDrawableRequest writes a CreateDrawable request to a byte slice.
func createDrawableRequest(c *xgb.Conn, Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// DestroyDrawableCookie is a cookie used only for DestroyDrawable requests.
type DestroyDrawableCookie struct {
	*xgb.Cookie
}

// DestroyDrawable sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DestroyDrawable(c *xgb.Conn, Screen uint32, Drawable uint32) DestroyDrawableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'DestroyDrawable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyDrawableRequest(c, Screen, Drawable), cookie)
	return DestroyDrawableCookie{cookie}
}

// DestroyDrawableChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroyDrawableCookie.Check()
func DestroyDrawableChecked(c *xgb.Conn, Screen uint32, Drawable uint32) DestroyDrawableCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'DestroyDrawable' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyDrawableRequest(c, Screen, Drawable), cookie)
	return DestroyDrawableCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroyDrawableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroyDrawable
// destroyDrawableRequest writes a DestroyDrawable request to a byte slice.
func destroyDrawableRequest(c *xgb.Conn, Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// GetDrawableInfoCookie is a cookie used only for GetDrawableInfo requests.
type GetDrawableInfoCookie struct {
	*xgb.Cookie
}

// GetDrawableInfo sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetDrawableInfoCookie.Reply()
func GetDrawableInfo(c *xgb.Conn, Screen uint32, Drawable uint32) GetDrawableInfoCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetDrawableInfo' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getDrawableInfoRequest(c, Screen, Drawable), cookie)
	return GetDrawableInfoCookie{cookie}
}

// GetDrawableInfoUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetDrawableInfoUnchecked(c *xgb.Conn, Screen uint32, Drawable uint32) GetDrawableInfoCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetDrawableInfo' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getDrawableInfoRequest(c, Screen, Drawable), cookie)
	return GetDrawableInfoCookie{cookie}
}

// GetDrawableInfoReply represents the data returned from a GetDrawableInfo request.
type GetDrawableInfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DrawableTableIndex uint32
	DrawableTableStamp uint32
	DrawableOriginX    int16
	DrawableOriginY    int16
	DrawableSizeW      int16
	DrawableSizeH      int16
	NumClipRects       uint32
	BackX              int16
	BackY              int16
	NumBackClipRects   uint32
	ClipRects          []DrmClipRect // size: xgb.Pad((int(NumClipRects) * 8))
	BackClipRects      []DrmClipRect // size: xgb.Pad((int(NumBackClipRects) * 8))
}

// Reply blocks and returns the reply data for a GetDrawableInfo request.
func (cook GetDrawableInfoCookie) Reply() (*GetDrawableInfoReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getDrawableInfoReply(buf), nil
}

// getDrawableInfoReply reads a byte slice into a GetDrawableInfoReply value.
func getDrawableInfoReply(buf []byte) *GetDrawableInfoReply {
	v := new(GetDrawableInfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.DrawableTableIndex = xgb.Get32(buf[b:])
	b += 4

	v.DrawableTableStamp = xgb.Get32(buf[b:])
	b += 4

	v.DrawableOriginX = int16(xgb.Get16(buf[b:]))
	b += 2

	v.DrawableOriginY = int16(xgb.Get16(buf[b:]))
	b += 2

	v.DrawableSizeW = int16(xgb.Get16(buf[b:]))
	b += 2

	v.DrawableSizeH = int16(xgb.Get16(buf[b:]))
	b += 2

	v.NumClipRects = xgb.Get32(buf[b:])
	b += 4

	v.BackX = int16(xgb.Get16(buf[b:]))
	b += 2

	v.BackY = int16(xgb.Get16(buf[b:]))
	b += 2

	v.NumBackClipRects = xgb.Get32(buf[b:])
	b += 4

	v.ClipRects = make([]DrmClipRect, v.NumClipRects)
	b += DrmClipRectReadList(buf[b:], v.ClipRects)

	v.BackClipRects = make([]DrmClipRect, v.NumBackClipRects)
	b += DrmClipRectReadList(buf[b:], v.BackClipRects)

	return v
}

// Write request to wire for GetDrawableInfo
// getDrawableInfoRequest writes a GetDrawableInfo request to a byte slice.
func getDrawableInfoRequest(c *xgb.Conn, Screen uint32, Drawable uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Drawable)
	b += 4

	return buf
}

// GetDeviceInfoCookie is a cookie used only for GetDeviceInfo requests.
type GetDeviceInfoCookie struct {
	*xgb.Cookie
}

// GetDeviceInfo sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetDeviceInfoCookie.Reply()
func GetDeviceInfo(c *xgb.Conn, Screen uint32) GetDeviceInfoCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetDeviceInfo' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getDeviceInfoRequest(c, Screen), cookie)
	return GetDeviceInfoCookie{cookie}
}

// GetDeviceInfoUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetDeviceInfoUnchecked(c *xgb.Conn, Screen uint32) GetDeviceInfoCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'GetDeviceInfo' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getDeviceInfoRequest(c, Screen), cookie)
	return GetDeviceInfoCookie{cookie}
}

// GetDeviceInfoReply represents the data returned from a GetDeviceInfo request.
type GetDeviceInfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	FramebufferHandleLow    uint32
	FramebufferHandleHigh   uint32
	FramebufferOriginOffset uint32
	FramebufferSize         uint32
	FramebufferStride       uint32
	DevicePrivateSize       uint32
	DevicePrivate           []uint32 // size: xgb.Pad((int(DevicePrivateSize) * 4))
}

// Reply blocks and returns the reply data for a GetDeviceInfo request.
func (cook GetDeviceInfoCookie) Reply() (*GetDeviceInfoReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getDeviceInfoReply(buf), nil
}

// getDeviceInfoReply reads a byte slice into a GetDeviceInfoReply value.
func getDeviceInfoReply(buf []byte) *GetDeviceInfoReply {
	v := new(GetDeviceInfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.FramebufferHandleLow = xgb.Get32(buf[b:])
	b += 4

	v.FramebufferHandleHigh = xgb.Get32(buf[b:])
	b += 4

	v.FramebufferOriginOffset = xgb.Get32(buf[b:])
	b += 4

	v.FramebufferSize = xgb.Get32(buf[b:])
	b += 4

	v.FramebufferStride = xgb.Get32(buf[b:])
	b += 4

	v.DevicePrivateSize = xgb.Get32(buf[b:])
	b += 4

	v.DevicePrivate = make([]uint32, v.DevicePrivateSize)
	for i := 0; i < int(v.DevicePrivateSize); i++ {
		v.DevicePrivate[i] = xgb.Get32(buf[b:])
		b += 4
	}
	b = xgb.Pad(b)

	return v
}

// Write request to wire for GetDeviceInfo
// getDeviceInfoRequest writes a GetDeviceInfo request to a byte slice.
func getDeviceInfoRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}

// AuthConnectionCookie is a cookie used only for AuthConnection requests.
type AuthConnectionCookie struct {
	*xgb.Cookie
}

// AuthConnection sends a checked request.
// If an error occurs, it will be returned with the reply by calling AuthConnectionCookie.Reply()
func AuthConnection(c *xgb.Conn, Screen uint32, Magic uint32) AuthConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'AuthConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(authConnectionRequest(c, Screen, Magic), cookie)
	return AuthConnectionCookie{cookie}
}

// AuthConnectionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func AuthConnectionUnchecked(c *xgb.Conn, Screen uint32, Magic uint32) AuthConnectionCookie {
	if _, ok := c.Extensions["XFREE86-DRI"]; !ok {
		panic("Cannot issue request 'AuthConnection' using the uninitialized extension 'XFree86-DRI'. xf86dri.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(authConnectionRequest(c, Screen, Magic), cookie)
	return AuthConnectionCookie{cookie}
}

// AuthConnectionReply represents the data returned from a AuthConnection request.
type AuthConnectionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Authenticated uint32
}

// Reply blocks and returns the reply data for a AuthConnection request.
func (cook AuthConnectionCookie) Reply() (*AuthConnectionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return authConnectionReply(buf), nil
}

// authConnectionReply reads a byte slice into a AuthConnectionReply value.
func authConnectionReply(buf []byte) *AuthConnectionReply {
	v := new(AuthConnectionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Authenticated = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for AuthConnection
// authConnectionRequest writes a AuthConnection request to a byte slice.
func authConnectionRequest(c *xgb.Conn, Screen uint32, Magic uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XFREE86-DRI"]
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	xgb.Put32(buf[b:], Magic)
	b += 4

	return buf
}