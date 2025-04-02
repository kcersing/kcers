package config

type ServerConfig struct {
	Name           string           `mapstructure:"Name" json:"Name"`
	Host           string           `mapstructure:"Host" json:"Host"`
	Port           int              `mapstructure:"Port" json:"Port"`
	Timeout        int              `mapstructure:"Timeout" json:"Timeout"`
	IsProd         bool             `mapstructure:"IsProd" json:"IsProd"`
	Domain         string           `mapstructure:"Domain" json:"Domain"`
	MySQLInfo      MySQLConfig      `mapstructure:"MySQL" json:"MySQL"`
	PostgreSQLInfo PostgreSQLConfig `mapstructure:"PostgreSQL" json:"PostgreSQL"`
	Captcha        Captcha          `mapstructure:"Captcha" json:"Captcha"`
	Auth           Auth             `mapstructure:"Auth" json:"Auth"`
	Redis          Redis            `mapstructure:"Redis" json:"Redis"`
	Casbin         CasbinConf       `mapstructure:"Casbin" json:"Casbin"`
	Payment        Payment          `mapstructure:"Payment" json:"Payment"`
	Minio          Minio            `mapstructure:"Minio" json:"Minio"`
	DouYin         DouYin           `mapstructure:"DouYin" json:"DouYin"`
	Wechat         Wechat           `mapstructure:"Wechat" json:"Wechat"`
	Aliyun         Aliyun           `mapstructure:"Aliyun" json:"Aliyun"`
}

type MySQLConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type PostgreSQLConfig struct {
	Host string `mapstructure:"Host" json:"Host"`
	Salt string `mapstructure:"Salt" json:"Salt"`
}
type Captcha struct {
	KeyLong   int `mapstructure:"KeyLong" json:"KeyLong"`
	ImgWidth  int `mapstructure:"ImgWidth" json:"ImgWidth"`
	ImgHeight int `mapstructure:"ImgHeight" json:"ImgHeight"`
}
type Auth struct {
	OAuthKey     string `mapstructure:"OAuthKey" json:"OAuthKey"`
	AccessSecret string `mapstructure:"AccessSecret" json:"AccessSecret"`
	AccessExpire int    `mapstructure:"AccessExpire" json:"AccessExpire"`
}

type Redis struct {
	Host     string `mapstructure:"Host" json:"Host"`
	Password string `mapstructure:"Password" json:"Password"`
}

type CasbinConf struct {
	ModelText string `mapstructure:"ModelText" json:"ModelText"`
}

type Payment struct {
	Alipay    *AliPay    `mapstructure:"Alipay" yaml:"Alipay"`
	Wechatpay *Wechatpay `mapstructure:"Wechatpay" json:"Wechatpay"`
}

type Wechatpay struct {
	Appid             string `mapstructure:"ModelText" yaml:"appid"`
	MchId             string `mapstructure:"ModelText" yaml:"mch_id"`
	ApiKey            string `mapstructure:"ModelText" yaml:"api_key"`
	ApiV3Key          string `mapstructure:"ModelText" yaml:"api_v3_key"`
	CertFileContent   string `mapstructure:"ModelText" yaml:"cert_file_content"`
	KeyFileContent    string `mapstructure:"ModelText" yaml:"key_file_content"`
	Pkcs12FileContent string `mapstructure:"ModelText" yaml:"pkcs12_file_content"`
	SerialNo          string `mapstructure:"ModelText" yaml:"serial_no"`
	NotifyUrl         string `mapstructure:"ModelText" yaml:"notify_url"`

	CertificateKeyPath string `mapstructure:"certificate_key_path" yaml:"certificate_key_path"`
	WechatPaySerial    string `mapstructure:"wechat_pay_serial" yaml:"wechat_pay_serial"`
	RSAPublicKeyPath   string `mapstructure:"rsa_public_key_path" yaml:"rsa_public_key_path"`
	SubMchID           string `mapstructure:"sub_mch_id" yaml:"sub_mch_id"`
	SubAppID           string `mapstructure:"sub_app_id" yaml:"sub_app_id"`
}

type AliPay struct {
	Appid                   string `mapstructure:"ModelText" yaml:"appid"`
	PrivateKey              string `mapstructure:"ModelText" yaml:"private_key"`
	AppPublicCertContent    string `mapstructure:"ModelText" yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `mapstructure:"ModelText" yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `mapstructure:"ModelText" yaml:"alipay_public_cert_content"`
}

type Minio struct {
	EndPoint        string `mapstructure:"EndPoint" yaml:"EndPoint"`
	AccessKeyID     string `mapstructure:"AccessKeyID" yaml:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey" yaml:"SecretAccessKey"`
	UseSSL          bool   `mapstructure:"UseSSL" yaml:"UseSSL"`

	VideoBucketName string `mapstructure:"VideoBucketName" yaml:"VideoBucketName"`
	ImgBucketName   string `mapstructure:"ImgBucketName" yaml:"ImgBucketName"`

	Url string `mapstructure:"Url" yaml:"Url"`
}
type DouYin struct {
	AppId     string `mapstructure:"AppId" yaml:"AppId"`
	AppSecret string `mapstructure:"AppSecret" yaml:"AppSecret"`
	AccountId string `mapstructure:"AccountId" yaml:"AccountId"`
}

type Wechat struct {
	WXMiniProgram MiniProgram `mapstructure:"MiniProgram" yaml:"MiniProgram"`
}
type MiniProgram struct {
	AppID  string `mapstructure:"appid" yaml:"appid"`
	Secret string `mapstructure:"secret" yaml:"secret"`
	Token  string `mapstructure:"token" yaml:"token"`
	AESKey string `mapstructure:"aes_key" yaml:"aes_key"`

	AppKey  string `mapstructure:"app_key" yaml:"app_key"`
	OfferID string `mapstructure:"offer_id" yaml:"offer_id"`
}

type Aliyun struct {
	Access Access `mapstructure:"Access" yaml:"Access"`
	Sms    Sms    `mapstructure:"Sms" yaml:"Sms"`
}
type Access struct {
	AccessKeyId     string `mapstructure:"AccessKeyId" yaml:"AccessKeyId"`
	AccessKeySecret string `mapstructure:"AccessKeySecret" yaml:"AccessKeySecret"`
}
type Sms struct {
	Captcha SmsTemplate `mapstructure:"Captcha" yaml:"Captcha"`
}
type SmsTemplate struct {
	SignName     string `mapstructure:"SignName" yaml:"SignName"`
	TemplateCode string `mapstructure:"TemplateCode" yaml:"TemplateCode"`
}
