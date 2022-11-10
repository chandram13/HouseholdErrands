// Marvish Chandra

package HouseholdErrands

import "fmt" 


func givenErrands(chore){
    if chore == "meal prep":
    fmt.Println("Make sure to cook for part or the entire week after getting your groceries.")
    if chore == "clean dishes":
    fmt.Println("Either use your dishwasher at the end of the day or routinely wash your dishes frequently.")
    if chore == "cleaning the floor":
    fmt.Println("Make sure to vaccuum the carpet and mop any hardwood floors in your place.")
    if chore == "cleaning the trash":
    fmt.Println("Routinely take out any trash bags you see around your place and safely dispose of them.")
    if chore == "cleaning the bathroom":
    fmt.Println("Clean the area around your bathroom with a mop and scrub your tub for any grime or built up bacteria.")
    if chore == "cleaning the windows":
    fmt.Println("Have a bottle of window cleaner and some paper towels to wipe each of the windows in your place.")
    else: fmt.Println("Either you have completed the necessary chores or the given chore is not required.")
}