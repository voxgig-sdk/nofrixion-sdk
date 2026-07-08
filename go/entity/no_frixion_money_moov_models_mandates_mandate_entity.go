package entity

import (
	"github.com/voxgig-sdk/nofrixion-sdk/go/core"

	vs "github.com/voxgig-sdk/nofrixion-sdk/go/utility/struct"
)

type NoFrixionMoneyMoovModelsMandatesMandateEntity struct {
	name    string
	client  *core.NofrixionSDK
	utility *core.Utility
	entopts map[string]any
	data    map[string]any
	match   map[string]any
	entctx  *core.Context
}

func NewNoFrixionMoneyMoovModelsMandatesMandateEntity(client *core.NofrixionSDK, entopts map[string]any) *NoFrixionMoneyMoovModelsMandatesMandateEntity {
	if entopts == nil {
		entopts = map[string]any{}
	}
	if _, ok := entopts["active"]; !ok {
		entopts["active"] = true
	} else if entopts["active"] == false {
		// keep false
	} else {
		entopts["active"] = true
	}

	e := &NoFrixionMoneyMoovModelsMandatesMandateEntity{
		name:    "no_frixion_money_moov_models_mandates_mandate",
		client:  client,
		utility: client.GetUtility(),
		entopts: entopts,
		data:    map[string]any{},
		match:   map[string]any{},
	}

	e.entctx = e.utility.MakeContext(map[string]any{
		"entity":  e,
		"entopts": entopts,
	}, client.GetRootCtx())

	e.utility.FeatureHook(e.entctx, "PostConstructEntity")

	return e
}

func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) GetName() string { return e.name }

func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Make() core.Entity {
	opts := map[string]any{}
	for k, v := range e.entopts {
		opts[k] = v
	}
	return NewNoFrixionMoneyMoovModelsMandatesMandateEntity(e.client, opts)
}

func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Data(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.data = core.ToMapAny(vs.Clone(args[0]))
		if e.data == nil {
			e.data = map[string]any{}
		}
		e.utility.FeatureHook(e.entctx, "SetData")
	}

	e.utility.FeatureHook(e.entctx, "GetData")
	out := vs.Clone(e.data)
	return out
}

func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Match(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.match = core.ToMapAny(vs.Clone(args[0]))
		if e.match == nil {
			e.match = map[string]any{}
		}
		e.utility.FeatureHook(e.entctx, "SetMatch")
	}

	e.utility.FeatureHook(e.entctx, "GetMatch")
	out := vs.Clone(e.match)
	return out
}

// DataTyped is the statically-typed accessor for this entity's data. With no
// argument it returns the current data as an NoFrixionMoneyMoovModelsMandatesMandate; with an argument it
// sets the data and returns the stored value. It delegates to the untyped Data
// (identical runtime) and converts at the typed boundary.
func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) DataTyped(data ...NoFrixionMoneyMoovModelsMandatesMandate) NoFrixionMoneyMoovModelsMandatesMandate {
	if len(data) > 0 {
		return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](e.Data(asMap(data[0])))
	}
	return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](e.Data())
}

// MatchTyped mirrors DataTyped for the entity's match filter. The match is a
// partial of the entity, so it round-trips through NoFrixionMoneyMoovModelsMandatesMandate (all fields
// optional at the wire level).
func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) MatchTyped(match ...NoFrixionMoneyMoovModelsMandatesMandate) NoFrixionMoneyMoovModelsMandatesMandate {
	if len(match) > 0 {
		return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](e.Match(asMap(match[0])))
	}
	return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](e.Match())
}


func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Load(reqmatch map[string]any, ctrl map[string]any) (any, error) {
	utility := e.utility
	ctx := utility.MakeContext(map[string]any{
		"opname":   "load",
		"ctrl":     ctrl,
		"match":    e.match,
		"data":     e.data,
		"reqmatch": reqmatch,
	}, e.entctx)

	return e.runOp(ctx, func() {
		if ctx.Result != nil {
			if ctx.Result.Resmatch != nil {
				e.match = ctx.Result.Resmatch
			}
			if ctx.Result.Resdata != nil {
				e.data = core.ToMapAny(vs.Clone(ctx.Result.Resdata))
				if e.data == nil {
					e.data = map[string]any{}
				}
			}
		}
	})
}

// LoadTyped is the statically-typed variant of Load: it takes an
// NoFrixionMoneyMoovModelsMandatesMandateLoadMatch and returns an NoFrixionMoneyMoovModelsMandatesMandate. It delegates to the untyped
// Load (identical runtime) and converts at the typed boundary.
func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) LoadTyped(reqmatch NoFrixionMoneyMoovModelsMandatesMandateLoadMatch, ctrl map[string]any) (NoFrixionMoneyMoovModelsMandatesMandate, error) {
	res, err := e.Load(asMap(reqmatch), ctrl)
	if err != nil {
		return NoFrixionMoneyMoovModelsMandatesMandate{}, err
	}
	return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](res), nil
}



func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) List(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("list", e.name)
}



func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Create(reqdata map[string]any, ctrl map[string]any) (any, error) {
	utility := e.utility
	ctx := utility.MakeContext(map[string]any{
		"opname":  "create",
		"ctrl":    ctrl,
		"match":   e.match,
		"data":    e.data,
		"reqdata": reqdata,
	}, e.entctx)

	return e.runOp(ctx, func() {
		if ctx.Result != nil {
			if ctx.Result.Resdata != nil {
				e.data = core.ToMapAny(vs.Clone(ctx.Result.Resdata))
				if e.data == nil {
					e.data = map[string]any{}
				}
			}
		}
	})
}

// CreateTyped is the statically-typed variant of Create: it takes an
// NoFrixionMoneyMoovModelsMandatesMandateCreateData and returns an NoFrixionMoneyMoovModelsMandatesMandate. It delegates to the untyped
// Create (identical runtime) and converts at the typed boundary.
func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) CreateTyped(reqdata NoFrixionMoneyMoovModelsMandatesMandateCreateData, ctrl map[string]any) (NoFrixionMoneyMoovModelsMandatesMandate, error) {
	res, err := e.Create(asMap(reqdata), ctrl)
	if err != nil {
		return NoFrixionMoneyMoovModelsMandatesMandate{}, err
	}
	return typedFrom[NoFrixionMoneyMoovModelsMandatesMandate](res), nil
}



func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Update(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("update", e.name)
}


func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) Remove(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("remove", e.name)
}


func (e *NoFrixionMoneyMoovModelsMandatesMandateEntity) runOp(ctx *core.Context, postDone func()) (any, error) {
	utility := e.utility

	utility.FeatureHook(ctx, "PrePoint")
	point, err := utility.MakePoint(ctx)
	ctx.Out["point"] = point
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreSpec")
	spec, err := utility.MakeSpec(ctx)
	ctx.Out["spec"] = spec
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreRequest")
	resp, err := utility.MakeRequest(ctx)
	ctx.Out["request"] = resp
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResponse")
	resp2, err := utility.MakeResponse(ctx)
	ctx.Out["response"] = resp2
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResult")
	result, err := utility.MakeResult(ctx)
	ctx.Out["result"] = result
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreDone")
	postDone()

	return utility.Done(ctx)
}
