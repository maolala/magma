/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

/*! \file mme_app_itti_messaging.h
  \brief
  \author Sebastien ROUX, Lionel Gauthier
  \company Eurecom
  \email: lionel.gauthier@eurecom.fr
*/

#ifndef FILE_MME_APP_ITTI_MESSAGING_SEEN
#define FILE_MME_APP_ITTI_MESSAGING_SEEN
#include <inttypes.h>
#include <string.h>

#include "log.h"
#include "timer.h"
#include "mme_config.h"
#include "3gpp_36.401.h"
#include "common_types.h"
#include "intertask_interface.h"
#include "intertask_interface_types.h"
#include "itti_types.h"
#include "mme_app_desc.h"
#include "mme_app_ue_context.h"
#include "s1ap_messages_types.h"

void mme_app_itti_ue_context_release(
  struct ue_mm_context_s *ue_context_p,
  enum s1cause cause);
int mme_app_notify_s1ap_ue_context_released(const mme_ue_s1ap_id_t ue_idP);
int mme_app_send_s11_release_access_bearers_req(
  struct ue_mm_context_s *const ue_mm_context,
  const pdn_cid_t pdn_index);
int mme_app_send_s11_create_session_req(mme_app_desc_t *mme_app_desc_p,
  struct ue_mm_context_s *const ue_mm_context,
  const pdn_cid_t pdn_cid);

static inline void mme_app_itti_ue_context_mod_for_csfb(
  struct ue_mm_context_s *ue_context_p)
{
  MessageDef *message_p;

  message_p =
    itti_alloc_new_message(TASK_MME_APP, S1AP_UE_CONTEXT_MODIFICATION_REQUEST);
  memset(
    (void *) &message_p->ittiMsg.s1ap_ue_context_mod_request,
    0,
    sizeof(itti_s1ap_ue_context_mod_req_t));
  S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).mme_ue_s1ap_id =
    ue_context_p->mme_ue_s1ap_id;
  S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).enb_ue_s1ap_id =
    ue_context_p->enb_ue_s1ap_id;
  if (
    (ue_context_p->sgs_context != NULL) &&
    ((ue_context_p->sgs_context->csfb_service_type == CSFB_SERVICE_MO_CALL) ||
     (ue_context_p->sgs_context->csfb_service_type == CSFB_SERVICE_MT_CALL))) {
    S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).presencemask =
      S1AP_UE_CONTEXT_MOD_LAI_PRESENT;
    mme_config_read_lock(&mme_config);
    S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).lai = mme_config.lai;
    mme_config_unlock(&mme_config);
    S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).presencemask |=
      S1AP_UE_CONTEXT_MOD_CSFB_INDICATOR_PRESENT;
    if (ue_context_p->sgs_context->is_emergency_call == true) {
      S1AP_UE_CONTEXT_MODIFICATION_REQUEST (message_p).cs_fallback_indicator =
      CSFB_HIGH_PRIORITY;
      ue_context_p->sgs_context->is_emergency_call = false;
    } else {
      S1AP_UE_CONTEXT_MODIFICATION_REQUEST (message_p).cs_fallback_indicator =
      CSFB_REQUIRED;
    }
  }
  OAILOG_INFO(
    LOG_MME_APP,
    "Sent S1AP_UE_CONTEXT_MODIFICATION_REQUEST mme_ue_s1ap_id %06" PRIX32 " \n",
    S1AP_UE_CONTEXT_MODIFICATION_REQUEST(message_p).mme_ue_s1ap_id);
  itti_send_msg_to_task(TASK_S1AP, INSTANCE_DEFAULT, message_p);

  /* Start timer to wait for UE Context Modification from eNB
   * If timer expires treat this as failure of ongoing procedure and abort corresponding NAS procedure
   * such as SERVICE REQUEST and Send Service Reject to eNB
   */
  if (
    timer_setup(
      ue_context_p->ue_context_modification_timer.sec,
      0,
      TASK_MME_APP,
      INSTANCE_DEFAULT,
      TIMER_ONE_SHOT,
      (void *) &(ue_context_p->mme_ue_s1ap_id),
      sizeof(mme_ue_s1ap_id_t),
      &(ue_context_p->ue_context_modification_timer.id)) < 0) {
    OAILOG_ERROR(
      LOG_MME_APP,
      "Failed to start UE context modification timer for UE id  %d \n",
      ue_context_p->mme_ue_s1ap_id);
    ue_context_p->ue_context_modification_timer.id = MME_APP_TIMER_INACTIVE_ID;
  } else {
    OAILOG_DEBUG(
      LOG_MME_APP,
      "MME APP :Sent UE context modification request and Started guard timer "
      "for UE id %d\n",
      ue_context_p->mme_ue_s1ap_id);
  }
  OAILOG_FUNC_OUT(LOG_MME_APP);
}

#endif /* FILE_MME_APP_ITTI_MESSAGING_SEEN */
