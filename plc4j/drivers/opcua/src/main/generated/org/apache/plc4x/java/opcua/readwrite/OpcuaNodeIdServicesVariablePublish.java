/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariablePublish {
  PublishSubscribeType_ConnectionName_Placeholder_PublisherId((int) 14418L),
  PublishSubscribeType_ConnectionName_Placeholder_Status_State((int) 14420L),
  PublishSubscribeType_ConnectionName_Placeholder_RemoveGroup_InputArguments((int) 14425L),
  PublishSubscribeType_RemoveConnection_InputArguments((int) 14433L),
  PublishSubscribeType_PublishedDataSets_AddPublishedDataItems_InputArguments((int) 14436L),
  PublishSubscribeType_PublishedDataSets_AddPublishedDataItems_OutputArguments((int) 14437L),
  PublishSubscribeType_PublishedDataSets_AddPublishedEvents_InputArguments((int) 14439L),
  PublishSubscribeType_PublishedDataSets_AddPublishedEvents_OutputArguments((int) 14440L),
  PublishSubscribeType_PublishedDataSets_RemovePublishedDataSet_InputArguments((int) 14442L),
  PublishSubscribe_GetSecurityKeys_InputArguments((int) 15216L),
  PublishSubscribe_GetSecurityKeys_OutputArguments((int) 15217L),
  PublishSubscribe_GetSecurityGroup_InputArguments((int) 15441L),
  PublishSubscribe_GetSecurityGroup_OutputArguments((int) 15442L),
  PublishSubscribe_SecurityGroups_AddSecurityGroup_InputArguments((int) 15445L),
  PublishSubscribe_SecurityGroups_AddSecurityGroup_OutputArguments((int) 15446L),
  PublishSubscribe_SecurityGroups_RemoveSecurityGroup_InputArguments((int) 15448L),
  PublishSubscribeType_ConnectionName_Placeholder_Address_NetworkInterface((int) 15533L),
  PublishSubscribeType_Status_State((int) 15845L),
  PublishSubscribeType_ConnectionName_Placeholder_AddWriterGroup_InputArguments((int) 16558L),
  PublishSubscribeType_ConnectionName_Placeholder_AddWriterGroup_OutputArguments((int) 16559L),
  PublishSubscribeType_ConnectionName_Placeholder_AddReaderGroup_InputArguments((int) 16561L),
  PublishSubscribeType_ConnectionName_Placeholder_AddReaderGroup_OutputArguments((int) 16571L),
  PublishSubscribeType_AddConnection_InputArguments((int) 16599L),
  PublishSubscribeType_AddConnection_OutputArguments((int) 16600L),
  PublishSubscribeType_PublishedDataSets_AddPublishedDataItemsTemplate_InputArguments((int) 16611L),
  PublishSubscribeType_PublishedDataSets_AddPublishedDataItemsTemplate_OutputArguments(
      (int) 16638L),
  PublishSubscribeType_PublishedDataSets_AddPublishedEventsTemplate_InputArguments((int) 16640L),
  PublishSubscribeType_PublishedDataSets_AddPublishedEventsTemplate_OutputArguments((int) 16641L),
  PublishSubscribeType_PublishedDataSets_AddDataSetFolder_InputArguments((int) 16678L),
  PublishSubscribeType_PublishedDataSets_AddDataSetFolder_OutputArguments((int) 16679L),
  PublishSubscribeType_PublishedDataSets_RemoveDataSetFolder_InputArguments((int) 16681L),
  PublishSubscribeType_ConnectionName_Placeholder_TransportProfileUri((int) 17292L),
  PublishSubscribeType_ConnectionName_Placeholder_TransportProfileUri_RestrictToList((int) 17295L),
  PublishSubscribeType_SetSecurityKeys_InputArguments((int) 17297L),
  PublishSubscribe_SetSecurityKeys_InputArguments((int) 17365L),
  PublishSubscribe_AddConnection_InputArguments((int) 17367L),
  PublishSubscribe_AddConnection_OutputArguments((int) 17368L),
  PublishSubscribe_RemoveConnection_InputArguments((int) 17370L),
  PublishSubscribe_PublishedDataSets_AddPublishedDataItems_InputArguments((int) 17373L),
  PublishSubscribe_PublishedDataSets_AddPublishedDataItems_OutputArguments((int) 17374L),
  PublishSubscribe_PublishedDataSets_AddPublishedEvents_InputArguments((int) 17376L),
  PublishSubscribe_PublishedDataSets_AddPublishedEvents_OutputArguments((int) 17377L),
  PublishSubscribe_PublishedDataSets_AddPublishedDataItemsTemplate_InputArguments((int) 17379L),
  PublishSubscribe_PublishedDataSets_AddPublishedDataItemsTemplate_OutputArguments((int) 17380L),
  PublishSubscribe_PublishedDataSets_AddPublishedEventsTemplate_InputArguments((int) 17382L),
  PublishSubscribe_PublishedDataSets_AddPublishedEventsTemplate_OutputArguments((int) 17383L),
  PublishSubscribe_PublishedDataSets_RemovePublishedDataSet_InputArguments((int) 17385L),
  PublishSubscribe_PublishedDataSets_AddDataSetFolder_InputArguments((int) 17401L),
  PublishSubscribe_PublishedDataSets_AddDataSetFolder_OutputArguments((int) 17402L),
  PublishSubscribe_PublishedDataSets_RemoveDataSetFolder_InputArguments((int) 17404L),
  PublishSubscribe_Status_State((int) 17406L),
  PublishSubscribe_Diagnostics_DiagnosticsLevel((int) 17410L),
  PublishSubscribe_Diagnostics_TotalInformation((int) 17411L),
  PublishSubscribe_Diagnostics_TotalInformation_Active((int) 17412L),
  PublishSubscribe_Diagnostics_TotalInformation_Classification((int) 17413L),
  PublishSubscribe_Diagnostics_TotalInformation_DiagnosticsLevel((int) 17414L),
  PublishSubscribe_Diagnostics_TotalInformation_TimeFirstChange((int) 17415L),
  PublishSubscribe_Diagnostics_TotalError((int) 17416L),
  PublishSubscribe_Diagnostics_TotalError_Active((int) 17417L),
  PublishSubscribe_Diagnostics_TotalError_Classification((int) 17418L),
  PublishSubscribe_Diagnostics_TotalError_DiagnosticsLevel((int) 17419L),
  PublishSubscribe_Diagnostics_TotalError_TimeFirstChange((int) 17420L),
  PublishSubscribe_Diagnostics_SubError((int) 17422L),
  PublishSubscribe_Diagnostics_Counters_StateError((int) 17424L),
  PublishSubscribe_Diagnostics_Counters_StateError_Active((int) 17425L),
  PublishSubscribe_Diagnostics_Counters_StateError_Classification((int) 17426L),
  PublishSubscribe_Diagnostics_Counters_StateError_DiagnosticsLevel((int) 17429L),
  PublishSubscribe_Diagnostics_Counters_StateError_TimeFirstChange((int) 17430L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByMethod((int) 17431L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByMethod_Active((int) 17432L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByMethod_Classification((int) 17433L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel((int) 17434L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange((int) 17435L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByParent((int) 17436L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByParent_Active((int) 17437L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByParent_Classification((int) 17438L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel((int) 17439L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange((int) 17440L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalFromError((int) 17441L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalFromError_Active((int) 17442L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalFromError_Classification((int) 17443L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel((int) 17444L),
  PublishSubscribe_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange((int) 17445L),
  PublishSubscribe_Diagnostics_Counters_StatePausedByParent((int) 17446L),
  PublishSubscribe_Diagnostics_Counters_StatePausedByParent_Active((int) 17447L),
  PublishSubscribe_Diagnostics_Counters_StatePausedByParent_Classification((int) 17448L),
  PublishSubscribe_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel((int) 17449L),
  PublishSubscribe_Diagnostics_Counters_StatePausedByParent_TimeFirstChange((int) 17450L),
  PublishSubscribe_Diagnostics_Counters_StateDisabledByMethod((int) 17451L),
  PublishSubscribe_Diagnostics_Counters_StateDisabledByMethod_Active((int) 17452L),
  PublishSubscribe_Diagnostics_Counters_StateDisabledByMethod_Classification((int) 17453L),
  PublishSubscribe_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel((int) 17454L),
  PublishSubscribe_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange((int) 17455L),
  PublishSubscribe_Diagnostics_LiveValues_ConfiguredDataSetWriters((int) 17458L),
  PublishSubscribe_Diagnostics_LiveValues_ConfiguredDataSetWriters_DiagnosticsLevel((int) 17459L),
  PublishSubscribe_Diagnostics_LiveValues_ConfiguredDataSetReaders((int) 17460L),
  PublishSubscribe_Diagnostics_LiveValues_ConfiguredDataSetReaders_DiagnosticsLevel((int) 17461L),
  PublishSubscribe_Diagnostics_LiveValues_OperationalDataSetWriters((int) 17462L),
  PublishSubscribe_Diagnostics_LiveValues_OperationalDataSetWriters_DiagnosticsLevel((int) 17463L),
  PublishSubscribe_Diagnostics_LiveValues_OperationalDataSetReaders((int) 17464L),
  PublishSubscribe_Diagnostics_LiveValues_OperationalDataSetReaders_DiagnosticsLevel((int) 17466L),
  PublishSubscribeType_ConnectionName_Placeholder_ConnectionProperties((int) 17478L),
  PublishSubscribeType_SupportedTransportProfiles((int) 17479L),
  PublishSubscribe_SupportedTransportProfiles((int) 17481L),
  PublishSubscribeType_ConnectionName_Placeholder_Address_NetworkInterface_Selections((int) 17503L),
  PublishSubscribeType_ConnectionName_Placeholder_Address_NetworkInterface_SelectionDescriptions(
      (int) 17504L),
  PublishSubscribeType_ConnectionName_Placeholder_Address_NetworkInterface_RestrictToList(
      (int) 17505L),
  PublishSubscribeType_ConnectionName_Placeholder_TransportProfileUri_Selections((int) 17706L),
  PublishSubscribeType_ConnectionName_Placeholder_TransportProfileUri_SelectionDescriptions(
      (int) 17707L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_DiagnosticsLevel((int) 18668L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalInformation((int) 18669L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalInformation_Active((int) 18670L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalInformation_Classification(
      (int) 18671L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalInformation_DiagnosticsLevel(
      (int) 18672L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalInformation_TimeFirstChange(
      (int) 18673L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalError((int) 18674L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalError_Active((int) 18675L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalError_Classification(
      (int) 18676L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalError_DiagnosticsLevel(
      (int) 18677L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_TotalError_TimeFirstChange(
      (int) 18678L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_SubError((int) 18680L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateError((int) 18682L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateError_Active(
      (int) 18683L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateError_Classification(
      (int) 18684L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateError_DiagnosticsLevel(
      (int) 18685L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateError_TimeFirstChange(
      (int) 18686L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByMethod(
      (int) 18687L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_Active(
      (int) 18688L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_Classification(
      (int) 18689L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel(
      (int) 18690L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange(
      (int) 18691L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByParent(
      (int) 18692L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByParent_Active(
      (int) 18693L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByParent_Classification(
      (int) 18694L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel(
      (int) 18695L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange(
      (int) 18696L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalFromError(
      (int) 18697L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalFromError_Active(
      (int) 18698L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalFromError_Classification(
      (int) 18699L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel(
      (int) 18700L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange(
      (int) 18701L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StatePausedByParent(
      (int) 18702L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StatePausedByParent_Active(
      (int) 18703L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StatePausedByParent_Classification(
      (int) 18704L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel(
      (int) 18705L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StatePausedByParent_TimeFirstChange(
      (int) 18706L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateDisabledByMethod(
      (int) 18707L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_Active(
      (int) 18708L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_Classification(
      (int) 18709L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel(
      (int) 18710L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange(
      (int) 18711L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_LiveValues_ResolvedAddress(
      (int) 18713L),
  PublishSubscribeType_ConnectionName_Placeholder_Diagnostics_LiveValues_ResolvedAddress_DiagnosticsLevel(
      (int) 18714L),
  PublishSubscribeType_Diagnostics_DiagnosticsLevel((int) 18716L),
  PublishSubscribeType_Diagnostics_TotalInformation((int) 18717L),
  PublishSubscribeType_Diagnostics_TotalInformation_Active((int) 18718L),
  PublishSubscribeType_Diagnostics_TotalInformation_Classification((int) 18719L),
  PublishSubscribeType_Diagnostics_TotalInformation_DiagnosticsLevel((int) 18720L),
  PublishSubscribeType_Diagnostics_TotalInformation_TimeFirstChange((int) 18721L),
  PublishSubscribeType_Diagnostics_TotalError((int) 18722L),
  PublishSubscribeType_Diagnostics_TotalError_Active((int) 18723L),
  PublishSubscribeType_Diagnostics_TotalError_Classification((int) 18724L),
  PublishSubscribeType_Diagnostics_TotalError_DiagnosticsLevel((int) 18725L),
  PublishSubscribeType_Diagnostics_TotalError_TimeFirstChange((int) 18726L),
  PublishSubscribeType_Diagnostics_SubError((int) 18728L),
  PublishSubscribeType_Diagnostics_Counters_StateError((int) 18730L),
  PublishSubscribeType_Diagnostics_Counters_StateError_Active((int) 18731L),
  PublishSubscribeType_Diagnostics_Counters_StateError_Classification((int) 18732L),
  PublishSubscribeType_Diagnostics_Counters_StateError_DiagnosticsLevel((int) 18733L),
  PublishSubscribeType_Diagnostics_Counters_StateError_TimeFirstChange((int) 18734L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByMethod((int) 18735L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByMethod_Active((int) 18736L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByMethod_Classification((int) 18737L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel((int) 18738L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange((int) 18739L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByParent((int) 18740L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByParent_Active((int) 18741L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByParent_Classification((int) 18742L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel((int) 18743L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange((int) 18744L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalFromError((int) 18745L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalFromError_Active((int) 18746L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalFromError_Classification((int) 18747L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel(
      (int) 18748L),
  PublishSubscribeType_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange((int) 18749L),
  PublishSubscribeType_Diagnostics_Counters_StatePausedByParent((int) 18750L),
  PublishSubscribeType_Diagnostics_Counters_StatePausedByParent_Active((int) 18751L),
  PublishSubscribeType_Diagnostics_Counters_StatePausedByParent_Classification((int) 18752L),
  PublishSubscribeType_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel((int) 18753L),
  PublishSubscribeType_Diagnostics_Counters_StatePausedByParent_TimeFirstChange((int) 18754L),
  PublishSubscribeType_Diagnostics_Counters_StateDisabledByMethod((int) 18755L),
  PublishSubscribeType_Diagnostics_Counters_StateDisabledByMethod_Active((int) 18756L),
  PublishSubscribeType_Diagnostics_Counters_StateDisabledByMethod_Classification((int) 18757L),
  PublishSubscribeType_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel((int) 18758L),
  PublishSubscribeType_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange((int) 18759L),
  PublishSubscribeType_Diagnostics_LiveValues_ConfiguredDataSetWriters((int) 18761L),
  PublishSubscribeType_Diagnostics_LiveValues_ConfiguredDataSetWriters_DiagnosticsLevel(
      (int) 18762L),
  PublishSubscribeType_Diagnostics_LiveValues_ConfiguredDataSetReaders((int) 18763L),
  PublishSubscribeType_Diagnostics_LiveValues_ConfiguredDataSetReaders_DiagnosticsLevel(
      (int) 18764L),
  PublishSubscribeType_Diagnostics_LiveValues_OperationalDataSetWriters((int) 18765L),
  PublishSubscribeType_Diagnostics_LiveValues_OperationalDataSetWriters_DiagnosticsLevel(
      (int) 18766L),
  PublishSubscribeType_Diagnostics_LiveValues_OperationalDataSetReaders((int) 18767L),
  PublishSubscribeType_Diagnostics_LiveValues_OperationalDataSetReaders_DiagnosticsLevel(
      (int) 18768L),
  PublishSubscribeType_SubscribedDataSets_AddDataSetFolder_InputArguments((int) 23638L),
  PublishSubscribeType_SubscribedDataSets_AddDataSetFolder_OutputArguments((int) 23639L),
  PublishSubscribeType_SubscribedDataSets_RemoveDataSetFolder_InputArguments((int) 23641L),
  PublishSubscribeType_PubSubCapablities_MaxPubSubConnections((int) 23643L),
  PublishSubscribeType_PubSubCapablities_MaxWriterGroups((int) 23644L),
  PublishSubscribeType_PubSubCapablities_MaxReaderGroups((int) 23645L),
  PublishSubscribeType_PubSubCapablities_MaxDataSetWriters((int) 23646L),
  PublishSubscribeType_PubSubCapablities_MaxDataSetReaders((int) 23647L),
  PublishSubscribeType_PubSubCapablities_MaxFieldsPerDataSet((int) 23648L),
  PublishSubscribe_SubscribedDataSets_AddDataSetFolder_InputArguments((int) 23674L),
  PublishSubscribe_SubscribedDataSets_AddDataSetFolder_OutputArguments((int) 23675L),
  PublishSubscribe_SubscribedDataSets_RemoveDataSetFolder_InputArguments((int) 23677L),
  PublishSubscribe_PubSubCapablities_MaxPubSubConnections((int) 23679L),
  PublishSubscribe_PubSubCapablities_MaxWriterGroups((int) 23680L),
  PublishSubscribe_PubSubCapablities_MaxReaderGroups((int) 23681L),
  PublishSubscribe_PubSubCapablities_MaxDataSetWriters((int) 23682L),
  PublishSubscribe_PubSubCapablities_MaxDataSetReaders((int) 23683L),
  PublishSubscribe_PubSubCapablities_MaxFieldsPerDataSet((int) 23684L),
  PublishSubscribeType_SubscribedDataSets_AddSubscribedDataSet_InputArguments((int) 24005L),
  PublishSubscribeType_SubscribedDataSets_AddSubscribedDataSet_OutputArguments((int) 24006L),
  PublishSubscribeType_SubscribedDataSets_RemoveSubscribedDataSet_InputArguments((int) 24008L),
  PublishSubscribeType_DataSetClasses_DataSetName_Placeholder((int) 24009L),
  PublishSubscribe_SubscribedDataSets_AddSubscribedDataSet_InputArguments((int) 24011L),
  PublishSubscribe_SubscribedDataSets_AddSubscribedDataSet_OutputArguments((int) 24012L),
  PublishSubscribe_SubscribedDataSets_RemoveSubscribedDataSet_InputArguments((int) 24014L),
  PublishSubscribeType_PubSubConfiguration_Size((int) 25404L),
  PublishSubscribeType_PubSubConfiguration_Writable((int) 25405L),
  PublishSubscribeType_PubSubConfiguration_UserWritable((int) 25406L),
  PublishSubscribeType_PubSubConfiguration_OpenCount((int) 25407L),
  PublishSubscribeType_PubSubConfiguration_MimeType((int) 25408L),
  PublishSubscribeType_PubSubConfiguration_MaxByteStringLength((int) 25409L),
  PublishSubscribeType_PubSubConfiguration_LastModifiedTime((int) 25410L),
  PublishSubscribeType_PubSubConfiguration_Open_InputArguments((int) 25412L),
  PublishSubscribeType_PubSubConfiguration_Open_OutputArguments((int) 25413L),
  PublishSubscribeType_PubSubConfiguration_Close_InputArguments((int) 25415L),
  PublishSubscribeType_PubSubConfiguration_Read_InputArguments((int) 25417L),
  PublishSubscribeType_PubSubConfiguration_Read_OutputArguments((int) 25418L),
  PublishSubscribeType_PubSubConfiguration_Write_InputArguments((int) 25420L),
  PublishSubscribeType_PubSubConfiguration_GetPosition_InputArguments((int) 25422L),
  PublishSubscribeType_PubSubConfiguration_GetPosition_OutputArguments((int) 25423L),
  PublishSubscribeType_PubSubConfiguration_SetPosition_InputArguments((int) 25425L),
  PublishSubscribeType_PubSubConfiguration_ReserveIds_InputArguments((int) 25427L),
  PublishSubscribeType_PubSubConfiguration_ReserveIds_OutputArguments((int) 25428L),
  PublishSubscribeType_PubSubConfiguration_CloseAndUpdate_InputArguments((int) 25430L),
  PublishSubscribeType_PubSubConfiguration_CloseAndUpdate_OutputArguments((int) 25431L),
  PublishSubscribeType_DefaultDatagramPublisherId((int) 25432L),
  PublishSubscribeType_ConfigurationVersion((int) 25433L),
  PublishSubscribe_SecurityGroups_AddSecurityGroupFolder_InputArguments((int) 25435L),
  PublishSubscribe_SecurityGroups_AddSecurityGroupFolder_OutputArguments((int) 25436L),
  PublishSubscribe_SecurityGroups_RemoveSecurityGroupFolder_InputArguments((int) 25438L),
  PublishSubscribe_SecurityGroups_SupportedSecurityPolicyUris((int) 25439L),
  PublishSubscribe_KeyPushTargets_AddPushTarget_InputArguments((int) 25442L),
  PublishSubscribe_KeyPushTargets_AddPushTarget_OutputArguments((int) 25443L),
  PublishSubscribe_KeyPushTargets_RemovePushTarget_InputArguments((int) 25445L),
  PublishSubscribe_KeyPushTargets_AddPushTargetFolder_InputArguments((int) 25447L),
  PublishSubscribe_KeyPushTargets_AddPushTargetFolder_OutputArguments((int) 25448L),
  PublishSubscribe_KeyPushTargets_RemovePushTargetFolder_InputArguments((int) 25450L),
  PublishSubscribe_PubSubConfiguration_Size((int) 25452L),
  PublishSubscribe_PubSubConfiguration_Writable((int) 25453L),
  PublishSubscribe_PubSubConfiguration_UserWritable((int) 25454L),
  PublishSubscribe_PubSubConfiguration_OpenCount((int) 25455L),
  PublishSubscribe_PubSubConfiguration_MimeType((int) 25456L),
  PublishSubscribe_PubSubConfiguration_MaxByteStringLength((int) 25457L),
  PublishSubscribe_PubSubConfiguration_LastModifiedTime((int) 25458L),
  PublishSubscribe_PubSubConfiguration_Open_InputArguments((int) 25460L),
  PublishSubscribe_PubSubConfiguration_Open_OutputArguments((int) 25461L),
  PublishSubscribe_PubSubConfiguration_Close_InputArguments((int) 25463L),
  PublishSubscribe_PubSubConfiguration_Read_InputArguments((int) 25465L),
  PublishSubscribe_PubSubConfiguration_Read_OutputArguments((int) 25466L),
  PublishSubscribe_PubSubConfiguration_Write_InputArguments((int) 25468L),
  PublishSubscribe_PubSubConfiguration_GetPosition_InputArguments((int) 25470L),
  PublishSubscribe_PubSubConfiguration_GetPosition_OutputArguments((int) 25471L),
  PublishSubscribe_PubSubConfiguration_SetPosition_InputArguments((int) 25473L),
  PublishSubscribe_PubSubConfiguration_ReserveIds_InputArguments((int) 25475L),
  PublishSubscribe_PubSubConfiguration_ReserveIds_OutputArguments((int) 25476L),
  PublishSubscribe_PubSubConfiguration_CloseAndUpdate_InputArguments((int) 25478L),
  PublishSubscribe_PubSubConfiguration_CloseAndUpdate_OutputArguments((int) 25479L),
  PublishSubscribe_DefaultDatagramPublisherId((int) 25480L),
  PublishSubscribe_ConfigurationVersion((int) 25481L),
  PublishSubscribeType_PubSubCapablities_MaxDataSetWritersPerGroup((int) 32391L),
  PublishSubscribeType_PubSubCapablities_MaxNetworkMessageSizeDatagram((int) 32392L),
  PublishSubscribeType_PubSubCapablities_MaxNetworkMessageSizeBroker((int) 32393L),
  PublishSubscribeType_PubSubCapablities_SupportSecurityKeyPull((int) 32394L),
  PublishSubscribeType_PubSubCapablities_SupportSecurityKeyPush((int) 32395L),
  PublishSubscribeType_DefaultSecurityKeyServices((int) 32396L),
  PublishSubscribeType_ConfigurationProperties((int) 32397L),
  PublishSubscribe_PubSubCapablities_MaxDataSetWritersPerGroup((int) 32398L),
  PublishSubscribe_PubSubCapablities_MaxNetworkMessageSizeDatagram((int) 32399L),
  PublishSubscribe_PubSubCapablities_MaxNetworkMessageSizeBroker((int) 32400L),
  PublishSubscribe_PubSubCapablities_SupportSecurityKeyPull((int) 32401L),
  PublishSubscribe_PubSubCapablities_SupportSecurityKeyPush((int) 32402L),
  PublishSubscribe_DefaultSecurityKeyServices((int) 32403L),
  PublishSubscribe_ConfigurationProperties((int) 32404L),
  PublishSubscribeType_PubSubCapablities_MaxSecurityGroups((int) 32834L),
  PublishSubscribeType_PubSubCapablities_MaxPushTargets((int) 32835L),
  PublishSubscribeType_PubSubCapablities_MaxPublishedDataSets((int) 32836L),
  PublishSubscribeType_PubSubCapablities_MaxStandaloneSubscribedDataSets((int) 32837L),
  PublishSubscribeType_PubSubCapablities_SupportSecurityKeyServer((int) 32838L),
  PublishSubscribe_PubSubCapablities_MaxSecurityGroups((int) 32839L),
  PublishSubscribe_PubSubCapablities_MaxPushTargets((int) 32840L),
  PublishSubscribe_PubSubCapablities_MaxPublishedDataSets((int) 32841L),
  PublishSubscribe_PubSubCapablities_MaxStandaloneSubscribedDataSets((int) 32842L),
  PublishSubscribe_PubSubCapablities_SupportSecurityKeyServer((int) 32843L);
  private static final Map<Integer, OpcuaNodeIdServicesVariablePublish> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariablePublish value : OpcuaNodeIdServicesVariablePublish.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariablePublish(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariablePublish enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
