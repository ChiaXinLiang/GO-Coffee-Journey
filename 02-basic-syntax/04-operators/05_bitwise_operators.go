package main

import "fmt"

// Store features as bit flags
const (
    FeatureNone       = 0
    FeatureWifi       = 1 << 0  // 1
    FeatureParking    = 1 << 1  // 2
    FeatureDriveThru  = 1 << 2  // 4
    FeatureOutdoor    = 1 << 3  // 8
    FeaturePetFriendly = 1 << 4 // 16
    FeatureDelivery   = 1 << 5  // 32
)

// Day flags for special offers
const (
    Monday    = 1 << 0 // 1
    Tuesday   = 1 << 1 // 2
    Wednesday = 1 << 2 // 4
    Thursday  = 1 << 3 // 8
    Friday    = 1 << 4 // 16
    Saturday  = 1 << 5 // 32
    Sunday    = 1 << 6 // 64
)

func main() {
    fmt.Println("=== GoCoffee Bitwise Operations ===\n")
    
    // Store features using bitwise OR
    downtownStore := FeatureWifi | FeatureParking | FeatureOutdoor
    airportStore := FeatureWifi | FeatureDriveThru | FeatureDelivery
    suburbStore := FeatureWifi | FeatureParking | FeatureDriveThru | FeaturePetFriendly
    
    fmt.Println("STORE FEATURES:")
    fmt.Printf("Downtown: %b (%d)\n", downtownStore, downtownStore)
    displayFeatures("Downtown", downtownStore)
    
    fmt.Printf("\nAirport: %b (%d)\n", airportStore, airportStore)
    displayFeatures("Airport", airportStore)
    
    fmt.Printf("\nSuburb: %b (%d)\n", suburbStore, suburbStore)
    displayFeatures("Suburb", suburbStore)
    
    // Check for specific features using AND
    fmt.Println("\nFEATURE QUERIES:")
    
    stores := map[string]int{
        "Downtown": downtownStore,
        "Airport":  airportStore,
        "Suburb":   suburbStore,
    }
    
    // Which stores have parking?
    fmt.Println("Stores with parking:")
    for name, features := range stores {
        if features&FeatureParking != 0 {
            fmt.Printf("  ✓ %s\n", name)
        }
    }
    
    // Which stores have both WiFi AND Drive-thru?
    fmt.Println("\nStores with WiFi AND Drive-thru:")
    requiredFeatures := FeatureWifi | FeatureDriveThru
    for name, features := range stores {
        if features&requiredFeatures == requiredFeatures {
            fmt.Printf("  ✓ %s\n", name)
        }
    }
    
    // Toggle features using XOR
    fmt.Println("\nTOGGLING FEATURES:")
    
    testStore := FeatureWifi | FeatureParking
    fmt.Printf("Original: %b\n", testStore)
    displayFeatures("Test Store", testStore)
    
    // Toggle outdoor seating
    testStore ^= FeatureOutdoor
    fmt.Printf("\nAfter adding outdoor: %b\n", testStore)
    displayFeatures("Test Store", testStore)
    
    // Toggle outdoor again (removes it)
    testStore ^= FeatureOutdoor
    fmt.Printf("\nAfter toggling outdoor again: %b\n", testStore)
    displayFeatures("Test Store", testStore)
    
    // Clear features using AND NOT
    fmt.Println("\nREMOVING FEATURES:")
    
    // Remove parking from suburb store
    modifiedSuburb := suburbStore &^ FeatureParking
    fmt.Printf("Suburb without parking: %b\n", modifiedSuburb)
    displayFeatures("Modified Suburb", modifiedSuburb)
    
    // Shift operations for time slots
    fmt.Println("\nTIME SLOT RESERVATIONS:")
    
    // Each bit represents a 30-minute slot (9 AM - 5 PM = 16 slots)
    var reservations uint16 = 0
    
    // Book slot 3 (10:30-11:00)
    reservations |= 1 << 3
    
    // Book slot 7 (12:30-1:00)
    reservations |= 1 << 7
    
    // Book slots 10-11 (2:00-3:00)
    reservations |= 1 << 10
    reservations |= 1 << 11
    
    fmt.Printf("Reservations: %016b\n", reservations)
    fmt.Println("Time slots:")
    displayTimeSlots(reservations)
    
    // Special offer days
    fmt.Println("\nSPECIAL OFFER DAYS:")
    
    happyHourDays := Monday | Tuesday | Wednesday | Thursday | Friday
    weekendBrunch := Saturday | Sunday
    seniorDiscount := Tuesday | Thursday
    
    fmt.Printf("Happy Hour (weekdays): %07b\n", happyHourDays)
    fmt.Printf("Weekend Brunch: %07b\n", weekendBrunch)
    fmt.Printf("Senior Discount: %07b\n", seniorDiscount)
    
    // Check what offers are available on Tuesday
    checkOffers("Tuesday", Tuesday, happyHourDays, weekendBrunch, seniorDiscount)
    checkOffers("Saturday", Saturday, happyHourDays, weekendBrunch, seniorDiscount)
}

func displayFeatures(storeName string, features int) {
    fmt.Printf("%s features:\n", storeName)
    
    featureMap := map[int]string{
        FeatureWifi:       "WiFi",
        FeatureParking:    "Parking",
        FeatureDriveThru:  "Drive-Thru",
        FeatureOutdoor:    "Outdoor Seating",
        FeaturePetFriendly: "Pet Friendly",
        FeatureDelivery:   "Delivery",
    }
    
    for flag, name := range featureMap {
        if features&flag != 0 {
            fmt.Printf("  ✓ %s\n", name)
        }
    }
}

func displayTimeSlots(reservations uint16) {
    startHour := 9
    for slot := 0; slot < 16; slot++ {
        if reservations&(1<<slot) != 0 {
            hour := startHour + slot/2
            minute := 0
            if slot%2 == 1 {
                minute = 30
            }
            fmt.Printf("  Slot %2d: %2d:%02d - %2d:%02d ✓\n", 
                slot, hour, minute, hour, minute+30)
        }
    }
}

func checkOffers(day string, dayFlag, happy, brunch, senior int) {
    fmt.Printf("\n%s offers:\n", day)
    if dayFlag&happy != 0 {
        fmt.Println("  ✓ Happy Hour (3-5 PM)")
    }
    if dayFlag&brunch != 0 {
        fmt.Println("  ✓ Weekend Brunch Special")
    }
    if dayFlag&senior != 0 {
        fmt.Println("  ✓ Senior Discount (10%)")
    }
}