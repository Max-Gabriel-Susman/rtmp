package rtmp

import "time"

type TlsCertificateLoaderConfig struct {
	CertificatePath   string
	KeyPath           string
	CheckReloadPeriod time.Duration
	OnReload          func()
	OnError           func(err error)
}
