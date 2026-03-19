# InlineResponse2007Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Coupon distribution record ID (coupon_send_issuing_log.id) | [optional] 
**DetailsId** | **int32** | User coupon detail table primary key (separate table per type). This field is 0 for task coupons | [optional] 
**CouponType** | **string** | Coupon type, enum values same as the &#x60;coupon_type&#x60; parameter | [optional] 
**Name** | **string** | Coupon display name (i18n translated) | [optional] 
**Amount** | **string** | Coupon denomination (formatted string with thousand separators). Meaning by type: point card &#x3D; balance, interest rate boost coupon &#x3D; rate percentage (e.g., &#39;5%&#39;), VIP trial card &#x3D; VIP level number, position voucher &#x3D; face value x leverage, others &#x3D; face value | [optional] 
**OriginAmount** | **string** | Original denomination (string with trailing zeros removed). Only returned for point card type, other types do not have this field | [optional] 
**Currency** | **string** | Denomination unit. Point card &#x3D; &#39;POINT&#39;, interest rate boost coupon &#x3D; &#39;APR&#39;, VIP trial card &#x3D; &#39;VIP&#39;, Alpha cash coupon &#x3D; base currency, others &#x3D; uppercase currency symbol (e.g., &#39;USDT&#39;/&#39;GT&#39;) | [optional] 
**RuleNew** | **string** | Coupon usage rule text (i18n translated). Only has value in the detail endpoint, fixed as empty string in the list endpoint | [optional] 
**Status** | **string** | Coupon status. Regular coupon: NOT_ACTIVE (pending activation), ACTIVATED (activated), TO_BE_USED (to be used), EXPIRED (expired), RECYCLED (recycled), INVALID (invalidated), USED (used), UNKNOWN (unknown), LOCKED (locked, P2P only). Task coupon: TASK_START (task not started), TASK_WAIT (task in progress), TASK_DONE (task completed, processing), TASK_EXPIRED (task not completed, expired), TASK_NOT_STARTED_EXPIRED (not started, expired), TASK_RECEIVE_SUCCESS (reward claimed successfully), TASK_RECEIVE_FAIL (reward claim failed) | [optional] 
**JumpUrl** | [**InlineResponse2006DataJumpUrl**](inline_response_200_6_data_jump_url.md) |  | [optional] 
**HelpUrl** | [**InlineResponse2006DataHelpUrl**](inline_response_200_6_data_help_url.md) |  | [optional] 
**ExpireTime** | **int32** | Expiration time (Unix timestamp). Some types replace this with actual expiration time after activation/use. Point card type returns 0. Rules by type: cash&#x3D;not_received_expired_timest; contract_bonus&#x3D;uses expired_timest after activation/use; contract_bonus_new&#x3D;always uses expired_timest; position_voucher/tradfi_position_voucher&#x3D;uses expire_time after use; robot_bonus&#x3D;not_received_expired_timest; commission_rebate&#x3D;uses use_deadline after use; crypto_loan_interest&#x3D;not_using_expired_timest; copy_trading&#x3D;not_using_expired_timest; alpha_voucher&#x3D;not_received_expired_timest; etf_voucher&#x3D;not_using_expired_timest | [optional] 
**ExpireTimeOrderBy** | **int32** | Sorting expiration time (from the original expiration time of the distribution record, unaffected by activation) | [optional] 
**ExpireSecond** | **int32** | Seconds remaining until expiration, returns 0 if expired or point card type | [optional] 
**HasUsageHistory** | **bool** | Whether there is a usage history. Fixed as true for point card type, determined by type for others | [optional] 
**HasProgress** | **bool** | Whether to display a progress bar. Only true for commission_rebate, interest_voucher, and qualifying task coupons | [optional] 
**ProgressConfig** | [**InlineResponse2006DataProgressConfig**](inline_response_200_6_data_progress_config.md) |  | [optional] 
**ActivationInfo** | [**map[string]interface{}**](.md) | Type-specific activation information. Types without specific fields return empty object {}. Fields by type: interest_voucher&#x3D;{supported_pairs (applicable trading pairs), transaction_type}; contract_bonus_new&#x3D;{received_expired_hour (valid hours after activation)}; contract_bonus&#x3D;{check_unified_account_mode, received_expired_days (valid days after activation), abtest}; commission_rebate&#x3D;{market (spot/margin/futures/alpha/etf/tradfi), market_name (market display name)}; robot_bonus&#x3D;{designated_bots (ENABLED &#x3D; designated strategies only / DISABLED &#x3D; unrestricted)}; position_voucher&#x3D;{symbols (applicable trading pairs, empty &#x3D; unrestricted), leverage, need_user_funds (0 &#x3D; no user funds required / 1 &#x3D; user funds required), user_funds_amount, position_bonus (original face value)}; tradfi_position_voucher&#x3D;{symbols, leverage, position_bonus}; etf_voucher&#x3D;{currency_markets (ETF market list, comma-separated), amount (original face value)} | [optional] 
**IsTaskCoupon** | **int32** | Whether it is a task coupon. &#x60;0&#x60; &#x3D; regular coupon; &#x60;1&#x60; &#x3D; task coupon | [optional] 
**UpgradeToast** | **bool** | Whether to prompt the user to upgrade the App (true when the app version is too old to support the coupon) | [optional] 
**FromTask** | **bool** | [Detail endpoint only] Whether this regular coupon was obtained by completing a task (a sub-coupon automatically issued after task completion). Regular coupons may be true, task coupons are always false | [optional] 
**TaskTitle** | **string** | Task title. Task coupons have value; regular coupons return empty string (but task_title is not assigned when from_task&#x3D;true) | [optional] 
**TaskDesc** | **string** | Task description. Task coupons have value; regular coupons return empty string | [optional] 
**TaskStartAt** | **int32** | Task start timestamp (Unix). Task coupon: has value when in TASK_EXPIRED status, otherwise 0; regular coupon (from_task&#x3D;true): start time of the source task | [optional] 
**TaskExpireAt** | **int32** | Task validity expiration timestamp (Unix). Task coupon: claim validity deadline (0 means unlimited); regular coupon: fixed at 0 | [optional] 
**TaskCompletedAt** | **int32** | Task completion timestamp (Unix). Task coupon: task completion time (0 means not completed); regular coupon (from_task&#x3D;true): completion time of the source task | [optional] 
**Extra** | [**[][]map[string]interface{}**](array.md) | [Detail endpoint only] Coupon detail attributes organized by display blocks. The frontend renders them as separate sections. Point card type always returns empty array []. Other types return a 2D array composed of blocks: block 1 &#x3D; coupon name/source/status, block 2 &#x3D; coupon core attributes, block 3 &#x3D; time information. Each item contains: type (display type: string/timestamp/day/hour/status/btn), key (label text, i18n translated), value (string or integer, timestamp type is Unix timestamp) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


