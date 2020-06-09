package tpl

const (
	TplRitchieJson = `{
  "id" : "ritchie",
  "realm" : "{{ keycloakRealm }}",
  "notBefore" : 1582556485,
  "revokeRefreshToken" : false,
  "refreshTokenMaxReuse" : 0,
  "accessTokenLifespan" : 36000,
  "accessTokenLifespanForImplicitFlow" : 900,
  "ssoSessionIdleTimeout" : 1800,
  "ssoSessionMaxLifespan" : 36000,
  "ssoSessionIdleTimeoutRememberMe" : 0,
  "ssoSessionMaxLifespanRememberMe" : 0,
  "offlineSessionIdleTimeout" : 2592000,
  "offlineSessionMaxLifespanEnabled" : false,
  "offlineSessionMaxLifespan" : 5184000,
  "accessCodeLifespan" : 60,
  "accessCodeLifespanUserAction" : 300,
  "accessCodeLifespanLogin" : 1800,
  "actionTokenGeneratedByAdminLifespan" : 259200,
  "actionTokenGeneratedByUserLifespan" : 259200,
  "enabled" : true,
  "sslRequired" : "external",
  "registrationAllowed" : false,
  "registrationEmailAsUsername" : false,
  "rememberMe" : false,
  "verifyEmail" : false,
  "loginWithEmailAllowed" : true,
  "duplicateEmailsAllowed" : false,
  "resetPasswordAllowed" : false,
  "editUsernameAllowed" : false,
  "bruteForceProtected" : false,
  "permanentLockout" : false,
  "maxFailureWaitSeconds" : 900,
  "minimumQuickLoginWaitSeconds" : 60,
  "waitIncrementSeconds" : 60,
  "quickLoginCheckMilliSeconds" : 1000,
  "maxDeltaTimeSeconds" : 43200,
  "failureFactor" : 30,
  "roles" : {
    "realm" : [ {
      "id" : "8dce51ee-02bd-4100-bcff-bf74b20539bd",
      "name" : "user",
      "description" : "user",
      "composite" : false,
      "clientRole" : false,
      "containerId" : "ritchie",
      "attributes" : { }
    }, {
      "id" : "82026633-bbd4-43a1-ba9d-dbc174f7b045",
      "name" : "offline_access",
      "description" : "${role_offline-access}",
      "composite" : false,
      "clientRole" : false,
      "containerId" : "ritchie",
      "attributes" : { }
    }, {
      "id" : "3c5e95a6-79c2-4266-9a3f-adab95536cca",
      "name" : "uma_authorization",
      "description" : "${role_uma_authorization}",
      "composite" : false,
      "clientRole" : false,
      "containerId" : "ritchie",
      "attributes" : { }
    }, {
      "id" : "cccf5c53-1a8d-4d64-a3e4-0e3aa79a438b",
      "name" : "admin",
      "description" : "Role Admin",
      "composite" : false,
      "clientRole" : false,
      "containerId" : "ritchie",
      "attributes" : { }
    } ],
    "client" : {
      "realm-management" : [ {
        "id" : "49dd3c90-0e6e-4c39-aa9b-41975eb4f003",
        "name" : "view-clients",
        "description" : "${role_view-clients}",
        "composite" : true,
        "composites" : {
          "client" : {
            "realm-management" : [ "query-clients" ]
          }
        },
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "f0ed0ae1-18f2-4c18-be7a-4f7f44f4187c",
        "name" : "manage-clients",
        "description" : "${role_manage-clients}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "53dc9cda-2209-4c9b-ac2a-2c5b80def7a8",
        "name" : "query-clients",
        "description" : "${role_query-clients}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "b7da787e-f660-4173-b021-48e2a15292a3",
        "name" : "view-authorization",
        "description" : "${role_view-authorization}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "a3ad9c5f-d2d4-4d71-9712-e73a086595df",
        "name" : "view-identity-providers",
        "description" : "${role_view-identity-providers}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "d1f2a343-2806-4a39-bda7-ed637285be43",
        "name" : "create-client",
        "description" : "${role_create-client}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "6a988925-efbc-4278-97e6-0319d62b3996",
        "name" : "manage-identity-providers",
        "description" : "${role_manage-identity-providers}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "2c807dd6-cbdb-47a2-9e8d-571abf94f82d",
        "name" : "query-groups",
        "description" : "${role_query-groups}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "e2a58636-2e74-4063-97f1-96b6b7464d14",
        "name" : "view-realm",
        "description" : "${role_view-realm}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "c12db5f4-a9ee-4b2a-a836-1343956f3c1d",
        "name" : "manage-authorization",
        "description" : "${role_manage-authorization}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "dfd5fa08-68c2-4f51-b2ee-e29c01914797",
        "name" : "view-events",
        "description" : "${role_view-events}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "242342ba-ce6f-48d4-8c65-374018c4b9a5",
        "name" : "query-realms",
        "description" : "${role_query-realms}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "646c486a-1a5b-4d6e-99c2-6f481e7e92cf",
        "name" : "view-users",
        "description" : "${role_view-users}",
        "composite" : true,
        "composites" : {
          "client" : {
            "realm-management" : [ "query-groups", "query-users" ]
          }
        },
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "3bf832fe-42bb-4fb8-b7bc-81cd53fca88d",
        "name" : "realm-admin",
        "description" : "${role_realm-admin}",
        "composite" : true,
        "composites" : {
          "client" : {
            "realm-management" : [ "view-clients", "query-clients", "manage-clients", "view-authorization", "view-identity-providers", "create-client", "manage-identity-providers", "query-groups", "manage-authorization", "view-realm", "query-realms", "view-events", "view-users", "manage-events", "impersonation", "manage-users", "manage-realm", "query-users" ]
          }
        },
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "7fe1d779-e1d7-41d7-96e2-fb3ed0878091",
        "name" : "manage-events",
        "description" : "${role_manage-events}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "b037e3a3-3014-4782-8d25-6a504f050e8e",
        "name" : "impersonation",
        "description" : "${role_impersonation}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "4ec43e12-c101-4196-a476-4b646c0afe0f",
        "name" : "manage-users",
        "description" : "${role_manage-users}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "b81986b1-ecf8-4d6f-acf1-ae62f3483012",
        "name" : "manage-realm",
        "description" : "${role_manage-realm}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      }, {
        "id" : "c15288c1-6f15-4ab7-b21c-bd6fb23d6d09",
        "name" : "query-users",
        "description" : "${role_query-users}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
        "attributes" : { }
      } ],
      "user-login-without-permission" : [ {
        "id" : "95c0aaa6-d33f-46ea-bfb9-a012214a7150",
        "name" : "uma_protection",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "f88a127c-ddff-45f2-b28f-a96e260ac900",
        "attributes" : { }
      } ],
      "security-admin-console" : [ ],
      "admin-cli" : [ ],
      "broker" : [ {
        "id" : "c8bbd702-f938-4c45-80c8-9ff1d5464157",
        "name" : "read-token",
        "description" : "${role_read-token}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "ac5f8fe1-b8a8-472f-ac5f-846c35c30f2d",
        "attributes" : { }
      } ],
      "user-login" : [ {
        "id" : "f787681b-1b31-4da7-939b-6534a02af9a1",
        "name" : "uma_protection",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "f2dde806-507f-4ec3-8853-15c687493775",
        "attributes" : { }
      } ],
      "oauth" : [ ],
      "account" : [ {
        "id" : "41612545-2321-4650-8b01-77aa9ca11540",
        "name" : "manage-account",
        "description" : "${role_manage-account}",
        "composite" : true,
        "composites" : {
          "client" : {
            "account" : [ "manage-account-links" ]
          }
        },
        "clientRole" : true,
        "containerId" : "5399842a-1cc0-44c2-ac5a-b70925b93685",
        "attributes" : { }
      }, {
        "id" : "17a2f6e2-c856-4f66-925c-6f35ae1d603c",
        "name" : "manage-account-links",
        "description" : "${role_manage-account-links}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "5399842a-1cc0-44c2-ac5a-b70925b93685",
        "attributes" : { }
      }, {
        "id" : "a5ad0929-e7f3-49cd-82e5-4efcdcdfee2b",
        "name" : "view-profile",
        "description" : "${role_view-profile}",
        "composite" : false,
        "clientRole" : true,
        "containerId" : "5399842a-1cc0-44c2-ac5a-b70925b93685",
        "attributes" : { }
      } ]
    }
  },
  "groups" : [ ],
  "defaultRoles" : [ "uma_authorization", "offline_access", "user" ],
  "requiredCredentials" : [ "password" ],
  "otpPolicyType" : "totp",
  "otpPolicyAlgorithm" : "HmacSHA1",
  "otpPolicyInitialCounter" : 0,
  "otpPolicyDigits" : 6,
  "otpPolicyLookAheadWindow" : 1,
  "otpPolicyPeriod" : 30,
  "otpSupportedApplications" : [ "FreeOTP", "Google Authenticator" ],
  "users" : [ {
    "id" : "85254853-90e9-4354-ba12-5e3391f9a098",
    "createdTimestamp" : 1585402544396,
    "username" : "service-account-user-login",
    "enabled" : true,
    "totp" : false,
    "emailVerified" : false,
    "email" : "service-account-user-login@placeholder.org",
    "serviceAccountClientId" : "user-login",
    "credentials" : [ ],
    "disableableCredentialTypes" : [ ],
    "requiredActions" : [ ],
    "realmRoles" : [ "user", "offline_access", "uma_authorization" ],
    "clientRoles" : {
      "realm-management" : [ "manage-users" ],
      "user-login" : [ "uma_protection" ],
      "account" : [ "manage-account", "view-profile" ]
    },
    "notBefore" : 0,
    "groups" : [ ]
  }, {
    "id" : "b1407044-6bb5-411c-b193-8d27c9bd122d",
    "createdTimestamp" : 1585847093361,
    "username" : "service-account-user-login-without-permission",
    "enabled" : true,
    "totp" : false,
    "emailVerified" : false,
    "email" : "service-account-user-login-without-permission@placeholder.org",
    "serviceAccountClientId" : "user-login-without-permission",
    "credentials" : [ ],
    "disableableCredentialTypes" : [ ],
    "requiredActions" : [ ],
    "realmRoles" : [ "user", "offline_access", "uma_authorization" ],
    "clientRoles" : {
      "realm-management" : [ "view-users", "query-users" ],
      "user-login-without-permission" : [ "uma_protection" ],
      "account" : [ "manage-account", "view-profile" ]
    },
    "notBefore" : 0,
    "groups" : [ ]
  }, {
    "id" : "184e44f9-6a9b-4c8e-a50e-8dfd2d07caf0",
    "createdTimestamp" : 1585402795434,
    "username" : "user",
    "enabled" : true,
    "totp" : false,
    "emailVerified" : true,
    "firstName" : "user",
    "lastName" : "user",
    "email" : "user@user.com",
    "credentials" : [ {
      "type" : "password",
      "hashedSaltedValue" : "x35zAkT0JEQCAt2DSgxt6ibNL6ArZj2O9TMokpQAMeqHBYnSs4nzGc3uozU+1zBfXvNmVugZAFmxIdeUkFRDpA==",
      "salt" : "8MeKTr5Ne6S6eKt0/3BC9Q==",
      "hashIterations" : 27500,
      "counter" : 0,
      "algorithm" : "pbkdf2-sha256",
      "digits" : 0,
      "period" : 0,
      "createdDate" : 1585402811800,
      "config" : { }
    } ],
    "disableableCredentialTypes" : [ "password" ],
    "requiredActions" : [ ],
    "realmRoles" : [ "user", "offline_access", "uma_authorization" ],
    "clientRoles" : {
      "account" : [ "manage-account", "view-profile" ]
    },
    "notBefore" : 0,
    "groups" : [ ]
  } ],
  "scopeMappings" : [ {
    "clientScope" : "offline_access",
    "roles" : [ "offline_access" ]
  } ],
  "clients" : [ {
    "id" : "5399842a-1cc0-44c2-ac5a-b70925b93685",
    "clientId" : "account",
    "name" : "${client_account}",
    "baseUrl" : "/auth/realms/ritchie/account",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "account",
    "defaultRoles" : [ "manage-account", "view-profile" ],
    "redirectUris" : [ "/auth/realms/ritchie/account/*" ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : false,
    "serviceAccountsEnabled" : false,
    "publicClient" : false,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : { },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : false,
    "nodeReRegistrationTimeout" : 0,
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "75c215b7-89aa-49a8-aeb1-05a765312b7f",
    "clientId" : "admin-cli",
    "name" : "${client_admin-cli}",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "admin-cli",
    "redirectUris" : [ ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : false,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : true,
    "serviceAccountsEnabled" : false,
    "publicClient" : true,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : { },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : false,
    "nodeReRegistrationTimeout" : 0,
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "ac5f8fe1-b8a8-472f-ac5f-846c35c30f2d",
    "clientId" : "broker",
    "name" : "${client_broker}",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "broker",
    "redirectUris" : [ ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : false,
    "serviceAccountsEnabled" : false,
    "publicClient" : false,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : { },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : false,
    "nodeReRegistrationTimeout" : 0,
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "6de8fb38-32ee-4032-b7a8-328507074738",
    "clientId" : "oauth",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "oauth",
    "redirectUris" : [ "http://localhost:8888/ritchie/callback" ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : true,
    "directAccessGrantsEnabled" : false,
    "serviceAccountsEnabled" : false,
    "publicClient" : true,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : {
      "saml.assertion.signature" : "false",
      "saml.force.post.binding" : "false",
      "saml.multivalued.roles" : "false",
      "saml.encrypt" : "false",
      "saml.server.signature" : "false",
      "saml.server.signature.keyinfo.ext" : "false",
      "exclude.session.state.from.auth.response" : "false",
      "saml_force_name_id_format" : "false",
      "saml.client.signature" : "false",
      "tls.client.certificate.bound.access.tokens" : "false",
      "saml.authnstatement" : "false",
      "display.on.consent.screen" : "false",
      "saml.onetimeuse.condition" : "false"
    },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : true,
    "nodeReRegistrationTimeout" : -1,
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "205f3f16-3d60-4a51-acbf-a52d4dd9ab06",
    "clientId" : "realm-management",
    "name" : "${client_realm-management}",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "realm-management",
    "redirectUris" : [ ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : true,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : false,
    "serviceAccountsEnabled" : false,
    "publicClient" : false,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : { },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : false,
    "nodeReRegistrationTimeout" : 0,
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "08672997-08c5-417e-a926-b027c99046e8",
    "clientId" : "security-admin-console",
    "name" : "${client_security-admin-console}",
    "baseUrl" : "/auth/admin/ritchie/console/index.html",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "security-admin-console",
    "redirectUris" : [ "/auth/admin/ritchie/console/*" ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : false,
    "serviceAccountsEnabled" : false,
    "publicClient" : true,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : { },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : false,
    "nodeReRegistrationTimeout" : 0,
    "protocolMappers" : [ {
      "id" : "a6584ed7-aa1a-4361-9bbb-87a1446321fc",
      "name" : "locale",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "locale",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "locale",
        "jsonType.label" : "String"
      }
    } ],
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ]
  }, {
    "id" : "f2dde806-507f-4ec3-8853-15c687493775",
    "clientId" : "{{ keycloakClientId }}",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "{{ keycloakClientSecret }}",
    "redirectUris" : [ ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : false,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : true,
    "serviceAccountsEnabled" : true,
    "authorizationServicesEnabled" : true,
    "publicClient" : false,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : {
      "saml.assertion.signature" : "false",
      "saml.force.post.binding" : "false",
      "saml.multivalued.roles" : "false",
      "saml.encrypt" : "false",
      "saml.server.signature" : "false",
      "saml.server.signature.keyinfo.ext" : "false",
      "exclude.session.state.from.auth.response" : "false",
      "saml_force_name_id_format" : "false",
      "saml.client.signature" : "false",
      "tls.client.certificate.bound.access.tokens" : "false",
      "saml.authnstatement" : "false",
      "display.on.consent.screen" : "false",
      "saml.onetimeuse.condition" : "false"
    },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : true,
    "nodeReRegistrationTimeout" : -1,
    "protocolMappers" : [ {
      "id" : "c73e178e-d9d6-4996-be1c-9d3ce0d60ec1",
      "name" : "Client IP Address",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientAddress",
        "userinfo.token.claim" : "true",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientAddress",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "53bdd9cb-28d0-47f6-845c-ea2fb5fdb1d7",
      "name" : "user-login-prd",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-audience-mapper",
      "consentRequired" : false,
      "config" : {
        "included.client.audience" : "user-login-prod",
        "id.token.claim" : "false",
        "access.token.claim" : "true",
        "userinfo.token.claim" : "false"
      }
    }, {
      "id" : "926eb4c1-a5af-4c2a-a87e-df4b0817ab5e",
      "name" : "Client Host",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientHost",
        "userinfo.token.claim" : "true",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientHost",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "aeda3743-8379-4e08-81e0-57afa4aaeed0",
      "name" : "Client ID",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientId",
        "userinfo.token.claim" : "true",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientId",
        "jsonType.label" : "String"
      }
    } ],
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ],
    "authorizationSettings" : {
      "allowRemoteResourceManagement" : true,
      "policyEnforcementMode" : "ENFORCING",
      "resources" : [ {
        "name" : "Default Resource",
        "type" : "urn:user-login-prod:resources:default",
        "ownerManagedAccess" : false,
        "attributes" : { },
        "_id" : "56a0f993-bfd9-46dd-84f6-eca170b33951",
        "uris" : [ "/*" ]
      } ],
      "policies" : [ {
        "id" : "f8ea3a55-20d5-4dac-92c5-00f1a1a2c9da",
        "name" : "Default Policy",
        "description" : "A policy that grants access only for users within this realm",
        "type" : "js",
        "logic" : "POSITIVE",
        "decisionStrategy" : "AFFIRMATIVE",
        "config" : {
          "code" : "// by default, grants any permission associated with this policy\n$evaluation.grant();\n"
        }
      }, {
        "id" : "53fdaf99-f161-4707-bc06-fdd90bf5d5f0",
        "name" : "Default Permission",
        "description" : "A permission that applies to the default resource type",
        "type" : "resource",
        "logic" : "POSITIVE",
        "decisionStrategy" : "UNANIMOUS",
        "config" : {
          "defaultResourceType" : "urn:user-login-prod:resources:default",
          "applyPolicies" : "[\"Default Policy\"]"
        }
      } ],
      "scopes" : [ ],
      "decisionStrategy" : "UNANIMOUS"
    }
  }, {
    "id" : "f88a127c-ddff-45f2-b28f-a96e260ac900",
    "clientId" : "user-login-without-permission",
    "surrogateAuthRequired" : false,
    "enabled" : true,
    "clientAuthenticatorType" : "client-secret",
    "secret" : "c0b1f8f1-c746-4573-9679-b7ecdf2d48cc",
    "redirectUris" : [ "http://localhost" ],
    "webOrigins" : [ ],
    "notBefore" : 0,
    "bearerOnly" : false,
    "consentRequired" : false,
    "standardFlowEnabled" : true,
    "implicitFlowEnabled" : false,
    "directAccessGrantsEnabled" : true,
    "serviceAccountsEnabled" : true,
    "authorizationServicesEnabled" : true,
    "publicClient" : false,
    "frontchannelLogout" : false,
    "protocol" : "openid-connect",
    "attributes" : {
      "saml.assertion.signature" : "false",
      "saml.force.post.binding" : "false",
      "saml.multivalued.roles" : "false",
      "saml.encrypt" : "false",
      "saml.server.signature" : "false",
      "saml.server.signature.keyinfo.ext" : "false",
      "exclude.session.state.from.auth.response" : "false",
      "saml_force_name_id_format" : "false",
      "saml.client.signature" : "false",
      "tls.client.certificate.bound.access.tokens" : "false",
      "saml.authnstatement" : "false",
      "display.on.consent.screen" : "false",
      "saml.onetimeuse.condition" : "false"
    },
    "authenticationFlowBindingOverrides" : { },
    "fullScopeAllowed" : true,
    "nodeReRegistrationTimeout" : -1,
    "protocolMappers" : [ {
      "id" : "a14b9547-d114-492c-9773-a254f1c7b1f0",
      "name" : "Client IP Address",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientAddress",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientAddress",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "22f54e5c-f51d-4ae9-9d06-f60829c712b6",
      "name" : "Client Host",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientHost",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientHost",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "530ae307-bc1f-482b-a7bd-335efdcb5c16",
      "name" : "Client ID",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usersessionmodel-note-mapper",
      "consentRequired" : false,
      "config" : {
        "user.session.note" : "clientId",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "clientId",
        "jsonType.label" : "String"
      }
    } ],
    "defaultClientScopes" : [ "web-origins", "role_list", "roles", "profile", "email" ],
    "optionalClientScopes" : [ "address", "phone", "offline_access", "microprofile-jwt" ],
    "authorizationSettings" : {
      "allowRemoteResourceManagement" : true,
      "policyEnforcementMode" : "ENFORCING",
      "resources" : [ {
        "name" : "Default Resource",
        "type" : "urn:user-login-without-permission:resources:default",
        "ownerManagedAccess" : false,
        "attributes" : { },
        "_id" : "d4f0a846-f135-4c27-af53-c690668970b8",
        "uris" : [ "/*" ]
      } ],
      "policies" : [ {
        "id" : "8afea607-ec63-404e-9bad-317321701b9d",
        "name" : "Default Policy",
        "description" : "A policy that grants access only for users within this realm",
        "type" : "js",
        "logic" : "POSITIVE",
        "decisionStrategy" : "AFFIRMATIVE",
        "config" : {
          "code" : "// by default, grants any permission associated with this policy\n$evaluation.grant();\n"
        }
      }, {
        "id" : "a83b3a1a-7148-4568-a279-9861cd8203f3",
        "name" : "Default Permission",
        "description" : "A permission that applies to the default resource type",
        "type" : "resource",
        "logic" : "POSITIVE",
        "decisionStrategy" : "UNANIMOUS",
        "config" : {
          "defaultResourceType" : "urn:user-login-without-permission:resources:default",
          "applyPolicies" : "[\"Default Policy\"]"
        }
      } ],
      "scopes" : [ ],
      "decisionStrategy" : "UNANIMOUS"
    }
  } ],
  "clientScopes" : [ {
    "id" : "71e23771-8f28-455d-b218-7f92be0e4c71",
    "name" : "address",
    "description" : "OpenID Connect built-in scope: address",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "true",
      "display.on.consent.screen" : "true",
      "consent.screen.text" : "${addressScopeConsentText}"
    },
    "protocolMappers" : [ {
      "id" : "7abde633-bd04-4123-83f3-56edbf881efb",
      "name" : "address",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-address-mapper",
      "consentRequired" : false,
      "config" : {
        "user.attribute.formatted" : "formatted",
        "user.attribute.country" : "country",
        "user.attribute.postal_code" : "postal_code",
        "userinfo.token.claim" : "true",
        "user.attribute.street" : "street",
        "id.token.claim" : "true",
        "user.attribute.region" : "region",
        "access.token.claim" : "true",
        "user.attribute.locality" : "locality"
      }
    } ]
  }, {
    "id" : "f17314a7-0a21-49c3-aef2-004c996ed838",
    "name" : "email",
    "description" : "OpenID Connect built-in scope: email",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "true",
      "display.on.consent.screen" : "true",
      "consent.screen.text" : "${emailScopeConsentText}"
    },
    "protocolMappers" : [ {
      "id" : "a5784ca5-1207-4894-9f93-067260e8109e",
      "name" : "email",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "email",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "email",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "b1ebd0fc-1f00-4287-8fdc-da641d157f02",
      "name" : "email verified",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "emailVerified",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "email_verified",
        "jsonType.label" : "boolean"
      }
    } ]
  }, {
    "id" : "c10e1a1f-3b4e-48de-adae-99edab979b29",
    "name" : "microprofile-jwt",
    "description" : "Microprofile - JWT built-in scope",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "true",
      "display.on.consent.screen" : "false"
    },
    "protocolMappers" : [ {
      "id" : "18846c4f-65f8-4dda-b1f9-c621854ffbbc",
      "name" : "groups",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-realm-role-mapper",
      "consentRequired" : false,
      "config" : {
        "multivalued" : "true",
        "userinfo.token.claim" : "true",
        "user.attribute" : "foo",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "groups",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "645db4bf-768f-4211-b18d-7cace41e3ac7",
      "name" : "upn",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "username",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "upn",
        "jsonType.label" : "String"
      }
    } ]
  }, {
    "id" : "7675ae06-d2a9-4959-8677-86b4451bee2c",
    "name" : "offline_access",
    "description" : "OpenID Connect built-in scope: offline_access",
    "protocol" : "openid-connect",
    "attributes" : {
      "consent.screen.text" : "${offlineAccessScopeConsentText}",
      "display.on.consent.screen" : "true"
    }
  }, {
    "id" : "b67f3e58-6fcc-4dda-bd7c-a5df26110679",
    "name" : "phone",
    "description" : "OpenID Connect built-in scope: phone",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "true",
      "display.on.consent.screen" : "true",
      "consent.screen.text" : "${phoneScopeConsentText}"
    },
    "protocolMappers" : [ {
      "id" : "88248348-bfd1-4a11-b298-b4b041b53c5a",
      "name" : "phone number verified",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "phoneNumberVerified",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "phone_number_verified",
        "jsonType.label" : "boolean"
      }
    }, {
      "id" : "24e8d5ef-3b13-44be-959a-83719d524b37",
      "name" : "phone number",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "phoneNumber",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "phone_number",
        "jsonType.label" : "String"
      }
    } ]
  }, {
    "id" : "fa9df969-8856-4b3d-ac59-548f48a9af55",
    "name" : "profile",
    "description" : "OpenID Connect built-in scope: profile",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "true",
      "display.on.consent.screen" : "true",
      "consent.screen.text" : "${profileScopeConsentText}"
    },
    "protocolMappers" : [ {
      "id" : "13bac1ce-045b-4fff-92e0-fc9cffc4484d",
      "name" : "full name",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-full-name-mapper",
      "consentRequired" : false,
      "config" : {
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "userinfo.token.claim" : "true"
      }
    }, {
      "id" : "908a69ed-6a92-45a9-9207-45c581de0c60",
      "name" : "username",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "username",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "preferred_username",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "20d7965f-9dfa-4009-928f-9e70f287f779",
      "name" : "family name",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "lastName",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "family_name",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "068e1879-930a-4f0b-b647-43fea27e5884",
      "name" : "given name",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-property-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "firstName",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "given_name",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "b9714319-42b1-4522-8dda-01dcf14c7145",
      "name" : "picture",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "picture",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "picture",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "6723b470-6451-414d-8260-88076f7d9cdd",
      "name" : "middle name",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "middleName",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "middle_name",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "d152f577-5869-4a6c-b14f-0d93ebe10386",
      "name" : "birthdate",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "birthdate",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "birthdate",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "1b5b015f-8fa6-46d2-b5f3-e2a6a187d846",
      "name" : "locale",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "locale",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "locale",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "21e454f6-9118-4ecf-b467-8baa53eddd79",
      "name" : "profile",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "profile",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "profile",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "59b204fa-ba4b-400a-97e1-2b9d010be048",
      "name" : "updated at",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "updatedAt",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "updated_at",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "9b01fd2d-2b20-44cf-bd59-3d9959f6f03c",
      "name" : "nickname",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "nickname",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "nickname",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "d13e48d7-6fde-47ed-978e-8ac4af3a25cd",
      "name" : "gender",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "gender",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "gender",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "eaa92e42-8441-4b4b-a6a7-ba9b7e93f936",
      "name" : "zoneinfo",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "zoneinfo",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "zoneinfo",
        "jsonType.label" : "String"
      }
    }, {
      "id" : "c905b390-f816-4039-adcf-fc851a61d290",
      "name" : "website",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-attribute-mapper",
      "consentRequired" : false,
      "config" : {
        "userinfo.token.claim" : "true",
        "user.attribute" : "website",
        "id.token.claim" : "true",
        "access.token.claim" : "true",
        "claim.name" : "website",
        "jsonType.label" : "String"
      }
    } ]
  }, {
    "id" : "865849f6-3e06-4d28-a88f-c2e560e9c1a1",
    "name" : "role_list",
    "description" : "SAML role list",
    "protocol" : "saml",
    "attributes" : {
      "consent.screen.text" : "${samlRoleListScopeConsentText}",
      "display.on.consent.screen" : "true"
    },
    "protocolMappers" : [ {
      "id" : "676c1641-f2a8-4458-a2cd-9d8f11e782f7",
      "name" : "role list",
      "protocol" : "saml",
      "protocolMapper" : "saml-role-list-mapper",
      "consentRequired" : false,
      "config" : {
        "single" : "false",
        "attribute.nameformat" : "Basic",
        "attribute.name" : "Role"
      }
    } ]
  }, {
    "id" : "27e6ff66-4949-4d85-bf52-8f87719d25bb",
    "name" : "roles",
    "description" : "OpenID Connect scope for add user roles to the access token",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "false",
      "display.on.consent.screen" : "true",
      "consent.screen.text" : "${rolesScopeConsentText}"
    },
    "protocolMappers" : [ {
      "id" : "703b6547-ee47-4841-a116-197264b6e283",
      "name" : "realm roles",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-realm-role-mapper",
      "consentRequired" : false,
      "config" : {
        "user.attribute" : "foo",
        "access.token.claim" : "true",
        "claim.name" : "realm_access.roles",
        "jsonType.label" : "String",
        "multivalued" : "true"
      }
    }, {
      "id" : "9da93ca5-631b-4c2f-ac9d-42b30b60f514",
      "name" : "audience resolve",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-audience-resolve-mapper",
      "consentRequired" : false,
      "config" : { }
    }, {
      "id" : "cd48689e-1e51-4f2b-a8e9-1d877789f73b",
      "name" : "client roles",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-usermodel-client-role-mapper",
      "consentRequired" : false,
      "config" : {
        "user.attribute" : "foo",
        "access.token.claim" : "true",
        "claim.name" : "resource_access.${client_id}.roles",
        "jsonType.label" : "String",
        "multivalued" : "true"
      }
    } ]
  }, {
    "id" : "3032c5a8-8181-40d4-9586-e4c9845c48e2",
    "name" : "web-origins",
    "description" : "OpenID Connect scope for add allowed web origins to the access token",
    "protocol" : "openid-connect",
    "attributes" : {
      "include.in.token.scope" : "false",
      "display.on.consent.screen" : "false",
      "consent.screen.text" : ""
    },
    "protocolMappers" : [ {
      "id" : "3f6633ce-b064-4ce6-8aba-02e970a46ef1",
      "name" : "allowed web origins",
      "protocol" : "openid-connect",
      "protocolMapper" : "oidc-allowed-origins-mapper",
      "consentRequired" : false,
      "config" : { }
    } ]
  } ],
  "defaultDefaultClientScopes" : [ "roles", "web-origins", "role_list", "email", "profile" ],
  "defaultOptionalClientScopes" : [ "address", "offline_access", "phone", "microprofile-jwt" ],
  "browserSecurityHeaders" : {
    "contentSecurityPolicyReportOnly" : "",
    "xContentTypeOptions" : "nosniff",
    "xRobotsTag" : "none",
    "xFrameOptions" : "SAMEORIGIN",
    "xXSSProtection" : "1; mode=block",
    "contentSecurityPolicy" : "frame-src 'self'; frame-ancestors 'self'; object-src 'none';",
    "strictTransportSecurity" : "max-age=31536000; includeSubDomains"
  },
  "smtpServer" : { },
  "eventsEnabled" : false,
  "eventsListeners" : [ "jboss-logging" ],
  "enabledEventTypes" : [ ],
  "adminEventsEnabled" : false,
  "adminEventsDetailsEnabled" : false,
  "identityProviders" : [ {
    "alias" : "google",
    "internalId" : "33cb47df-e232-4980-a5f2-86ab60ee7d6e",
    "providerId" : "google",
    "enabled" : true,
    "updateProfileFirstLoginMode" : "on",
    "trustEmail" : true,
    "storeToken" : true,
    "addReadTokenRoleOnCreate" : true,
    "authenticateByDefault" : false,
    "linkOnly" : false,
    "firstBrokerLoginFlowAlias" : "first broker login",
    "config" : {
      "hideOnLoginPage" : "",
      "offlineAccess" : "",
      "acceptsPromptNoneForwardFromClient" : "",
      "clientId" : "943232805725-28q9f5shsiaevia93d5ohbl2btf7iovc.apps.googleusercontent.com",
      "disableUserInfo" : "",
      "userIp" : "",
      "clientSecret" : "**********",
      "defaultScope" : "openid email profile",
      "useJwksUrl" : "true"
    }
  } ],
  "components" : {
    "org.keycloak.services.clientregistration.policy.ClientRegistrationPolicy" : [ {
      "id" : "da0372fd-663d-461d-bbe0-afaef4de3424",
      "name" : "Allowed Client Scopes",
      "providerId" : "allowed-client-templates",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : {
        "allow-default-scopes" : [ "true" ]
      }
    }, {
      "id" : "b260cfd0-be1e-47ed-8603-6bfba770d0f8",
      "name" : "Allowed Protocol Mapper Types",
      "providerId" : "allowed-protocol-mappers",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : {
        "allowed-protocol-mapper-types" : [ "saml-role-list-mapper", "oidc-full-name-mapper", "saml-user-property-mapper", "oidc-address-mapper", "oidc-sha256-pairwise-sub-mapper", "oidc-usermodel-attribute-mapper", "saml-user-attribute-mapper", "oidc-usermodel-property-mapper" ]
      }
    }, {
      "id" : "b9fa6a05-3ac6-48e9-95f3-1dd305a01f09",
      "name" : "Max Clients Limit",
      "providerId" : "max-clients",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : {
        "max-clients" : [ "200" ]
      }
    }, {
      "id" : "05d237fd-c69c-4021-9d79-11750af5f207",
      "name" : "Trusted Hosts",
      "providerId" : "trusted-hosts",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : {
        "host-sending-registration-request-must-match" : [ "true" ],
        "client-uris-must-match" : [ "true" ]
      }
    }, {
      "id" : "703e540f-7493-4012-8a0a-a7e8331277c5",
      "name" : "Allowed Client Scopes",
      "providerId" : "allowed-client-templates",
      "subType" : "authenticated",
      "subComponents" : { },
      "config" : {
        "allow-default-scopes" : [ "true" ]
      }
    }, {
      "id" : "2b045f4b-11bc-41e3-9813-3907b0b884f4",
      "name" : "Full Scope Disabled",
      "providerId" : "scope",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : { }
    }, {
      "id" : "6dd49602-cd96-44ab-9b28-5e3acfefafed",
      "name" : "Consent Required",
      "providerId" : "consent-required",
      "subType" : "anonymous",
      "subComponents" : { },
      "config" : { }
    }, {
      "id" : "0c373ef5-1a8e-4d05-88d5-83c170cac07a",
      "name" : "Allowed Protocol Mapper Types",
      "providerId" : "allowed-protocol-mappers",
      "subType" : "authenticated",
      "subComponents" : { },
      "config" : {
        "allowed-protocol-mapper-types" : [ "oidc-usermodel-property-mapper", "saml-role-list-mapper", "oidc-usermodel-attribute-mapper", "saml-user-property-mapper", "oidc-address-mapper", "oidc-full-name-mapper", "saml-user-attribute-mapper", "oidc-sha256-pairwise-sub-mapper" ]
      }
    } ],
    "org.keycloak.keys.KeyProvider" : [ {
      "id" : "2a0ae8cd-8ffd-4751-8748-37020f9a43c4",
      "name" : "aes-generated",
      "providerId" : "aes-generated",
      "subComponents" : { },
      "config" : {
        "kid" : [ "0d52149c-4a17-44d1-8b24-5dd6c725b182" ],
        "secret" : [ "n6FG7vJsO9ziQ3-rQ2KeoA" ],
        "priority" : [ "100" ]
      }
    }, {
      "id" : "47dbdd5e-3d7b-4c81-88f4-3a0e5a676fdb",
      "name" : "hmac-generated",
      "providerId" : "hmac-generated",
      "subComponents" : { },
      "config" : {
        "kid" : [ "610e1e46-4aa1-4e25-a8c4-a23d4ea1db60" ],
        "secret" : [ "c1yXe2WyQCX45iCKdWLyBcuZmU4up5uujcsu_7UgokAO-NSjSnrYcwa1tRDSJt-0Pp2B8voK_K3hc7zBPiQpyA" ],
        "priority" : [ "100" ],
        "algorithm" : [ "HS256" ]
      }
    }, {
      "id" : "d564cf01-aa90-47a1-9485-6164e0c95b1e",
      "name" : "rsa-generated",
      "providerId" : "rsa-generated",
      "subComponents" : { },
      "config" : {
        "privateKey" : [ "MIIEpAIBAAKCAQEAktD1gylMHUNkGg7PAEmfPHsiMijm2D2XsDgk7V/49p+x/1ibTHwsBm66IevFKoZqEs3CuGqJRboYhfS4oe3/OG5ciM/7N7TYg0fj0DOa726jvnd+gCwOvCugD9nZtBhxdrXrY5dZss3phvCI4kO49Nm5EM6Zk5MdmFp38OuND49BqQFcx/s1ekwtchOuzmuc9KBx34BMPnwd3X9lk/opmudUpH4QfIuQEJV0kK8JKZRvyCkLb7JTwndLtnfQ/XhexoVCq2e4809y1s5qBcZmqH1D3wHKAOoOA6SKqlCsMm6nAXw9RVfH8F7/pdKjnwAVaTwWnrVupOuG2F7/s1wplQIDAQABAoIBAAhFVU9qSCtt3HnLU8YEX8Acf7SrsTWFYsI+p3Xn7jKWlIR6DWWlqSVDn6DBk33tzJP6m3mgJtOSxEbnnm8g5Tvcm8HpVZGx0nCy4BaekZ/jb74PJUU7NilLm9zIvQb0SZASd4xIvnjgOMOtUUjN3+GLUm2r9eLmashmcFNWMcAqj7uoDiVYLeDGchT5bnR7MGvPUSqCdFN7O9BCGv2z7gYWewuEk3hBZ/j51G1/2kWXPUHhZE9ZEmoLH+iiJLZP8MQ/IEkRI0dL0jeJi8GFg7TfLGQmcolFjtNsW7NZ+stVpwzNroZeZS8xCpO/SYENNjQ8YM9RKujGSurjKgKg0QECgYEA48kLCi5ov5AxQPVH0SHH94w0jp+Fx6uI5YRRKLETxJ+I0fpJCCHczpTRTtGjz4PgCrthT7liG4wL1lnTzYqkjuomTetHx6UQm/NeDdyxr2R1noJAFjXNn6v8cyhqJe5e6oG10xJhvKDsLce7UVNfMv2Ap9dFOS77XBF2ydlWdiUCgYEApQBtdgEkZ2z+tnF1XsozzxStjbKNmPbUNIAnoiJ9S5pCazp1cxQsYehdbaDK5PQXfagxiBKTxbnHrsqb4HhUSrd9baodbCT7Vj1Rn1hruHywMOtl/Qj0fCRm+UuzEJNavA0tFVMSFpJqGV7uvZvfnl49cL3e/z3ULOn8rCiYcrECgYEAix/T2lgWKkqLir1FK/qSBCRiRWoxb47ZgjWazFu7UzCNhJJ14L2t+47xeE6a2rgVMvjXqrXzsheLc1RSBQxVDeT9mEHICaxLxLSUEatl61ZUcIflVKtv98I34q4ghyJcXqbywwRLJcaz+hIv2rAtSRuEZP8ajJ6hb85K22YfZPECgYAFvOERlHKWoiStZX6mYDOzBM4QOmBQLG0usjIuojJnOaMLpYnGCbKNaQ4urMPBeHeEqB+o0rSO0KoP+v22rZiIzJ6w/JRvWU3gLtuyxdlhHfzmhpkSVohh9MYb7zT345fpKDAkKLJdVJQZ88Irl8sL780FKaUijRzevvhJaFP80QKBgQCNRfLlrLItj6NWqBUfJux4h1umzuoNWnze5kSf1sjAP1qKY82RuI909836s857TXc9QdpBSyo8m3fHtc8gMV1ZcQ3AtQnXj+5J2MKFWNtvpOony4NqgvlQ2+WSR7rbs9UJtgV2GE3ET8OYb6giWjmMKhYjsm3VWwa8qudW40jxlw==" ],
        "certificate" : [ "MIICnTCCAYUCBgFxIVsAmjANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDDAdyaXRjaGllMB4XDTIwMDMyODEzMzQwNFoXDTMwMDMyODEzMzU0NFowEjEQMA4GA1UEAwwHcml0Y2hpZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJLQ9YMpTB1DZBoOzwBJnzx7IjIo5tg9l7A4JO1f+Pafsf9Ym0x8LAZuuiHrxSqGahLNwrhqiUW6GIX0uKHt/zhuXIjP+ze02INH49Azmu9uo753foAsDrwroA/Z2bQYcXa162OXWbLN6YbwiOJDuPTZuRDOmZOTHZhad/DrjQ+PQakBXMf7NXpMLXITrs5rnPSgcd+ATD58Hd1/ZZP6KZrnVKR+EHyLkBCVdJCvCSmUb8gpC2+yU8J3S7Z30P14XsaFQqtnuPNPctbOagXGZqh9Q98BygDqDgOkiqpQrDJupwF8PUVXx/Be/6XSo58AFWk8Fp61bqTrhthe/7NcKZUCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAAP1GiZAhpUYg50Iij/7xpKr+ULyCClpH2G7rZw6RettvXlBEYx4u7OOI3iyUOP9UYUUUZZxRt6QUuVKSDxlCTZT0UJ2WxLXzDM5ctbLPIqW+YKupvy6s82x7MK5CK3CkoAxXl1BV1djUZqubKq5yL/6arbFIPAo8+7D3ynQjtoYEgHF7piVmPK9I4roh0z/o2cjlSGZRnwk6BnVLoj7PgCdtVPfhjDGzYI6+M9mjza65NfSU4UR19nJRQ7vS4XENNco7Lvl1qjQCBELV6tslcFr66PBBQZcZwr0jLlUZv/Ke92MTwLdPD4S0Ewr46rliwx4pf3TAhcHSEbjYokbDKw==" ],
        "priority" : [ "100" ]
      }
    } ]
  },
  "internationalizationEnabled" : false,
  "supportedLocales" : [ ],
  "authenticationFlows" : [ {
    "id" : "2bd21bd8-c9f0-4c09-b42c-bb9499d255dc",
    "alias" : "Handle Existing Account",
    "description" : "Handle what to do if there is existing account with same email/username like authenticated identity provider",
    "providerId" : "basic-flow",
    "topLevel" : false,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "idp-confirm-link",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "idp-email-verification",
      "requirement" : "ALTERNATIVE",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "requirement" : "ALTERNATIVE",
      "priority" : 30,
      "flowAlias" : "Verify Existing Account by Re-authentication",
      "userSetupAllowed" : false,
      "autheticatorFlow" : true
    } ]
  }, {
    "id" : "6ea70ee8-778e-436b-8def-cfb9ef8270c4",
    "alias" : "Verify Existing Account by Re-authentication",
    "description" : "Reauthentication of existing account",
    "providerId" : "basic-flow",
    "topLevel" : false,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "idp-username-password-form",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "auth-otp-form",
      "requirement" : "OPTIONAL",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "e041cee9-e5a1-4a04-b96e-945921b8a354",
    "alias" : "browser",
    "description" : "browser based authentication",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "auth-cookie",
      "requirement" : "ALTERNATIVE",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "auth-spnego",
      "requirement" : "DISABLED",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "identity-provider-redirector",
      "requirement" : "ALTERNATIVE",
      "priority" : 25,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "requirement" : "ALTERNATIVE",
      "priority" : 30,
      "flowAlias" : "forms",
      "userSetupAllowed" : false,
      "autheticatorFlow" : true
    } ]
  }, {
    "id" : "bcd79a3b-670c-48fa-aa6f-ba4c9772505b",
    "alias" : "clients",
    "description" : "Base authentication for clients",
    "providerId" : "client-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "client-secret",
      "requirement" : "ALTERNATIVE",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "client-jwt",
      "requirement" : "ALTERNATIVE",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "client-secret-jwt",
      "requirement" : "ALTERNATIVE",
      "priority" : 30,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "client-x509",
      "requirement" : "ALTERNATIVE",
      "priority" : 40,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "90bc9043-11c8-424b-a2b3-c00bdb66e6d8",
    "alias" : "direct grant",
    "description" : "OpenID Connect Resource Owner Grant",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "direct-grant-validate-username",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "direct-grant-validate-password",
      "requirement" : "REQUIRED",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "direct-grant-validate-otp",
      "requirement" : "OPTIONAL",
      "priority" : 30,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "cf2acc30-2051-4aba-97ec-9e829a03c336",
    "alias" : "docker auth",
    "description" : "Used by Docker clients to authenticate against the IDP",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "docker-http-basic-authenticator",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "0da67b8b-9e34-493b-8aaf-d72d16e604e2",
    "alias" : "first broker login",
    "description" : "Actions taken after first broker login with identity provider account, which is not yet linked to any Keycloak account",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticatorConfig" : "review profile config",
      "authenticator" : "idp-review-profile",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticatorConfig" : "create unique user config",
      "authenticator" : "idp-create-user-if-unique",
      "requirement" : "ALTERNATIVE",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "requirement" : "ALTERNATIVE",
      "priority" : 30,
      "flowAlias" : "Handle Existing Account",
      "userSetupAllowed" : false,
      "autheticatorFlow" : true
    } ]
  }, {
    "id" : "fc6972df-cb94-405d-a931-3acf2949fbec",
    "alias" : "forms",
    "description" : "Username, password, otp and other auth forms.",
    "providerId" : "basic-flow",
    "topLevel" : false,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "auth-username-password-form",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "auth-otp-form",
      "requirement" : "OPTIONAL",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "ac7cd5e9-4884-4985-9226-0371291e8b91",
    "alias" : "http challenge",
    "description" : "An authentication flow based on challenge-response HTTP Authentication Schemes",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "no-cookie-redirect",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "basic-auth",
      "requirement" : "REQUIRED",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "basic-auth-otp",
      "requirement" : "DISABLED",
      "priority" : 30,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "auth-spnego",
      "requirement" : "DISABLED",
      "priority" : 40,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "e58081b9-05f3-40fa-965f-edb38ce77f1b",
    "alias" : "registration",
    "description" : "registration flow",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "registration-page-form",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "flowAlias" : "registration form",
      "userSetupAllowed" : false,
      "autheticatorFlow" : true
    } ]
  }, {
    "id" : "0d18e9ed-7f3d-4f10-af6c-e0cb3a06bdd0",
    "alias" : "registration form",
    "description" : "registration form",
    "providerId" : "form-flow",
    "topLevel" : false,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "registration-user-creation",
      "requirement" : "REQUIRED",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "registration-profile-action",
      "requirement" : "REQUIRED",
      "priority" : 40,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "registration-password-action",
      "requirement" : "REQUIRED",
      "priority" : 50,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "registration-recaptcha-action",
      "requirement" : "DISABLED",
      "priority" : 60,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "88ce1f18-d6ad-410c-9a98-648edcc0bf88",
    "alias" : "reset credentials",
    "description" : "Reset credentials for a user if they forgot their password or something",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "reset-credentials-choose-user",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "reset-credential-email",
      "requirement" : "REQUIRED",
      "priority" : 20,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "reset-password",
      "requirement" : "REQUIRED",
      "priority" : 30,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    }, {
      "authenticator" : "reset-otp",
      "requirement" : "OPTIONAL",
      "priority" : 40,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  }, {
    "id" : "6046199f-0c1d-4d1a-95a3-e52534daf9ec",
    "alias" : "saml ecp",
    "description" : "SAML ECP Profile Authentication Flow",
    "providerId" : "basic-flow",
    "topLevel" : true,
    "builtIn" : true,
    "authenticationExecutions" : [ {
      "authenticator" : "http-basic-authenticator",
      "requirement" : "REQUIRED",
      "priority" : 10,
      "userSetupAllowed" : false,
      "autheticatorFlow" : false
    } ]
  } ],
  "authenticatorConfig" : [ {
    "id" : "69ddafd6-b43a-4128-987b-77c38fa6c9cc",
    "alias" : "create unique user config",
    "config" : {
      "require.password.update.after.registration" : "false"
    }
  }, {
    "id" : "23eb7edb-a04e-434d-b497-8e04cedfddbc",
    "alias" : "review profile config",
    "config" : {
      "update.profile.on.first.login" : "missing"
    }
  } ],
  "requiredActions" : [ {
    "alias" : "CONFIGURE_TOTP",
    "name" : "Configure OTP",
    "providerId" : "CONFIGURE_TOTP",
    "enabled" : true,
    "defaultAction" : false,
    "priority" : 10,
    "config" : { }
  }, {
    "alias" : "terms_and_conditions",
    "name" : "Terms and Conditions",
    "providerId" : "terms_and_conditions",
    "enabled" : false,
    "defaultAction" : false,
    "priority" : 20,
    "config" : { }
  }, {
    "alias" : "UPDATE_PASSWORD",
    "name" : "Update Password",
    "providerId" : "UPDATE_PASSWORD",
    "enabled" : true,
    "defaultAction" : false,
    "priority" : 30,
    "config" : { }
  }, {
    "alias" : "UPDATE_PROFILE",
    "name" : "Update Profile",
    "providerId" : "UPDATE_PROFILE",
    "enabled" : true,
    "defaultAction" : false,
    "priority" : 40,
    "config" : { }
  }, {
    "alias" : "VERIFY_EMAIL",
    "name" : "Verify Email",
    "providerId" : "VERIFY_EMAIL",
    "enabled" : true,
    "defaultAction" : false,
    "priority" : 50,
    "config" : { }
  } ],
  "browserFlow" : "browser",
  "registrationFlow" : "registration",
  "directGrantFlow" : "direct grant",
  "resetCredentialsFlow" : "reset credentials",
  "clientAuthenticationFlow" : "clients",
  "dockerAuthenticationFlow" : "docker auth",
  "attributes" : {
    "_browser_header.xXSSProtection" : "1; mode=block",
    "_browser_header.xFrameOptions" : "SAMEORIGIN",
    "_browser_header.strictTransportSecurity" : "max-age=31536000; includeSubDomains",
    "permanentLockout" : "false",
    "quickLoginCheckMilliSeconds" : "1000",
    "_browser_header.xRobotsTag" : "none",
    "maxFailureWaitSeconds" : "900",
    "minimumQuickLoginWaitSeconds" : "60",
    "failureFactor" : "30",
    "actionTokenGeneratedByUserLifespan" : "259200",
    "maxDeltaTimeSeconds" : "43200",
    "_browser_header.xContentTypeOptions" : "nosniff",
    "offlineSessionMaxLifespan" : "5184000",
    "actionTokenGeneratedByAdminLifespan" : "259200",
    "_browser_header.contentSecurityPolicyReportOnly" : "",
    "bruteForceProtected" : "false",
    "_browser_header.contentSecurityPolicy" : "frame-src 'self'; frame-ancestors 'self'; object-src 'none';",
    "waitIncrementSeconds" : "60",
    "offlineSessionMaxLifespanEnabled" : "false"
  },
  "keycloakVersion" : "7.0.0",
  "userManagedAccessAllowed" : false
}`
	TplKeycloakValues = `init:
  image:
    repository: busybox
    tag: 1.31
    pullPolicy: IfNotPresent
  resources: {}
clusterDomain: cluster.local
keycloak:
  replicas: 1
  image:
    repository: jboss/keycloak
    tag: 9.0.2
    pullPolicy: IfNotPresent
    pullSecrets: []
  hostAliases: []
  proxyAddressForwarding: true
  enableServiceLinks: false
  podManagementPolicy: Parallel
  restartPolicy: Always
  serviceAccount:
    create: false
    name:
  securityContext:
    fsGroup: 1000
  containerSecurityContext:
    runAsUser: 1000
    runAsNonRoot: true
  basepath: auth
  extraInitContainers: |
  extraContainers: |
  lifecycleHooks: |
    # postStart:
    #   exec:
    #     command: ["/bin/sh", "-c", "ls"]
  terminationGracePeriodSeconds: 60
  extraArgs:  -Dkeycloak.import=/realm/realm-secret -Dkeycloak.profile.feature.scripts=enabled -Dkeycloak.profile.feature.upload_scripts=enabled
  username: keycloak
  password: "{{ keycloakAdminPassword }}"
  existingSecret: ""
  existingSecretKey: password
  jgroups:
    discoveryProtocol: dns.DNS_PING
    discoveryProperties: >
      "dns_query={{ template "keycloak.fullname" . }}-headless.{{ .Release.Namespace }}.svc.{{ .Values.clusterDomain }}"
  javaToolOptions: >-
    -XX:+UseContainerSupport
    -XX:MaxRAMPercentage=50.0
  extraEnv: |
    - name: DB_USER
      value: "{{ keycloakDbUser }}"
    - name: DB_PASSWORD
      value: "{{ keycloakDbPassword }}"
    - name: DB_VENDOR
      value: "POSTGRES"
    - name: DB_PORT
      value: "5432"
    - name: DB_ADDR
      value: "{{ keycloakDbAddress }}"
    - name: DB_DATABASE
      value: "{{ keycloakDbDatabase }}"
    - name: DB_SCHEMA
      value: "public"
    - name: KEYCLOAK_USER
      value: "admin"
    - name: PROXY_ADDRESS_FORWARDING
      value: "true"
  affinity: |
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchLabels:
              {{- include "keycloak.selectorLabels" . | nindent 10 }}
            matchExpressions:
              - key: role
                operator: NotIn
                values:
                  - test
          topologyKey: kubernetes.io/hostname
      preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 100
          podAffinityTerm:
            labelSelector:
              matchLabels:
                {{- include "keycloak.selectorLabels" . | nindent 12 }}
              matchExpressions:
                - key: role
                  operator: NotIn
                  values:
                    - test
            topologyKey: failure-domain.beta.kubernetes.io/zone
  nodeSelector: {}
  priorityClassName: ""
  tolerations: []
  podLabels: {}
  podAnnotations: {}
  livenessProbe: |
    httpGet:
      path: {{ if ne .Values.keycloak.basepath "" }}/{{ .Values.keycloak.basepath }}{{ end }}/
      port: http
    initialDelaySeconds: 300
    timeoutSeconds: 5
  readinessProbe: |
    httpGet:
      path: {{ if ne .Values.keycloak.basepath "" }}/{{ .Values.keycloak.basepath }}{{ end }}/realms/master
      port: http
    initialDelaySeconds: 30
    timeoutSeconds: 1
  resources:
    limits:
      cpu: "1024m"
      memory: "1024Mi"
    requests:
      cpu: "256m"
      memory: "512Mi"
  cli:
    enabled: true
    nodeIdentifier: |
      {{ .Files.Get "scripts/node-identifier.cli" }}
    logging: |
      {{ .Files.Get "scripts/logging.cli" }}
    ha: |
      {{ .Files.Get "scripts/ha.cli" }}
    datasource: |
      {{ .Files.Get "scripts/datasource.cli" }}
    custom: |
  startupScripts: {}
  extraVolumes: |
    - name: realm-secret
      secret:
        secretName: realm-secret
  extraVolumeMounts: |
    - name: realm-secret
      mountPath: "/realm/"
      readOnly: true
  extraPorts: |
  podDisruptionBudget: {}
  service:
    annotations: {}
    labels: {}
    type: ClusterIP
    httpPort: 80
    httpsPort: 8443
    jgroupsPort: 7600
  ingress:
    enabled: false
    path: /
    annotations: {}
    labels: {}
    hosts:
      - keycloak.example.com
    tls: []
  route:
    enabled: false
    path: /
    annotations: {}
    labels: {}
    host:
    tls:
      enabled: true
      insecureEdgeTerminationPolicy: Redirect
      termination: edge
  persistence:
    deployPostgres: false
    dbVendor: h2
    dbName: keycloak
    dbHost: mykeycloak
    dbPort: 5432
    existingSecret: ""
    existingSecretPasswordKey: ""
    existingSecretUsernameKey: ""
    dbUser: keycloak
    dbPassword: ""
postgresql:
  postgresqlUsername: keycloak
  postgresqlPassword: ""
  postgresqlDatabase: keycloak
  persistence:
    enabled: false
test:
  enabled: true
  image:
    repository: unguiculus/docker-python3-phantomjs-selenium
    tag: v1
    pullPolicy: IfNotPresent
  securityContext:
    fsGroup: 1000
  containerSecurityContext:
    runAsUser: 1000
    runAsNonRoot: true
prometheus:
  operator:
    enabled: false
    serviceMonitor:
      namespace: ""
      selector:
        release: prometheus
      interval: 10s
      scrapeTimeout: 10s
      path: /auth/realms/master/metrics
    prometheusRules:
      enabled: false
      selector:
        app: prometheus-operator
        release: prometheus
      rules: {}`
	TplRitchieValues = `envVars:
  - name: VAULT_ADDR
    value: "http://{{ vaultRelease }}:8200"
  - name: VAULT_AUTHENTICATION
    value: "KUBERNETES"
  - name: VAULT_K8S_ROLE
    value: "ritchie_credential_role"
  - name: FILE_CONFIG
    value: "/data/conf/ritchie.conf"
replicaCount: 1
image:
  repository: ritchie-server
  tag: qa-160
  pullPolicy: IfNotPresent
imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""
serviceaccount: vault-auth
podSecurityContext: {}
securityContext: {}
resources:
  limits:
    cpu: 512m
    memory: 1024Mi
  requests:
    cpu: 100m
    memory: 128Mi
nodeSelector: {}
tolerations: []
affinity: {}
livenessProbe:
  enabled: true
  failureThreshold: 3
  httpGet:
    path: /health
    port: 3000
    scheme: HTTP
  initialDelaySeconds: 10
  periodSeconds: 30
  successThreshold: 1
  timeoutSeconds: 1
readinessProbe:
  failureThreshold: 3
  httpGet:
    path: /health
    port: 3000
    scheme: HTTP
  initialDelaySeconds: 5
  periodSeconds: 10
  successThreshold: 1
  timeoutSeconds: 1
service:
  name: ritchie-server
  type: NodePort
  ports:
  - name: http
    port: 3000
    target_port: 3000
ritchie:
  orgs:
    - name: {{ ritchieOrganization }}
      keycloak:
        url: "{{ keycloakExternalAddress }}"
        realm: "{{ keycloakRealm }}"
        clientId: "{{ keycloakClientId }}"
        clientSecret: "{{ keycloakClientSecret }}"
      repositories: |
        [
          {
            "name": "commons-repo",
            "priority": 0,
            "treePath": "https://commons-repo.ritchiecli.io/tree/tree.json"
          }
        ]

      cliVersionPath:
        provider: "{{ ritchieCliVersionProvider }}"
        url: "{{ ritchieCliVersionURL }}"
      oauth :
        url: "{{ keycloakExternalAddress }}/auth/realms/{{ keycloakRealm }}"
        clientId: "oauth"
      credentials: |

        {
          "github": [
            {
              "field": "username",
              "type": "text"
            },
            {
              "field": "token",
              "type": "password"
            }
          ],
          "aws": [
            {
              "field": "accessKeyId",
              "type": "text"
            },
            {
              "field": "secretAccessKey",
              "type": "password"
            }
          ],
          "kubeconfig": [
            {
                "field": "base64config",
                "type" : "text"
            }
          ]
        }
imageCredentials:
  registry: 237930432518.dkr.ecr.sa-east-1.amazonaws.com
`

	TplVaultValues = `global:
  enabled: true
  image: "vault:1.1.0"
  imagePullPolicy: IfNotPresent
  imagePullSecrets: []
  tlsDisable: true
server:
  securityContext:
    readOnlyRootFilesystem: true
  resources:
     requests:
       memory: 256Mi
       cpu: 250m
     limits:
       memory: 256Mi
       cpu: 250m
  ingress:
    enabled: false
    annotations: {}
    hosts:
      - host: chart-example.local
        paths: []
    tls: []
  authDelegator:
    enabled: false
  extraEnvironmentVars: {}
  extraSecretEnvironmentVars: []
  extraVolumes: []
  affinity: |
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        - labelSelector:
            matchLabels:
              app.kubernetes.io/name: {{ template "vault.name" . }}
              app.kubernetes.io/instance: "{{ .Release.Name }}"
              component: server
          topologyKey: kubernetes.io/hostname
  tolerations: {}
  nodeSelector: {}
  extraLabels: {}
  annotations: {}
  service:
    enabled: true
    port: 8200
    targetPort: 8200
    annotations: {}
  dataStorage:
    enabled: true
    size: 10Gi
    storageClass: null
    accessMode: ReadWriteOnce
  auditStorage:
    enabled: false
    size: 10Gi
    storageClass: null
    accessMode: ReadWriteOnce
  dev:
    enabled: false
  standalone:
    enabled: false
    config: |
      ui = true

      listener "tcp" {
        tls_disable = 1
        address = "[::]:8200"
        cluster_address = "[::]:8201"
      }
      storage "file" {
        path = "/vault/data"
      }
  ha:
    enabled: true
    replicas: 2
    config: |
      ui = true

      listener "tcp" {
        tls_disable = 1
        address = "[::]:8200"
        cluster_address = "[::]:8201"
      }
      storage "consul" {
        path = "vault"
        address = "{{ consulRelease }}:8500"
      }
    disruptionBudget:
      enabled: true
      maxUnavailable: null
  serviceaccount:
    annotations: {}
ui:
  enabled: false
  serviceType: "ClusterIP"
  serviceNodePort: null
  externalPort: 8200
  annotations: {}`
)
