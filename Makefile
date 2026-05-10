
generate:
	openapi-generator generate -i $(SPEC) -g go --package-name $(PACKAGE) --git-host github.com --git-user-id CiscoDevNet --git-repo-id go-ciscosecureaccess -o $(PACKAGE) --additional-properties=isGoSubmodule=true --type-mappings integer=int64,int=int64,number=int64 --openapi-normalizer KEEP_ONLY_FIRST_TAG_IN_OPERATION=true --additional-properties=disallowAdditionalPropertiesIfNotPresent=false $(CLI_EXTRA)
	rm -f $(PACKAGE)/go.mod $(PACKAGE)/go.sum

generate-destinationlists:
	make generate PACKAGE=destinationlists SPEC=./specs/cisco_secure_access_destination_lists_api_1_0_0.yaml

generate-internaldomains:
	make generate PACKAGE=internaldomains SPEC=./specs/cisco_secure_access_internal_domains_api_2_0_0.yaml

generate-internalnetworks:
	make generate PACKAGE=internalnetworks SPEC=./specs/cisco_secure_access_internal_networks_api_2_0_0.yaml

generate-networks:
	make generate PACKAGE=networks SPEC=./specs/cisco_secure_access_networks_api_2_0_0.yaml

generate-ntg:
	make generate PACKAGE=ntg SPEC=./specs/cisco_secure_access_network_tunnel_groups_and_regions_api_1_0_0.yaml CLI_EXTRA="$$(cat specs/ntg_name_mapping.txt)"

generate-privateapps:
	make generate PACKAGE=privateapps SPEC=./specs/cisco_secure_access_private_resources_and_resource_groups_api_1_0_0.yaml

generate-reports:
	make generate PACKAGE=reports SPEC=./specs/cisco_secure_access_reports.yaml CLI_EXTRA="--additional-properties=useTags=false --additional-properties=disallowAdditionalPropertiesIfNotPresent=false --additional-properties=enumClassPrefix=true --openapi-generator-ignore-list api_secure_access.go"

generate-resconn:
	make generate PACKAGE=resconn SPEC=./specs/cisco_secure_access_resource_connector_groups_and_resource_connectors_api_1_0_0.yaml

generate-roaming:
	make generate PACKAGE=roaming SPEC=./specs/cisco_secure_access_roaming_computers_api_2_0_0.yaml

generate-rules:
	make generate PACKAGE=rules SPEC=./specs/cisco_secure_access_policy_rules_api_1_0_0.yaml CLI_EXTRA="--additional-properties=enumClassPrefix=true"

generate-swg:
	make generate PACKAGE=swg SPEC=./specs/cisco_secure_web_gateway_device_settings_api_2_0_0.yaml

generate-sites:
	make generate PACKAGE=sites SPEC=./specs/cisco_secure_access_sites_api_2_0_0.yaml

generate-virtualappliances:
	make generate PACKAGE=virtualappliances SPEC=./specs/cisco_secure_access_virtual_appliances_api_2_0_0.yaml

generate-all:
	for spec in destinationlists internaldomains internalnetworks networks ntg privateapps reports resconn roaming rules swg virtualappliances; do \
	make generate-$${spec} ;\
	done

test:
	GO_CISCOSECUREACCESS_RUN_FUNCTIONAL=true go test -v -tags=functional ./tests/*_test.go -count 1
