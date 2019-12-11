package main
import (
    "flag"
    "os"
    "strings"
)
var (
    // EnvSpecificValue is an env specific value
    EnvSpecificValue string
)
func init() {
    flag.Parse()
    environmentVar(&EnvSpecificValue, "ENV_SPECIFIC_VALUE", "Local")
}
func environmentVar(v *string, key, defaultValue string) {
    *v = os.Getenv(key)
    if "" == strings.TrimSpace(*v) {
        *v = defaultValue
    }
}