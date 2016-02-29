The following LDAP ACI must be applied to the bind user for the LDAP Portal to function:

```
dn: ou=People,dc=qixalite,dc=com
changetype: modify
add: aci
aci: (targetattr = "userPassword||givenName||sn")(version 3.0; acl "Allow user modification for LDAP-Portal"; allow(read, write)(userdn = "ldap:///cn=Gogs,ou=Applications,dc=qixalite,dc=com");)
```

