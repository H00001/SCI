package service

import (
	"context"
	"fmt"
)
import "gunplan.top/SCI/protocol"

type VerifyImpl struct {

}

func (v *VerifyImpl)AcquireMagic(cc context.Context,cv *connection.MagicInbound) (*connection.MagicOutbound	, error){
	fmt.Printf("UUID:%s\n",cv.Uuid)
	return &connection.MagicOutbound{ Magic:cv.Uuid+"-10-24-10-24"},nil
}

func (v *VerifyImpl)AcquireCalc(cc context.Context,cv *connection.Base64Result) (*connection.VerifyResult, error){
	return &connection.VerifyResult{Token:"00000-00000-0000-00000-00000"},nil
}

