# frozen_string_literal: true

# Typed models for the Nofrixion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Account entity data model.
#
# @!attribute [rw] accountBalances
#   @return [Array, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountIdentifications
#   @return [Array, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountNames
#   @return [Array, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] accountType
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] consolidatedAccountInformation
#   @return [Hash, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] details
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] fromDate
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] nickname
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] roleIDs
#   @return [Array, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierPhysicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] toDate
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] usageType
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
Account = Struct.new(
  :accountBalances,
  :accountID,
  :accountIdentifications,
  :accountName,
  :accountNames,
  :accountSupplierName,
  :accountType,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :consolidatedAccountInformation,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :description,
  :details,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :format,
  :fromDate,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :nickname,
  :physicalAccountID,
  :roleIDs,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierPhysicalAccountID,
  :supplierSepaInstantStatus,
  :toDate,
  :type,
  :usageType,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Account#load.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
AccountLoadMatch = Struct.new(
  :account_id,
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Account#list.
#
# @!attribute [rw] accountBalances
#   @return [Array, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountIdentifications
#   @return [Array, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountNames
#   @return [Array, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] accountType
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] consolidatedAccountInformation
#   @return [Hash, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] details
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] fromDate
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] nickname
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] roleIDs
#   @return [Array, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierPhysicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] toDate
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] usageType
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
AccountListMatch = Struct.new(
  :accountBalances,
  :accountID,
  :accountIdentifications,
  :accountName,
  :accountNames,
  :accountSupplierName,
  :accountType,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :consolidatedAccountInformation,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :description,
  :details,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :format,
  :fromDate,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :nickname,
  :physicalAccountID,
  :roleIDs,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierPhysicalAccountID,
  :supplierSepaInstantStatus,
  :toDate,
  :type,
  :usageType,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Account#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String]
#
# @!attribute [rw] accountBalances
#   @return [Array, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountIdentifications
#   @return [Array, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountNames
#   @return [Array, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] accountType
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] consolidatedAccountInformation
#   @return [Hash, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] details
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] fromDate
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] nickname
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] roleIDs
#   @return [Array, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierPhysicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] toDate
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] usageType
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
AccountCreateData = Struct.new(
  :account_id,
  :currency,
  :accountBalances,
  :accountID,
  :accountIdentifications,
  :accountName,
  :accountNames,
  :accountSupplierName,
  :accountType,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :consolidatedAccountInformation,
  :createdBy,
  :createdByDisplayName,
  :defaultPaymentRail,
  :description,
  :details,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :format,
  :fromDate,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :nickname,
  :physicalAccountID,
  :roleIDs,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierPhysicalAccountID,
  :supplierSepaInstantStatus,
  :toDate,
  :type,
  :usageType,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Account#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountBalances
#   @return [Array, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountIdentifications
#   @return [Array, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountNames
#   @return [Array, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] accountType
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] consolidatedAccountInformation
#   @return [Hash, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] details
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] fromDate
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] nickname
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] roleIDs
#   @return [Array, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierPhysicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] toDate
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] usageType
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
AccountUpdateData = Struct.new(
  :id,
  :accountBalances,
  :accountID,
  :accountIdentifications,
  :accountName,
  :accountNames,
  :accountSupplierName,
  :accountType,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :consolidatedAccountInformation,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :description,
  :details,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :format,
  :fromDate,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :nickname,
  :physicalAccountID,
  :roleIDs,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierPhysicalAccountID,
  :supplierSepaInstantStatus,
  :toDate,
  :type,
  :usageType,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Account#remove.
#
# @!attribute [rw] id
#   @return [String]
AccountRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Batch entity data model.
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
Batch = Struct.new(
  :approveUrl,
  :id,
  :payouts,
  keyword_init: true
)

# Request payload for Batch#load.
#
# @!attribute [rw] id
#   @return [String]
BatchLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Batch#create.
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
BatchCreateData = Struct.new(
  :approveUrl,
  :id,
  :payouts,
  keyword_init: true
)

# Beneficiary entity data model.
#
# @!attribute [rw] approvalCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] beneficiaries
#   @return [Array, nil]
#
# @!attribute [rw] beneficiaryEvents
#   @return [Array, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] failedBeneficiaries
#   @return [Hash, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIDs
#   @return [Array, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
Beneficiary = Struct.new(
  :approvalCallbackUrl,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :beneficiaries,
  :beneficiaryEvents,
  :canAuthorise,
  :canUpdate,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :destination,
  :failedBeneficiaries,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :sourceAccountIDs,
  :sourceAccounts,
  :theirReference,
  keyword_init: true
)

# Request payload for Beneficiary#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
BeneficiaryLoadMatch = Struct.new(
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Beneficiary#list.
#
# @!attribute [rw] approvalCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] beneficiaries
#   @return [Array, nil]
#
# @!attribute [rw] beneficiaryEvents
#   @return [Array, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] failedBeneficiaries
#   @return [Hash, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIDs
#   @return [Array, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
BeneficiaryListMatch = Struct.new(
  :approvalCallbackUrl,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :beneficiaries,
  :beneficiaryEvents,
  :canAuthorise,
  :canUpdate,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :destination,
  :failedBeneficiaries,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :sourceAccountIDs,
  :sourceAccounts,
  :theirReference,
  keyword_init: true
)

# Request payload for Beneficiary#create.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] approvalCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] beneficiaries
#   @return [Array, nil]
#
# @!attribute [rw] beneficiaryEvents
#   @return [Array, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] failedBeneficiaries
#   @return [Hash, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIDs
#   @return [Array, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
BeneficiaryCreateData = Struct.new(
  :id,
  :approvalCallbackUrl,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :beneficiaries,
  :beneficiaryEvents,
  :canAuthorise,
  :canUpdate,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :destination,
  :failedBeneficiaries,
  :hasCurrentUserAuthorised,
  :inserted,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :sourceAccountIDs,
  :sourceAccounts,
  :theirReference,
  keyword_init: true
)

# Request payload for Beneficiary#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] approvalCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] beneficiaries
#   @return [Array, nil]
#
# @!attribute [rw] beneficiaryEvents
#   @return [Array, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] failedBeneficiaries
#   @return [Hash, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIDs
#   @return [Array, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
BeneficiaryUpdateData = Struct.new(
  :id,
  :approvalCallbackUrl,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :beneficiaries,
  :beneficiaryEvents,
  :canAuthorise,
  :canUpdate,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :destination,
  :failedBeneficiaries,
  :hasCurrentUserAuthorised,
  :inserted,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :sourceAccountIDs,
  :sourceAccounts,
  :theirReference,
  keyword_init: true
)

# Request payload for Beneficiary#remove.
#
# @!attribute [rw] id
#   @return [String]
BeneficiaryRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# BeneficiaryGroup entity data model.
#
# @!attribute [rw] groupMembers
#   @return [Array, nil]
#
# @!attribute [rw] groupName
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String]
BeneficiaryGroup = Struct.new(
  :groupMembers,
  :groupName,
  :id,
  :inserted,
  :lastUpdated,
  :merchantID,
  keyword_init: true
)

# Request payload for BeneficiaryGroup#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
BeneficiaryGroupListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Card entity data model.
#
# @!attribute [rw] authorizedAmount
#   @return [String, nil]
#
# @!attribute [rw] currencyCode
#   @return [String, nil]
#
# @!attribute [rw] isPayerAuthenticationRequired
#   @return [Boolean, nil]
#
# @!attribute [rw] isSoftDecline
#   @return [Boolean, nil]
#
# @!attribute [rw] payerAuthenticationAccessToken
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationMerchantData
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationUrl
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationWindowHeight
#   @return [Integer, nil]
#
# @!attribute [rw] payerAuthenticationWindowWidth
#   @return [Integer, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] requestID
#   @return [String, nil]
#
# @!attribute [rw] responseCode
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] threeDSRedirectUrl
#   @return [String, nil]
#
# @!attribute [rw] transactionID
#   @return [String, nil]
Card = Struct.new(
  :authorizedAmount,
  :currencyCode,
  :isPayerAuthenticationRequired,
  :isSoftDecline,
  :payerAuthenticationAccessToken,
  :payerAuthenticationMerchantData,
  :payerAuthenticationUrl,
  :payerAuthenticationWindowHeight,
  :payerAuthenticationWindowWidth,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :requestID,
  :responseCode,
  :responseType,
  :status,
  :threeDSRedirectUrl,
  :transactionID,
  keyword_init: true
)

# Request payload for Card#create.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
#
# @!attribute [rw] authorizedAmount
#   @return [String, nil]
#
# @!attribute [rw] currencyCode
#   @return [String, nil]
#
# @!attribute [rw] isPayerAuthenticationRequired
#   @return [Boolean, nil]
#
# @!attribute [rw] isSoftDecline
#   @return [Boolean, nil]
#
# @!attribute [rw] payerAuthenticationAccessToken
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationMerchantData
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationUrl
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationWindowHeight
#   @return [Integer, nil]
#
# @!attribute [rw] payerAuthenticationWindowWidth
#   @return [Integer, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] requestID
#   @return [String, nil]
#
# @!attribute [rw] responseCode
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] threeDSRedirectUrl
#   @return [String, nil]
#
# @!attribute [rw] transactionID
#   @return [String, nil]
CardCreateData = Struct.new(
  :paymentrequest_id,
  :authorizedAmount,
  :currencyCode,
  :isPayerAuthenticationRequired,
  :isSoftDecline,
  :payerAuthenticationAccessToken,
  :payerAuthenticationMerchantData,
  :payerAuthenticationUrl,
  :payerAuthenticationWindowHeight,
  :payerAuthenticationWindowWidth,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :requestID,
  :responseCode,
  :responseType,
  :status,
  :threeDSRedirectUrl,
  :transactionID,
  keyword_init: true
)

# CardCustomerToken entity data model.
#
# @!attribute [rw] cardType
#   @return [String, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] expiryMonth
#   @return [String, nil]
#
# @!attribute [rw] expiryYear
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lastFourDigits
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] maskedCardNumber
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
CardCustomerToken = Struct.new(
  :cardType,
  :customerEmailAddress,
  :expiryMonth,
  :expiryYear,
  :id,
  :inserted,
  :lastFourDigits,
  :lastUpdated,
  :maskedCardNumber,
  :merchantID,
  :paymentRequestID,
  keyword_init: true
)

# Request payload for CardCustomerToken#load.
#
# @!attribute [rw] customer_email_address
#   @return [String]
CardCustomerTokenLoadMatch = Struct.new(
  :customer_email_address,
  keyword_init: true
)

# Request payload for CardCustomerToken#list.
#
# @!attribute [rw] customer_email_address
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String]
CardCustomerTokenListMatch = Struct.new(
  :customer_email_address,
  :merchant_id,
  keyword_init: true
)

# Request payload for CardCustomerToken#remove.
#
# @!attribute [rw] id
#   @return [String]
CardCustomerTokenRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# CardPayment entity data model.
#
# @!attribute [rw] authorizedAmount
#   @return [String, nil]
#
# @!attribute [rw] currencyCode
#   @return [String, nil]
#
# @!attribute [rw] isPayerAuthenticationRequired
#   @return [Boolean, nil]
#
# @!attribute [rw] isSoftDecline
#   @return [Boolean, nil]
#
# @!attribute [rw] payerAuthenticationAccessToken
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationMerchantData
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationUrl
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationWindowHeight
#   @return [Integer, nil]
#
# @!attribute [rw] payerAuthenticationWindowWidth
#   @return [Integer, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] requestID
#   @return [String, nil]
#
# @!attribute [rw] responseCode
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] threeDSRedirectUrl
#   @return [String, nil]
#
# @!attribute [rw] transactionID
#   @return [String, nil]
CardPayment = Struct.new(
  :authorizedAmount,
  :currencyCode,
  :isPayerAuthenticationRequired,
  :isSoftDecline,
  :payerAuthenticationAccessToken,
  :payerAuthenticationMerchantData,
  :payerAuthenticationUrl,
  :payerAuthenticationWindowHeight,
  :payerAuthenticationWindowWidth,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :requestID,
  :responseCode,
  :responseType,
  :status,
  :threeDSRedirectUrl,
  :transactionID,
  keyword_init: true
)

# Request payload for CardPayment#create.
#
# @!attribute [rw] partial_refund_amount
#   @return [Float, nil]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
#
# @!attribute [rw] authorizedAmount
#   @return [String, nil]
#
# @!attribute [rw] currencyCode
#   @return [String, nil]
#
# @!attribute [rw] isPayerAuthenticationRequired
#   @return [Boolean, nil]
#
# @!attribute [rw] isSoftDecline
#   @return [Boolean, nil]
#
# @!attribute [rw] payerAuthenticationAccessToken
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationMerchantData
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationUrl
#   @return [String, nil]
#
# @!attribute [rw] payerAuthenticationWindowHeight
#   @return [Integer, nil]
#
# @!attribute [rw] payerAuthenticationWindowWidth
#   @return [Integer, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] requestID
#   @return [String, nil]
#
# @!attribute [rw] responseCode
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] threeDSRedirectUrl
#   @return [String, nil]
#
# @!attribute [rw] transactionID
#   @return [String, nil]
CardPaymentCreateData = Struct.new(
  :partial_refund_amount,
  :paymentrequest_id,
  :authorizedAmount,
  :currencyCode,
  :isPayerAuthenticationRequired,
  :isSoftDecline,
  :payerAuthenticationAccessToken,
  :payerAuthenticationMerchantData,
  :payerAuthenticationUrl,
  :payerAuthenticationWindowHeight,
  :payerAuthenticationWindowWidth,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :requestID,
  :responseCode,
  :responseType,
  :status,
  :threeDSRedirectUrl,
  :transactionID,
  keyword_init: true
)

