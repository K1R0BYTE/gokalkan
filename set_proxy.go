package gokalkan

import (
	"net/url"

	"github.com/K1R0BYTE/gokalkan/ckalkan"
)

// Включает использование прокси сервера.
func (cli *Client) SetProxyOn(proxyURL string) error {
	url, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	flags := ckalkan.FlagProxyOn
	return cli.kc.SetProxy(flags, url)
}

// Отключает использование прокси сервера.
func (cli *Client) SetProxyOff(proxyURL string) error {
	url, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	flags := ckalkan.FlagProxyOff
	return cli.kc.SetProxy(flags, url)
}
