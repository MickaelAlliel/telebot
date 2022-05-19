package decoders

import "net/url"

type URLDecoder url.URL

func (urld *URLDecoder) Decode(value string) error {
	url.Parse(value)
	u, err := url.Parse(value)
	if err != nil {
		return err
	}
	*urld = URLDecoder(*u)
	return nil
}
