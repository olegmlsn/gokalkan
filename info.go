package gokalkan

import "fmt"

func (cli *Client) GetInfo(alias string) (string, error) {
	result, err := cli.kc.X509ExportCertificateFromStore(alias)
	if err != nil {
		return "", fmt.Errorf("get info err: %w", err)
	}
	return result, nil
}