# CardPublicKey entity data model.
#
# @!attribute [rw] jwt
#   @return [String, nil]
CardPublicKey = Struct.new(
  :jwt,
  keyword_init: true
)

# Request payload for CardPublicKey#load.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
CardPublicKeyLoadMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Consent entity data model.
#
# @!attribute [rw] authorisationUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] institutionID
#   @return [String, nil]
#
# @!attribute [rw] isConnectedAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] provider
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
Consent = Struct.new(
  :authorisationUrl,
  :callbackUrl,
  :consentID,
  :emailAddress,
  :expiryDate,
  :failureCallbackUrl,
  :id,
  :inserted,
  :institutionID,
  :isConnectedAccounts,
  :isEnabled,
  :merchantID,
  :provider,
  :successWebHookUrl,
  keyword_init: true
)

# Request payload for Consent#load.
#
# @!attribute [rw] id
#   @return [String]
ConsentLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Consent#list.
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String]
ConsentListMatch = Struct.new(
  :email,
  :merchant_id,
  keyword_init: true
)

# Request payload for Consent#create.
#
# @!attribute [rw] authorisationUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] institutionID
#   @return [String, nil]
#
# @!attribute [rw] isConnectedAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] provider
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
ConsentCreateData = Struct.new(
  :authorisationUrl,
  :callbackUrl,
  :consentID,
  :emailAddress,
  :expiryDate,
  :failureCallbackUrl,
  :id,
  :inserted,
  :institutionID,
  :isConnectedAccounts,
  :isEnabled,
  :merchantID,
  :provider,
  :successWebHookUrl,
  keyword_init: true
)

# Request payload for Consent#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authorisationUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] institutionID
#   @return [String, nil]
#
# @!attribute [rw] isConnectedAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] provider
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
ConsentUpdateData = Struct.new(
  :id,
  :authorisationUrl,
  :callbackUrl,
  :consentID,
  :emailAddress,
  :expiryDate,
  :failureCallbackUrl,
  :inserted,
  :institutionID,
  :isConnectedAccounts,
  :isEnabled,
  :merchantID,
  :provider,
  :successWebHookUrl,
  keyword_init: true
)

# Request payload for Consent#remove.
#
# @!attribute [rw] id
#   @return [String]
ConsentRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] decimals
#   @return [Integer, nil]
#
# @!attribute [rw] isFiat
#   @return [Boolean, nil]
#
# @!attribute [rw] iso4217AlphaCode
#   @return [String, nil]
#
# @!attribute [rw] iso4217NumericCode
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
Currency = Struct.new(
  :code,
  :decimals,
  :isFiat,
  :iso4217AlphaCode,
  :iso4217NumericCode,
  :symbol,
  keyword_init: true
)

# Request payload for Currency#list.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] decimals
#   @return [Integer, nil]
#
# @!attribute [rw] isFiat
#   @return [Boolean, nil]
#
# @!attribute [rw] iso4217AlphaCode
#   @return [String, nil]
#
# @!attribute [rw] iso4217NumericCode
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
CurrencyListMatch = Struct.new(
  :code,
  :decimals,
  :isFiat,
  :iso4217AlphaCode,
  :iso4217NumericCode,
  :symbol,
  keyword_init: true
)

# DirectDebitBatchSubmit entity data model.
#
# @!attribute [rw] failedSubmissions
#   @return [Hash, nil]
#
# @!attribute [rw] successfulSubmissions
#   @return [Array, nil]
DirectDebitBatchSubmit = Struct.new(
  :failedSubmissions,
  :successfulSubmissions,
  keyword_init: true
)

# Request payload for DirectDebitBatchSubmit#create.
#
# @!attribute [rw] failedSubmissions
#   @return [Hash, nil]
#
# @!attribute [rw] successfulSubmissions
#   @return [Array, nil]
DirectDebitBatchSubmitCreateData = Struct.new(
  :failedSubmissions,
  :successfulSubmissions,
  keyword_init: true
)

# FxRate entity data model.
#
# @!attribute [rw] destinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] exchangeRate
#   @return [Float, nil]
#
# @!attribute [rw] expiryTime
#   @return [String, nil]
#
# @!attribute [rw] quoteID
#   @return [String, nil]
#
# @!attribute [rw] sourceCurrency
#   @return [String, nil]
FxRate = Struct.new(
  :destinationCurrency,
  :exchangeRate,
  :expiryTime,
  :quoteID,
  :sourceCurrency,
  keyword_init: true
)

# Request payload for FxRate#load.
#
# @!attribute [rw] destination
#   @return [String]
#
# @!attribute [rw] source
#   @return [String]
#
# @!attribute [rw] valid_for_minute
#   @return [Integer]
FxRateLoadMatch = Struct.new(
  :destination,
  :source,
  :valid_for_minute,
  keyword_init: true
)

# Request payload for FxRate#list.
#
# @!attribute [rw] destination
#   @return [String]
#
# @!attribute [rw] source
#   @return [String]
FxRateListMatch = Struct.new(
  :destination,
  :source,
  keyword_init: true
)

# IPayment entity data model.
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
IPayment = Struct.new(
  :paymentRequestID,
  :responseType,
  keyword_init: true
)

# Request payload for IPayment#create.
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
IPaymentCreateData = Struct.new(
  :paymentRequestID,
  :responseType,
  keyword_init: true
)

# Mandate entity data model.
#
# @!attribute [rw] accountNumber
#   @return [String, nil]
#
# @!attribute [rw] addressLine1
#   @return [String]
#
# @!attribute [rw] addressLine2
#   @return [String, nil]
#
# @!attribute [rw] approvedAt
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] countryCode
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customerAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] customerCity
#   @return [String, nil]
#
# @!attribute [rw] customerCountryCode
#   @return [String, nil]
#
# @!attribute [rw] customerCountryName
#   @return [String, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerFirstName
#   @return [String, nil]
#
# @!attribute [rw] customerIban
#   @return [String, nil]
#
# @!attribute [rw] customerLastName
#   @return [String, nil]
#
# @!attribute [rw] customerSortCode
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String]
#
# @!attribute [rw] firstName
#   @return [String]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isRecurring
#   @return [Boolean, nil]
#
# @!attribute [rw] lastName
#   @return [String]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] postalCode
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] sortCode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplierBankAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierCustomerID
#   @return [String, nil]
#
# @!attribute [rw] supplierMandateID
#   @return [String, nil]
#
# @!attribute [rw] supplierName
#   @return [String, nil]
#
# @!attribute [rw] supplierStatus
#   @return [String, nil]
Mandate = Struct.new(
  :accountNumber,
  :addressLine1,
  :addressLine2,
  :approvedAt,
  :city,
  :countryCode,
  :currency,
  :customerAccountNumber,
  :customerCity,
  :customerCountryCode,
  :customerCountryName,
  :customerEmailAddress,
  :customerFirstName,
  :customerIban,
  :customerLastName,
  :customerSortCode,
  :emailAddress,
  :firstName,
  :iban,
  :id,
  :inserted,
  :isRecurring,
  :lastName,
  :lastUpdated,
  :merchantID,
  :postalCode,
  :reference,
  :sortCode,
  :status,
  :supplierBankAccountID,
  :supplierCustomerID,
  :supplierMandateID,
  :supplierName,
  :supplierStatus,
  keyword_init: true
)

# Request payload for Mandate#load.
#
# @!attribute [rw] id
#   @return [String]
MandateLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Mandate#create.
#
# @!attribute [rw] accountNumber
#   @return [String, nil]
#
# @!attribute [rw] addressLine1
#   @return [String]
#
# @!attribute [rw] addressLine2
#   @return [String, nil]
#
# @!attribute [rw] approvedAt
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] countryCode
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customerAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] customerCity
#   @return [String, nil]
#
# @!attribute [rw] customerCountryCode
#   @return [String, nil]
#
# @!attribute [rw] customerCountryName
#   @return [String, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerFirstName
#   @return [String, nil]
#
# @!attribute [rw] customerIban
#   @return [String, nil]
#
# @!attribute [rw] customerLastName
#   @return [String, nil]
#
# @!attribute [rw] customerSortCode
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String]
#
# @!attribute [rw] firstName
#   @return [String]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isRecurring
#   @return [Boolean, nil]
#
# @!attribute [rw] lastName
#   @return [String]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] postalCode
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] sortCode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplierBankAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierCustomerID
#   @return [String, nil]
#
# @!attribute [rw] supplierMandateID
#   @return [String, nil]
#
# @!attribute [rw] supplierName
#   @return [String, nil]
#
# @!attribute [rw] supplierStatus
#   @return [String, nil]
MandateCreateData = Struct.new(
  :accountNumber,
  :addressLine1,
  :addressLine2,
  :approvedAt,
  :city,
  :countryCode,
  :currency,
  :customerAccountNumber,
  :customerCity,
  :customerCountryCode,
  :customerCountryName,
  :customerEmailAddress,
  :customerFirstName,
  :customerIban,
  :customerLastName,
  :customerSortCode,
  :emailAddress,
  :firstName,
  :iban,
  :id,
  :inserted,
  :isRecurring,
  :lastName,
  :lastUpdated,
  :merchantID,
  :postalCode,
  :reference,
  :sortCode,
  :status,
  :supplierBankAccountID,
  :supplierCustomerID,
  :supplierMandateID,
  :supplierName,
  :supplierStatus,
  keyword_init: true
)

# Merchant entity data model.
#
# @!attribute [rw] accountCurrencies
#   @return [Array, nil]
#
# @!attribute [rw] canHaveTrustAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] cardPaymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] companyID
#   @return [String, nil]
#
# @!attribute [rw] displayQrOnHostedPay
#   @return [Boolean, nil]
#
# @!attribute [rw] hostedPayVersion
#   @return [Integer, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isBlocked
#   @return [Boolean, nil]
#
# @!attribute [rw] isExited
#   @return [Boolean, nil]
#
# @!attribute [rw] isSuspended
#   @return [Boolean, nil]
#
# @!attribute [rw] jurisdiction
#   @return [String, nil]
#
# @!attribute [rw] logoUrlPng
#   @return [String, nil]
#
# @!attribute [rw] logoUrlSvg
#   @return [String, nil]
#
# @!attribute [rw] merchantCategoryCode
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] parentMerchant
#   @return [Hash, nil]
#
# @!attribute [rw] paymentAccountLimit
#   @return [Integer, nil]
#
# @!attribute [rw] paymentAccounts
#   @return [Array, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] shortName
#   @return [String, nil]
#
# @!attribute [rw] supportedPaymentMethodsList
#   @return [Array, nil]
#
# @!attribute [rw] suspensionReason
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] tradingName
#   @return [String, nil]
#
# @!attribute [rw] webHookLimit
#   @return [Integer, nil]
#
# @!attribute [rw] yourRoleName
#   @return [String, nil]
Merchant = Struct.new(
  :accountCurrencies,
  :canHaveTrustAccounts,
  :cardPaymentProcessor,
  :companyID,
  :displayQrOnHostedPay,
  :hostedPayVersion,
  :id,
  :inserted,
  :isBlocked,
  :isExited,
  :isSuspended,
  :jurisdiction,
  :logoUrlPng,
  :logoUrlSvg,
  :merchantCategoryCode,
  :name,
  :notes,
  :parentMerchant,
  :paymentAccountLimit,
  :paymentAccounts,
  :reason,
  :shortName,
  :supportedPaymentMethodsList,
  :suspensionReason,
  :tags,
  :timeZoneId,
  :tradingName,
  :webHookLimit,
  :yourRoleName,
  keyword_init: true
)

# Request payload for Merchant#load.
#
# @!attribute [rw] id
#   @return [String]
MerchantLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Merchant#list.
#
# @!attribute [rw] accountCurrencies
#   @return [Array, nil]
#
# @!attribute [rw] canHaveTrustAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] cardPaymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] companyID
#   @return [String, nil]
#
# @!attribute [rw] displayQrOnHostedPay
#   @return [Boolean, nil]
#
# @!attribute [rw] hostedPayVersion
#   @return [Integer, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isBlocked
#   @return [Boolean, nil]
#
# @!attribute [rw] isExited
#   @return [Boolean, nil]
#
# @!attribute [rw] isSuspended
#   @return [Boolean, nil]
#
# @!attribute [rw] jurisdiction
#   @return [String, nil]
#
# @!attribute [rw] logoUrlPng
#   @return [String, nil]
#
# @!attribute [rw] logoUrlSvg
#   @return [String, nil]
#
# @!attribute [rw] merchantCategoryCode
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] parentMerchant
#   @return [Hash, nil]
#
# @!attribute [rw] paymentAccountLimit
#   @return [Integer, nil]
#
# @!attribute [rw] paymentAccounts
#   @return [Array, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] shortName
#   @return [String, nil]
#
# @!attribute [rw] supportedPaymentMethodsList
#   @return [Array, nil]
#
# @!attribute [rw] suspensionReason
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] tradingName
#   @return [String, nil]
#
# @!attribute [rw] webHookLimit
#   @return [Integer, nil]
#
# @!attribute [rw] yourRoleName
#   @return [String, nil]
MerchantListMatch = Struct.new(
  :accountCurrencies,
  :canHaveTrustAccounts,
  :cardPaymentProcessor,
  :companyID,
  :displayQrOnHostedPay,
  :hostedPayVersion,
  :id,
  :inserted,
  :isBlocked,
  :isExited,
  :isSuspended,
  :jurisdiction,
  :logoUrlPng,
  :logoUrlSvg,
  :merchantCategoryCode,
  :name,
  :notes,
  :parentMerchant,
  :paymentAccountLimit,
  :paymentAccounts,
  :reason,
  :shortName,
  :supportedPaymentMethodsList,
  :suspensionReason,
  :tags,
  :timeZoneId,
  :tradingName,
  :webHookLimit,
  :yourRoleName,
  keyword_init: true
)

