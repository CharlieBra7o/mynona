package main

const MigaduAPIEndpoint string = "https://api.migadu.com/v1/"

// Mainboxes
type Mailboxes struct {
	Mailboxes []Mailbox `json:"mailboxes"`
}
type Mailbox struct {
	Address               string        `json:"address"`
	AutorespondActive     bool          `json:"autorespond_active"`
	AutorespondBody       string        `json:"autorespond_body"`
	AutorespondExpiresOn  interface{}   `json:"autorespond_expires_on"`
	AutorespondSubject    string        `json:"autorespond_subject"`
	ChangedAt             string        `json:"changed_at"`
	Delegations           []interface{} `json:"delegations"`
	DomainName            string        `json:"domain_name"`
	Expireable            bool          `json:"expireable"`
	ExpiresOn             interface{}   `json:"expires_on"`
	FooterActive          bool          `json:"footer_active"`
	FooterHTMLBody        interface{}   `json:"footer_html_body"`
	FooterPlainBody       interface{}   `json:"footer_plain_body"`
	Identities            []Identity    `json:"identities"`
	IsInternal            bool          `json:"is_internal"`
	LastLoginAt           string        `json:"last_login_at"`
	LocalPart             string        `json:"local_part"`
	MayAccessIMAP         bool          `json:"may_access_imap"`
	MayAccessManagesieve  bool          `json:"may_access_managesieve"`
	MayAccessPop3         bool          `json:"may_access_pop3"`
	MayReceive            bool          `json:"may_receive"`
	MaySend               bool          `json:"may_send"`
	Name                  string        `json:"name"`
	PasswordRecoveryEmail interface{}   `json:"password_recovery_email"`
	RecipientDenylist     []interface{} `json:"recipient_denylist"`
	RemoveUponExpiry      bool          `json:"remove_upon_expiry"`
	SenderAllowlist       []interface{} `json:"sender_allowlist"`
	SenderDenylist        []interface{} `json:"sender_denylist"`
	SpamAction            string        `json:"spam_action"`
	SpamAggressiveness    string        `json:"spam_aggressiveness"`
	StorageUsage          float64       `json:"storage_usage"`
}

// Identities
type Identities struct {
	Identities []Identity `json:"identities"`
}
type Identity struct {
	Address              string      `json:"address"`
	DomainName           string      `json:"domain_name"`
	FooterActive         bool        `json:"footer_active"`
	FooterHTMLBody       interface{} `json:"footer_html_body"`
	FooterPlainBody      interface{} `json:"footer_plain_body"`
	LocalPart            string      `json:"local_part"`
	MayAccessIMAP        bool        `json:"may_access_imap"`
	MayAccessManagesieve bool        `json:"may_access_managesieve"`
	MayAccessPop3        bool        `json:"may_access_pop3"`
	MayReceive           bool        `json:"may_receive"`
	MaySend              bool        `json:"may_send"`
	Name                 string      `json:"name"`
}

// Aliases
type AddressAliases struct {
	Aliases []AddressAlias `json:"address_aliases"`
}
type AddressAlias struct {
	Address          string      `json:"address"`
	Destinations     []string    `json:"destinations"`
	DomainName       string      `json:"domain_name"`
	Expireable       bool        `json:"expireable"`
	ExpiresOn        interface{} `json:"expires_on"`
	IsInternal       bool        `json:"is_internal"`
	LocalPart        string      `json:"local_part"`
	RemoveUponExpiry bool        `json:"remove_upon_expiry"`
}

// Rewrites
type Rewrites struct {
	Rewrites []Rewrite `json:"rewrites"`
}
type Rewrite struct {
	Destinations  []string `json:"destinations"`
	LocalPartRule string   `json:"local_part_rule"`
	Name          string   `json:"name"`
	OrderNum      int64    `json:"order_num"`
}
