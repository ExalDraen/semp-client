// Code generated by go-swagger; DO NOT EDIT.

package msg_vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new msg vpn API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for msg vpn API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateMsgVpn creates a message v p n object

Creates a Message VPN object. Any attribute missing from the request will be set to its default value.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
replicationBridgeAuthenticationBasicPassword||||x|
replicationBridgeAuthenticationClientCertContent||||x|
replicationBridgeAuthenticationClientCertPassword||||x|
replicationEnabledQueueBehavior||||x|



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
EventThreshold|clearPercent|setPercent|clearValue, setValue
EventThreshold|clearValue|setValue|clearPercent, setPercent
EventThreshold|setPercent|clearPercent|clearValue, setValue
EventThreshold|setValue|clearValue|clearPercent, setPercent
EventThresholdByValue|clearValue|setValue|
EventThresholdByValue|setValue|clearValue|
MsgVpn|authenticationBasicProfileName|authenticationBasicType|
MsgVpn|authorizationProfileName|authorizationType|
MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled|
MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled|
MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword|
MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername|
MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent|
MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|



A SEMP client authorized with a minimum access scope/level of "global/readwrite" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) CreateMsgVpn(params *CreateMsgVpnParams, authInfo runtime.ClientAuthInfoWriter) (*CreateMsgVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMsgVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMsgVpn",
		Method:             "POST",
		PathPattern:        "/msgVpns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMsgVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateMsgVpnOK), nil

}

/*
CreateMsgVpnSequencedTopic creates a sequenced topic object

Creates a Sequenced Topic object. Any attribute missing from the request will be set to its default value.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x||x||
sequencedTopic|x|x|||



A SEMP client authorized with a minimum access scope/level of "vpn/readwrite" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) CreateMsgVpnSequencedTopic(params *CreateMsgVpnSequencedTopicParams, authInfo runtime.ClientAuthInfoWriter) (*CreateMsgVpnSequencedTopicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMsgVpnSequencedTopicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMsgVpnSequencedTopic",
		Method:             "POST",
		PathPattern:        "/msgVpns/{msgVpnName}/sequencedTopics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMsgVpnSequencedTopicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateMsgVpnSequencedTopicOK), nil

}

/*
DeleteMsgVpn deletes a message v p n object

Deletes a Message VPN object.

A SEMP client authorized with a minimum access scope/level of "global/readwrite" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) DeleteMsgVpn(params *DeleteMsgVpnParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMsgVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMsgVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMsgVpn",
		Method:             "DELETE",
		PathPattern:        "/msgVpns/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMsgVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMsgVpnOK), nil

}

/*
DeleteMsgVpnSequencedTopic deletes a sequenced topic object

Deletes a Sequenced Topic object.

A SEMP client authorized with a minimum access scope/level of "vpn/readwrite" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) DeleteMsgVpnSequencedTopic(params *DeleteMsgVpnSequencedTopicParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMsgVpnSequencedTopicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMsgVpnSequencedTopicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMsgVpnSequencedTopic",
		Method:             "DELETE",
		PathPattern:        "/msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMsgVpnSequencedTopicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMsgVpnSequencedTopicOK), nil

}

/*
GetMsgVpn gets a message v p n object

Gets a Message VPN object.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
replicationBridgeAuthenticationBasicPassword||x|
replicationBridgeAuthenticationClientCertContent||x|
replicationBridgeAuthenticationClientCertPassword||x|
replicationEnabledQueueBehavior||x|



A SEMP client authorized with a minimum access scope/level of "vpn/readonly" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) GetMsgVpn(params *GetMsgVpnParams, authInfo runtime.ClientAuthInfoWriter) (*GetMsgVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMsgVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMsgVpn",
		Method:             "GET",
		PathPattern:        "/msgVpns/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMsgVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMsgVpnOK), nil

}

/*
GetMsgVpnSequencedTopic gets a sequenced topic object

Gets a Sequenced Topic object.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
sequencedTopic|x||



A SEMP client authorized with a minimum access scope/level of "vpn/readonly" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) GetMsgVpnSequencedTopic(params *GetMsgVpnSequencedTopicParams, authInfo runtime.ClientAuthInfoWriter) (*GetMsgVpnSequencedTopicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMsgVpnSequencedTopicParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMsgVpnSequencedTopic",
		Method:             "GET",
		PathPattern:        "/msgVpns/{msgVpnName}/sequencedTopics/{sequencedTopic}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMsgVpnSequencedTopicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMsgVpnSequencedTopicOK), nil

}

/*
GetMsgVpnSequencedTopics gets a list of sequenced topic objects

Gets a list of Sequenced Topic objects.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
sequencedTopic|x||



A SEMP client authorized with a minimum access scope/level of "vpn/readonly" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) GetMsgVpnSequencedTopics(params *GetMsgVpnSequencedTopicsParams, authInfo runtime.ClientAuthInfoWriter) (*GetMsgVpnSequencedTopicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMsgVpnSequencedTopicsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMsgVpnSequencedTopics",
		Method:             "GET",
		PathPattern:        "/msgVpns/{msgVpnName}/sequencedTopics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMsgVpnSequencedTopicsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMsgVpnSequencedTopicsOK), nil

}

/*
GetMsgVpns gets a list of message v p n objects

Gets a list of Message VPN objects.


Attribute|Identifying|Write-Only|Deprecated
:---|:---:|:---:|:---:
msgVpnName|x||
replicationBridgeAuthenticationBasicPassword||x|
replicationBridgeAuthenticationClientCertContent||x|
replicationBridgeAuthenticationClientCertPassword||x|
replicationEnabledQueueBehavior||x|



A SEMP client authorized with a minimum access scope/level of "vpn/readonly" is required to perform this operation.

This has been available since 2.0.
*/
func (a *Client) GetMsgVpns(params *GetMsgVpnsParams, authInfo runtime.ClientAuthInfoWriter) (*GetMsgVpnsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMsgVpnsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMsgVpns",
		Method:             "GET",
		PathPattern:        "/msgVpns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMsgVpnsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMsgVpnsOK), nil

}