# Request payload for Merchant#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountCurrencies
#   @return [Array, nil]
#
# @!attribute [rw] canHaveTrustAccounts
#   @return [Boolean, nil]
#
# @!attribute [rw] cardPaymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] companyID
#   @return [String, nil]
#
# @!attribute [rw] displayQrOnHostedPay
#   @return [Boolean, nil]
#
# @!attribute [rw] hostedPayVersion
#   @return [Integer, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isBlocked
#   @return [Boolean, nil]
#
# @!attribute [rw] isExited
#   @return [Boolean, nil]
#
# @!attribute [rw] isSuspended
#   @return [Boolean, nil]
#
# @!attribute [rw] jurisdiction
#   @return [String, nil]
#
# @!attribute [rw] logoUrlPng
#   @return [String, nil]
#
# @!attribute [rw] logoUrlSvg
#   @return [String, nil]
#
# @!attribute [rw] merchantCategoryCode
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] parentMerchant
#   @return [Hash, nil]
#
# @!attribute [rw] paymentAccountLimit
#   @return [Integer, nil]
#
# @!attribute [rw] paymentAccounts
#   @return [Array, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] shortName
#   @return [String, nil]
#
# @!attribute [rw] supportedPaymentMethodsList
#   @return [Array, nil]
#
# @!attribute [rw] suspensionReason
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] tradingName
#   @return [String, nil]
#
# @!attribute [rw] webHookLimit
#   @return [Integer, nil]
#
# @!attribute [rw] yourRoleName
#   @return [String, nil]
MerchantUpdateData = Struct.new(
  :id,
  :accountCurrencies,
  :canHaveTrustAccounts,
  :cardPaymentProcessor,
  :companyID,
  :displayQrOnHostedPay,
  :hostedPayVersion,
  :inserted,
  :isBlocked,
  :isExited,
  :isSuspended,
  :jurisdiction,
  :logoUrlPng,
  :logoUrlSvg,
  :merchantCategoryCode,
  :name,
  :notes,
  :parentMerchant,
  :paymentAccountLimit,
  :paymentAccounts,
  :reason,
  :shortName,
  :supportedPaymentMethodsList,
  :suspensionReason,
  :tags,
  :timeZoneId,
  :tradingName,
  :webHookLimit,
  :yourRoleName,
  keyword_init: true
)

# Request payload for Merchant#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] user_id
#   @return [String]
MerchantRemoveMatch = Struct.new(
  :id,
  :user_id,
  keyword_init: true
)

# MerchantAuthorisationSetting entity data model.
#
# @!attribute [rw] amountLower
#   @return [Float, nil]
#
# @!attribute [rw] amountUpper
#   @return [Float, nil]
#
# @!attribute [rw] authorisationType
#   @return [String, nil]
#
# @!attribute [rw] beneficiariesOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lastEditorCantAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] numberOfAuthorisers
#   @return [Integer, nil]
#
# @!attribute [rw] roleSettings
#   @return [Array, nil]
MerchantAuthorisationSetting = Struct.new(
  :amountLower,
  :amountUpper,
  :authorisationType,
  :beneficiariesOnly,
  :id,
  :inserted,
  :lastEditorCantAuthorise,
  :lastUpdated,
  :merchantID,
  :numberOfAuthorisers,
  :roleSettings,
  keyword_init: true
)

# Request payload for MerchantAuthorisationSetting#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantAuthorisationSettingListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# MerchantDirectDebitMandatePage entity data model.
#
# @!attribute [rw] approvedAt
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customerAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] customerCity
#   @return [String, nil]
#
# @!attribute [rw] customerCountryCode
#   @return [String, nil]
#
# @!attribute [rw] customerCountryName
#   @return [String, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerFirstName
#   @return [String, nil]
#
# @!attribute [rw] customerIban
#   @return [String, nil]
#
# @!attribute [rw] customerLastName
#   @return [String, nil]
#
# @!attribute [rw] customerSortCode
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isRecurring
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplierBankAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierCustomerID
#   @return [String, nil]
#
# @!attribute [rw] supplierMandateID
#   @return [String, nil]
#
# @!attribute [rw] supplierName
#   @return [String, nil]
#
# @!attribute [rw] supplierStatus
#   @return [String, nil]
MerchantDirectDebitMandatePage = Struct.new(
  :approvedAt,
  :currency,
  :customerAccountNumber,
  :customerCity,
  :customerCountryCode,
  :customerCountryName,
  :customerEmailAddress,
  :customerFirstName,
  :customerIban,
  :customerLastName,
  :customerSortCode,
  :id,
  :inserted,
  :isRecurring,
  :lastUpdated,
  :merchantID,
  :reference,
  :status,
  :supplierBankAccountID,
  :supplierCustomerID,
  :supplierMandateID,
  :supplierName,
  :supplierStatus,
  keyword_init: true
)

# Request payload for MerchantDirectDebitMandatePage#list.
#
# @!attribute [rw] approvedAt
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customerAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] customerCity
#   @return [String, nil]
#
# @!attribute [rw] customerCountryCode
#   @return [String, nil]
#
# @!attribute [rw] customerCountryName
#   @return [String, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerFirstName
#   @return [String, nil]
#
# @!attribute [rw] customerIban
#   @return [String, nil]
#
# @!attribute [rw] customerLastName
#   @return [String, nil]
#
# @!attribute [rw] customerSortCode
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isRecurring
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplierBankAccountID
#   @return [String, nil]
#
# @!attribute [rw] supplierCustomerID
#   @return [String, nil]
#
# @!attribute [rw] supplierMandateID
#   @return [String, nil]
#
# @!attribute [rw] supplierName
#   @return [String, nil]
#
# @!attribute [rw] supplierStatus
#   @return [String, nil]
MerchantDirectDebitMandatePageListMatch = Struct.new(
  :approvedAt,
  :currency,
  :customerAccountNumber,
  :customerCity,
  :customerCountryCode,
  :customerCountryName,
  :customerEmailAddress,
  :customerFirstName,
  :customerIban,
  :customerLastName,
  :customerSortCode,
  :id,
  :inserted,
  :isRecurring,
  :lastUpdated,
  :merchantID,
  :reference,
  :status,
  :supplierBankAccountID,
  :supplierCustomerID,
  :supplierMandateID,
  :supplierName,
  :supplierStatus,
  keyword_init: true
)

# MerchantPayByBankSetting entity data model.
#
# @!attribute [rw] bankCountryCodes
#   @return [Array, nil]
#
# @!attribute [rw] bankID
#   @return [String, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] businessInstitutionID
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] logo
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] messageImageUrl
#   @return [String, nil]
#
# @!attribute [rw] order
#   @return [Integer, nil]
#
# @!attribute [rw] personalInstitutionID
#   @return [String, nil]
#
# @!attribute [rw] processor
#   @return [String, nil]
#
# @!attribute [rw] warningHeading
#   @return [String, nil]
#
# @!attribute [rw] warningMessage
#   @return [String, nil]
MerchantPayByBankSetting = Struct.new(
  :bankCountryCodes,
  :bankID,
  :bankName,
  :businessInstitutionID,
  :currency,
  :logo,
  :message,
  :messageImageUrl,
  :order,
  :personalInstitutionID,
  :processor,
  :warningHeading,
  :warningMessage,
  keyword_init: true
)

# Request payload for MerchantPayByBankSetting#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantPayByBankSettingListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# MerchantPaymentRequestTemplate entity data model.
#
# @!attribute [rw] bankPaymentOptions
#   @return [Hash, nil]
#
# @!attribute [rw] cardPaymentAddressOptions
#   @return [Hash, nil]
#
# @!attribute [rw] cardPaymentCaptureOptions
#   @return [Hash, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] defaultFields
#   @return [Array, nil]
#
# @!attribute [rw] description
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] notificationOptions
#   @return [Hash, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Hash, nil]
#
# @!attribute [rw] paymentTerms
#   @return [Hash, nil]
#
# @!attribute [rw] priorityBankOptions
#   @return [Hash, nil]
#
# @!attribute [rw] template
#   @return [Hash]
MerchantPaymentRequestTemplate = Struct.new(
  :bankPaymentOptions,
  :cardPaymentAddressOptions,
  :cardPaymentCaptureOptions,
  :customFields,
  :defaultFields,
  :description,
  :id,
  :inserted,
  :lastUpdated,
  :merchantID,
  :name,
  :notificationOptions,
  :paymentMethods,
  :paymentTerms,
  :priorityBankOptions,
  :template,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
MerchantPaymentRequestTemplateLoadMatch = Struct.new(
  :id,
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantPaymentRequestTemplateListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
#
# @!attribute [rw] bankPaymentOptions
#   @return [Hash, nil]
#
# @!attribute [rw] cardPaymentAddressOptions
#   @return [Hash, nil]
#
# @!attribute [rw] cardPaymentCaptureOptions
#   @return [Hash, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] defaultFields
#   @return [Array, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notificationOptions
#   @return [Hash, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Hash, nil]
#
# @!attribute [rw] paymentTerms
#   @return [Hash, nil]
#
# @!attribute [rw] priorityBankOptions
#   @return [Hash, nil]
#
# @!attribute [rw] template
#   @return [Hash, nil]
MerchantPaymentRequestTemplateUpdateData = Struct.new(
  :id,
  :paymentrequest_id,
  :bankPaymentOptions,
  :cardPaymentAddressOptions,
  :cardPaymentCaptureOptions,
  :customFields,
  :defaultFields,
  :description,
  :inserted,
  :lastUpdated,
  :merchantID,
  :name,
  :notificationOptions,
  :paymentMethods,
  :paymentTerms,
  :priorityBankOptions,
  :template,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
MerchantPaymentRequestTemplateRemoveMatch = Struct.new(
  :id,
  :paymentrequest_id,
  keyword_init: true
)

# MerchantToken entity data model.
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] expiresAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] hmacAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] ipAddressWhitelist
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] permissionTypes
#   @return [Array, nil]
#
# @!attribute [rw] requestSignatureVersion
#   @return [Integer, nil]
#
# @!attribute [rw] sharedSecretAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] sharedSecretBase64
#   @return [String, nil]
#
# @!attribute [rw] token
#   @return [String, nil]
MerchantToken = Struct.new(
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :description,
  :expiresAt,
  :hasCurrentUserAuthorised,
  :hmacAlgorithm,
  :id,
  :inserted,
  :ipAddressWhitelist,
  :isArchived,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :nonce,
  :permissionTypes,
  :requestSignatureVersion,
  :sharedSecretAlgorithm,
  :sharedSecretBase64,
  :token,
  keyword_init: true
)

# Request payload for MerchantToken#load.
#
# @!attribute [rw] id
#   @return [String]
MerchantTokenLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for MerchantToken#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantTokenListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for MerchantToken#create.
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] expiresAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] hmacAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] ipAddressWhitelist
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] permissionTypes
#   @return [Array, nil]
#
# @!attribute [rw] requestSignatureVersion
#   @return [Integer, nil]
#
# @!attribute [rw] sharedSecretAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] sharedSecretBase64
#   @return [String, nil]
#
# @!attribute [rw] token
#   @return [String, nil]
MerchantTokenCreateData = Struct.new(
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :description,
  :expiresAt,
  :hasCurrentUserAuthorised,
  :hmacAlgorithm,
  :id,
  :inserted,
  :ipAddressWhitelist,
  :isArchived,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :nonce,
  :permissionTypes,
  :requestSignatureVersion,
  :sharedSecretAlgorithm,
  :sharedSecretBase64,
  :token,
  keyword_init: true
)

# Request payload for MerchantToken#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] expiresAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] hmacAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] ipAddressWhitelist
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastAuthorised
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] permissionTypes
#   @return [Array, nil]
#
# @!attribute [rw] requestSignatureVersion
#   @return [Integer, nil]
#
# @!attribute [rw] sharedSecretAlgorithm
#   @return [String, nil]
#
# @!attribute [rw] sharedSecretBase64
#   @return [String, nil]
#
# @!attribute [rw] token
#   @return [String, nil]
MerchantTokenUpdateData = Struct.new(
  :id,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :description,
  :expiresAt,
  :hasCurrentUserAuthorised,
  :hmacAlgorithm,
  :inserted,
  :ipAddressWhitelist,
  :isArchived,
  :isEnabled,
  :lastAuthorised,
  :lastUpdated,
  :merchantID,
  :nonce,
  :permissionTypes,
  :requestSignatureVersion,
  :sharedSecretAlgorithm,
  :sharedSecretBase64,
  :token,
  keyword_init: true
)

# Metadata entity data model.
class Metadata
end

# Request payload for Metadata#load.
class MetadataLoadMatch
end

# NoFrixionVersion entity data model.
#
# @!attribute [rw] buildVersion
#   @return [Integer, nil]
#
# @!attribute [rw] majorVersion
#   @return [Integer, nil]
#
# @!attribute [rw] minorVersion
#   @return [Integer, nil]
#
# @!attribute [rw] releaseName
#   @return [String, nil]
NoFrixionVersion = Struct.new(
  :buildVersion,
  :majorVersion,
  :minorVersion,
  :releaseName,
  keyword_init: true
)

# Request payload for NoFrixionVersion#load.
#
# @!attribute [rw] buildVersion
#   @return [Integer, nil]
#
# @!attribute [rw] majorVersion
#   @return [Integer, nil]
#
# @!attribute [rw] minorVersion
#   @return [Integer, nil]
#
# @!attribute [rw] releaseName
#   @return [String, nil]
NoFrixionVersionLoadMatch = Struct.new(
  :buildVersion,
  :majorVersion,
  :minorVersion,
  :releaseName,
  keyword_init: true
)

# OpenBanking entity data model.
class OpenBanking
end

