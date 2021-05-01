package sms

import (
	"fmt"

	"github.com/coretrix/hitrix/pkg/entity"
	"github.com/coretrix/hitrix/service/component/clock"
	"github.com/latolukasz/orm"
)

type ISender interface {
	SendOTPSMS(*OTP) error
	SendOTPCallout(*OTP) error
	SendMessage(*Message) error
}

type Sender struct {
	OrmService     *orm.Engine
	Clock          clock.Clock
	GatewayFactory map[string]Gateway
}

func (s *Sender) SendOTPSMS(otp *OTP) error {
	primaryProvider := otp.Provider.Primary
	secondaryProvider := otp.Provider.Secondary

	primaryGateway, ok := s.GatewayFactory[primaryProvider]
	if !ok {
		return fmt.Errorf("primary provider not supported")
	}

	otp.OTP = fmt.Sprintf(otp.Template, otp.OTP)

	smsTrackerEntity := &entity.SmsTrackerEntity{
		To:                 otp.Number,
		Text:               otp.OTP,
		FromPrimaryGateway: primaryProvider,
		SentAt:             s.Clock.Now(),
	}

	trySecondaryProvider := false
	status, err := primaryGateway.SendOTPSMS(otp)
	if err != nil {
		trySecondaryProvider = true
		smsTrackerEntity.PrimaryGatewayError = err.Error()
	}

	if trySecondaryProvider {
		secondaryGateway, ok := s.GatewayFactory[secondaryProvider]
		if !ok {
			return fmt.Errorf("secondary provider not supported")
		}

		smsTrackerEntity.FromSecondaryGateway = secondaryProvider

		status, err = secondaryGateway.SendOTPSMS(otp)
		if err != nil {
			smsTrackerEntity.SecondaryGatewayError = err.Error()
		}
	}

	smsTrackerEntity.Status = status
	s.OrmService.Flush(smsTrackerEntity)
	return nil
}

func (s *Sender) SendOTPCallout(otp *OTP) error {
	primaryProvider := otp.Provider.Primary
	secondaryProvider := otp.Provider.Secondary

	primaryGateway, ok := s.GatewayFactory[primaryProvider]
	if !ok {
		return fmt.Errorf("primary provider not supported")
	}

	otp.OTP = fmt.Sprintf(otp.Template, otp.OTP)

	smsTrackerEntity := &entity.SmsTrackerEntity{
		To:                 otp.Number,
		Text:               otp.OTP,
		FromPrimaryGateway: primaryProvider,
		SentAt:             s.Clock.Now(),
	}

	trySecondaryProvider := false
	status, err := primaryGateway.SendOTPCallout(otp)
	if err != nil {
		trySecondaryProvider = true
		smsTrackerEntity.PrimaryGatewayError = err.Error()
	}

	if trySecondaryProvider {
		secondaryGateway, ok := s.GatewayFactory[secondaryProvider]
		if !ok {
			return fmt.Errorf("secondary provider not supported")
		}

		smsTrackerEntity.FromSecondaryGateway = secondaryProvider

		status, err = secondaryGateway.SendOTPCallout(otp)
		if err != nil {
			smsTrackerEntity.SecondaryGatewayError = err.Error()
		}
	}

	smsTrackerEntity.Status = status
	s.OrmService.Flush(smsTrackerEntity)
	return nil
}

func (s *Sender) SendMessage(message *Message) error {
	primaryProvider := message.Provider.Primary
	secondaryProvider := message.Provider.Secondary

	primaryGateway, ok := s.GatewayFactory[primaryProvider]
	if !ok {
		return fmt.Errorf("primary provider not supported")
	}

	smsTrackerEntity := &entity.SmsTrackerEntity{
		To:                 message.Number,
		Text:               message.Text,
		FromPrimaryGateway: primaryProvider,
		SentAt:             s.Clock.Now(),
	}

	trySecondaryProvider := false
	status, err := primaryGateway.SendSMSMessage(message)
	if err != nil {
		trySecondaryProvider = true
		smsTrackerEntity.PrimaryGatewayError = err.Error()
	}

	if trySecondaryProvider {
		secondaryGateway, ok := s.GatewayFactory[secondaryProvider]
		if !ok {
			return fmt.Errorf("secondary provider not supported")
		}

		smsTrackerEntity.FromSecondaryGateway = secondaryProvider

		status, err = secondaryGateway.SendSMSMessage(message)
		if err != nil {
			smsTrackerEntity.SecondaryGatewayError = err.Error()
		}
	}

	smsTrackerEntity.Status = status
	s.OrmService.Flush(smsTrackerEntity)
	return nil
}
