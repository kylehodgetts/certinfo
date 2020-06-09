package main

import (
	"flag"
	"os"
	"strconv"
)

type Flags struct {
	Expiry   bool
	Insecure bool
	Version  bool
	Args     []string
}

func ParseFlags() (Flags, error) {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	expiry := f.Bool("expiry", getBoolEnv("CERTINFO_EXPIRY", false),
		"print expiry of certificates")
	insecure := f.Bool("insecure", getBoolEnv("CERTINFO_INSECURE", false),
		"whether a client verifies the server's certificate chain and host name")
	version := f.Bool("version", getBoolEnv("CERTINFO_VERSION", false),
		"certinfo version")

	if err := f.Parse(os.Args[1:]); err != nil {
		return Flags{}, err
	}

	return Flags{
		Expiry:   boolValue(expiry),
		Insecure: boolValue(insecure),
		Version:  boolValue(version),
		Args:     f.Args(),
	}, nil
}

func getBoolEnv(envName string, defaultValue bool) bool {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if intValue, err := strconv.ParseBool(env); err == nil {
		return intValue
	}
	return defaultValue
}

func boolValue(v *bool) bool {

	if v == nil {
		return false
	}
	return *v
}