# Request payload for OpenBanking#create.
#
# @!attribute [rw] account_id
#   @return [String]
OpenBankingCreateData = Struct.new(
  :account_id,
  keyword_init: true
)

# Request payload for OpenBanking#remove.
#
# @!attribute [rw] account_id
#   @return [String]
OpenBankingRemoveMatch = Struct.new(
  :account_id,
  keyword_init: true
)

# Payeeverification entity data model.
#
# @!attribute [rw] accountName
#   @return [String]
#
# @!attribute [rw] accountNumber
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [String]
#
# @!attribute [rw] payeeVerifiedAccountName
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
#
# @!attribute [rw] secondaryIdentification
#   @return [String, nil]
#
# @!attribute [rw] sortCode
#   @return [String, nil]
Payeeverification = Struct.new(
  :accountName,
  :accountNumber,
  :iban,
  :payeeVerifiedAccountName,
  :result,
  :secondaryIdentification,
  :sortCode,
  keyword_init: true
)

# Request payload for Payeeverification#create.
#
# @!attribute [rw] accountName
#   @return [String]
#
# @!attribute [rw] accountNumber
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [String]
#
# @!attribute [rw] payeeVerifiedAccountName
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
#
# @!attribute [rw] secondaryIdentification
#   @return [String, nil]
#
# @!attribute [rw] sortCode
#   @return [String, nil]
PayeeverificationCreateData = Struct.new(
  :accountName,
  :accountNumber,
  :iban,
  :payeeVerifiedAccountName,
  :result,
  :secondaryIdentification,
  :sortCode,
  keyword_init: true
)

# Payment entity data model.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardNoPayerAuthentication
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] cardTransmitRawDetails
#   @return [Boolean, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] shippingAddressCity
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCountryCode
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCounty
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine1
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine2
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressPostCode
#   @return [String, nil]
#
# @!attribute [rw] shippingEmail
#   @return [String, nil]
#
# @!attribute [rw] shippingFirstName
#   @return [String, nil]
#
# @!attribute [rw] shippingLastName
#   @return [String, nil]
#
# @!attribute [rw] shippingPhone
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
Payment = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardNoPayerAuthentication,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :cardTransmitRawDetails,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :dueDate,
  :events,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentMethods,
  :paymentProcessor,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :shippingAddressCity,
  :shippingAddressCountryCode,
  :shippingAddressCounty,
  :shippingAddressLine1,
  :shippingAddressLine2,
  :shippingAddressPostCode,
  :shippingEmail,
  :shippingFirstName,
  :shippingLastName,
  :shippingPhone,
  :status,
  :successWebHookUrl,
  :tagIds,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for Payment#load.
#
# @!attribute [rw] id
#   @return [String]
PaymentLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payment#create.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardNoPayerAuthentication
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] cardTransmitRawDetails
#   @return [Boolean, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] shippingAddressCity
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCountryCode
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCounty
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine1
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine2
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressPostCode
#   @return [String, nil]
#
# @!attribute [rw] shippingEmail
#   @return [String, nil]
#
# @!attribute [rw] shippingFirstName
#   @return [String, nil]
#
# @!attribute [rw] shippingLastName
#   @return [String, nil]
#
# @!attribute [rw] shippingPhone
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentCreateData = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardNoPayerAuthentication,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :cardTransmitRawDetails,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :dueDate,
  :events,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentMethods,
  :paymentProcessor,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :shippingAddressCity,
  :shippingAddressCountryCode,
  :shippingAddressCounty,
  :shippingAddressLine1,
  :shippingAddressLine2,
  :shippingAddressPostCode,
  :shippingEmail,
  :shippingFirstName,
  :shippingLastName,
  :shippingPhone,
  :status,
  :successWebHookUrl,
  :tagIds,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for Payment#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardNoPayerAuthentication
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] cardTransmitRawDetails
#   @return [Boolean, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] shippingAddressCity
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCountryCode
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressCounty
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine1
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressLine2
#   @return [String, nil]
#
# @!attribute [rw] shippingAddressPostCode
#   @return [String, nil]
#
# @!attribute [rw] shippingEmail
#   @return [String, nil]
#
# @!attribute [rw] shippingFirstName
#   @return [String, nil]
#
# @!attribute [rw] shippingLastName
#   @return [String, nil]
#
# @!attribute [rw] shippingPhone
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentUpdateData = Struct.new(
  :id,
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardNoPayerAuthentication,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :cardTransmitRawDetails,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :dueDate,
  :events,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentMethods,
  :paymentProcessor,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :shippingAddressCity,
  :shippingAddressCountryCode,
  :shippingAddressCounty,
  :shippingAddressLine1,
  :shippingAddressLine2,
  :shippingAddressPostCode,
  :shippingEmail,
  :shippingFirstName,
  :shippingLastName,
  :shippingPhone,
  :status,
  :successWebHookUrl,
  :tagIds,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# PaymentAccount entity data model.
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
PaymentAccount = Struct.new(
  :accountName,
  :accountSupplierName,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :physicalAccountID,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierSepaInstantStatus,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for PaymentAccount#list.
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
PaymentAccountListMatch = Struct.new(
  :accountName,
  :accountSupplierName,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :physicalAccountID,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierSepaInstantStatus,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# PaymentAccountMinimal entity data model.
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
PaymentAccountMinimal = Struct.new(
  :accountName,
  :availableBalance,
  :balance,
  :balanceMinorUnits,
  :currency,
  :id,
  :identifier,
  :isArchived,
  :isConnectedAccount,
  :merchantID,
  :submittedPayoutsBalance,
  keyword_init: true
)

# Request payload for PaymentAccountMinimal#list.
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
PaymentAccountMinimalListMatch = Struct.new(
  :accountName,
  :availableBalance,
  :balance,
  :balanceMinorUnits,
  :currency,
  :id,
  :identifier,
  :isArchived,
  :isConnectedAccount,
  :merchantID,
  :submittedPayoutsBalance,
  keyword_init: true
)

# PaymentInitiation entity data model.
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] redirectUrl
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] specificErrorMessage
#   @return [String, nil]
PaymentInitiation = Struct.new(
  :paymentInitiationID,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :redirectUrl,
  :responseType,
  :specificErrorMessage,
  keyword_init: true
)

# Request payload for PaymentInitiation#create.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] redirectUrl
#   @return [String, nil]
#
# @!attribute [rw] responseType
#   @return [String, nil]
#
# @!attribute [rw] specificErrorMessage
#   @return [String, nil]
PaymentInitiationCreateData = Struct.new(
  :paymentrequest_id,
  :paymentInitiationID,
  :paymentRequestCallbackUrl,
  :paymentRequestID,
  :redirectUrl,
  :responseType,
  :specificErrorMessage,
  keyword_init: true
)

# PaymentRequest entity data model.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] doSimulateSettlementFailure
#   @return [Boolean, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] errorDescription
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPaymentRequests
#   @return [Hash, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRequests
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentRequest = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :doSimulateSettlementFailure,
  :dueDate,
  :errorDescription,
  :events,
  :failedPaymentRequests,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :institution,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentInitiationID,
  :paymentMethods,
  :paymentProcessor,
  :paymentRequests,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :status,
  :successWebHookUrl,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for PaymentRequest#load.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] doSimulateSettlementFailure
#   @return [Boolean, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] errorDescription
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPaymentRequests
#   @return [Hash, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRequests
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentRequestLoadMatch = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :doSimulateSettlementFailure,
  :dueDate,
  :errorDescription,
  :events,
  :failedPaymentRequests,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :institution,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentInitiationID,
  :paymentMethods,
  :paymentProcessor,
  :paymentRequests,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :status,
  :successWebHookUrl,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for PaymentRequest#list.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] doSimulateSettlementFailure
#   @return [Boolean, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] errorDescription
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPaymentRequests
#   @return [Hash, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRequests
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentRequestListMatch = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :doSimulateSettlementFailure,
  :dueDate,
  :errorDescription,
  :events,
  :failedPaymentRequests,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :institution,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentInitiationID,
  :paymentMethods,
  :paymentProcessor,
  :paymentRequests,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :status,
  :successWebHookUrl,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for PaymentRequest#create.
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] doSimulateSettlementFailure
#   @return [Boolean, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] errorDescription
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPaymentRequests
#   @return [Hash, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRequests
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentRequestCreateData = Struct.new(
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :doSimulateSettlementFailure,
  :dueDate,
  :errorDescription,
  :events,
  :failedPaymentRequests,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :institution,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentInitiationID,
  :paymentMethods,
  :paymentProcessor,
  :paymentRequests,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :status,
  :successWebHookUrl,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for PaymentRequest#update.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
#
# @!attribute [rw] addresses
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] autoSendReceipt
#   @return [Boolean, nil]
#
# @!attribute [rw] baseOriginUrl
#   @return [String, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizeOnly
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateToken
#   @return [Boolean, nil]
#
# @!attribute [rw] cardCreateTokenMode
#   @return [String, nil]
#
# @!attribute [rw] cardIgnoreCVN
#   @return [Boolean, nil]
#
# @!attribute [rw] cardProcessorMerchantID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentID
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] createdByUser
#   @return [Hash, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFields
#   @return [Array, nil]
#
# @!attribute [rw] customerEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] customerName
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destinationAccount
#   @return [Hash, nil]
#
# @!attribute [rw] directDebitPayment
#   @return [Hash, nil]
#
# @!attribute [rw] doSimulateSettlementFailure
#   @return [Boolean, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] errorDescription
#   @return [String, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPaymentRequests
#   @return [Hash, nil]
#
# @!attribute [rw] failureCallbackUrl
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] hostedPayCheckoutUrl
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignoreAddressVerification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] insertedSortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoiceExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] merchantDirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] notificationEmailAddresses
#   @return [String, nil]
#
# @!attribute [rw] notificationRoleIDs
#   @return [Array, nil]
#
# @!attribute [rw] orderID
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentSteps
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] paymentMethods
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRequests
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] pispAccountID
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandboxSettleDelayInSeconds
#   @return [Integer, nil]
#
# @!attribute [rw] shippingAddress
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] successWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenisedCards
#   @return [Array, nil]
#
# @!attribute [rw] transactions
#   @return [Array, nil]
#
# @!attribute [rw] useHostedPaymentPage
#   @return [Boolean, nil]
PaymentRequestUpdateData = Struct.new(
  :paymentrequest_id,
  :addresses,
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :autoSendReceipt,
  :baseOriginUrl,
  :callbackUrl,
  :cardAuthorizeOnly,
  :cardCreateToken,
  :cardCreateTokenMode,
  :cardIgnoreCVN,
  :cardProcessorMerchantID,
  :cardStripePaymentIntentID,
  :cardStripePaymentIntentSecret,
  :createdByUser,
  :currency,
  :customFields,
  :customerEmailAddress,
  :customerID,
  :customerName,
  :description,
  :destinationAccount,
  :directDebitPayment,
  :doSimulateSettlementFailure,
  :dueDate,
  :errorDescription,
  :events,
  :failedPaymentRequests,
  :failureCallbackUrl,
  :fieldDisplaySettings,
  :formattedAmount,
  :hostedPayCheckoutUrl,
  :id,
  :ignoreAddressVerification,
  :inserted,
  :insertedSortable,
  :institution,
  :isArchived,
  :jwk,
  :lastUpdated,
  :lightningInvoice,
  :lightningInvoiceExpiresAt,
  :merchantDirectDebitMandateID,
  :merchantID,
  :merchantTokenDescription,
  :notificationEmailAddresses,
  :notificationRoleIDs,
  :orderID,
  :partialPaymentMethod,
  :partialPaymentSteps,
  :paymentAttempts,
  :paymentInitiationID,
  :paymentMethods,
  :paymentProcessor,
  :paymentRequests,
  :payrunID,
  :pispAccountID,
  :priorityBankID,
  :result,
  :sandboxSettleDelayInSeconds,
  :shippingAddress,
  :status,
  :successWebHookUrl,
  :tags,
  :title,
  :tokenisedCards,
  :transactions,
  :useHostedPaymentPage,
  keyword_init: true
)

# Request payload for PaymentRequest#remove.
#
# @!attribute [rw] id
#   @return [String]
PaymentRequestRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# PaymentRequestEvent entity data model.
#
# @!attribute [rw] amount
#   @return [Float]
#
# @!attribute [rw] applePayTransactionID
#   @return [String, nil]
#
# @!attribute [rw] cardAuthorizationResponseID
#   @return [String, nil]
#
# @!attribute [rw] cardExpiryMonth
#   @return [Integer, nil]
#
# @!attribute [rw] cardExpiryYear
#   @return [Integer, nil]
#
# @!attribute [rw] cardIssuer
#   @return [String, nil]
#
# @!attribute [rw] cardIssuerCountry
#   @return [String, nil]
#
# @!attribute [rw] cardLastFourDigits
#   @return [String, nil]
#
# @!attribute [rw] cardRequestID
#   @return [String, nil]
#
# @!attribute [rw] cardScheme
#   @return [String, nil]
#
# @!attribute [rw] cardTokenCustomerID
#   @return [String, nil]
#
# @!attribute [rw] cardTransactionID
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] directDebitPaymentID
#   @return [String, nil]
#
# @!attribute [rw] directDebitPaymentReference
#   @return [String, nil]
#
# @!attribute [rw] drirectDebitMandateID
#   @return [String, nil]
#
# @!attribute [rw] errorMessage
#   @return [String, nil]
#
# @!attribute [rw] errorReason
#   @return [String, nil]
#
# @!attribute [rw] eventType
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lightningInvoice
#   @return [String, nil]
#
# @!attribute [rw] lightningRHash
#   @return [String, nil]
#
# @!attribute [rw] originUrl
#   @return [String, nil]
#
# @!attribute [rw] paymentMethodType
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessorName
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] pispBankStatus
#   @return [String, nil]
#
# @!attribute [rw] pispPaymentInitiationID
#   @return [String, nil]
#
# @!attribute [rw] pispPaymentInstitutionName
#   @return [String, nil]
#
# @!attribute [rw] pispPaymentServiceProviderID
#   @return [String, nil]
#
# @!attribute [rw] pispRedirectUrl
#   @return [String, nil]
#
# @!attribute [rw] reconciledTransactionID
#   @return [String, nil]
#
# @!attribute [rw] refundPayoutID
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] walletName
#   @return [String, nil]
PaymentRequestEvent = Struct.new(
  :amount,
  :applePayTransactionID,
  :cardAuthorizationResponseID,
  :cardExpiryMonth,
  :cardExpiryYear,
  :cardIssuer,
  :cardIssuerCountry,
  :cardLastFourDigits,
  :cardRequestID,
  :cardScheme,
  :cardTokenCustomerID,
  :cardTransactionID,
  :currency,
  :directDebitPaymentID,
  :directDebitPaymentReference,
  :drirectDebitMandateID,
  :errorMessage,
  :errorReason,
  :eventType,
  :id,
  :inserted,
  :lightningInvoice,
  :lightningRHash,
  :originUrl,
  :paymentMethodType,
  :paymentProcessorName,
  :paymentRequestID,
  :pispBankStatus,
  :pispPaymentInitiationID,
  :pispPaymentInstitutionName,
  :pispPaymentServiceProviderID,
  :pispRedirectUrl,
  :reconciledTransactionID,
  :refundPayoutID,
  :status,
  :walletName,
  keyword_init: true
)

