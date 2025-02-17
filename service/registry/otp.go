package registry

import (
	"errors"
	"fmt"
	"strings"

	"github.com/latolukasz/beeorm"
	"github.com/sarulabs/di"

	"github.com/coretrix/hitrix/pkg/entity"
	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/config"
	"github.com/coretrix/hitrix/service/component/generator"
	"github.com/coretrix/hitrix/service/component/otp"
)

func ServiceProviderOTP(forceProviders ...string) *service.DefinitionGlobal {
	return &service.DefinitionGlobal{
		Name: service.OTPService,
		Build: func(ctn di.Container) (interface{}, error) {
			configService := ctn.Get(service.ConfigService).(config.IConfig)
			generatorService := ctn.Get(service.GeneratorService).(generator.IGenerator)

			providers := make([]otp.IOTPSMSGateway, 0)
			if len(forceProviders) > 0 {
				for _, forceProvider := range forceProviders {
					builderFunc, ok := smsOTPProviderBuilderFactory[forceProvider]
					if !ok {
						return nil, fmt.Errorf("unknown provider: %v", forceProvider)
					}

					provider, err := builderFunc(configService, generatorService, nil)
					if err != nil {
						return nil, err
					}

					providers = append(providers, provider)
				}
			} else {
				ormService := ctn.Get(service.ORMEngineGlobalService).(*beeorm.Engine)

				q := &beeorm.RedisSearchQuery{}
				q.FilterString("Key", "otp_sms_provider")

				settingsEntity := &entity.SettingsEntity{}
				if has := ormService.RedisSearchOne(settingsEntity, q); !has {
					return nil, errors.New("otp_sms_provider not found in settings")
				}

				providersWithPhonePrefixes := strings.Split(settingsEntity.Value, ";")
				for _, providerWithPhonePrefixes := range providersWithPhonePrefixes {
					providerNameWithPhonePrefixes := strings.Split(providerWithPhonePrefixes, ":")
					providerName := providerNameWithPhonePrefixes[0]

					builderFunc, ok := smsOTPProviderBuilderFactory[providerName]
					if !ok {
						return nil, fmt.Errorf("unknown otp_sms_provider: %v", providerName)
					}

					var phonePrefixes []string
					if len(providerNameWithPhonePrefixes) > 1 {
						phonePrefixes = make([]string, 0)
						phonePrefixesSplit := strings.Split(providerNameWithPhonePrefixes[1], ",")
						if len(phonePrefixesSplit) != 0 {
							phonePrefixes = phonePrefixesSplit
						}
					}

					provider, err := builderFunc(configService, generatorService, phonePrefixes)
					if err != nil {
						return nil, err
					}

					providers = append(providers, provider)
				}
			}

			if len(providers) == 0 {
				return nil, errors.New("must provide otp_sms_provider in settings or at least 1 forceProviders")
			}

			retry, ok := configService.Bool("sms.retry")
			if !ok {
				return nil, errors.New("missing sms.retry")
			}

			return otp.NewOTP(retry, providers...), nil
		},
	}
}

var smsOTPProviderBuilderFactory = map[string]func(
	configService config.IConfig,
	generatorService generator.IGenerator,
	phonePrefixes []string,
) (otp.IOTPSMSGateway, error){
	otp.SMSOTPProviderTwilio: twilioSMSOTPProviderBuilder,
	otp.SMSOTPProviderSinch:  sinchSMSOTPProviderBuilder,
	otp.SMSOTPProviderMada:   madaSMSOTPProviderBuilder,
}

func twilioSMSOTPProviderBuilder(configService config.IConfig, _ generator.IGenerator, _ []string) (otp.IOTPSMSGateway, error) {
	sid, ok := configService.String("sms.twilio.sid")
	if !ok {
		return nil, errors.New("missing sms.twilio.sid")
	}

	token, ok := configService.String("sms.twilio.token")

	if !ok {
		return nil, errors.New("missing sms.twilio.token")
	}

	verifySID, _ := configService.String("sms.twilio.verify_sid")

	return otp.NewTwilioSMSOTPProvider(sid, token, verifySID), nil
}

func sinchSMSOTPProviderBuilder(configService config.IConfig, _ generator.IGenerator, _ []string) (otp.IOTPSMSGateway, error) {
	appID, ok := configService.String("sms.sinch.app_id")
	if !ok {
		return nil, errors.New("missing sms.sinch.app_id")
	}

	appSecret, ok := configService.String("sms.sinch.app_secret")
	if !ok {
		return nil, errors.New("missing sms.sinch.app_secret")
	}

	verificationURL, ok := configService.String("sms.sinch.verification_url")
	if !ok {
		return nil, errors.New("missing sms.sinch.verification_url")
	}

	return otp.NewSinchSMSOTPProvider(appID, appSecret, verificationURL), nil
}

func madaSMSOTPProviderBuilder(
	configService config.IConfig,
	generatorService generator.IGenerator,
	phonePrefixes []string,
) (otp.IOTPSMSGateway, error) {
	username, ok := configService.String("sms.mada.username")
	if !ok {
		return nil, errors.New("missing sms.mada.username")
	}

	password, ok := configService.String("sms.mada.password")
	if !ok {
		return nil, errors.New("missing sms.mada.password")
	}

	url, ok := configService.String("sms.mada.url")
	if !ok {
		return nil, errors.New("missing sms.mada.url")
	}

	sourceName, ok := configService.String("sms.mada.source_name")
	if !ok {
		return nil, errors.New("missing sms.mada.source_name")
	}

	var otpLength int

	otpLengthConfig, ok := configService.Int("authentication.otp_length")
	if ok && otpLengthConfig > 0 {
		otpLength = otpLengthConfig
	}

	return otp.NewMadaSMSOTPProvider(username, password, url, sourceName, otpLength, phonePrefixes, generatorService), nil
}