/*
ReplaceMsgVpn replaces a message v p n object

Replaces a Message VPN object. Any attribute missing from the request will be set to its default value, unless the user is not authorized to change its value in which case the missing attribute will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
replicationBridgeAuthenticationBasicPassword|||x||
replicationBridgeAuthenticationClientCertContent|||x||
replicationBridgeAuthenticationClientCertPassword|||x||
replicationEnabledQueueBehavior|||x||



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
EventThreshold|clearPercent|setPercent|clearValue, setValue
EventThreshold|clearValue|setValue|clearPercent, setPercent
EventThreshold|setPercent|clearPercent|clearValue, setValue
EventThreshold|setValue|clearValue|clearPercent, setPercent
EventThresholdByValue|clearValue|setValue|
EventThresholdByValue|setValue|clearValue|
MsgVpn|authenticationBasicProfileName|authenticationBasicType|
MsgVpn|authorizationProfileName|authorizationType|
MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled|
MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled|
MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword|
MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername|
MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent|
MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|



A SEMP client authorized with a minimum access scope/level of "vpn/readwrite" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
authenticationBasicEnabled|global/readwrite
authenticationBasicProfileName|global/readwrite
authenticationBasicRadiusDomain|global/readwrite
authenticationBasicType|global/readwrite
authenticationClientCertAllowApiProvidedUsernameEnabled|global/readwrite
authenticationClientCertEnabled|global/readwrite
authenticationClientCertMaxChainDepth|global/readwrite
authenticationClientCertRevocationCheckMode|global/readwrite
authenticationClientCertUsernameSource|global/readwrite
authenticationClientCertValidateDateEnabled|global/readwrite
authenticationKerberosAllowApiProvidedUsernameEnabled|global/readwrite
authenticationKerberosEnabled|global/readwrite
bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite
bridgingTlsServerCertMaxChainDepth|global/readwrite
bridgingTlsServerCertValidateDateEnabled|global/readwrite
dmrEnabled|global/readwrite
exportSubscriptionsEnabled|global/readwrite
maxConnectionCount|global/readwrite
maxEgressFlowCount|global/readwrite
maxEndpointCount|global/readwrite
maxIngressFlowCount|global/readwrite
maxMsgSpoolUsage|global/readwrite
maxSubscriptionCount|global/readwrite
maxTransactedSessionCount|global/readwrite
maxTransactionCount|global/readwrite
replicationBridgeAuthenticationBasicClientUsername|global/readwrite
replicationBridgeAuthenticationBasicPassword|global/readwrite
replicationBridgeAuthenticationClientCertContent|global/readwrite
replicationBridgeAuthenticationClientCertPassword|global/readwrite
replicationBridgeAuthenticationScheme|global/readwrite
replicationBridgeCompressedDataEnabled|global/readwrite
replicationBridgeEgressFlowWindowSize|global/readwrite
replicationBridgeRetryDelay|global/readwrite
replicationBridgeTlsEnabled|global/readwrite
replicationBridgeUnidirectionalClientProfileName|global/readwrite
replicationEnabled|global/readwrite
replicationEnabledQueueBehavior|global/readwrite
replicationQueueMaxMsgSpoolUsage|global/readwrite
replicationRole|global/readwrite
restTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite
restTlsServerCertMaxChainDepth|global/readwrite
restTlsServerCertValidateDateEnabled|global/readwrite
sempOverMsgBusAdminClientEnabled|global/readwrite
sempOverMsgBusAdminDistributedCacheEnabled|global/readwrite
sempOverMsgBusAdminEnabled|global/readwrite
sempOverMsgBusEnabled|global/readwrite
sempOverMsgBusShowEnabled|global/readwrite
serviceRestIncomingMaxConnectionCount|global/readwrite
serviceRestIncomingPlainTextListenPort|global/readwrite
serviceRestIncomingTlsListenPort|global/readwrite
serviceRestOutgoingMaxConnectionCount|global/readwrite
serviceSmfMaxConnectionCount|global/readwrite
serviceWebMaxConnectionCount|global/readwrite



This has been available since 2.0.
*/
func (a *Client) ReplaceMsgVpn(params *ReplaceMsgVpnParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceMsgVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceMsgVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceMsgVpn",
		Method:             "PUT",
		PathPattern:        "/msgVpns/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReplaceMsgVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceMsgVpnOK), nil

}