# Request payload for PaymentRequestEvent#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestEventListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# PaymentRequestMetric entity data model.
class PaymentRequestMetric
end

# Request payload for PaymentRequestMetric#load.
class PaymentRequestMetricLoadMatch
end

# PaymentRequestMinimal entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] callbackUrl
#   @return [String, nil]
#
# @!attribute [rw] cardStripePaymentIntentSecret
#   @return [String, nil]
#
# @!attribute [rw] countryCode
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customFieldsToDisplay
#   @return [Array, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] dueDate
#   @return [String, nil]
#
# @!attribute [rw] fieldDisplaySettings
#   @return [Array, nil]
#
# @!attribute [rw] googlePayMerchantID
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantLogoUrlPng
#   @return [String, nil]
#
# @!attribute [rw] merchantLogoUrlSvg
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] merchantShortName
#   @return [String, nil]
#
# @!attribute [rw] partialPaymentMethod
#   @return [String, nil]
#
# @!attribute [rw] paymentAttempts
#   @return [Array, nil]
#
# @!attribute [rw] paymentMethodsList
#   @return [Array, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessorKey
#   @return [String, nil]
#
# @!attribute [rw] pispError
#   @return [String, nil]
#
# @!attribute [rw] priorityBankID
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] stripeAccountID
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
PaymentRequestMinimal = Struct.new(
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :callbackUrl,
  :cardStripePaymentIntentSecret,
  :countryCode,
  :currency,
  :customFieldsToDisplay,
  :description,
  :dueDate,
  :fieldDisplaySettings,
  :googlePayMerchantID,
  :id,
  :jwk,
  :merchantID,
  :merchantLogoUrlPng,
  :merchantLogoUrlSvg,
  :merchantName,
  :merchantShortName,
  :partialPaymentMethod,
  :paymentAttempts,
  :paymentMethodsList,
  :paymentProcessor,
  :paymentProcessorKey,
  :pispError,
  :priorityBankID,
  :status,
  :stripeAccountID,
  :title,
  keyword_init: true
)

# Request payload for PaymentRequestMinimal#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestMinimalListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# PaymentRequestResult entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountPending
#   @return [Float, nil]
#
# @!attribute [rw] amountReceived
#   @return [Float, nil]
#
# @!attribute [rw] amountRefunded
#   @return [Float, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customerID
#   @return [String, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] payments
#   @return [Array, nil]
#
# @!attribute [rw] pispAuthorizations
#   @return [Array, nil]
#
# @!attribute [rw] requestedAmount
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
PaymentRequestResult = Struct.new(
  :amount,
  :amountPending,
  :amountReceived,
  :amountRefunded,
  :currency,
  :customerID,
  :paymentRequestID,
  :payments,
  :pispAuthorizations,
  :requestedAmount,
  :result,
  keyword_init: true
)

# Request payload for PaymentRequestResult#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestResultListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Payout entity data model.
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] allowIncomplete
#   @return [Boolean, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] approvePayoutUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash]
#
# @!attribute [rw] beneficiaryID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canProcess
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] chargeBearer
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [String, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currentUserID
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] documents
#   @return [Array, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPayouts
#   @return [Hash, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedFxDestinationAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedSchedule
#   @return [String, nil]
#
# @!attribute [rw] formattedScheduleDayOnly
#   @return [String, nil]
#
# @!attribute [rw] formattedSourceAccountAvailableBalance
#   @return [String, nil]
#
# @!attribute [rw] fxDestinationAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxDestinationAmountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] fxDestinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteID
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] fxUseDestinationAmount
#   @return [Boolean, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoiceID
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isFailed
#   @return [Boolean, nil]
#
# @!attribute [rw] isSettled
#   @return [Boolean, nil]
#
# @!attribute [rw] isSubmitted
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRail
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] payrunName
#   @return [String, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] sourceAccountAvailableBalance
#   @return [Float, nil]
#
# @!attribute [rw] sourceAccountAvailableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] sourceAccountBic
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountCurrency
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIban
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIdentifier
#   @return [Hash]
#
# @!attribute [rw] sourceAccountName
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountSortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] topupPayrunID
#   @return [String, nil]
#
# @!attribute [rw] transactedAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxRate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
Payout = Struct.new(
  :accountID,
  :allowIncomplete,
  :amount,
  :amountMinorUnits,
  :approvePayoutUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :beneficiary,
  :beneficiaryID,
  :canAuthorise,
  :canProcess,
  :canUpdate,
  :chargeBearer,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :currentUserID,
  :description,
  :destination,
  :documents,
  :events,
  :failedPayouts,
  :formattedAmount,
  :formattedFxDestinationAmount,
  :formattedSchedule,
  :formattedScheduleDayOnly,
  :formattedSourceAccountAvailableBalance,
  :fxDestinationAmount,
  :fxDestinationAmountMinorUnits,
  :fxDestinationCurrency,
  :fxQuoteExpiresAt,
  :fxQuoteID,
  :fxRate,
  :fxUseDestinationAmount,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :invoiceID,
  :isArchived,
  :isFailed,
  :isSettled,
  :isSubmitted,
  :lastUpdated,
  :merchantID,
  :merchantTokenDescription,
  :nonce,
  :paymentProcessor,
  :paymentRail,
  :payouts,
  :payrunID,
  :payrunName,
  :reason,
  :rule,
  :scheduleDate,
  :scheduled,
  :sourceAccountAvailableBalance,
  :sourceAccountAvailableBalanceMinorUnits,
  :sourceAccountBic,
  :sourceAccountCurrency,
  :sourceAccountIban,
  :sourceAccountIdentifier,
  :sourceAccountName,
  :sourceAccountNumber,
  :sourceAccountSortcode,
  :status,
  :tagIds,
  :tags,
  :theirReference,
  :topupPayrunID,
  :transactedAmount,
  :transactedFxAmount,
  :transactedFxRate,
  :type,
  :userID,
  :yourReference,
  keyword_init: true
)

# Request payload for Payout#load.
#
# @!attribute [rw] id
#   @return [String]
PayoutLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payout#list.
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] allowIncomplete
#   @return [Boolean, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] approvePayoutUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash, nil]
#
# @!attribute [rw] beneficiaryID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canProcess
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] chargeBearer
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [String, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currentUserID
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] documents
#   @return [Array, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPayouts
#   @return [Hash, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedFxDestinationAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedSchedule
#   @return [String, nil]
#
# @!attribute [rw] formattedScheduleDayOnly
#   @return [String, nil]
#
# @!attribute [rw] formattedSourceAccountAvailableBalance
#   @return [String, nil]
#
# @!attribute [rw] fxDestinationAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxDestinationAmountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] fxDestinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteID
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] fxUseDestinationAmount
#   @return [Boolean, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoiceID
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isFailed
#   @return [Boolean, nil]
#
# @!attribute [rw] isSettled
#   @return [Boolean, nil]
#
# @!attribute [rw] isSubmitted
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRail
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] payrunName
#   @return [String, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] sourceAccountAvailableBalance
#   @return [Float, nil]
#
# @!attribute [rw] sourceAccountAvailableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] sourceAccountBic
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountCurrency
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIban
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIdentifier
#   @return [Hash, nil]
#
# @!attribute [rw] sourceAccountName
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountSortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] topupPayrunID
#   @return [String, nil]
#
# @!attribute [rw] transactedAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxRate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
PayoutListMatch = Struct.new(
  :accountID,
  :allowIncomplete,
  :amount,
  :amountMinorUnits,
  :approvePayoutUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :beneficiary,
  :beneficiaryID,
  :canAuthorise,
  :canProcess,
  :canUpdate,
  :chargeBearer,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :currentUserID,
  :description,
  :destination,
  :documents,
  :events,
  :failedPayouts,
  :formattedAmount,
  :formattedFxDestinationAmount,
  :formattedSchedule,
  :formattedScheduleDayOnly,
  :formattedSourceAccountAvailableBalance,
  :fxDestinationAmount,
  :fxDestinationAmountMinorUnits,
  :fxDestinationCurrency,
  :fxQuoteExpiresAt,
  :fxQuoteID,
  :fxRate,
  :fxUseDestinationAmount,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :invoiceID,
  :isArchived,
  :isFailed,
  :isSettled,
  :isSubmitted,
  :lastUpdated,
  :merchantID,
  :merchantTokenDescription,
  :nonce,
  :paymentProcessor,
  :paymentRail,
  :payouts,
  :payrunID,
  :payrunName,
  :reason,
  :rule,
  :scheduleDate,
  :scheduled,
  :sourceAccountAvailableBalance,
  :sourceAccountAvailableBalanceMinorUnits,
  :sourceAccountBic,
  :sourceAccountCurrency,
  :sourceAccountIban,
  :sourceAccountIdentifier,
  :sourceAccountName,
  :sourceAccountNumber,
  :sourceAccountSortcode,
  :status,
  :tagIds,
  :tags,
  :theirReference,
  :topupPayrunID,
  :transactedAmount,
  :transactedFxAmount,
  :transactedFxRate,
  :type,
  :userID,
  :yourReference,
  keyword_init: true
)

# Request payload for Payout#create.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] allowIncomplete
#   @return [Boolean, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] approvePayoutUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash]
#
# @!attribute [rw] beneficiaryID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canProcess
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] chargeBearer
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [String, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currentUserID
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] documents
#   @return [Array, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPayouts
#   @return [Hash, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedFxDestinationAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedSchedule
#   @return [String, nil]
#
# @!attribute [rw] formattedScheduleDayOnly
#   @return [String, nil]
#
# @!attribute [rw] formattedSourceAccountAvailableBalance
#   @return [String, nil]
#
# @!attribute [rw] fxDestinationAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxDestinationAmountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] fxDestinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteID
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] fxUseDestinationAmount
#   @return [Boolean, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoiceID
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isFailed
#   @return [Boolean, nil]
#
# @!attribute [rw] isSettled
#   @return [Boolean, nil]
#
# @!attribute [rw] isSubmitted
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRail
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] payrunName
#   @return [String, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] sourceAccountAvailableBalance
#   @return [Float, nil]
#
# @!attribute [rw] sourceAccountAvailableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] sourceAccountBic
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountCurrency
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIban
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIdentifier
#   @return [Hash]
#
# @!attribute [rw] sourceAccountName
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountSortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] topupPayrunID
#   @return [String, nil]
#
# @!attribute [rw] transactedAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxRate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
PayoutCreateData = Struct.new(
  :id,
  :accountID,
  :allowIncomplete,
  :amount,
  :amountMinorUnits,
  :approvePayoutUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :beneficiary,
  :beneficiaryID,
  :canAuthorise,
  :canProcess,
  :canUpdate,
  :chargeBearer,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :currentUserID,
  :description,
  :destination,
  :documents,
  :events,
  :failedPayouts,
  :formattedAmount,
  :formattedFxDestinationAmount,
  :formattedSchedule,
  :formattedScheduleDayOnly,
  :formattedSourceAccountAvailableBalance,
  :fxDestinationAmount,
  :fxDestinationAmountMinorUnits,
  :fxDestinationCurrency,
  :fxQuoteExpiresAt,
  :fxQuoteID,
  :fxRate,
  :fxUseDestinationAmount,
  :hasCurrentUserAuthorised,
  :inserted,
  :invoiceID,
  :isArchived,
  :isFailed,
  :isSettled,
  :isSubmitted,
  :lastUpdated,
  :merchantID,
  :merchantTokenDescription,
  :nonce,
  :paymentProcessor,
  :paymentRail,
  :payouts,
  :payrunID,
  :payrunName,
  :reason,
  :rule,
  :scheduleDate,
  :scheduled,
  :sourceAccountAvailableBalance,
  :sourceAccountAvailableBalanceMinorUnits,
  :sourceAccountBic,
  :sourceAccountCurrency,
  :sourceAccountIban,
  :sourceAccountIdentifier,
  :sourceAccountName,
  :sourceAccountNumber,
  :sourceAccountSortcode,
  :status,
  :tagIds,
  :tags,
  :theirReference,
  :topupPayrunID,
  :transactedAmount,
  :transactedFxAmount,
  :transactedFxRate,
  :type,
  :userID,
  :yourReference,
  keyword_init: true
)

