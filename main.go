package main

import (
	"fmt"

	"github.com/go-ldap/ldap"
)

const (
	BindUsername = "administrator@xksec.com"
	// BindUsername = "buddy@xksec.com"
	BindPassword = "password-Changeme"
	BaseDN       = "ou=JDCLOUD,dc=xksec,dc=com"
	// BaseDN = "cn=,dc=xksec,dc=com"
	// Filter = "(sAMAccountName=*)"
	// Filter = "(objectclass=*)"
	Filter = "(|(objectclass=user)(objectclass=person)(objectclass=inetOrgPerson)(objectclass=organizationalPerson))"
	// Filter = "(|(objectclass=group)(objectclass=groupofnames)(objectclass=groupofuniquenames)(objectclass=organizationalUnit))"
)

func main() {
	l, err := ldap.DialURL("ldap://116.198.40.159:389")
	if err != nil {
		panic(err)
	}
	// fmt.Println(l)

	if err := l.Bind(BindUsername, BindPassword); err != nil {
		panic(err)
	}

	searchReq := ldap.NewSearchRequest(
		BaseDN,
		// ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.ScopeSingleLevel,
		// ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)

	for _, v := range result.Entries {
		fmt.Printf("Entry: %+v\n", v)
		for _, vv := range v.Attributes {
			fmt.Printf("Attribute: %+v -> %+v\n", vv.Name, vv.Values)
		}
	}
}
