# InlineResponse2007DataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Coupon distribution record ID (coupon_send_issuing_log.id), used as &#x60;last_id&#x60; for cursor-based pagination | [optional] 
**DetailsId** | **int32** | User coupon detail table primary key (separate table per type). This field is 0 for task coupons, has value for regular coupons | [optional] 
**CouponType** | **string** | Coupon type, enum values same as the &#x60;type&#x60; parameter | [optional] 
**Name** | **string** | Coupon display name (i18n translated) | [optional] 
**Amount** | **string** | Coupon denomination (formatted string with thousand separators). Meaning by type: point card &#x3D; balance, interest rate boost coupon &#x3D; interest rate percentage (e.g., &#39;5%&#39;), VIP trial card &#x3D; VIP level number, position voucher &#x3D; face value × leverage, others &#x3D; face value | [optional] 
**OriginAmount** | **string** | Original denomination (string with trailing zeros removed). Only returned for point card type, other types do not have this field | [optional] 
**Currency** | **string** | Denomination unit. Point card &#x3D; &#39;POINT&#39;, interest rate boost coupon &#x3D; &#39;APR&#39;, VIP trial card &#x3D; &#39;VIP&#39;, Alpha cash coupon &#x3D; base currency, others &#x3D; uppercase currency symbol (e.g., &#39;USDT&#39;/&#39;GT&#39;) | [optional] 
**RuleNew** | **string** | Coupon usage rule text. List endpoint always returns empty string \&quot;\&quot;, only detail endpoint returns actual value | [optional] 
**Status** | **string** | Coupon status. Regular coupon: NOT_ACTIVE (pending activation), ACTIVATED (activated), TO_BE_USED (to be used), EXPIRED (expired), RECYCLED (recycled), INVALID (invalidated), USED (used), UNKNOWN (unknown), LOCKED (locked, P2P only). Task coupon: TASK_START (task not started), TASK_WAIT (task in progress), TASK_DONE (task completed, processing), TASK_EXPIRED (task not completed, expired), TASK_NOT_STARTED_EXPIRED (not started, expired), TASK_RECEIVE_SUCCESS (reward claimed successfully), TASK_RECEIVE_FAIL (reward claim failed) | [optional] 
**JumpUrl** | [**InlineResponse2007DataJumpUrl**](inline_response_200_7_data_jump_url.md) |  | [optional] 
**HelpUrl** | [**InlineResponse2007DataHelpUrl**](inline_response_200_7_data_help_url.md) |  | [optional] 
**ExpireTime** | **int32** | Expiration time (Unix timestamp). Some types replace this with actual expiration time after activation or use (e.g., contract_bonus uses expired_timest after activation). Point card type returns 0 | [optional] 
**ExpireTimeOrderBy** | **int32** | Sorting expiration time (from the original expiration time of the distribution record, unaffected by activation). Used as the &#x60;expire_time&#x60; parameter for the next request in cursor-based pagination | [optional] 
**ExpireSecond** | **int32** | Seconds remaining until expiration. Returns 0 for expired or Point Card types | [optional] 
**HasUsageHistory** | **bool** | Whether there is a usage history. Fixed as true for point card type, determined by type for others | [optional] 
**HasProgress** | **bool** | Whether to display a progress bar. Only true for commission_rebate, interest_voucher, and qualifying task coupons | [optional] 
**ProgressConfig** | [**InlineResponse2007DataProgressConfig**](inline_response_200_7_data_progress_config.md) |  | [optional] 
**ActivationInfo** | [**map[string]interface{}**](.md) | Type-specific activation information. Types without specific fields return empty object {}. Fields by type: interest_voucher&#x3D;{supported_pairs,transaction_type}; contract_bonus_new&#x3D;{received_expired_hour}; contract_bonus&#x3D;{check_unified_account_mode,received_expired_days,abtest}; commission_rebate&#x3D;{market,market_name}; robot_bonus&#x3D;{designated_bots}; position_voucher&#x3D;{symbols,leverage,need_user_funds,user_funds_amount,position_bonus}; tradfi_position_voucher&#x3D;{symbols,leverage,position_bonus}; etf_voucher&#x3D;{currency_markets,amount} | [optional] 
**IsTaskCoupon** | **int32** | Whether it is a task coupon. &#x60;0&#x60; &#x3D; regular coupon; &#x60;1&#x60; &#x3D; task coupon | [optional] 
**UpgradeToast** | **bool** | Whether to prompt the user to upgrade the App (true when app version is too old to support the coupon). Triggered types: copy_trading/alpha_voucher (Android&lt;7320000/iOS&lt;202507320000), commission_rebate subtype tradfi (Android&lt;8040000/iOS&lt;202608040000), etf_voucher (Android&lt;8090000/iOS&lt;202608090000), tradfi_position_voucher (Android&lt;8100000/iOS&lt;202508240000) | [optional] 
**TaskTitle** | **string** | Task title (only task coupons have value, regular coupons return empty string \&quot;\&quot;) | [optional] 
**TaskDesc** | **string** | Task description (only task coupons have value, regular coupons return empty string \&quot;\&quot;) | [optional] 
**TaskStartAt** | **int32** | Task start timestamp (Unix). Only has value for task coupons in TASK_EXPIRED status, otherwise 0 | [optional] 
**TaskExpireAt** | **int32** | Task expiration timestamp (Unix). Currently fixed at 0, reserved field | [optional] 
**TaskCompletedAt** | **int32** | Task completion timestamp (Unix). Currently fixed at 0, reserved field | [optional] 
**Extra** | **[]map[string]interface{}** | Extension fields. List endpoint always returns empty array [], only the detail endpoint has values | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