# Request payload for Payout#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] allowIncomplete
#   @return [Boolean, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] approvePayoutUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash, nil]
#
# @!attribute [rw] beneficiaryID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canProcess
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] chargeBearer
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [String, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currentUserID
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] documents
#   @return [Array, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] failedPayouts
#   @return [Hash, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedFxDestinationAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedSchedule
#   @return [String, nil]
#
# @!attribute [rw] formattedScheduleDayOnly
#   @return [String, nil]
#
# @!attribute [rw] formattedSourceAccountAvailableBalance
#   @return [String, nil]
#
# @!attribute [rw] fxDestinationAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxDestinationAmountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] fxDestinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteID
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] fxUseDestinationAmount
#   @return [Boolean, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoiceID
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isFailed
#   @return [Boolean, nil]
#
# @!attribute [rw] isSettled
#   @return [Boolean, nil]
#
# @!attribute [rw] isSubmitted
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRail
#   @return [String, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] payrunName
#   @return [String, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] sourceAccountAvailableBalance
#   @return [Float, nil]
#
# @!attribute [rw] sourceAccountAvailableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] sourceAccountBic
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountCurrency
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIban
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIdentifier
#   @return [Hash, nil]
#
# @!attribute [rw] sourceAccountName
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountSortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tagIds
#   @return [Array, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] topupPayrunID
#   @return [String, nil]
#
# @!attribute [rw] transactedAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxRate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
PayoutUpdateData = Struct.new(
  :id,
  :accountID,
  :allowIncomplete,
  :amount,
  :amountMinorUnits,
  :approvePayoutUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :beneficiary,
  :beneficiaryID,
  :canAuthorise,
  :canProcess,
  :canUpdate,
  :chargeBearer,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :currentUserID,
  :description,
  :destination,
  :documents,
  :events,
  :failedPayouts,
  :formattedAmount,
  :formattedFxDestinationAmount,
  :formattedSchedule,
  :formattedScheduleDayOnly,
  :formattedSourceAccountAvailableBalance,
  :fxDestinationAmount,
  :fxDestinationAmountMinorUnits,
  :fxDestinationCurrency,
  :fxQuoteExpiresAt,
  :fxQuoteID,
  :fxRate,
  :fxUseDestinationAmount,
  :hasCurrentUserAuthorised,
  :inserted,
  :invoiceID,
  :isArchived,
  :isFailed,
  :isSettled,
  :isSubmitted,
  :lastUpdated,
  :merchantID,
  :merchantTokenDescription,
  :nonce,
  :paymentProcessor,
  :paymentRail,
  :payouts,
  :payrunID,
  :payrunName,
  :reason,
  :rule,
  :scheduleDate,
  :scheduled,
  :sourceAccountAvailableBalance,
  :sourceAccountAvailableBalanceMinorUnits,
  :sourceAccountBic,
  :sourceAccountCurrency,
  :sourceAccountIban,
  :sourceAccountIdentifier,
  :sourceAccountName,
  :sourceAccountNumber,
  :sourceAccountSortcode,
  :status,
  :tagIds,
  :tags,
  :theirReference,
  :topupPayrunID,
  :transactedAmount,
  :transactedFxAmount,
  :transactedFxRate,
  :type,
  :userID,
  :yourReference,
  keyword_init: true
)

# Request payload for Payout#remove.
#
# @!attribute [rw] id
#   @return [String]
PayoutRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# PayoutKeysetPage entity data model.
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] approvePayoutUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canProcess
#   @return [Boolean, nil]
#
# @!attribute [rw] canUpdate
#   @return [Boolean, nil]
#
# @!attribute [rw] chargeBearer
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [String, nil]
#
# @!attribute [rw] createdByEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currentUserID
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] documents
#   @return [Array, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] formattedAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedFxDestinationAmount
#   @return [String, nil]
#
# @!attribute [rw] formattedSchedule
#   @return [String, nil]
#
# @!attribute [rw] formattedScheduleDayOnly
#   @return [String, nil]
#
# @!attribute [rw] formattedSourceAccountAvailableBalance
#   @return [String, nil]
#
# @!attribute [rw] fxDestinationAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxDestinationAmountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] fxDestinationCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteExpiresAt
#   @return [String, nil]
#
# @!attribute [rw] fxQuoteID
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] fxUseDestinationAmount
#   @return [Boolean, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoiceID
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isFailed
#   @return [Boolean, nil]
#
# @!attribute [rw] isSettled
#   @return [Boolean, nil]
#
# @!attribute [rw] isSubmitted
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantTokenDescription
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] paymentProcessor
#   @return [String, nil]
#
# @!attribute [rw] paymentRail
#   @return [String, nil]
#
# @!attribute [rw] payrunID
#   @return [String, nil]
#
# @!attribute [rw] payrunName
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] sourceAccountAvailableBalance
#   @return [Float, nil]
#
# @!attribute [rw] sourceAccountAvailableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] sourceAccountBic
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountCurrency
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIban
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountIdentifier
#   @return [Hash]
#
# @!attribute [rw] sourceAccountName
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountNumber
#   @return [String, nil]
#
# @!attribute [rw] sourceAccountSortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] topupPayrunID
#   @return [String, nil]
#
# @!attribute [rw] transactedAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxAmount
#   @return [Float, nil]
#
# @!attribute [rw] transactedFxRate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
PayoutKeysetPage = Struct.new(
  :accountID,
  :amount,
  :amountMinorUnits,
  :approvePayoutUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :beneficiary,
  :canAuthorise,
  :canProcess,
  :canUpdate,
  :chargeBearer,
  :createdBy,
  :createdByEmailAddress,
  :currency,
  :currentUserID,
  :description,
  :destination,
  :documents,
  :events,
  :formattedAmount,
  :formattedFxDestinationAmount,
  :formattedSchedule,
  :formattedScheduleDayOnly,
  :formattedSourceAccountAvailableBalance,
  :fxDestinationAmount,
  :fxDestinationAmountMinorUnits,
  :fxDestinationCurrency,
  :fxQuoteExpiresAt,
  :fxQuoteID,
  :fxRate,
  :fxUseDestinationAmount,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :invoiceID,
  :isArchived,
  :isFailed,
  :isSettled,
  :isSubmitted,
  :lastUpdated,
  :merchantID,
  :merchantTokenDescription,
  :nonce,
  :paymentProcessor,
  :paymentRail,
  :payrunID,
  :payrunName,
  :rule,
  :scheduleDate,
  :scheduled,
  :sourceAccountAvailableBalance,
  :sourceAccountAvailableBalanceMinorUnits,
  :sourceAccountBic,
  :sourceAccountCurrency,
  :sourceAccountIban,
  :sourceAccountIdentifier,
  :sourceAccountName,
  :sourceAccountNumber,
  :sourceAccountSortcode,
  :status,
  :tags,
  :theirReference,
  :topupPayrunID,
  :transactedAmount,
  :transactedFxAmount,
  :transactedFxRate,
  :type,
  :userID,
  :yourReference,
  keyword_init: true
)

# Request payload for PayoutKeysetPage#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
PayoutKeysetPageListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# PayoutMetric entity data model.
class PayoutMetric
end

# Request payload for PayoutMetric#load.
class PayoutMetricLoadMatch
end

# Payrun entity data model.
#
# @!attribute [rw] authorisationDate
#   @return [String, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canDelete
#   @return [Boolean, nil]
#
# @!attribute [rw] canEdit
#   @return [Boolean, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoices
#   @return [Array, nil]
#
# @!attribute [rw] invoicesMinimal
#   @return [Array, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lastUpdatedBy
#   @return [Hash]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] payments
#   @return [Array, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payoutsCount
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduledDate
#   @return [String, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] totalEur
#   @return [Float, nil]
#
# @!attribute [rw] totalGbp
#   @return [Float, nil]
#
# @!attribute [rw] totalUsd
#   @return [Float, nil]
Payrun = Struct.new(
  :authorisationDate,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :canAuthorise,
  :canDelete,
  :canEdit,
  :events,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :invoices,
  :invoicesMinimal,
  :isArchived,
  :lastUpdated,
  :lastUpdatedBy,
  :merchantID,
  :name,
  :nonce,
  :notes,
  :payments,
  :payouts,
  :payoutsCount,
  :reason,
  :scheduleDate,
  :scheduledDate,
  :sourceAccounts,
  :status,
  :totalEur,
  :totalGbp,
  :totalUsd,
  keyword_init: true
)

# Request payload for Payrun#load.
#
# @!attribute [rw] id
#   @return [String]
PayrunLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payrun#list.
#
# @!attribute [rw] authorisationDate
#   @return [String, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canDelete
#   @return [Boolean, nil]
#
# @!attribute [rw] canEdit
#   @return [Boolean, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoices
#   @return [Array, nil]
#
# @!attribute [rw] invoicesMinimal
#   @return [Array, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lastUpdatedBy
#   @return [Hash, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] payments
#   @return [Array, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payoutsCount
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduledDate
#   @return [String, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] totalEur
#   @return [Float, nil]
#
# @!attribute [rw] totalGbp
#   @return [Float, nil]
#
# @!attribute [rw] totalUsd
#   @return [Float, nil]
PayrunListMatch = Struct.new(
  :authorisationDate,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :canAuthorise,
  :canDelete,
  :canEdit,
  :events,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :invoices,
  :invoicesMinimal,
  :isArchived,
  :lastUpdated,
  :lastUpdatedBy,
  :merchantID,
  :name,
  :nonce,
  :notes,
  :payments,
  :payouts,
  :payoutsCount,
  :reason,
  :scheduleDate,
  :scheduledDate,
  :sourceAccounts,
  :status,
  :totalEur,
  :totalGbp,
  :totalUsd,
  keyword_init: true
)

# Request payload for Payrun#create.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authorisationDate
#   @return [String, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canDelete
#   @return [Boolean, nil]
#
# @!attribute [rw] canEdit
#   @return [Boolean, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoices
#   @return [Array, nil]
#
# @!attribute [rw] invoicesMinimal
#   @return [Array, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lastUpdatedBy
#   @return [Hash]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] payments
#   @return [Array, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payoutsCount
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduledDate
#   @return [String, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] totalEur
#   @return [Float, nil]
#
# @!attribute [rw] totalGbp
#   @return [Float, nil]
#
# @!attribute [rw] totalUsd
#   @return [Float, nil]
PayrunCreateData = Struct.new(
  :id,
  :authorisationDate,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :canAuthorise,
  :canDelete,
  :canEdit,
  :events,
  :hasCurrentUserAuthorised,
  :inserted,
  :invoices,
  :invoicesMinimal,
  :isArchived,
  :lastUpdated,
  :lastUpdatedBy,
  :merchantID,
  :name,
  :nonce,
  :notes,
  :payments,
  :payouts,
  :payoutsCount,
  :reason,
  :scheduleDate,
  :scheduledDate,
  :sourceAccounts,
  :status,
  :totalEur,
  :totalGbp,
  :totalUsd,
  keyword_init: true
)

# Request payload for Payrun#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authorisationDate
#   @return [String, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] batchPayoutID
#   @return [String, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] canDelete
#   @return [Boolean, nil]
#
# @!attribute [rw] canEdit
#   @return [Boolean, nil]
#
# @!attribute [rw] events
#   @return [Array, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoices
#   @return [Array, nil]
#
# @!attribute [rw] invoicesMinimal
#   @return [Array, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] lastUpdatedBy
#   @return [Hash, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] notes
#   @return [String, nil]
#
# @!attribute [rw] payments
#   @return [Array, nil]
#
# @!attribute [rw] payouts
#   @return [Array, nil]
#
# @!attribute [rw] payoutsCount
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] scheduleDate
#   @return [String, nil]
#
# @!attribute [rw] scheduledDate
#   @return [String, nil]
#
# @!attribute [rw] sourceAccounts
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] totalEur
#   @return [Float, nil]
#
# @!attribute [rw] totalGbp
#   @return [Float, nil]
#
# @!attribute [rw] totalUsd
#   @return [Float, nil]
PayrunUpdateData = Struct.new(
  :id,
  :authorisationDate,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :batchPayoutID,
  :canAuthorise,
  :canDelete,
  :canEdit,
  :events,
  :hasCurrentUserAuthorised,
  :inserted,
  :invoices,
  :invoicesMinimal,
  :isArchived,
  :lastUpdated,
  :lastUpdatedBy,
  :merchantID,
  :name,
  :nonce,
  :notes,
  :payments,
  :payouts,
  :payoutsCount,
  :reason,
  :scheduleDate,
  :scheduledDate,
  :sourceAccounts,
  :status,
  :totalEur,
  :totalGbp,
  :totalUsd,
  keyword_init: true
)

# Request payload for Payrun#remove.
#
# @!attribute [rw] id
#   @return [String]
PayrunRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Report entity data model.
class Report
end

# Request payload for Report#update.
#
# @!attribute [rw] id
#   @return [String]
ReportUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# ReportResult entity data model.
#
# @!attribute [rw] contentType
#   @return [String, nil]
#
# @!attribute [rw] contents
#   @return [String, nil]
#
# @!attribute [rw] lastCompletedAt
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] reportName
#   @return [String, nil]
#
# @!attribute [rw] reportType
#   @return [String, nil]
#
# @!attribute [rw] statementNumber
#   @return [Integer, nil]
ReportResult = Struct.new(
  :contentType,
  :contents,
  :lastCompletedAt,
  :merchantID,
  :reportName,
  :reportType,
  :statementNumber,
  keyword_init: true
)

# Request payload for ReportResult#load.
#
# @!attribute [rw] id
#   @return [Integer]
#
# @!attribute [rw] report_id
#   @return [String]
ReportResultLoadMatch = Struct.new(
  :id,
  :report_id,
  keyword_init: true
)

