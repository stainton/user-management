/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/stainton/user-management/cmd"

func main() {
	runner := cmd.NewUserManager()
	runner.Execute()
}
