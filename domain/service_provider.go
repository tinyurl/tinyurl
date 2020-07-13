package domain

type URLStore interface {
	Create(url *URL)
	Update(url *URL)
	GetByOriginURL(originURL string) *URL
	GetByShortPath(shortPath string) *URL
	DropDatabase()
	SaveSenderWorker(sender *SenderWorker)
	UpdateSenderWorker(sender *SenderWorker)
	GetSenderWorker() *SenderWorker
}

// ServiceProvider object hold service which server need
type ServiceProvider struct {
	StoreClient  URLStore
	KeyGenerater KeyGenerater
	GlobalConfig *GlobalConfig
}
