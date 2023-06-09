// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	thrift "thrift/src"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type RpcService interface {
	// Parameters:
	//  - CallTime
	//  - FunCode
	//  - ParamMap
	FunCall(ctx context.Context, callTime int64, funCode string, paramMap map[string]string) (r []string, err error)
}

type RpcServiceClient struct {
	c thrift.TClient
}

func NewRpcServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RpcServiceClient {
	return &RpcServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewRpcServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RpcServiceClient {
	return &RpcServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewRpcServiceClient(c thrift.TClient) *RpcServiceClient {
	return &RpcServiceClient{
		c: c,
	}
}

func (p *RpcServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - CallTime
//  - FunCode
//  - ParamMap
func (p *RpcServiceClient) FunCall(ctx context.Context, callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	var _args0 RpcServiceFunCallArgs
	_args0.CallTime = callTime
	_args0.FunCode = funCode
	_args0.ParamMap = paramMap
	var _result1 RpcServiceFunCallResult
	if err = p.Client_().Call(ctx, "funCall", &_args0, &_result1); err != nil {
		return
	}
	return _result1.GetSuccess(), nil
}

type RpcServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      RpcService
}

func (p *RpcServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RpcServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RpcServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRpcServiceProcessor(handler RpcService) *RpcServiceProcessor {

	self2 := &RpcServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["funCall"] = &rpcServiceProcessorFunCall{handler: handler}
	return self2
}

func (p *RpcServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3

}

type rpcServiceProcessorFunCall struct {
	handler RpcService
}

func (p *rpcServiceProcessorFunCall) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RpcServiceFunCallArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("funCall", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	result := RpcServiceFunCallResult{}
	var retval []string
	var err2 error
	if retval, err2 = p.handler.FunCall(ctx, args.CallTime, args.FunCode, args.ParamMap); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing funCall: "+err2.Error())
		oprot.WriteMessageBegin("funCall", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("funCall", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - CallTime
//  - FunCode
//  - ParamMap
type RpcServiceFunCallArgs struct {
	CallTime int64             `thrift:"callTime,1" db:"callTime" json:"callTime"`
	FunCode  string            `thrift:"funCode,2" db:"funCode" json:"funCode"`
	ParamMap map[string]string `thrift:"paramMap,3" db:"paramMap" json:"paramMap"`
}

func NewRpcServiceFunCallArgs() *RpcServiceFunCallArgs {
	return &RpcServiceFunCallArgs{}
}

func (p *RpcServiceFunCallArgs) GetCallTime() int64 {
	return p.CallTime
}

func (p *RpcServiceFunCallArgs) GetFunCode() string {
	return p.FunCode
}

func (p *RpcServiceFunCallArgs) GetParamMap() map[string]string {
	return p.ParamMap
}
func (p *RpcServiceFunCallArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.CallTime = v
	}
	return nil
}

func (p *RpcServiceFunCallArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.FunCode = v
	}
	return nil
}

func (p *RpcServiceFunCallArgs) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.ParamMap = tMap
	for i := 0; i < size; i++ {
		var _key4 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key4 = v
		}
		var _val5 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val5 = v
		}
		p.ParamMap[_key4] = _val5
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("funCall_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("callTime", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:callTime: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.CallTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.callTime (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:callTime: ", p), err)
	}
	return err
}

func (p *RpcServiceFunCallArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("funCode", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:funCode: ", p), err)
	}
	if err := oprot.WriteString(string(p.FunCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.funCode (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:funCode: ", p), err)
	}
	return err
}

func (p *RpcServiceFunCallArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("paramMap", thrift.MAP, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:paramMap: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.ParamMap)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.ParamMap {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:paramMap: ", p), err)
	}
	return err
}

func (p *RpcServiceFunCallArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceFunCallArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RpcServiceFunCallResult struct {
	Success []string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRpcServiceFunCallResult() *RpcServiceFunCallResult {
	return &RpcServiceFunCallResult{}
}

var RpcServiceFunCallResult_Success_DEFAULT []string

func (p *RpcServiceFunCallResult) GetSuccess() []string {
	return p.Success
}
func (p *RpcServiceFunCallResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RpcServiceFunCallResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField0(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		var _elem6 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem6 = v
		}
		p.Success = append(p.Success, _elem6)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("funCall_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RpcServiceFunCallResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceFunCallResult(%+v)", *p)
}