# Role entity data model.
#
# @!attribute [rw] failedRoles
#   @return [Hash, nil]
#
# @!attribute [rw] roles
#   @return [Array, nil]
Role = Struct.new(
  :failedRoles,
  :roles,
  keyword_init: true
)

# Request payload for Role#create.
#
# @!attribute [rw] merchant_id
#   @return [String]
#
# @!attribute [rw] failedRoles
#   @return [Hash, nil]
#
# @!attribute [rw] roles
#   @return [Array, nil]
RoleCreateData = Struct.new(
  :merchant_id,
  :failedRoles,
  :roles,
  keyword_init: true
)

# Rule entity data model.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] endAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isDisabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastExecutedAt
#   @return [String, nil]
#
# @!attribute [rw] lastRunAtTransactionDate
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] onApprovedWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionErrorWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionSuccessWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] startAt
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweepAction
#   @return [Hash, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] triggerCronExpression
#   @return [String, nil]
#
# @!attribute [rw] triggerOnPayIn
#   @return [Boolean, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] webHookSecret
#   @return [String, nil]
Rule = Struct.new(
  :account,
  :accountID,
  :approveUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :createdBy,
  :description,
  :endAt,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :isDisabled,
  :lastExecutedAt,
  :lastRunAtTransactionDate,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :onApprovedWebHookUrl,
  :onExecutionErrorWebHookUrl,
  :onExecutionSuccessWebHookUrl,
  :startAt,
  :status,
  :sweepAction,
  :timeZoneId,
  :triggerCronExpression,
  :triggerOnPayIn,
  :userID,
  :webHookSecret,
  keyword_init: true
)

# Request payload for Rule#load.
#
# @!attribute [rw] id
#   @return [String]
RuleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Rule#list.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] endAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isDisabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastExecutedAt
#   @return [String, nil]
#
# @!attribute [rw] lastRunAtTransactionDate
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] onApprovedWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionErrorWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionSuccessWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] startAt
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweepAction
#   @return [Hash, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] triggerCronExpression
#   @return [String, nil]
#
# @!attribute [rw] triggerOnPayIn
#   @return [Boolean, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] webHookSecret
#   @return [String, nil]
RuleListMatch = Struct.new(
  :account,
  :accountID,
  :approveUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :createdBy,
  :description,
  :endAt,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :isDisabled,
  :lastExecutedAt,
  :lastRunAtTransactionDate,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :onApprovedWebHookUrl,
  :onExecutionErrorWebHookUrl,
  :onExecutionSuccessWebHookUrl,
  :startAt,
  :status,
  :sweepAction,
  :timeZoneId,
  :triggerCronExpression,
  :triggerOnPayIn,
  :userID,
  :webHookSecret,
  keyword_init: true
)

# Request payload for Rule#create.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] endAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isDisabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastExecutedAt
#   @return [String, nil]
#
# @!attribute [rw] lastRunAtTransactionDate
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] onApprovedWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionErrorWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionSuccessWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] startAt
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweepAction
#   @return [Hash, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] triggerCronExpression
#   @return [String, nil]
#
# @!attribute [rw] triggerOnPayIn
#   @return [Boolean, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] webHookSecret
#   @return [String, nil]
RuleCreateData = Struct.new(
  :account,
  :accountID,
  :approveUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :createdBy,
  :description,
  :endAt,
  :hasCurrentUserAuthorised,
  :id,
  :inserted,
  :isDisabled,
  :lastExecutedAt,
  :lastRunAtTransactionDate,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :onApprovedWebHookUrl,
  :onExecutionErrorWebHookUrl,
  :onExecutionSuccessWebHookUrl,
  :startAt,
  :status,
  :sweepAction,
  :timeZoneId,
  :triggerCronExpression,
  :triggerOnPayIn,
  :userID,
  :webHookSecret,
  keyword_init: true
)

# Request payload for Rule#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] approveUrl
#   @return [String, nil]
#
# @!attribute [rw] approverID
#   @return [String, nil]
#
# @!attribute [rw] authenticationMethods
#   @return [Array, nil]
#
# @!attribute [rw] authorisations
#   @return [Array, nil]
#
# @!attribute [rw] authorisersCompletedCount
#   @return [Integer, nil]
#
# @!attribute [rw] authorisersRequiredCount
#   @return [Integer, nil]
#
# @!attribute [rw] canAuthorise
#   @return [Boolean, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] endAt
#   @return [String, nil]
#
# @!attribute [rw] hasCurrentUserAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isDisabled
#   @return [Boolean, nil]
#
# @!attribute [rw] lastExecutedAt
#   @return [String, nil]
#
# @!attribute [rw] lastRunAtTransactionDate
#   @return [String, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] onApprovedWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionErrorWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] onExecutionSuccessWebHookUrl
#   @return [String, nil]
#
# @!attribute [rw] startAt
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweepAction
#   @return [Hash, nil]
#
# @!attribute [rw] timeZoneId
#   @return [String, nil]
#
# @!attribute [rw] triggerCronExpression
#   @return [String, nil]
#
# @!attribute [rw] triggerOnPayIn
#   @return [Boolean, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] webHookSecret
#   @return [String, nil]
RuleUpdateData = Struct.new(
  :id,
  :account,
  :accountID,
  :approveUrl,
  :approverID,
  :authenticationMethods,
  :authorisations,
  :authorisersCompletedCount,
  :authorisersRequiredCount,
  :canAuthorise,
  :createdBy,
  :description,
  :endAt,
  :hasCurrentUserAuthorised,
  :inserted,
  :isDisabled,
  :lastExecutedAt,
  :lastRunAtTransactionDate,
  :lastUpdated,
  :merchantID,
  :name,
  :nonce,
  :onApprovedWebHookUrl,
  :onExecutionErrorWebHookUrl,
  :onExecutionSuccessWebHookUrl,
  :startAt,
  :status,
  :sweepAction,
  :timeZoneId,
  :triggerCronExpression,
  :triggerOnPayIn,
  :userID,
  :webHookSecret,
  keyword_init: true
)

# Request payload for Rule#remove.
#
# @!attribute [rw] id
#   @return [String]
RuleRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# RuleEvent entity data model.
#
# @!attribute [rw] errorMessage
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isAuthoriseToEnable
#   @return [Boolean, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] rawResponse
#   @return [String, nil]
#
# @!attribute [rw] ruleEventType
#   @return [String, nil]
#
# @!attribute [rw] ruleID
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
RuleEvent = Struct.new(
  :errorMessage,
  :id,
  :inserted,
  :isAuthoriseToEnable,
  :message,
  :rawResponse,
  :ruleEventType,
  :ruleID,
  :user,
  keyword_init: true
)

# Request payload for RuleEvent#list.
#
# @!attribute [rw] id
#   @return [String]
RuleEventListMatch = Struct.new(
  :id,
  keyword_init: true
)

# Tag entity data model.
#
# @!attribute [rw] colourHex
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String]
#
# @!attribute [rw] name
#   @return [String]
Tag = Struct.new(
  :colourHex,
  :description,
  :id,
  :merchantID,
  :name,
  keyword_init: true
)

# Request payload for Tag#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
TagListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Tag#create.
#
# @!attribute [rw] merchant_id
#   @return [String]
#
# @!attribute [rw] colourHex
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String]
#
# @!attribute [rw] name
#   @return [String]
TagCreateData = Struct.new(
  :merchant_id,
  :colourHex,
  :description,
  :id,
  :merchantID,
  :name,
  keyword_init: true
)

# Token entity data model.
class Token
end

# Request payload for Token#create.
#
# @!attribute [rw] id
#   @return [String]
TokenCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Token#remove.
#
# @!attribute [rw] id
#   @return [String]
TokenRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Transaction entity data model.
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSequenceNumber
#   @return [Integer, nil]
#
# @!attribute [rw] addressDetails
#   @return [Hash, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bookingDateTime
#   @return [String, nil]
#
# @!attribute [rw] chargeDetails
#   @return [Hash, nil]
#
# @!attribute [rw] content
#   @return [Array, nil]
#
# @!attribute [rw] counterparty
#   @return [Hash, nil]
#
# @!attribute [rw] counterpartySummary
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currencyExchange
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] enrichment
#   @return [Hash, nil]
#
# @!attribute [rw] fxAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] grossAmount
#   @return [Hash]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isoBankTransactionCode
#   @return [Hash, nil]
#
# @!attribute [rw] merchant
#   @return [Hash, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] pageNumber
#   @return [Integer, nil]
#
# @!attribute [rw] pageSize
#   @return [Integer, nil]
#
# @!attribute [rw] payeeDetails
#   @return [Hash]
#
# @!attribute [rw] payerDetails
#   @return [Hash]
#
# @!attribute [rw] paymentRequestCustomFields
#   @return [Hash, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] payoutID
#   @return [String, nil]
#
# @!attribute [rw] proprietaryBankTransactionCode
#   @return [Hash, nil]
#
# @!attribute [rw] rawReference
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] ruleID
#   @return [String, nil]
#
# @!attribute [rw] statementReferences
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplementaryData
#   @return [Object, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] totalPages
#   @return [Integer, nil]
#
# @!attribute [rw] totalSize
#   @return [Integer, nil]
#
# @!attribute [rw] transactionAmount
#   @return [Hash]
#
# @!attribute [rw] transactionDate
#   @return [String, nil]
#
# @!attribute [rw] transactionInformation
#   @return [Array, nil]
#
# @!attribute [rw] transactionMutability
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] valueDateTime
#   @return [String, nil]
#
# @!attribute [rw] virtualIBAN
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
Transaction = Struct.new(
  :accountID,
  :accountName,
  :accountSequenceNumber,
  :addressDetails,
  :amount,
  :amountMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bookingDateTime,
  :chargeDetails,
  :content,
  :counterparty,
  :counterpartySummary,
  :currency,
  :currencyExchange,
  :date,
  :description,
  :enrichment,
  :fxAmount,
  :fxCurrency,
  :fxRate,
  :grossAmount,
  :id,
  :inserted,
  :isoBankTransactionCode,
  :merchant,
  :merchantID,
  :pageNumber,
  :pageSize,
  :payeeDetails,
  :payerDetails,
  :paymentRequestCustomFields,
  :paymentRequestID,
  :payoutID,
  :proprietaryBankTransactionCode,
  :rawReference,
  :reference,
  :ruleID,
  :statementReferences,
  :status,
  :supplementaryData,
  :tags,
  :theirReference,
  :totalPages,
  :totalSize,
  :transactionAmount,
  :transactionDate,
  :transactionInformation,
  :transactionMutability,
  :type,
  :valueDateTime,
  :virtualIBAN,
  :yourReference,
  keyword_init: true
)

# Request payload for Transaction#load.
#
# @!attribute [rw] id
#   @return [String]
TransactionLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Transaction#list.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
TransactionListMatch = Struct.new(
  :account_id,
  :id,
  keyword_init: true
)

# Request payload for Transaction#create.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountID
#   @return [String, nil]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSequenceNumber
#   @return [Integer, nil]
#
# @!attribute [rw] addressDetails
#   @return [Hash, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amountMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bookingDateTime
#   @return [String, nil]
#
# @!attribute [rw] chargeDetails
#   @return [Hash, nil]
#
# @!attribute [rw] content
#   @return [Array, nil]
#
# @!attribute [rw] counterparty
#   @return [Hash, nil]
#
# @!attribute [rw] counterpartySummary
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currencyExchange
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] enrichment
#   @return [Hash, nil]
#
# @!attribute [rw] fxAmount
#   @return [Float, nil]
#
# @!attribute [rw] fxCurrency
#   @return [String, nil]
#
# @!attribute [rw] fxRate
#   @return [Float, nil]
#
# @!attribute [rw] grossAmount
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isoBankTransactionCode
#   @return [Hash, nil]
#
# @!attribute [rw] merchant
#   @return [Hash, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] pageNumber
#   @return [Integer, nil]
#
# @!attribute [rw] pageSize
#   @return [Integer, nil]
#
# @!attribute [rw] payeeDetails
#   @return [Hash]
#
# @!attribute [rw] payerDetails
#   @return [Hash]
#
# @!attribute [rw] paymentRequestCustomFields
#   @return [Hash, nil]
#
# @!attribute [rw] paymentRequestID
#   @return [String, nil]
#
# @!attribute [rw] payoutID
#   @return [String, nil]
#
# @!attribute [rw] proprietaryBankTransactionCode
#   @return [Hash, nil]
#
# @!attribute [rw] rawReference
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] ruleID
#   @return [String, nil]
#
# @!attribute [rw] statementReferences
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplementaryData
#   @return [Object, nil]
#
# @!attribute [rw] tags
#   @return [Array, nil]
#
# @!attribute [rw] theirReference
#   @return [String, nil]
#
# @!attribute [rw] totalPages
#   @return [Integer, nil]
#
# @!attribute [rw] totalSize
#   @return [Integer, nil]
#
# @!attribute [rw] transactionAmount
#   @return [Hash]
#
# @!attribute [rw] transactionDate
#   @return [String, nil]
#
# @!attribute [rw] transactionInformation
#   @return [Array, nil]
#
# @!attribute [rw] transactionMutability
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] valueDateTime
#   @return [String, nil]
#
# @!attribute [rw] virtualIBAN
#   @return [String, nil]
#
# @!attribute [rw] yourReference
#   @return [String, nil]
TransactionCreateData = Struct.new(
  :id,
  :accountID,
  :accountName,
  :accountSequenceNumber,
  :addressDetails,
  :amount,
  :amountMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bookingDateTime,
  :chargeDetails,
  :content,
  :counterparty,
  :counterpartySummary,
  :currency,
  :currencyExchange,
  :date,
  :description,
  :enrichment,
  :fxAmount,
  :fxCurrency,
  :fxRate,
  :grossAmount,
  :inserted,
  :isoBankTransactionCode,
  :merchant,
  :merchantID,
  :pageNumber,
  :pageSize,
  :payeeDetails,
  :payerDetails,
  :paymentRequestCustomFields,
  :paymentRequestID,
  :payoutID,
  :proprietaryBankTransactionCode,
  :rawReference,
  :reference,
  :ruleID,
  :statementReferences,
  :status,
  :supplementaryData,
  :tags,
  :theirReference,
  :totalPages,
  :totalSize,
  :transactionAmount,
  :transactionDate,
  :transactionInformation,
  :transactionMutability,
  :type,
  :valueDateTime,
  :virtualIBAN,
  :yourReference,
  keyword_init: true
)

