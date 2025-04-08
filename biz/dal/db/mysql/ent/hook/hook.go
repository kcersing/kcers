// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"kcers/biz/dal/db/mysql/ent"
)

// The APIFunc type is an adapter to allow the use of ordinary
// function as API mutator.
type APIFunc func(context.Context, *ent.APIMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f APIFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.APIMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.APIMutation", m)
}

// The BannerFunc type is an adapter to allow the use of ordinary
// function as Banner mutator.
type BannerFunc func(context.Context, *ent.BannerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannerMutation", m)
}

// The ContractFunc type is an adapter to allow the use of ordinary
// function as Contract mutator.
type ContractFunc func(context.Context, *ent.ContractMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ContractFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ContractMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ContractMutation", m)
}

// The DictionaryFunc type is an adapter to allow the use of ordinary
// function as Dictionary mutator.
type DictionaryFunc func(context.Context, *ent.DictionaryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DictionaryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DictionaryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DictionaryMutation", m)
}

// The DictionaryDetailFunc type is an adapter to allow the use of ordinary
// function as DictionaryDetail mutator.
type DictionaryDetailFunc func(context.Context, *ent.DictionaryDetailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DictionaryDetailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DictionaryDetailMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DictionaryDetailMutation", m)
}

// The EntryLogsFunc type is an adapter to allow the use of ordinary
// function as EntryLogs mutator.
type EntryLogsFunc func(context.Context, *ent.EntryLogsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EntryLogsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EntryLogsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EntryLogsMutation", m)
}

// The FaceFunc type is an adapter to allow the use of ordinary
// function as Face mutator.
type FaceFunc func(context.Context, *ent.FaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FaceMutation", m)
}

// The LogsFunc type is an adapter to allow the use of ordinary
// function as Logs mutator.
type LogsFunc func(context.Context, *ent.LogsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LogsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LogsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LogsMutation", m)
}

// The MemberFunc type is an adapter to allow the use of ordinary
// function as Member mutator.
type MemberFunc func(context.Context, *ent.MemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberMutation", m)
}

// The MemberContractFunc type is an adapter to allow the use of ordinary
// function as MemberContract mutator.
type MemberContractFunc func(context.Context, *ent.MemberContractMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberContractFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberContractMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberContractMutation", m)
}

// The MemberContractContentFunc type is an adapter to allow the use of ordinary
// function as MemberContractContent mutator.
type MemberContractContentFunc func(context.Context, *ent.MemberContractContentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberContractContentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberContractContentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberContractContentMutation", m)
}

// The MemberDetailsFunc type is an adapter to allow the use of ordinary
// function as MemberDetails mutator.
type MemberDetailsFunc func(context.Context, *ent.MemberDetailsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberDetailsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberDetailsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberDetailsMutation", m)
}

// The MemberNoteFunc type is an adapter to allow the use of ordinary
// function as MemberNote mutator.
type MemberNoteFunc func(context.Context, *ent.MemberNoteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberNoteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberNoteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberNoteMutation", m)
}

// The MemberProductFunc type is an adapter to allow the use of ordinary
// function as MemberProduct mutator.
type MemberProductFunc func(context.Context, *ent.MemberProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberProductMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberProductMutation", m)
}

// The MemberProductPropertyFunc type is an adapter to allow the use of ordinary
// function as MemberProductProperty mutator.
type MemberProductPropertyFunc func(context.Context, *ent.MemberProductPropertyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberProductPropertyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberProductPropertyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberProductPropertyMutation", m)
}

// The MemberProfileFunc type is an adapter to allow the use of ordinary
// function as MemberProfile mutator.
type MemberProfileFunc func(context.Context, *ent.MemberProfileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberProfileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MemberProfileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberProfileMutation", m)
}

// The MenuFunc type is an adapter to allow the use of ordinary
// function as Menu mutator.
type MenuFunc func(context.Context, *ent.MenuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MenuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MenuMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MenuMutation", m)
}

// The MenuParamFunc type is an adapter to allow the use of ordinary
// function as MenuParam mutator.
type MenuParamFunc func(context.Context, *ent.MenuParamMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MenuParamFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MenuParamMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MenuParamMutation", m)
}

// The MessagesFunc type is an adapter to allow the use of ordinary
// function as Messages mutator.
type MessagesFunc func(context.Context, *ent.MessagesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MessagesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MessagesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MessagesMutation", m)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *ent.OrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderMutation", m)
}

// The OrderAmountFunc type is an adapter to allow the use of ordinary
// function as OrderAmount mutator.
type OrderAmountFunc func(context.Context, *ent.OrderAmountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderAmountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderAmountMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderAmountMutation", m)
}

// The OrderItemFunc type is an adapter to allow the use of ordinary
// function as OrderItem mutator.
type OrderItemFunc func(context.Context, *ent.OrderItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderItemMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderItemMutation", m)
}

// The OrderPayFunc type is an adapter to allow the use of ordinary
// function as OrderPay mutator.
type OrderPayFunc func(context.Context, *ent.OrderPayMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderPayFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderPayMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderPayMutation", m)
}

// The OrderSalesFunc type is an adapter to allow the use of ordinary
// function as OrderSales mutator.
type OrderSalesFunc func(context.Context, *ent.OrderSalesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderSalesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrderSalesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderSalesMutation", m)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *ent.ProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductMutation", m)
}

// The ProductPropertyFunc type is an adapter to allow the use of ordinary
// function as ProductProperty mutator.
type ProductPropertyFunc func(context.Context, *ent.ProductPropertyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductPropertyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductPropertyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductPropertyMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
}

// The ScheduleFunc type is an adapter to allow the use of ordinary
// function as Schedule mutator.
type ScheduleFunc func(context.Context, *ent.ScheduleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScheduleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ScheduleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScheduleMutation", m)
}

// The ScheduleCoachFunc type is an adapter to allow the use of ordinary
// function as ScheduleCoach mutator.
type ScheduleCoachFunc func(context.Context, *ent.ScheduleCoachMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScheduleCoachFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ScheduleCoachMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScheduleCoachMutation", m)
}

// The ScheduleMemberFunc type is an adapter to allow the use of ordinary
// function as ScheduleMember mutator.
type ScheduleMemberFunc func(context.Context, *ent.ScheduleMemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScheduleMemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ScheduleMemberMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScheduleMemberMutation", m)
}

// The SmsFunc type is an adapter to allow the use of ordinary
// function as Sms mutator.
type SmsFunc func(context.Context, *ent.SmsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SmsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SmsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SmsMutation", m)
}

// The SmsLogFunc type is an adapter to allow the use of ordinary
// function as SmsLog mutator.
type SmsLogFunc func(context.Context, *ent.SmsLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SmsLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SmsLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SmsLogMutation", m)
}

// The TokenFunc type is an adapter to allow the use of ordinary
// function as Token mutator.
type TokenFunc func(context.Context, *ent.TokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TokenMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The VenueFunc type is an adapter to allow the use of ordinary
// function as Venue mutator.
type VenueFunc func(context.Context, *ent.VenueMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VenueFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VenueMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VenueMutation", m)
}

// The VenuePlaceFunc type is an adapter to allow the use of ordinary
// function as VenuePlace mutator.
type VenuePlaceFunc func(context.Context, *ent.VenuePlaceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VenuePlaceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VenuePlaceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VenuePlaceMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
