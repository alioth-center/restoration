package service

import (
	"context"
	"encoding/json"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/network/rpc"
	"github.com/alioth-center/infrastructure/trace"
	"github.com/alioth-center/infrastructure/utils/generate"
	"github.com/alioth-center/infrastructure/utils/timezone"
	"github.com/alioth-center/restoration/proto/pbg"
)

func NewRestorationService(log logger.Logger) *RestorationService {
	s := &RestorationService{log: log}
	s.chain = rpc.NewChain(s.RestorationRequestPreprocess, s.RestorationLoggingHandler)
	return s
}

type RestorationService struct {
	log   logger.Logger
	chain rpc.Chain[*pbg.RestorationCollectionRequest, *pbg.RestorationCollectionResponse]

	pbg.UnimplementedRestorationServiceServer
}

func (s RestorationService) RestorationCollection(ctx context.Context, request *pbg.RestorationCollectionRequest) (response *pbg.RestorationCollectionResponse, err error) {
	return rpc.ExecuteChain(s.RestorationCollectionHandlerChain(), ctx, request, response)
}

func (s RestorationService) RestorationCollectionHandlerChain() rpc.Chain[*pbg.RestorationCollectionRequest, *pbg.RestorationCollectionResponse] {
	return s.chain
}

func (s RestorationService) RestorationRequestPreprocess(ctx *rpc.Context[*pbg.RestorationCollectionRequest, *pbg.RestorationCollectionResponse]) {
	request := ctx.GetRequest()

	// check level, if level incorrect, set level to info
	if _, exist := levels[logger.Level(request.Level)]; !exist {
		request.Level = string(logger.LevelInfo)
	}

	// check trace id, if trace id is empty, generate a new one
	if request.TraceId == "" {
		request.TraceId = generate.UUID()
	}

	// check identification, if identification is empty, use its ip, if ip is empty, set identification to unknown
	if request.Identification == "" {
		ip := trace.GetClientIPFromPeer(ctx.GetContext())
		if ip == "" {
			request.Identification = "unknown"
		} else {
			request.Identification = ip
		}
	}

	// set request to context
	ctx.SetRequest(request)
}

func (s RestorationService) RestorationLoggingHandler(ctx *rpc.Context[*pbg.RestorationCollectionRequest, *pbg.RestorationCollectionResponse]) {
	request := ctx.GetRequest()

	// build logger fields
	entry := &logger.Entry{
		File:    request.File,
		Level:   request.Level,
		Service: request.Service,
		TraceID: request.TraceId,
		Message: request.Message,
	}

	// try to unmarshal data
	var data any
	unmarshalErr := json.Unmarshal(request.Data, &data)
	if unmarshalErr != nil {
		// cannot unmarshal data, use raw message
		entry.Data = string(request.Data)
	} else {
		// unmarshal data success, use unmarshalled data
		entry.Data = data
	}

	level := logger.Level(request.Level)
	field := logger.NewFieldsFromEntry(entry).
		WithCallTime(timezone.UnixMilliTimestampInZero(request.CallTime)).
		WithField(restorationIdentificationKey, request.Identification)
	for k, v := range request.Extra {
		if k == restorationIdentificationKey {
			// identification is a reserved key, skip it
			continue
		}

		var fieldValue any
		unmarshalErr := json.Unmarshal(v, &fieldValue)
		if unmarshalErr != nil {
			// unmarshal extra field success, use unmarshalled data
			field = field.WithField(k, string(v))
		} else {
			// cannot unmarshal extra field, use raw message
			field = field.WithField(k, fieldValue)
		}
	}

	// execute log collection
	s.log.Log(level, field)

	ctx.SetResponse(&pbg.RestorationCollectionResponse{})
}
