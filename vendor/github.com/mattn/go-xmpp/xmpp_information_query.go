package xmpp

import (
	"fmt"
	"strconv"
)

const IQTypeGet = "get"
const IQTypeSet = "set"
const IQTypeResult = "result"

func (c *Client) Discovery() (string, error) {
	const namespace = "http://jabber.org/protocol/disco#items"
	// use getCookie for a pseudo random id.
	reqID := strconv.FormatUint(uint64(getCookie()), 10)
	return c.RawInformationQuery(c.jid, c.domain, reqID, IQTypeGet, namespace, "")
}

// RawInformationQuery sends an information query request to the server.
func (c *Client) RawInformationQuery(from, to, id, iqType, requestNamespace, body string) (string, error) {
	const xmlIQ = "<iq from='%s' to='%s' id='%s' type='%s'><query xmlns='%s'>%s</query></iq>"
	_, err := fmt.Fprintf(c.conn, xmlIQ, xmlEscape(from), xmlEscape(to), id, iqType, requestNamespace, body)
	return id, err
}