# Request payload for Transaction#remove.
#
# @!attribute [rw] id
#   @return [String]
TransactionRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# User entity data model.
#
# @!attribute [rw] clientSessionTimeouts
#   @return [Array, nil]
#
# @!attribute [rw] emailAddress
#   @return [String]
#
# @!attribute [rw] firstName
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] lastName
#   @return [String]
#
# @!attribute [rw] passkeyAdded
#   @return [Boolean, nil]
#
# @!attribute [rw] permissions
#   @return [Hash, nil]
#
# @!attribute [rw] profile
#   @return [String, nil]
#
# @!attribute [rw] rolesWithScope
#   @return [Array, nil]
#
# @!attribute [rw] twoFactorEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] userInviteID
#   @return [String, nil]
User = Struct.new(
  :clientSessionTimeouts,
  :emailAddress,
  :firstName,
  :id,
  :lastName,
  :passkeyAdded,
  :permissions,
  :profile,
  :rolesWithScope,
  :twoFactorEnabled,
  :userInviteID,
  keyword_init: true
)

# Request payload for User#list.
#
# @!attribute [rw] clientSessionTimeouts
#   @return [Array, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] firstName
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] lastName
#   @return [String, nil]
#
# @!attribute [rw] passkeyAdded
#   @return [Boolean, nil]
#
# @!attribute [rw] permissions
#   @return [Hash, nil]
#
# @!attribute [rw] profile
#   @return [String, nil]
#
# @!attribute [rw] rolesWithScope
#   @return [Array, nil]
#
# @!attribute [rw] twoFactorEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] userInviteID
#   @return [String, nil]
UserListMatch = Struct.new(
  :clientSessionTimeouts,
  :emailAddress,
  :firstName,
  :id,
  :lastName,
  :passkeyAdded,
  :permissions,
  :profile,
  :rolesWithScope,
  :twoFactorEnabled,
  :userInviteID,
  keyword_init: true
)

# Request payload for User#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] clientSessionTimeouts
#   @return [Array, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] firstName
#   @return [String, nil]
#
# @!attribute [rw] lastName
#   @return [String, nil]
#
# @!attribute [rw] passkeyAdded
#   @return [Boolean, nil]
#
# @!attribute [rw] permissions
#   @return [Hash, nil]
#
# @!attribute [rw] profile
#   @return [String, nil]
#
# @!attribute [rw] rolesWithScope
#   @return [Array, nil]
#
# @!attribute [rw] twoFactorEnabled
#   @return [Boolean, nil]
#
# @!attribute [rw] userInviteID
#   @return [String, nil]
UserUpdateData = Struct.new(
  :id,
  :clientSessionTimeouts,
  :emailAddress,
  :firstName,
  :lastName,
  :passkeyAdded,
  :permissions,
  :profile,
  :rolesWithScope,
  :twoFactorEnabled,
  :userInviteID,
  keyword_init: true
)

# UserInvite entity data model.
#
# @!attribute [rw] authorisationStatus
#   @return [Hash, nil]
#
# @!attribute [rw] failedUserInvites
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] initialRoleID
#   @return [String, nil]
#
# @!attribute [rw] inviteeEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviteeFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviteeLastName
#   @return [String, nil]
#
# @!attribute [rw] inviterEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviterFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviterLastName
#   @return [String, nil]
#
# @!attribute [rw] isAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] isInviteeRegistered
#   @return [Boolean, nil]
#
# @!attribute [rw] lastInvited
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] registrationUrl
#   @return [String, nil]
#
# @!attribute [rw] sendInviteEmail
#   @return [Boolean, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] userInvites
#   @return [Array, nil]
UserInvite = Struct.new(
  :authorisationStatus,
  :failedUserInvites,
  :id,
  :initialRoleID,
  :inviteeEmailAddress,
  :inviteeFirstName,
  :inviteeLastName,
  :inviterEmailAddress,
  :inviterFirstName,
  :inviterLastName,
  :isAuthorised,
  :isInviteeRegistered,
  :lastInvited,
  :merchantID,
  :merchantName,
  :message,
  :registrationUrl,
  :sendInviteEmail,
  :status,
  :user,
  :userID,
  :userInvites,
  keyword_init: true
)

# Request payload for UserInvite#load.
#
# @!attribute [rw] id
#   @return [String]
UserInviteLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for UserInvite#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
UserInviteListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for UserInvite#create.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authorisationStatus
#   @return [Hash, nil]
#
# @!attribute [rw] failedUserInvites
#   @return [Hash, nil]
#
# @!attribute [rw] initialRoleID
#   @return [String, nil]
#
# @!attribute [rw] inviteeEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviteeFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviteeLastName
#   @return [String, nil]
#
# @!attribute [rw] inviterEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviterFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviterLastName
#   @return [String, nil]
#
# @!attribute [rw] isAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] isInviteeRegistered
#   @return [Boolean, nil]
#
# @!attribute [rw] lastInvited
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] registrationUrl
#   @return [String, nil]
#
# @!attribute [rw] sendInviteEmail
#   @return [Boolean, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] userInvites
#   @return [Array, nil]
UserInviteCreateData = Struct.new(
  :id,
  :authorisationStatus,
  :failedUserInvites,
  :initialRoleID,
  :inviteeEmailAddress,
  :inviteeFirstName,
  :inviteeLastName,
  :inviterEmailAddress,
  :inviterFirstName,
  :inviterLastName,
  :isAuthorised,
  :isInviteeRegistered,
  :lastInvited,
  :merchantID,
  :merchantName,
  :message,
  :registrationUrl,
  :sendInviteEmail,
  :status,
  :user,
  :userID,
  :userInvites,
  keyword_init: true
)

# Request payload for UserInvite#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] authorisationStatus
#   @return [Hash, nil]
#
# @!attribute [rw] failedUserInvites
#   @return [Hash, nil]
#
# @!attribute [rw] initialRoleID
#   @return [String, nil]
#
# @!attribute [rw] inviteeEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviteeFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviteeLastName
#   @return [String, nil]
#
# @!attribute [rw] inviterEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] inviterFirstName
#   @return [String, nil]
#
# @!attribute [rw] inviterLastName
#   @return [String, nil]
#
# @!attribute [rw] isAuthorised
#   @return [Boolean, nil]
#
# @!attribute [rw] isInviteeRegistered
#   @return [Boolean, nil]
#
# @!attribute [rw] lastInvited
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] registrationUrl
#   @return [String, nil]
#
# @!attribute [rw] sendInviteEmail
#   @return [Boolean, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash, nil]
#
# @!attribute [rw] userID
#   @return [String, nil]
#
# @!attribute [rw] userInvites
#   @return [Array, nil]
UserInviteUpdateData = Struct.new(
  :id,
  :authorisationStatus,
  :failedUserInvites,
  :initialRoleID,
  :inviteeEmailAddress,
  :inviteeFirstName,
  :inviteeLastName,
  :inviterEmailAddress,
  :inviterFirstName,
  :inviterLastName,
  :isAuthorised,
  :isInviteeRegistered,
  :lastInvited,
  :merchantID,
  :merchantName,
  :message,
  :registrationUrl,
  :sendInviteEmail,
  :status,
  :user,
  :userID,
  :userInvites,
  keyword_init: true
)

# Request payload for UserInvite#remove.
#
# @!attribute [rw] id
#   @return [String]
UserInviteRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Virtual entity data model.
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
Virtual = Struct.new(
  :accountName,
  :accountSupplierName,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :name,
  :physicalAccountID,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierSepaInstantStatus,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Virtual#create.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
VirtualCreateData = Struct.new(
  :account_id,
  :accountName,
  :accountSupplierName,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :id,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :name,
  :physicalAccountID,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierSepaInstantStatus,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Request payload for Virtual#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] accountName
#   @return [String, nil]
#
# @!attribute [rw] accountSupplierName
#   @return [String, nil]
#
# @!attribute [rw] availableBalance
#   @return [Float, nil]
#
# @!attribute [rw] availableBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] bankName
#   @return [String, nil]
#
# @!attribute [rw] consentID
#   @return [String, nil]
#
# @!attribute [rw] createdBy
#   @return [Hash, nil]
#
# @!attribute [rw] createdByDisplayName
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] defaultPaymentRail
#   @return [String, nil]
#
# @!attribute [rw] displayName
#   @return [String, nil]
#
# @!attribute [rw] expiryDate
#   @return [String, nil]
#
# @!attribute [rw] externalAccountIcon
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] isArchived
#   @return [Boolean, nil]
#
# @!attribute [rw] isConnectedAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isDefault
#   @return [Boolean, nil]
#
# @!attribute [rw] isTrustAccount
#   @return [Boolean, nil]
#
# @!attribute [rw] isVirtual
#   @return [Boolean, nil]
#
# @!attribute [rw] lastTransaction
#   @return [Hash, nil]
#
# @!attribute [rw] lastUpdated
#   @return [String, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] merchantName
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] physicalAccountID
#   @return [String, nil]
#
# @!attribute [rw] rules
#   @return [Array, nil]
#
# @!attribute [rw] submittedPayoutsBalance
#   @return [Float, nil]
#
# @!attribute [rw] submittedPayoutsBalanceMinorUnits
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplierSepaInstantStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedConnectionStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedLastSyncedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailedAt
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncLastFailureReason
#   @return [String, nil]
#
# @!attribute [rw] xeroBankFeedSyncStatus
#   @return [String, nil]
#
# @!attribute [rw] xeroUnsynchronisedTransactionsCount
#   @return [Integer, nil]
VirtualUpdateData = Struct.new(
  :account_id,
  :id,
  :accountName,
  :accountSupplierName,
  :availableBalance,
  :availableBalanceMinorUnits,
  :balance,
  :balanceMinorUnits,
  :bankName,
  :consentID,
  :createdBy,
  :createdByDisplayName,
  :currency,
  :defaultPaymentRail,
  :displayName,
  :expiryDate,
  :externalAccountIcon,
  :identifier,
  :inserted,
  :isArchived,
  :isConnectedAccount,
  :isDefault,
  :isTrustAccount,
  :isVirtual,
  :lastTransaction,
  :lastUpdated,
  :merchantID,
  :merchantName,
  :name,
  :physicalAccountID,
  :rules,
  :submittedPayoutsBalance,
  :submittedPayoutsBalanceMinorUnits,
  :summary,
  :supplierSepaInstantStatus,
  :xeroBankFeedConnectionStatus,
  :xeroBankFeedLastSyncedAt,
  :xeroBankFeedSyncLastFailedAt,
  :xeroBankFeedSyncLastFailureReason,
  :xeroBankFeedSyncStatus,
  :xeroUnsynchronisedTransactionsCount,
  keyword_init: true
)

# Webhook entity data model.
#
# @!attribute [rw] destinationUrl
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] failedNotificationEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] isActive
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] notificationMethod
#   @return [String, nil]
#
# @!attribute [rw] resourceTypes
#   @return [Array, nil]
#
# @!attribute [rw] retry
#   @return [Boolean, nil]
#
# @!attribute [rw] secret
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [Integer, nil]
Webhook = Struct.new(
  :destinationUrl,
  :emailAddress,
  :failedNotificationEmailAddress,
  :id,
  :isActive,
  :merchantID,
  :notificationMethod,
  :resourceTypes,
  :retry,
  :secret,
  :version,
  keyword_init: true
)

# Request payload for Webhook#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
WebhookLoadMatch = Struct.new(
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Webhook#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
WebhookListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Webhook#create.
#
# @!attribute [rw] destinationUrl
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] failedNotificationEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] isActive
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] notificationMethod
#   @return [String, nil]
#
# @!attribute [rw] resourceTypes
#   @return [Array, nil]
#
# @!attribute [rw] retry
#   @return [Boolean, nil]
#
# @!attribute [rw] secret
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [Integer, nil]
WebhookCreateData = Struct.new(
  :destinationUrl,
  :emailAddress,
  :failedNotificationEmailAddress,
  :id,
  :isActive,
  :merchantID,
  :notificationMethod,
  :resourceTypes,
  :retry,
  :secret,
  :version,
  keyword_init: true
)

# Request payload for Webhook#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] destinationUrl
#   @return [String, nil]
#
# @!attribute [rw] emailAddress
#   @return [String, nil]
#
# @!attribute [rw] failedNotificationEmailAddress
#   @return [String, nil]
#
# @!attribute [rw] isActive
#   @return [Boolean, nil]
#
# @!attribute [rw] merchantID
#   @return [String, nil]
#
# @!attribute [rw] notificationMethod
#   @return [String, nil]
#
# @!attribute [rw] resourceTypes
#   @return [Array, nil]
#
# @!attribute [rw] retry
#   @return [Boolean, nil]
#
# @!attribute [rw] secret
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [Integer, nil]
WebhookUpdateData = Struct.new(
  :id,
  :destinationUrl,
  :emailAddress,
  :failedNotificationEmailAddress,
  :isActive,
  :merchantID,
  :notificationMethod,
  :resourceTypes,
  :retry,
  :secret,
  :version,
  keyword_init: true
)

# Request payload for Webhook#remove.
#
# @!attribute [rw] id
#   @return [String]
WebhookRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

