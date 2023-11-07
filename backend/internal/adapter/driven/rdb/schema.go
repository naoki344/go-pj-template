package rdbadapter

type Customer struct {
	ID      int64  `bun:"id,pk,autoincrement"`
	Name string `bun:"name"`
	NameKana *string `bun:"name_kana"`
	Telephone string `bun:"telephone"`
	Email string `bun:"email"`
	PersonInChargeName string `bun:"person_in_charge_name"`
	PersonInChargeNameKana *string `bun:"person_in_charge_name_kana"`
	PostalCode string `bun:"postal_code"`
	PrefID int64 `bun:"pref_id"`
	Address1 string `bun:"address1"`
	Address2 string `bun:"address2"`
}
