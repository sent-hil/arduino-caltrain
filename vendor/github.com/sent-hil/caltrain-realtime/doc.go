// caltrain provides realtime train timings for various train stations in both
// directions. It does so by scraping the Caltrain mobile page.
//
// Note, Caltrain alerts are not reflected in train timings.
//
// Please don't abuse this api.
//
// Example:
//
//   package main
//
//   import (
//     "github.com/sent-hil/caltrain-realtime"
//   )
//
//   func main() {
//     timings, err := caltrain.GetRealTimings(caltrain.SanFrancisco, caltrain.SouthBound)
//   }
package caltrain
