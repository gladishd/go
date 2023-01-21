// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationdiscoveryservice provides the client and types for making API
// requests to AWS Application Discovery Service.
//
// AWS Application Discovery Service helps you plan application migration projects.
// It automatically identifies servers, virtual machines (VMs), and network
// dependencies in your on-premises data centers. For more information, see
// the AWS Application Discovery Service FAQ (http://aws.amazon.com/application-discovery/faqs/).
// Application Discovery Service offers three ways of performing discovery and
// collecting data about your on-premises servers:
//
//    * Agentless discovery is recommended for environments that use VMware
//    vCenter Server. This mode doesn't require you to install an agent on each
//    host. It does not work in non-VMware environments. Agentless discovery
//    gathers server information regardless of the operating systems, which
//    minimizes the time required for initial on-premises infrastructure assessment.
//    Agentless discovery doesn't collect information about network dependencies,
//    only agent-based discovery collects that information.
//
//    * Agent-based discovery collects a richer set of data than agentless discovery
//    by using the AWS Application Discovery Agent, which you install on one
//    or more hosts in your data center. The agent captures infrastructure and
//    application information, including an inventory of running processes,
//    system performance information, resource utilization, and network dependencies.
//    The information collected by agents is secured at rest and in transit
//    to the Application Discovery Service database in the cloud.
//
//    * AWS Partner Network (APN) solutions integrate with Application Discovery
//    Service, enabling you to import details of your on-premises environment
//    directly into Migration Hub without using the discovery connector or discovery
//    agent. Third-party application discovery tools can query AWS Application
//    Discovery Service, and they can write to the Application Discovery Service
//    database using the public API. In this way, you can import data into Migration
//    Hub and view it, so that you can associate applications with servers and
//    track migrations.
//
// Recommendations
//
// We recommend that you use agent-based discovery for non-VMware environments,
// and whenever you want to collect information about network dependencies.
// You can run agent-based and agentless discovery simultaneously. Use agentless
// discovery to complete the initial infrastructure assessment quickly, and
// then install agents on select hosts to collect additional information.
//
// Working With This Guide
//
// This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for Application Discovery Service. The
// topic for each action shows the API request parameters and the response.
// Alternatively, you can use one of the AWS SDKs to access an API that is tailored
// to the programming language or platform that you're using. For more information,
// see AWS SDKs (http://aws.amazon.com/tools/#SDKs).
//
//    * Remember that you must set your Migration Hub home region before you
//    call any of these APIs.
//
//    * You must make API calls for write actions (create, notify, associate,
//    disassociate, import, or put) while in your home region, or a HomeRegionNotSetException
//    error is returned.
//
//    * API calls for read actions (list, describe, stop, and delete) are permitted
//    outside of your home region.
//
//    * Although it is unlikely, the Migration Hub home region could change.
//    If you call APIs outside the home region, an InvalidInputException is
//    returned.
//
//    * You must call GetHomeRegion to obtain the latest Migration Hub home
//    region.
//
// This guide is intended for use with the AWS Application Discovery Service
// User Guide (http://docs.aws.amazon.com/application-discovery/latest/userguide/).
//
// All data is handled according to the AWS Privacy Policy (http://aws.amazon.com/privacy/).
// You can operate Application Discovery Service offline to inspect collected
// data before it is shared with the service.
//
// See https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01 for more information on this service.
//
// See applicationdiscoveryservice package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationdiscoveryservice/
//
// Using the Client
//
// To contact AWS Application Discovery Service with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Application Discovery Service client ApplicationDiscoveryService for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationdiscoveryservice/#New
package applicationdiscoveryservice
