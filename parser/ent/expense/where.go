// Code generated by entc, DO NOT EDIT.

package expense

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"mickaelalliel.com/telebot/parser/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// OwnerName applies equality check predicate on the "ownerName" field. It's identical to OwnerNameEQ.
func OwnerName(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerName), v))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategory), v))
	})
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCategory), v...))
	})
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCategory), v...))
	})
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategory), v))
	})
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategory), v))
	})
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategory), v))
	})
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategory), v))
	})
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategory), v))
	})
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategory), v))
	})
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategory), v))
	})
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategory), v))
	})
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategory), v))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMethod), v))
	})
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMethod), v))
	})
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMethod), v))
	})
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMethod), v))
	})
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMethod), v))
	})
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMethod), v))
	})
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMethod), v))
	})
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMethod), v))
	})
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMethod), v))
	})
}

// OwnerNameEQ applies the EQ predicate on the "ownerName" field.
func OwnerNameEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerName), v))
	})
}

// OwnerNameNEQ applies the NEQ predicate on the "ownerName" field.
func OwnerNameNEQ(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwnerName), v))
	})
}

// OwnerNameIn applies the In predicate on the "ownerName" field.
func OwnerNameIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwnerName), v...))
	})
}

// OwnerNameNotIn applies the NotIn predicate on the "ownerName" field.
func OwnerNameNotIn(vs ...string) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwnerName), v...))
	})
}

// OwnerNameGT applies the GT predicate on the "ownerName" field.
func OwnerNameGT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwnerName), v))
	})
}

// OwnerNameGTE applies the GTE predicate on the "ownerName" field.
func OwnerNameGTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwnerName), v))
	})
}

// OwnerNameLT applies the LT predicate on the "ownerName" field.
func OwnerNameLT(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwnerName), v))
	})
}

// OwnerNameLTE applies the LTE predicate on the "ownerName" field.
func OwnerNameLTE(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwnerName), v))
	})
}

// OwnerNameContains applies the Contains predicate on the "ownerName" field.
func OwnerNameContains(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOwnerName), v))
	})
}

// OwnerNameHasPrefix applies the HasPrefix predicate on the "ownerName" field.
func OwnerNameHasPrefix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOwnerName), v))
	})
}

// OwnerNameHasSuffix applies the HasSuffix predicate on the "ownerName" field.
func OwnerNameHasSuffix(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOwnerName), v))
	})
}

// OwnerNameEqualFold applies the EqualFold predicate on the "ownerName" field.
func OwnerNameEqualFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOwnerName), v))
	})
}

// OwnerNameContainsFold applies the ContainsFold predicate on the "ownerName" field.
func OwnerNameContainsFold(v string) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOwnerName), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Expense {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expense(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Expense) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Expense) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Expense) predicate.Expense {
	return predicate.Expense(func(s *sql.Selector) {
		p(s.Not())
	})
}
