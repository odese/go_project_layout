package utils

// WorkEnvironment ...
// OPTINAL VALUES: "LOCAL_DEV", "LOCAL_PROD", "DEV", "PROD"
// 'LOCAL' changes only the output format of the logger etc. Nothing else.
var WorkEnvironment string = GetEnvValue(WorkEnvironmentKey)
