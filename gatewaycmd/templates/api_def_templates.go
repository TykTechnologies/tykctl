package templates

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/TykTechnologies/gateway-sdk/pkg/apim"
	"github.com/TykTechnologies/tykctl/util"
)

func createLeanKeylessAPIDefinition() *apim.APIDefinition {
	api := apim.APIDefinition{}
	api.Active = util.GetBoolPtr(true)
	api.Name = util.GetStrPtr("test1")
	api.ApiId = util.GetStrPtr("test1")
	api.Slug = util.GetStrPtr("test1")
	api.UseKeyless = util.GetBoolPtr(true)
	api.Proxy = new(apim.APIDefinitionProxy)
	api.Proxy.ListenPath = util.GetStrPtr("/test1/")
	api.Proxy.TargetUrl = util.GetStrPtr("https://httpbin.org")
	api.Proxy.StripListenPath = util.GetBoolPtr(true)
	api.Proxy.PreserveHostHeader = util.GetBoolPtr(true)
	api.OrgId = util.GetStrPtr("default")
	api.VersionData = new(apim.APIDefinitionVersionData)
	api.VersionData.NotVersioned = util.GetBoolPtr(true)
	dt := map[string]apim.VersionInfo{
		"Default": {
			Name: util.GetStrPtr("Default"),
		},
	}
	api.VersionData.Versions = &dt

	return &api
}

