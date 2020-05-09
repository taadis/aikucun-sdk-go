package sdk

type Options struct {
	Url       string
	AppId     string
	AppSecret string
}

type Option func(*Options)

// WithUrl set default url for the client
func WithUrl(url string) Option {
	return func(options *Options) {
		options.Url = url
	}
}

// WithAppId set default url for the client
func WithAppId(appId string) Option {
	return func(options *Options) {
		options.AppId = appId
	}
}

// WithAppSecret set default url for the client
func WithAppSecret(appSecret string) Option {
	return func(options *Options) {
		options.AppSecret = appSecret
	}
}
