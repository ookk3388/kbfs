// Auto-generated by avdl-compiler v1.3.26 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/notify.avdl

package stellar1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PaymentNotificationArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
	PaymentID PaymentID `codec:"paymentID" json:"paymentID"`
}

type PaymentStatusNotificationArg struct {
	AccountID AccountID `codec:"accountID" json:"accountID"`
	PaymentID PaymentID `codec:"paymentID" json:"paymentID"`
}

type RequestStatusNotificationArg struct {
	ReqID KeybaseRequestID `codec:"reqID" json:"reqID"`
}

type NotifyInterface interface {
	PaymentNotification(context.Context, PaymentNotificationArg) error
	PaymentStatusNotification(context.Context, PaymentStatusNotificationArg) error
	RequestStatusNotification(context.Context, KeybaseRequestID) error
}

func NotifyProtocol(i NotifyInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "stellar.1.notify",
		Methods: map[string]rpc.ServeHandlerDescription{
			"paymentNotification": {
				MakeArg: func() interface{} {
					var ret [1]PaymentNotificationArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PaymentNotificationArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PaymentNotificationArg)(nil), args)
						return
					}
					err = i.PaymentNotification(ctx, typedArgs[0])
					return
				},
			},
			"paymentStatusNotification": {
				MakeArg: func() interface{} {
					var ret [1]PaymentStatusNotificationArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]PaymentStatusNotificationArg)
					if !ok {
						err = rpc.NewTypeError((*[1]PaymentStatusNotificationArg)(nil), args)
						return
					}
					err = i.PaymentStatusNotification(ctx, typedArgs[0])
					return
				},
			},
			"requestStatusNotification": {
				MakeArg: func() interface{} {
					var ret [1]RequestStatusNotificationArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]RequestStatusNotificationArg)
					if !ok {
						err = rpc.NewTypeError((*[1]RequestStatusNotificationArg)(nil), args)
						return
					}
					err = i.RequestStatusNotification(ctx, typedArgs[0].ReqID)
					return
				},
			},
		},
	}
}

type NotifyClient struct {
	Cli rpc.GenericClient
}

func (c NotifyClient) PaymentNotification(ctx context.Context, __arg PaymentNotificationArg) (err error) {
	err = c.Cli.Notify(ctx, "stellar.1.notify.paymentNotification", []interface{}{__arg})
	return
}

func (c NotifyClient) PaymentStatusNotification(ctx context.Context, __arg PaymentStatusNotificationArg) (err error) {
	err = c.Cli.Notify(ctx, "stellar.1.notify.paymentStatusNotification", []interface{}{__arg})
	return
}

func (c NotifyClient) RequestStatusNotification(ctx context.Context, reqID KeybaseRequestID) (err error) {
	__arg := RequestStatusNotificationArg{ReqID: reqID}
	err = c.Cli.Notify(ctx, "stellar.1.notify.requestStatusNotification", []interface{}{__arg})
	return
}
