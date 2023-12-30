/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package flutter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	projectName string
	screenName  string
)

// flutterCmd represents the flutter command
var FlutterCmd = &cobra.Command{
	Use:   "flutter",
	Short: "A brief description of your command",
	Long:  `A Flutter CLI for making dope shit`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		flutterCLI("--version", "")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		flutterCLI("create", projectName)
	},
}

var createScreen = &cobra.Command{
	Use:   "screen",
	Short: "Create A New Screen Component",
	Long:  `This creates a new flutter bloc module`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Command value name: %v \n", screenName)

		createScreenFile(screenName)
		// createCubitFile(screenName)
	},
}

func createScreenFile(screenName string) {
	screenNameLower := strings.ToLower(screenName)
	cmd := exec.Command("mkdir", screenNameLower)
	_, err := cmd.CombinedOutput()
	check(err)

	screenData := fmt.Sprintf(`
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class %vScreen extends StatelessWidget {
	const %vScreen({super.key});

	@override
	Widget build(BuildContext context) {
		return BlocBuilder<%vCubit, %vVM>(
			builder: (context, state) {
				final cubit = context.read<%vCubit>();
				return Scaffold(
					body: Center(
						child: Text("%v")
					)
				);
			}
		);
	}
}
	`, screenName, screenName, screenName, screenName, screenName, screenName)

	cubitData := fmt.Sprintf(`
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:supabase_flutter/supabase_flutter.dart';

class %vCubit extends Cubit<%vVM> {
	final supabase = Supabase.instance.client;
	
	%vCubit() : super(%vVM());
	
}

class %vVM {
	%vVM.init()

	%vVM({})
}
	`, screenName, screenName, screenName, screenName, screenName, screenName, screenName)

	screen := []byte(screenData)
	cubit := []byte(cubitData)
	screenFile := fmt.Sprintf("%s/%s_screen.dart", screenNameLower, screenNameLower)
	cubitFile := fmt.Sprintf("%s/%s_cubit.dart", screenNameLower, screenNameLower)
	err1 := os.WriteFile(screenFile, screen, 0644)
	check(err1)
	err2 := os.WriteFile(cubitFile, cubit, 0644)
	check(err2)

	// f, err := os.Create(cubit)
	// check(err)
	// defer f.Close()

	// n, err := f.WriteString(`
	// 	Screen....
	// `)
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n)
}

func createCubitFile(sceenName string) {

	cubit := fmt.Sprintf("/%s/%s_cubit.dart", sceenName, sceenName)
	f, err := os.Create(cubit)
	check(err)
	defer f.Close()

	n, err := f.WriteString(`
		Cubit.... 
	`)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
}

func flutterCLI(command string, args string) {
	cmd := exec.Command("flutter", command, args)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Flutter function failed: %v", err)
	}
	fmt.Printf("%s\n", b)
}

func check(e error) {
	if e != nil {
		if e != nil {
			log.Printf("Failed to launch failed: %v", e)
		}
		panic(e)
	}
}

func init() {
	FlutterCmd.AddCommand(versionCmd)
	FlutterCmd.AddCommand(createCmd)
	FlutterCmd.AddCommand(createScreen)

	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "The name of the project")
	createScreen.Flags().StringVarP(&screenName, "screen", "s", "", "The screen name")

	if err := createCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flutterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flutterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