/*
UpdateMsgVpn updates a message v p n object

Updates a Message VPN object. Any attribute missing from the request will be left unchanged.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated
:---|:---:|:---:|:---:|:---:|:---:
msgVpnName|x|x|||
replicationBridgeAuthenticationBasicPassword|||x||
replicationBridgeAuthenticationClientCertContent|||x||
replicationBridgeAuthenticationClientCertPassword|||x||
replicationEnabledQueueBehavior|||x||



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
EventThreshold|clearPercent|setPercent|clearValue, setValue
EventThreshold|clearValue|setValue|clearPercent, setPercent
EventThreshold|setPercent|clearPercent|clearValue, setValue
EventThreshold|setValue|clearValue|clearPercent, setPercent
EventThresholdByValue|clearValue|setValue|
EventThresholdByValue|setValue|clearValue|
MsgVpn|authenticationBasicProfileName|authenticationBasicType|
MsgVpn|authorizationProfileName|authorizationType|
MsgVpn|eventPublishTopicFormatMqttEnabled|eventPublishTopicFormatSmfEnabled|
MsgVpn|eventPublishTopicFormatSmfEnabled|eventPublishTopicFormatMqttEnabled|
MsgVpn|replicationBridgeAuthenticationBasicClientUsername|replicationBridgeAuthenticationBasicPassword|
MsgVpn|replicationBridgeAuthenticationBasicPassword|replicationBridgeAuthenticationBasicClientUsername|
MsgVpn|replicationBridgeAuthenticationClientCertPassword|replicationBridgeAuthenticationClientCertContent|
MsgVpn|replicationEnabledQueueBehavior|replicationEnabled|



A SEMP client authorized with a minimum access scope/level of "vpn/readwrite" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
authenticationBasicEnabled|global/readwrite
authenticationBasicProfileName|global/readwrite
authenticationBasicRadiusDomain|global/readwrite
authenticationBasicType|global/readwrite
authenticationClientCertAllowApiProvidedUsernameEnabled|global/readwrite
authenticationClientCertEnabled|global/readwrite
authenticationClientCertMaxChainDepth|global/readwrite
authenticationClientCertRevocationCheckMode|global/readwrite
authenticationClientCertUsernameSource|global/readwrite
authenticationClientCertValidateDateEnabled|global/readwrite
authenticationKerberosAllowApiProvidedUsernameEnabled|global/readwrite
authenticationKerberosEnabled|global/readwrite
bridgingTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite
bridgingTlsServerCertMaxChainDepth|global/readwrite
bridgingTlsServerCertValidateDateEnabled|global/readwrite
dmrEnabled|global/readwrite
exportSubscriptionsEnabled|global/readwrite
maxConnectionCount|global/readwrite
maxEgressFlowCount|global/readwrite
maxEndpointCount|global/readwrite
maxIngressFlowCount|global/readwrite
maxMsgSpoolUsage|global/readwrite
maxSubscriptionCount|global/readwrite
maxTransactedSessionCount|global/readwrite
maxTransactionCount|global/readwrite
replicationBridgeAuthenticationBasicClientUsername|global/readwrite
replicationBridgeAuthenticationBasicPassword|global/readwrite
replicationBridgeAuthenticationClientCertContent|global/readwrite
replicationBridgeAuthenticationClientCertPassword|global/readwrite
replicationBridgeAuthenticationScheme|global/readwrite
replicationBridgeCompressedDataEnabled|global/readwrite
replicationBridgeEgressFlowWindowSize|global/readwrite
replicationBridgeRetryDelay|global/readwrite
replicationBridgeTlsEnabled|global/readwrite
replicationBridgeUnidirectionalClientProfileName|global/readwrite
replicationEnabled|global/readwrite
replicationEnabledQueueBehavior|global/readwrite
replicationQueueMaxMsgSpoolUsage|global/readwrite
replicationRole|global/readwrite
restTlsServerCertEnforceTrustedCommonNameEnabled|global/readwrite
restTlsServerCertMaxChainDepth|global/readwrite
restTlsServerCertValidateDateEnabled|global/readwrite
sempOverMsgBusAdminClientEnabled|global/readwrite
sempOverMsgBusAdminDistributedCacheEnabled|global/readwrite
sempOverMsgBusAdminEnabled|global/readwrite
sempOverMsgBusEnabled|global/readwrite
sempOverMsgBusShowEnabled|global/readwrite
serviceRestIncomingMaxConnectionCount|global/readwrite
serviceRestIncomingPlainTextListenPort|global/readwrite
serviceRestIncomingTlsListenPort|global/readwrite
serviceRestOutgoingMaxConnectionCount|global/readwrite
serviceSmfMaxConnectionCount|global/readwrite
serviceWebMaxConnectionCount|global/readwrite



This has been available since 2.0.
*/
func (a *Client) UpdateMsgVpn(params *UpdateMsgVpnParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateMsgVpnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMsgVpnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMsgVpn",
		Method:             "PATCH",
		PathPattern:        "/msgVpns/{msgVpnName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateMsgVpnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateMsgVpnOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}