func createFullProtectedAPIDefinition() (*apim.APIDefinition, error) {
	definitions := `{
		"api_id": "4a77d8bfe76f41ad7ae5875b2259df3f",
		"jwt_issued_at_validation_skew": 0,
		"upstream_certificates": {},
		"use_keyless": true,
		"enable_coprocess_auth": false,
		"base_identity_provided_by": "auth_token",
		"custom_middleware": {
		  "pre": [],
		  "post": [],
		  "post_key_auth": [],
		  "auth_check": {
			"name": "",
			"path": "",
			"require_session": false,
			"raw_body_only": false
		  },
		  "response": [],
		  "driver": "",
		  "id_extractor": {
			"extract_from": "",
			"extract_with": "",
			"extractor_config": {}
		  }
		},
		"disable_quota": false,
		"custom_middleware_bundle": "",
		"cache_options": {
		  "cache_timeout": 60,
		  "enable_cache": true,
		  "cache_all_safe_requests": false,
		  "cache_response_codes": [],
		  "enable_upstream_cache_control": false,
		  "cache_control_ttl_header": "",
		  "cache_by_headers": []
		},
		"enable_ip_blacklisting": false,
		"tag_headers": [],
		"jwt_scope_to_policy_mapping": {},
		"pinned_public_keys": {},
		"expire_analytics_after": 0,
		"domain": "",
		"openid_options": {
		  "providers": [],
		  "segregate_by_client": false
		},
		"jwt_policy_field_name": "",
		"enable_proxy_protocol": false,
		"jwt_default_policies": [],
		"active": true,
		"jwt_expires_at_validation_skew": 0,
		"config_data": {},
		"notifications": {
		  "shared_secret": "",
		  "oauth_on_keychange_url": ""
		},
		"jwt_client_base_field": "",
		"auth": {
		  "use_param": false,
		  "param_name": "",
		  "use_cookie": false,
		  "cookie_name": "",
		  "auth_header_name": "Authorization",
		  "use_certificate": false,
		  "validate_signature": false,
		  "signature": {
			"algorithm": "",
			"header": "",
			"secret": "",
			"allowed_clock_skew": 0,
			"error_code": 0,
			"error_message": ""
		  }
		},
		"check_host_against_uptime_tests": false,
		"auth_provider": {
		  "name": "",
		  "storage_engine": "",
		  "meta": {}
		},
		"blacklisted_ips": [],
		"graphql": {
		  "enabled": false,
		  "execution_mode": "proxyOnly",
		  "version": "2",
		  "schema": "",
		  "type_field_configurations": [],
		  "playground": {
			"enabled": false,
			"path": ""
		  },
		  "engine": {
			"field_configs": [],
			"data_sources": []
		  }
		},
		"hmac_allowed_clock_skew": -1,
		"dont_set_quota_on_create": false,
		"uptime_tests": {
		  "check_list": [],
		  "config": {
			"expire_utime_after": 0,
			"service_discovery": {
			  "use_discovery_service": false,
			  "query_endpoint": "",
			  "use_nested_query": false,
			  "parent_data_path": "",
			  "data_path": "",
			  "cache_timeout": 60
			},
			"recheck_wait": 0
		  }
		},
		"enable_jwt": false,
		"do_not_track": false,
		"name": "mTls",
		"slug": "mtls",
		"oauth_meta": {
		  "allowed_access_types": [],
		  "allowed_authorize_types": [],
		  "auth_login_redirect": ""
		},
		"CORS": {
		  "enable": false,
		  "max_age": 24,
		  "allow_credentials": false,
		  "exposed_headers": [],
		  "allowed_headers": [
			"Origin",
			"Accept",
			"Content-Type",
			"X-Requested-With",
			"Authorization"
		  ],
		  "options_passthrough": false,
		  "debug": false,
		  "allowed_origins": [
			"*"
		  ],
		  "allowed_methods": [
			"GET",
			"POST",
			"HEAD"
		  ]
		},
		"event_handlers": {
		  "events": {}
		},
		"proxy": {
		  "target_url": "https://httpbin.org",
		  "service_discovery": {
			"endpoint_returns_list": false,
			"cache_timeout": 0,
			"parent_data_path": "",
			"query_endpoint": "",
			"use_discovery_service": false,
			"_sd_show_port_path": false,
			"target_path": "",
			"use_target_list": false,
			"use_nested_query": false,
			"data_path": "",
			"port_data_path": ""
		  },
		  "check_host_against_uptime_tests": false,
		  "transport": {
			"ssl_insecure_skip_verify": false,
			"ssl_min_version": 0,
			"proxy_url": "",
			"ssl_ciphers": []
		  },
		  "target_list": [],
		  "preserve_host_header": false,
		  "strip_listen_path": true,
		  "enable_load_balancing": false,
		  "listen_path": "/mtls/",
		  "disable_strip_slash": true
		},
		"client_certificates": [
		  "certs/concat.pem"
		],
		"use_basic_auth": false,
		"version_data": {
		  "not_versioned": true,
		  "default_version": "",
		  "versions": {
			"Default": {
			  "name": "Default",
			  "expires": "",
			  "paths": {
				"ignored": [],
				"white_list": [],
				"black_list": []
			  },
			  "use_extended_paths": true,
			  "extended_paths": {
				"ignored": [],
				"white_list": [],
				"black_list": [],
				"transform": [],
				"transform_response": [],
				"transform_jq": [],
				"transform_jq_response": [],
				"transform_headers": [],
				"transform_response_headers": [],
				"hard_timeouts": [],
				"circuit_breakers": [],
				"url_rewrites": [],
				"virtual": [],
				"size_limits": [],
				"method_transforms": [],
				"track_endpoints": [],
				"do_not_track_endpoints": [],
				"validate_json": [],
				"internal": []
			  },
			  "global_headers": {},
			  "global_headers_remove": [],
			  "global_response_headers": {},
			  "global_response_headers_remove": [],
			  "ignore_endpoint_case": false,
			  "global_size_limit": 0,
			  "override_target": ""
			}
		  }
		},
		"jwt_scope_claim_name": "",
		"use_standard_auth": false,
		"session_lifetime": 0,
		"hmac_allowed_algorithms": [],
		"disable_rate_limit": false,
		"definition": {
		  "location": "header",
		  "key": "x-api-version",
		  "strip_path": false
		},
		"use_oauth2": false,
		"jwt_source": "",
		"jwt_signing_method": "",
		"jwt_not_before_validation_skew": 0,
		"use_go_plugin_auth": false,
		"jwt_identity_base_field": "",
		"allowed_ips": [],
		"request_signing": {
		  "is_enabled": false,
		  "secret": "",
		  "key_id": "",
		  "algorithm": "",
		  "header_list": [],
		  "certificate_id": "",
		  "signature_header": ""
		},
		"org_id": "5e9d9544a1dcd60001d0ed20",
		"enable_ip_whitelisting": false,
		"global_rate_limit": {
		  "rate": 0,
		  "per": 0
		},
		"protocol": "",
		"enable_context_vars": false,
		"tags": [],
		"basic_auth": {
		  "disable_caching": false,
		  "cache_ttl": 0,
		  "extract_from_body": false,
		  "body_user_regexp": "",
		  "body_password_regexp": ""
		},
		"listen_port": 0,
		"session_provider": {
		  "name": "",
		  "storage_engine": "",
		  "meta": {}
		},
		"auth_configs": {
		  "authToken": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "basic": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "coprocess": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "hmac": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "jwt": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "oauth": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  },
		  "oidc": {
			"use_param": false,
			"param_name": "",
			"use_cookie": false,
			"cookie_name": "",
			"auth_header_name": "Authorization",
			"use_certificate": false,
			"validate_signature": false,
			"signature": {
			  "algorithm": "",
			  "header": "",
			  "secret": "",
			  "allowed_clock_skew": 0,
			  "error_code": 0,
			  "error_message": ""
			}
		  }
		},
		"strip_auth_data": false,
		"id": "60522b32a74aee000171cbbb",
		"certificates": [],
		"enable_signature_checking": false,
		"use_openid": false,
		"internal": false,
		"jwt_skip_kid": false,
		"enable_batch_request_support": false,
		"enable_detailed_recording": false,
		"response_processors": [],
		"use_mutual_tls_auth": true
	  }`

	api := apim.NewAPIDefinition()

	err := json.Unmarshal([]byte(definitions), &api)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func SaveTemplateToFile(name, dir string, body interface{}) error {
	fileName := name + ".json"
	path := filepath.Join(dir, fileName)

	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("template %s already exists", name)
	}

	if !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	bytes, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)

	return err
}
