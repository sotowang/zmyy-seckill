package ip

import "testing"

func TestFetchIp(t *testing.T) {
	FetchIp()
}
func TestProxyTest(t *testing.T) {
	ip := "59.56.74.51:9999"
	proxyTest(ip, nil)
}
